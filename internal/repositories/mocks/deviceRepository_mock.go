package mocks

import (
	"reflect"

	"github.com/wptest/internal/models"
	gomock "github.com/golang/mock/gomock"
)

type DeviceRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *DeviceRepositoryMockRecorder
}

type DeviceRepositoryMockRecorder struct {
	mock *DeviceRepositoryMock
}

func NewDeviceRepositoryMock(ctrl *gomock.Controller) *DeviceRepositoryMock {
	mock := &DeviceRepositoryMock{ctrl: ctrl}
	mock.recorder = &DeviceRepositoryMockRecorder{mock}
	return mock
}

func (m *DeviceRepositoryMock) EXPECT() *DeviceRepositoryMockRecorder {
	return m.recorder
}

func (m *DeviceRepositoryMock) GetAll() ([]models.Device, error) {
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *DeviceRepositoryMockRecorder) GetAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*DeviceRepositoryMock)(nil).GetAll))
}

func (m *DeviceRepositoryMock) GetById(id uint64) (models.Device, error) {
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *DeviceRepositoryMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*DeviceRepositoryMock)(nil).GetById), arg0)
}

func (m *DeviceRepositoryMock) Store(device models.Device) (err error) {
	ret := m.ctrl.Call(m, "Store", device)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *DeviceRepositoryMockRecorder) Store(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*DeviceRepositoryMock)(nil).Store), arg0)
}
