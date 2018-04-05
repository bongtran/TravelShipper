package controllers

import (
	"encoding/json"
	"TravelShipper/constants"
	"net/http"
	"TravelShipper/model"
	"TravelShipper/common"
	"TravelShipper/store"
	"time"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var dataResource model.Item
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

	dataResource.UserID = session.(model.UserLite).ID
	dataResource.CreatedDate = time.Now().UTC()
	dataResource.ModifiedDate = dataResource.CreatedDate
	dataResource.ItemStatus = 1

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("items")
	itemStore := store.ItemStore{C: col}

	err, status := itemStore.CreateItem(dataResource)

	//TO DO notification
	//dataStore = common.NewDataStore()
	//defer dataStore.Close()
	//col = dataStore.Collection("users")
	//userStore := store.UserStore{C: col}
	//
	//userStore.UpdateLocation(dataResource, session.(model.UserSession).ID.Hex())

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

func MyItems(w http.ResponseWriter, r *http.Request) {
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
	col := dataStore.Collection("items")
	itemStore := store.ItemStore{C: col}

	items := itemStore.MyItem(session.(model.UserLite).ID)

	data := model.ResponseModel{
		StatusCode: constants.Successful.V(),
		Data:       items,
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

func ItemDetail(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("item")
	itemStore := store.ItemStore{C: col}

	item, err, status := itemStore.GetItemDetail(id)

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.GetItemDetailFail:
		data.Error = status.T()
		break
	case constants.Successful:
		if item.UserID == session.(model.UserLite).ID {
			data.Data = item
		} else {
			itemLite := model.ItemLite{
				ID:          item.ID,
				UserID:      item.UserID,
				CreatedDate: item.CreatedDate,
				ItemStatus:  item.ItemStatus,
				Name:        item.Name,
				Price:       item.Price,
				Quantity:    item.Quantity,
				Description: item.Description,
				Longitude:   item.Longitude,
				Latitude:    item.Latitude,
				Province:    item.Province,
				Country:     item.Country,
				ItemUrl:     item.ItemUrl,
				Address:     item.Address,
				CountryCode: item.CountryCode,
			}
			data.Data = itemLite
		}

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

func MySuggestItems(w http.ResponseWriter, r *http.Request) {
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
	col := dataStore.Collection("items")
	itemStore := store.ItemStore{C: col}

	items := itemStore.SuggestItem(session.(model.UserLite).ID.Hex())

	data := model.ResponseModel{
		StatusCode: constants.Successful.V(),
		Data:       items,
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
