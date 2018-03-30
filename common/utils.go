package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"TravelShipper/model"
	"TravelShipper/constants"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server, MongoDBHost, DBUser, DBPwd, Database string
		LogLevel                                     int
	}
)

// DisplayAppError provides app specific error in JSON
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HTTPStatus: code,
	}
	//log.Printf("AppError]: %s\n", handlerError)
	Error.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func ResponseError(w http.ResponseWriter, handlerError error, message string, code int) {
	data := model.ResponseModel{
		StatusCode: code,
		Error:      handlerError.Error(),
		Data:       message,
	}

	j, err := json.Marshal(data)
	if err != nil {
		DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			constants.InternalError.V(),
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func ResponseErrorString(w http.ResponseWriter, errString string , message string, code int) {
	data := model.ResponseModel{
		StatusCode: code,
		Error:      errString,
		Data:       message,
	}

	j, err := json.Marshal(data)
	if err != nil {
		DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			constants.InternalError.V(),
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
