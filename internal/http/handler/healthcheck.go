package handler

import (
	"net/http"

	"github.com/wptest/pkg/responses"
)

// GetHealthCheck this function for get health check
func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := responses.APIOK
	responses.Write(w, resp)
}
