// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=mock_interfaces.go -package=store -typed
//

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	meta "github.com/jialeicui/feedpilot/pkg/meta"
	gomock "go.uber.org/mock/gomock"
)

// MockStringer is a mock of Stringer interface.
type MockStringer struct {
	ctrl     *gomock.Controller
	recorder *MockStringerMockRecorder
	isgomock struct{}
}

// MockStringerMockRecorder is the mock recorder for MockStringer.
type MockStringerMockRecorder struct {
	mock *MockStringer
}

// NewMockStringer creates a new mock instance.
func NewMockStringer(ctrl *gomock.Controller) *MockStringer {
	mock := &MockStringer{ctrl: ctrl}
	mock.recorder = &MockStringerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStringer) EXPECT() *MockStringerMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockStringer) Load(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockStringerMockRecorder) Load(arg0 any) *MockStringerLoadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockStringer)(nil).Load), arg0)
	return &MockStringerLoadCall{Call: call}
}

// MockStringerLoadCall wrap *gomock.Call
type MockStringerLoadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringerLoadCall) Return(arg0 error) *MockStringerLoadCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringerLoadCall) Do(f func(string) error) *MockStringerLoadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringerLoadCall) DoAndReturn(f func(string) error) *MockStringerLoadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// String mocks base method.
func (m *MockStringer) String() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// String indicates an expected call of String.
func (mr *MockStringerMockRecorder) String() *MockStringerStringCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockStringer)(nil).String))
	return &MockStringerStringCall{Call: call}
}

// MockStringerStringCall wrap *gomock.Call
type MockStringerStringCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringerStringCall) Return(arg0 string, arg1 error) *MockStringerStringCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringerStringCall) Do(f func() (string, error)) *MockStringerStringCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringerStringCall) DoAndReturn(f func() (string, error)) *MockStringerStringCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUserStore is a mock of UserStore interface.
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
	isgomock struct{}
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore.
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance.
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserStore) Delete(ID meta.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserStoreMockRecorder) Delete(ID any) *MockUserStoreDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserStore)(nil).Delete), ID)
	return &MockUserStoreDeleteCall{Call: call}
}

// MockUserStoreDeleteCall wrap *gomock.Call
type MockUserStoreDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserStoreDeleteCall) Return(arg0 error) *MockUserStoreDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserStoreDeleteCall) Do(f func(meta.UserID) error) *MockUserStoreDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserStoreDeleteCall) DoAndReturn(f func(meta.UserID) error) *MockUserStoreDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockUserStore) Get(ID meta.UserID) (*meta.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ID)
	ret0, _ := ret[0].(*meta.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserStoreMockRecorder) Get(ID any) *MockUserStoreGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserStore)(nil).Get), ID)
	return &MockUserStoreGetCall{Call: call}
}

// MockUserStoreGetCall wrap *gomock.Call
type MockUserStoreGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserStoreGetCall) Return(arg0 *meta.User, arg1 error) *MockUserStoreGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserStoreGetCall) Do(f func(meta.UserID) (*meta.User, error)) *MockUserStoreGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserStoreGetCall) DoAndReturn(f func(meta.UserID) (*meta.User, error)) *MockUserStoreGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockUserStore) List(offset, limit int) ([]*meta.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", offset, limit)
	ret0, _ := ret[0].([]*meta.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserStoreMockRecorder) List(offset, limit any) *MockUserStoreListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserStore)(nil).List), offset, limit)
	return &MockUserStoreListCall{Call: call}
}

// MockUserStoreListCall wrap *gomock.Call
type MockUserStoreListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserStoreListCall) Return(arg0 []*meta.User, arg1 error) *MockUserStoreListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserStoreListCall) Do(f func(int, int) ([]*meta.User, error)) *MockUserStoreListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserStoreListCall) DoAndReturn(f func(int, int) ([]*meta.User, error)) *MockUserStoreListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Put mocks base method.
func (m *MockUserStore) Put(ID meta.UserID, user *meta.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockUserStoreMockRecorder) Put(ID, user any) *MockUserStorePutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockUserStore)(nil).Put), ID, user)
	return &MockUserStorePutCall{Call: call}
}

