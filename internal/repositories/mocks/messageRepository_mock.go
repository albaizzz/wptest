package mocks

import (
	"reflect"

	"github.com/wptest/internal/models"
	gomock "github.com/golang/mock/gomock"
)

type MessageRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *MessageRepositoryMockRecorder
}

type MessageRepositoryMockRecorder struct {
	mock *MessageRepositoryMock
}

func NewMessageRepositoryMock(ctrl *gomock.Controller) *MessageRepositoryMock {
	mock := &MessageRepositoryMock{ctrl: ctrl}
	mock.recorder = &MessageRepositoryMockRecorder{mock}
	return mock
}

func (m *MessageRepositoryMock) EXPECT() *MessageRepositoryMockRecorder {
	return m.recorder
}

func (m *MessageRepositoryMock) Publish(request models.Device) error {
	ret := m.ctrl.Call(m, "Publish", request)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MessageRepositoryMockRecorder) Publish(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MessageRepositoryMock)(nil).Publish), arg0)
}
