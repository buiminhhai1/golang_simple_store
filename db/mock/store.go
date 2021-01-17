// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thuhangpham/simplestores/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	db "github.com/thuhangpham/simplestores/db/sqlc"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBook mocks base method
func (m *MockStore) CreateBook(arg0 context.Context, arg1 db.CreateBookParams) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook
func (mr *MockStoreMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockStore)(nil).CreateBook), arg0, arg1)
}

// CreateCategory mocks base method
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.BookCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.BookCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateOrder mocks base method
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 db.CreateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreateOrderItem mocks base method
func (m *MockStore) CreateOrderItem(arg0 context.Context, arg1 db.CreateOrderItemParams) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItem indicates an expected call of CreateOrderItem
func (mr *MockStoreMockRecorder) CreateOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockStore)(nil).CreateOrderItem), arg0, arg1)
}

// CreateUser mocks base method
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteBook mocks base method
func (m *MockStore) DeleteBook(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook
func (mr *MockStoreMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockStore)(nil).DeleteBook), arg0, arg1)
}

// DeleteCategory mocks base method
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteOrderById mocks base method
func (m *MockStore) DeleteOrderById(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderById indicates an expected call of DeleteOrderById
func (mr *MockStoreMockRecorder) DeleteOrderById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderById", reflect.TypeOf((*MockStore)(nil).DeleteOrderById), arg0, arg1)
}

// DeleteOrderItem mocks base method
func (m *MockStore) DeleteOrderItem(arg0 context.Context, arg1 db.DeleteOrderItemParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderItem indicates an expected call of DeleteOrderItem
func (mr *MockStoreMockRecorder) DeleteOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderItem", reflect.TypeOf((*MockStore)(nil).DeleteOrderItem), arg0, arg1)
}

// DeleteUser mocks base method
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetBookById mocks base method
func (m *MockStore) GetBookById(arg0 context.Context, arg1 int32) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById
func (mr *MockStoreMockRecorder) GetBookById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockStore)(nil).GetBookById), arg0, arg1)
}

// GetBooksByCategoryId mocks base method
func (m *MockStore) GetBooksByCategoryId(arg0 context.Context, arg1 db.GetBooksByCategoryIdParams) ([]db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByCategoryId", arg0, arg1)
	ret0, _ := ret[0].([]db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByCategoryId indicates an expected call of GetBooksByCategoryId
func (mr *MockStoreMockRecorder) GetBooksByCategoryId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByCategoryId", reflect.TypeOf((*MockStore)(nil).GetBooksByCategoryId), arg0, arg1)
}

// GetCategory mocks base method
func (m *MockStore) GetCategory(arg0 context.Context, arg1 int32) (db.BookCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.BookCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetOrderById mocks base method
func (m *MockStore) GetOrderById(arg0 context.Context, arg1 int32) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById
func (mr *MockStoreMockRecorder) GetOrderById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockStore)(nil).GetOrderById), arg0, arg1)
}

// GetOrderItemByPrimaryKey mocks base method
func (m *MockStore) GetOrderItemByPrimaryKey(arg0 context.Context, arg1 db.GetOrderItemByPrimaryKeyParams) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemByPrimaryKey", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemByPrimaryKey indicates an expected call of GetOrderItemByPrimaryKey
func (mr *MockStoreMockRecorder) GetOrderItemByPrimaryKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemByPrimaryKey", reflect.TypeOf((*MockStore)(nil).GetOrderItemByPrimaryKey), arg0, arg1)
}

