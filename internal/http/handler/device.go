package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/wptest/internal/models"
	"github.com/wptest/internal/services"
	"github.com/wptest/pkg/responses"
	"github.com/gorilla/mux"
)

type DeviceHandler struct {
	DeviceService  services.IDeviceService
	MessageService services.IMessageService
}

type IDeviceHandler interface {
	InsertDevice(w http.ResponseWriter, r *http.Request)
	GetDevice(w http.ResponseWriter, r *http.Request)
	GetDevices(w http.ResponseWriter, r *http.Request)
}

func NewDeviceHandler(deviceService services.IDeviceService) *DeviceHandler {
	return &DeviceHandler{
		DeviceService: deviceService,
	}
}

func (d *DeviceHandler) InsertDevice(w http.ResponseWriter, r *http.Request) {

	var body models.Device

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err)
		responses.Write(w, responses.APIErrorValidation.WithMessage(err))
	}

	isValid, errValidation := govalidator.ValidateStruct(body)
	if !isValid {
		responses.Write(w, responses.APIErrorValidation.WithMessage(errValidation))
		return
	}

	err := d.DeviceService.Publish(body)
	if err != nil {
		responses.Write(w, responses.APIErrorUnknown)
		return
	}
	responses.Write(w, responses.APICreated)
}

func (d *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.ParseUint(param["id"], 10, 64)
	device, err := d.DeviceService.GetById(id)
	if err != nil {
		responses.Write(w, responses.APIErrorUnknown)
	}
	responses.Write(w, responses.APIOK.WithData(device))
}

func (d *DeviceHandler) GetDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := d.DeviceService.GetAll()
	if err != nil {
		responses.Write(w, responses.APIErrorUnknown)
	}
	responses.Write(w, responses.APIOK.WithData(devices))
}
