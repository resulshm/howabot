package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/sirupsen/logrus"
)

func SendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(statusCode)
	resp := models.Response{}
	if statusCode == 200 {
		resp.Success = true
	}

	resp.Data = data

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.WithError(err).Error("Error in encoding response")
	}
}