// MockUserStorePutCall wrap *gomock.Call
type MockUserStorePutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserStorePutCall) Return(arg0 error) *MockUserStorePutCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserStorePutCall) Do(f func(meta.UserID, *meta.User) error) *MockUserStorePutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserStorePutCall) DoAndReturn(f func(meta.UserID, *meta.User) error) *MockUserStorePutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockPostStore is a mock of PostStore interface.
type MockPostStore struct {
	ctrl     *gomock.Controller
	recorder *MockPostStoreMockRecorder
	isgomock struct{}
}

// MockPostStoreMockRecorder is the mock recorder for MockPostStore.
type MockPostStoreMockRecorder struct {
	mock *MockPostStore
}

// NewMockPostStore creates a new mock instance.
func NewMockPostStore(ctrl *gomock.Controller) *MockPostStore {
	mock := &MockPostStore{ctrl: ctrl}
	mock.recorder = &MockPostStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostStore) EXPECT() *MockPostStoreMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPostStore) Delete(ID meta.PostID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostStoreMockRecorder) Delete(ID any) *MockPostStoreDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostStore)(nil).Delete), ID)
	return &MockPostStoreDeleteCall{Call: call}
}

// MockPostStoreDeleteCall wrap *gomock.Call
type MockPostStoreDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPostStoreDeleteCall) Return(arg0 error) *MockPostStoreDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPostStoreDeleteCall) Do(f func(meta.PostID) error) *MockPostStoreDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPostStoreDeleteCall) DoAndReturn(f func(meta.PostID) error) *MockPostStoreDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockPostStore) Get(ID meta.PostID) (*meta.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ID)
	ret0, _ := ret[0].(*meta.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostStoreMockRecorder) Get(ID any) *MockPostStoreGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostStore)(nil).Get), ID)
	return &MockPostStoreGetCall{Call: call}
}

// MockPostStoreGetCall wrap *gomock.Call
type MockPostStoreGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPostStoreGetCall) Return(arg0 *meta.Post, arg1 error) *MockPostStoreGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPostStoreGetCall) Do(f func(meta.PostID) (*meta.Post, error)) *MockPostStoreGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPostStoreGetCall) DoAndReturn(f func(meta.PostID) (*meta.Post, error)) *MockPostStoreGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockPostStore) List(offset, limit int) ([]*meta.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", offset, limit)
	ret0, _ := ret[0].([]*meta.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPostStoreMockRecorder) List(offset, limit any) *MockPostStoreListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPostStore)(nil).List), offset, limit)
	return &MockPostStoreListCall{Call: call}
}

// MockPostStoreListCall wrap *gomock.Call
type MockPostStoreListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPostStoreListCall) Return(arg0 []*meta.Post, arg1 error) *MockPostStoreListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPostStoreListCall) Do(f func(int, int) ([]*meta.Post, error)) *MockPostStoreListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPostStoreListCall) DoAndReturn(f func(int, int) ([]*meta.Post, error)) *MockPostStoreListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Put mocks base method.
func (m *MockPostStore) Put(ID meta.PostID, post *meta.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ID, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockPostStoreMockRecorder) Put(ID, post any) *MockPostStorePutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPostStore)(nil).Put), ID, post)
	return &MockPostStorePutCall{Call: call}
}

// MockPostStorePutCall wrap *gomock.Call
type MockPostStorePutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPostStorePutCall) Return(arg0 error) *MockPostStorePutCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPostStorePutCall) Do(f func(meta.PostID, *meta.Post) error) *MockPostStorePutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPostStorePutCall) DoAndReturn(f func(meta.PostID, *meta.Post) error) *MockPostStorePutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockObjectStore is a mock of ObjectStore interface.
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
	isgomock struct{}
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore.
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance.
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockObjectStore) Get(key string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockObjectStoreMockRecorder) Get(key any) *MockObjectStoreGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockObjectStore)(nil).Get), key)
	return &MockObjectStoreGetCall{Call: call}
}