// GetOrderItemsByBookID mocks base method
func (m *MockStore) GetOrderItemsByBookID(arg0 context.Context, arg1 db.GetOrderItemsByBookIDParams) ([]db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemsByBookID", arg0, arg1)
	ret0, _ := ret[0].([]db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemsByBookID indicates an expected call of GetOrderItemsByBookID
func (mr *MockStoreMockRecorder) GetOrderItemsByBookID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemsByBookID", reflect.TypeOf((*MockStore)(nil).GetOrderItemsByBookID), arg0, arg1)
}

// GetOrderItemsByOrderID mocks base method
func (m *MockStore) GetOrderItemsByOrderID(arg0 context.Context, arg1 db.GetOrderItemsByOrderIDParams) ([]db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemsByOrderID", arg0, arg1)
	ret0, _ := ret[0].([]db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemsByOrderID indicates an expected call of GetOrderItemsByOrderID
func (mr *MockStoreMockRecorder) GetOrderItemsByOrderID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemsByOrderID", reflect.TypeOf((*MockStore)(nil).GetOrderItemsByOrderID), arg0, arg1)
}

// GetOrdersByUserId mocks base method
func (m *MockStore) GetOrdersByUserId(arg0 context.Context, arg1 db.GetOrdersByUserIdParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByUserId", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByUserId indicates an expected call of GetOrdersByUserId
func (mr *MockStoreMockRecorder) GetOrdersByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByUserId", reflect.TypeOf((*MockStore)(nil).GetOrdersByUserId), arg0, arg1)
}

// GetUser mocks base method
func (m *MockStore) GetUser(arg0 context.Context, arg1 int32) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByUserName mocks base method
func (m *MockStore) GetUserByUserName(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUserName", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUserName indicates an expected call of GetUserByUserName
func (mr *MockStoreMockRecorder) GetUserByUserName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUserName", reflect.TypeOf((*MockStore)(nil).GetUserByUserName), arg0, arg1)
}

// ListBookCategories mocks base method
func (m *MockStore) ListBookCategories(arg0 context.Context, arg1 db.ListBookCategoriesParams) ([]db.BookCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBookCategories", arg0, arg1)
	ret0, _ := ret[0].([]db.BookCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBookCategories indicates an expected call of ListBookCategories
func (mr *MockStoreMockRecorder) ListBookCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookCategories", reflect.TypeOf((*MockStore)(nil).ListBookCategories), arg0, arg1)
}

// ListBooks mocks base method
func (m *MockStore) ListBooks(arg0 context.Context, arg1 db.ListBooksParams) ([]db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBooks", arg0, arg1)
	ret0, _ := ret[0].([]db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBooks indicates an expected call of ListBooks
func (mr *MockStoreMockRecorder) ListBooks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBooks", reflect.TypeOf((*MockStore)(nil).ListBooks), arg0, arg1)
}

// ListOrders mocks base method
func (m *MockStore) ListOrders(arg0 context.Context, arg1 db.ListOrdersParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders
func (mr *MockStoreMockRecorder) ListOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockStore)(nil).ListOrders), arg0, arg1)
}

// ListUsers mocks base method
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// OrderItemTx mocks base method
func (m *MockStore) OrderItemTx(arg0 context.Context, arg1 db.OrderItemTxParams) (db.OrderItemTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderItemTx", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItemTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderItemTx indicates an expected call of OrderItemTx
func (mr *MockStoreMockRecorder) OrderItemTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderItemTx", reflect.TypeOf((*MockStore)(nil).OrderItemTx), arg0, arg1)
}

// UpdateBook mocks base method
func (m *MockStore) UpdateBook(arg0 context.Context, arg1 db.UpdateBookParams) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook
func (mr *MockStoreMockRecorder) UpdateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockStore)(nil).UpdateBook), arg0, arg1)
}

// UpdateCategory mocks base method
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (db.BookCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.BookCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateOrderPriceInvoice mocks base method
func (m *MockStore) UpdateOrderPriceInvoice(arg0 context.Context, arg1 db.UpdateOrderPriceInvoiceParams) (db.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderPriceInvoice", arg0, arg1)
	ret0, _ := ret[0].(db.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderPriceInvoice indicates an expected call of UpdateOrderPriceInvoice
func (mr *MockStoreMockRecorder) UpdateOrderPriceInvoice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderPriceInvoice", reflect.TypeOf((*MockStore)(nil).UpdateOrderPriceInvoice), arg0, arg1)
}

// UpdateOrderStatus mocks base method
func (m *MockStore) UpdateOrderStatus(arg0 context.Context, arg1 db.UpdateOrderStatusParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus
func (mr *MockStoreMockRecorder) UpdateOrderStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockStore)(nil).UpdateOrderStatus), arg0, arg1)
}

// UpdateQuantityBook mocks base method
func (m *MockStore) UpdateQuantityBook(arg0 context.Context, arg1 db.UpdateQuantityBookParams) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityBook", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuantityBook indicates an expected call of UpdateQuantityBook
func (mr *MockStoreMockRecorder) UpdateQuantityBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityBook", reflect.TypeOf((*MockStore)(nil).UpdateQuantityBook), arg0, arg1)
}

// UpdateUser mocks base method
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
