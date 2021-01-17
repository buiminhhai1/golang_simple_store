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