// MockObjectStoreGetCall wrap *gomock.Call
type MockObjectStoreGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockObjectStoreGetCall) Return(arg0 []byte) *MockObjectStoreGetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockObjectStoreGetCall) Do(f func(string) []byte) *MockObjectStoreGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockObjectStoreGetCall) DoAndReturn(f func(string) []byte) *MockObjectStoreGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Put mocks base method.
func (m *MockObjectStore) Put(key string, value []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", key, value)
}

// Put indicates an expected call of Put.
func (mr *MockObjectStoreMockRecorder) Put(key, value any) *MockObjectStorePutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockObjectStore)(nil).Put), key, value)
	return &MockObjectStorePutCall{Call: call}
}

// MockObjectStorePutCall wrap *gomock.Call
type MockObjectStorePutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockObjectStorePutCall) Return() *MockObjectStorePutCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockObjectStorePutCall) Do(f func(string, []byte)) *MockObjectStorePutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockObjectStorePutCall) DoAndReturn(f func(string, []byte)) *MockObjectStorePutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockKvStore is a mock of KvStore interface.
type MockKvStore struct {
	ctrl     *gomock.Controller
	recorder *MockKvStoreMockRecorder
	isgomock struct{}
}

// MockKvStoreMockRecorder is the mock recorder for MockKvStore.
type MockKvStoreMockRecorder struct {
	mock *MockKvStore
}

// NewMockKvStore creates a new mock instance.
func NewMockKvStore(ctrl *gomock.Controller) *MockKvStore {
	mock := &MockKvStore{ctrl: ctrl}
	mock.recorder = &MockKvStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKvStore) EXPECT() *MockKvStoreMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockKvStore) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockKvStoreMockRecorder) Delete(key any) *MockKvStoreDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKvStore)(nil).Delete), key)
	return &MockKvStoreDeleteCall{Call: call}
}

// MockKvStoreDeleteCall wrap *gomock.Call
type MockKvStoreDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKvStoreDeleteCall) Return(arg0 error) *MockKvStoreDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKvStoreDeleteCall) Do(f func(string) error) *MockKvStoreDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKvStoreDeleteCall) DoAndReturn(f func(string) error) *MockKvStoreDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockKvStore) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKvStoreMockRecorder) Get(key any) *MockKvStoreGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKvStore)(nil).Get), key)
	return &MockKvStoreGetCall{Call: call}
}

// MockKvStoreGetCall wrap *gomock.Call
type MockKvStoreGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKvStoreGetCall) Return(arg0 string, arg1 error) *MockKvStoreGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKvStoreGetCall) Do(f func(string) (string, error)) *MockKvStoreGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKvStoreGetCall) DoAndReturn(f func(string) (string, error)) *MockKvStoreGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockKvStore) List(offset, limit int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", offset, limit)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockKvStoreMockRecorder) List(offset, limit any) *MockKvStoreListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKvStore)(nil).List), offset, limit)
	return &MockKvStoreListCall{Call: call}
}

// MockKvStoreListCall wrap *gomock.Call
type MockKvStoreListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKvStoreListCall) Return(arg0 []string, arg1 error) *MockKvStoreListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKvStoreListCall) Do(f func(int, int) ([]string, error)) *MockKvStoreListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKvStoreListCall) DoAndReturn(f func(int, int) ([]string, error)) *MockKvStoreListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Put mocks base method.
func (m *MockKvStore) Put(key string, value Stringer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockKvStoreMockRecorder) Put(key, value any) *MockKvStorePutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockKvStore)(nil).Put), key, value)
	return &MockKvStorePutCall{Call: call}
}

// MockKvStorePutCall wrap *gomock.Call
type MockKvStorePutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKvStorePutCall) Return(arg0 error) *MockKvStorePutCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKvStorePutCall) Do(f func(string, Stringer) error) *MockKvStorePutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKvStorePutCall) DoAndReturn(f func(string, Stringer) error) *MockKvStorePutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
