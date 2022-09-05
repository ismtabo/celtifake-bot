// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ismtabo/phrases-of-the-year/pkg/controller (interfaces: TelegramBotApiController)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	telebot "gopkg.in/telebot.v3"
)

// MockTelegramBotApiController is a mock of TelegramBotApiController interface.
type MockTelegramBotApiController struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramBotApiControllerMockRecorder
}

// MockTelegramBotApiControllerMockRecorder is the mock recorder for MockTelegramBotApiController.
type MockTelegramBotApiControllerMockRecorder struct {
	mock *MockTelegramBotApiController
}

// NewMockTelegramBotApiController creates a new mock instance.
func NewMockTelegramBotApiController(ctrl *gomock.Controller) *MockTelegramBotApiController {
	mock := &MockTelegramBotApiController{ctrl: ctrl}
	mock.recorder = &MockTelegramBotApiControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramBotApiController) EXPECT() *MockTelegramBotApiControllerMockRecorder {
	return m.recorder
}

// Help mocks base method.
func (m *MockTelegramBotApiController) Help(arg0 context.Context, arg1 telebot.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Help", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Help indicates an expected call of Help.
func (mr *MockTelegramBotApiControllerMockRecorder) Help(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Help", reflect.TypeOf((*MockTelegramBotApiController)(nil).Help), arg0, arg1)
}

// New mocks base method.
func (m *MockTelegramBotApiController) New(arg0 context.Context, arg1 telebot.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockTelegramBotApiControllerMockRecorder) New(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockTelegramBotApiController)(nil).New), arg0, arg1)
}

// Search mocks base method.
func (m *MockTelegramBotApiController) Search(arg0 context.Context, arg1 telebot.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Search indicates an expected call of Search.
func (mr *MockTelegramBotApiControllerMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockTelegramBotApiController)(nil).Search), arg0, arg1)
}

// Start mocks base method.
func (m *MockTelegramBotApiController) Start(arg0 context.Context, arg1 telebot.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockTelegramBotApiControllerMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTelegramBotApiController)(nil).Start), arg0, arg1)
}
