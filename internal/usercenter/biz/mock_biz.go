// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/costa92/k8s-krm-go/internal/usercenter/biz (interfaces: IBiz)

// Package biz is a generated GoMock package.
package biz

import (
	gomock "github.com/golang/mock/gomock"
)

// MockIBiz is a mock of IBiz interface.
type MockIBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIBizMockRecorder
}

// MockIBizMockRecorder is the mock recorder for MockIBiz.
type MockIBizMockRecorder struct {
	mock *MockIBiz
}

// NewMockIBiz creates a new mock instance.
func NewMockIBiz(ctrl *gomock.Controller) *MockIBiz {
	mock := &MockIBiz{ctrl: ctrl}
	mock.recorder = &MockIBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBiz) EXPECT() *MockIBizMockRecorder {
	return m.recorder
}
