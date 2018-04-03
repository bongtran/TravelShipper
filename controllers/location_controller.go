package controllers

import (
	"encoding/json"
	"TravelShipper/constants"
	"net/http"
	"log"
	"TravelShipper/model"
	"TravelShipper/common"
	"TravelShipper/store"
	"time"
)

func SetLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("Location")
	var dataResource model.Location
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.ResponseError(
			w,
			err,
			"Invalid Login data",
			http.StatusInternalServerError,
		)
		return
	}

	err = dataResource.Validate()
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid data",
			http.StatusBadRequest,
		)
		return
	}

	session := r.Context().Value("user")
	if session == nil {
		common.ResponseErrorString(
			w,
			constants.InternalError.T(),
			"Invalid Token Data",
			http.StatusInternalServerError,
		)
		return
	}

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("locations")
	locationStore := store.LocationStore{C: col}
	if dataResource.BeginTime.After(time.Now().UTC()){
		dataResource.CurrentLocation = true
	}
	err, status := locationStore.SetLocation(dataResource)


	dataStore = common.NewDataStore()
	defer dataStore.Close()
	col = dataStore.Collection("users")
	userStore := store.UserStore{C: col}

	userStore.UpdateLocation(dataResource, session.(model.User).ID.Hex())

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.SetLocationFail:
		data.Error = status.T()
		//if err != nil {
		//	data.Data = constants.SetLocationFail.T()
		//	data.Error = err.Error()
		//}
		break
	case constants.Successful:
		data.Data = ""
		break
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(data)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			constants.InternalError.V(),
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetMyLocation(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value("user")
	if session == nil {
		common.ResponseErrorString(
			w,
			constants.InternalError.T(),
			"Invalid Token Data",
			http.StatusInternalServerError,
		)
		return
	}

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("locations")
	locationStore := store.LocationStore{C: col}

	location, err, status := locationStore.GetLocation(session.(model.User).ID.Hex())

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.Fail:
		data.Error = status.T()
		//if err != nil {
		//	data.Data = constants.SetLocationFail.T()
		//	data.Error = err.Error()
		//}
		break
	case constants.Successful:
		data.Data = location
		break
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(data)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			constants.InternalError.V(),
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

