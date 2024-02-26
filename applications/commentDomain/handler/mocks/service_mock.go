// Code generated by MockGen. DO NOT EDIT.
// Source: typedef.go

// Package HandlerServiceMocks is a generated GoMock package.
package HandlerServiceMocks

import (
	context "context"
	model "github.com/TremblingV5/DouTok/applications/commentDomain/dal/hbModel"
	reflect "reflect"

	entity "github.com/TremblingV5/DouTok/kitex_gen/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockCommentDomainService is a mock of CommentDomainService interface.
type MockCommentDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockCommentDomainServiceMockRecorder
}

// MockCommentDomainServiceMockRecorder is the mock recorder for MockCommentDomainService.
type MockCommentDomainServiceMockRecorder struct {
	mock *MockCommentDomainService
}

// NewMockCommentDomainService creates a new mock instance.
func NewMockCommentDomainService(ctrl *gomock.Controller) *MockCommentDomainService {
	mock := &MockCommentDomainService{ctrl: ctrl}
	mock.recorder = &MockCommentDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentDomainService) EXPECT() *MockCommentDomainServiceMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockCommentDomainService) AddComment(ctx context.Context, videoId, userId, conversationId, lastId, toUserId int64, content string) (*entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", ctx, videoId, userId, conversationId, lastId, toUserId, content)
	ret0, _ := ret[0].(*entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockCommentDomainServiceMockRecorder) AddComment(ctx, videoId, userId, conversationId, lastId, toUserId, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockCommentDomainService)(nil).AddComment), ctx, videoId, userId, conversationId, lastId, toUserId, content)
}

// CountComments mocks base method.
func (m *MockCommentDomainService) CountComments(ctx context.Context, videoId ...int64) (map[int64]int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range videoId {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountComments", varargs...)
	ret0, _ := ret[0].(map[int64]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountComments indicates an expected call of CountComments.
func (mr *MockCommentDomainServiceMockRecorder) CountComments(ctx interface{}, videoId ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, videoId...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountComments", reflect.TypeOf((*MockCommentDomainService)(nil).CountComments), varargs...)
}

// ListComment mocks base method.
func (m *MockCommentDomainService) ListComment(ctx context.Context, videoId int64) ([]*model.CommentInHB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComment", ctx, videoId)
	ret0, _ := ret[0].([]*model.CommentInHB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComment indicates an expected call of ListComment.
func (mr *MockCommentDomainServiceMockRecorder) ListComment(ctx, videoId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComment", reflect.TypeOf((*MockCommentDomainService)(nil).ListComment), ctx, videoId)
}

// RemoveComment mocks base method.
func (m *MockCommentDomainService) RemoveComment(ctx context.Context, userId, commentId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveComment", ctx, userId, commentId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveComment indicates an expected call of RemoveComment.
func (mr *MockCommentDomainServiceMockRecorder) RemoveComment(ctx, userId, commentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveComment", reflect.TypeOf((*MockCommentDomainService)(nil).RemoveComment), ctx, userId, commentId)
}