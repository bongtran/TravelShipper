package controllers

import (
	"net/http"
	"TravelShipper/stellar"
	"github.com/gorilla/mux"
	"TravelShipper/common"
	"encoding/json"
	"TravelShipper/store"
	"TravelShipper/model"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	status, result := stellar.GetBalance()

	w.Write(result)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

func GetBalanceFromID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	status, result := stellar.GetBalanceFromID(id)

	w.Write(result)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	result, pair := stellar.InitAccount()
	//pairByte, err := json.Marshal(pair)

	//userKeyPair :=
	dataStore := common.NewDataStore()
	defer dataStore.Close()

	col := dataStore.Collection("userkeypair")
	keyStore := store.KeyStore{col}

	//userID := bson.ObjectId()
	userPair := model.UserKeyPair{Address: pair.Address(),
		Seed: pair.Seed(),}
	user := r.Context().Value("user")
	if user != nil {
		userPair.UserID = user.(model.User).ID
	}
	err := keyStore.Create(userPair)

	//account, err := horizon.DefaultTestNetClient.LoadAccount(userPair.Seed)
	data, err := json.Marshal(userPair)

	if !result || err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func TrustAsset(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}

	var result string

	info := transferResource.Data

	result, err = stellar.TrustAnAsset(info.To, info.Issuer, info.AssetCode, info.Amount)

	data, _ := json.Marshal(result)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func TransferToThirdPerson(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}

	var result string

	info := transferResource.Data

	result, err = stellar.TrustAnAsset(info.To, info.Issuer, info.AssetCode, "1000000000")
	if err != nil {
		result, err = stellar.SendAnAssetPayment(info.From, info.To, info.Issuer, info.AssetCode, info.Amount)
	}

	data, _ := json.Marshal(result)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}

	var result string

	info := transferResource.Data
	if info.AssetCode == "" {
		result, err = stellar.SendAPayment(info.From, info.To, info.Amount)
	} else {
		result, err = stellar.SendAnAssetPayment(info.From, info.To, info.Issuer, info.AssetCode, info.Amount)
	}

	data, _ := json.Marshal(result)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func CreateICO(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	info := transferResource.Data
	result, err := stellar.ICO(info.From, info.To, info.AssetCode, info.Amount)

	var data []byte

	if err == nil {
		data, err = json.Marshal(result)
	}
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func CheckTrust(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	info := transferResource.Data
	result, err := stellar.CheckTrust(info.To, info.Issuer, info.AssetCode)

	var data []byte

	if err == nil {
		data, err = json.Marshal(result)
	}
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func AddInformation(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	info := transferResource.Data
	result, err := stellar.AddInformation(info.Issuer, info.Amount)

	var data []byte

	if err == nil {
		data, err = json.Marshal(result)
	}
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func LockAccount(w http.ResponseWriter, r *http.Request) {
	var transferResource TransferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	info := transferResource.Data
	result, err := stellar.LockAccount(info.Issuer)

	var data []byte

	if err == nil {
		data, err = json.Marshal(result)
	}
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	var transferResource OfferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	info := transferResource.Data
	result, err := stellar.CreateOffer(info)

	var data []byte

	if err == nil {
		data, err = json.Marshal(result)
	}
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}

func CreatePassiveOffer(w http.ResponseWriter, r *http.Request) {
	var transferResource OfferResource

	err := json.NewDecoder(r.Body).Decode(&transferResource)
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	info := transferResource.Data
	result, err := stellar.CreateAssetOffer(info)

	var data []byte

	if err == nil {
		data, err = json.Marshal(result)
	}
	if err != nil {
		common.DisplayAppError(w, err, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}