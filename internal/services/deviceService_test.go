package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/wptest/internal/models"
	repomock "github.com/wptest/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func Test_Store(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	deviceRepository := repomock.NewDeviceRepositoryMock(mockCtrl)
	messageRepository := repomock.NewMessageRepositoryMock(mockCtrl)
	svc := NewDeviceService(deviceRepository, messageRepository)

	t.Run("Success Inserted", func(t *testing.T) {
		device := models.Device{
			ID:        0,
			Device:    "SGD-12344",
			Value:     1234.266,
			UpdatedAt: time.Unix(1008910273, 0),
		}
		deviceRepository.EXPECT().Store(device).Return(nil).Times(1)
		err := svc.ReceiveDevice()
		assert.Equal(t, err, nil)
	})

	t.Run("Fail Inserted", func(t *testing.T) {
		err := svc.Store([]byte(xmlFail))
		fmt.Println(err)
		assert.Equal(t, err, fmt.Errorf("Unmatch data request"))
	})
}

func Test_GetById(t *testing.T) {

}

func Test_GetAll(t *testing.T) {

}
