package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/thuhangpham/simplestores/db/mock"
	db "github.com/thuhangpham/simplestores/db/sqlc"
	"github.com/thuhangpham/simplestores/util"
)

func randomUser() db.User {
	return db.User{
		ID:        int32(util.RandomInt(1, 1000)),
		Username:  util.RandomString(20),
		Password:  util.RandomString(8),
		Role:      db.Role(util.RandomRole()),
		Email:     util.RandomEmail(),
		FullName:  util.RandomString(20),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestGetUserAPI(t *testing.T) {
	user := randomUser()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// build stubs
	store.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(user.ID)).
		Times(1).
		Return(user, nil)

	// start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%v", user.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)

	// check responsive
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchUser(t, recorder.Body, user)
}

func TestVerifyUser(t *testing.T) {
	// testCases := []struct {
	// 	name          string
	// 	body          gin.H
	// 	buildStubs    func(store *mockdb.MockStore)
	// 	checkResponse func(recoder *httptest.ResponseRecorder)
	// }{
	// 	{
	// 		name: "OK",
	// 		body: gin.H{
	// 			"owner":    account.Owner,
	// 			"currency": account.Currency,
	// 		},
	// 		buildStubs: func(store *mockdb.MockStore) {
	// 			arg := db.CreateAccountParams{
	// 				Owner:    account.Owner,
	// 				Currency: account.Currency,
	// 				Balance:  0,
	// 			}

	// 			store.EXPECT().
	// 				CreateAccount(gomock.Any(), gomock.Eq(arg)).
	// 				Times(1).
	// 				Return(account, nil)
	// 		},
	// 		checkResponse: func(recorder *httptest.ResponseRecorder) {
	// 			require.Equal(t, http.StatusOK, recorder.Code)
	// 			requireBodyMatchAccount(t, recorder.Body, account)
	// 		},
	// 	},
	// 	{
	// 		name: "InternalError",
	// 		body: gin.H{
	// 			"owner":    account.Owner,
	// 			"currency": account.Currency,
	// 		},
	// 		buildStubs: func(store *mockdb.MockStore) {
	// 			store.EXPECT().
	// 				CreateAccount(gomock.Any(), gomock.Any()).
	// 				Times(1).
	// 				Return(db.Account{}, sql.ErrConnDone)
	// 		},
	// 		checkResponse: func(recorder *httptest.ResponseRecorder) {
	// 			require.Equal(t, http.StatusInternalServerError, recorder.Code)
	// 		},
	// 	},
	// 	{
	// 		name: "InvalidCurrency",
	// 		body: gin.H{
	// 			"owner":    account.Owner,
	// 			"currency": "invalid",
	// 		},
	// 		buildStubs: func(store *mockdb.MockStore) {
	// 			store.EXPECT().
	// 				CreateAccount(gomock.Any(), gomock.Any()).
	// 				Times(0)
	// 		},
	// 		checkResponse: func(recorder *httptest.ResponseRecorder) {
	// 			require.Equal(t, http.StatusBadRequest, recorder.Code)
	// 		},
	// 	},
	// 	{
	// 		name: "InvalidOwner",
	// 		body: gin.H{
	// 			"owner":    "",
	// 			"currency": account.Currency,
	// 		},
	// 		buildStubs: func(store *mockdb.MockStore) {
	// 			store.EXPECT().
	// 				CreateAccount(gomock.Any(), gomock.Any()).
	// 				Times(0)
	// 		},
	// 		checkResponse: func(recorder *httptest.ResponseRecorder) {
	// 			require.Equal(t, http.StatusBadRequest, recorder.Code)
	// 		},
	// 	},
	// }

	// for i := range testCases {
	// 	tc := testCases[i]

	// 	t.Run(tc.name, func(t *testing.T) {
	// 		ctrl := gomock.NewController(t)
	// 		defer ctrl.Finish()

	// 		store := mockdb.NewMockStore(ctrl)
	// 		tc.buildStubs(store)

	// 		server := NewServer(store)
	// 		recorder := httptest.NewRecorder()

	// 		// Marshal body data to JSON
	// 		data, err := json.Marshal(tc.body)
	// 		require.NoError(t, err)

	// 		url := "/accounts"
	// 		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	// 		require.NoError(t, err)

	// 		server.router.ServeHTTP(recorder, request)
	// 		tc.checkResponse(recorder)
	// 	})
	// }
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	require.Equal(t, user.ID, gotUser.ID)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Password, gotUser.Password)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Role, gotUser.Role)

}

func TestGenerateToken(t *testing.T) {
	config, err := util.LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	jwtWrapper := JwtWrapper{
		SecretKey:       config.JwtSecret,
		Issuer:          config.JwtSecret,
		ExpirationHours: 48,
	}

	user := randomUser()
	require.NotEmpty(t, user)

	generatedToken, err := jwtWrapper.GenerateToken(user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, generatedToken)

	claims, err := jwtWrapper.ValidateToken(generatedToken)
	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.Equal(t, claims.Username, user.Username)
}
