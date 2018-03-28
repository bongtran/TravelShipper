package controllers

import (
	"encoding/json"
	"net/http"
	"log"
	"TravelShipper/common"
	"TravelShipper/store"
	"TravelShipper/model"
	"TravelShipper/emails"
	"TravelShipper/utils"
)

// Register add a new User document
// Handler for HTTP Post - "/users/register"
func Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register")
	var dataResource model.RegisterResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}

	log.Println("email: " + dataResource.Email)
	code := utils.RandStringBytesMaskImprSrc(6)
	log.Println(code)

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	user := model.User{
		Email: dataResource.Email,
	}

	// Insert User document
	statusCode, err := userStore.Create(user, dataResource.Password)

	response := model.ResponseModel{
		StatusCode: statusCode,
	}

	emails.SendVerifyEmail(dataResource.Email, code)

	switch statusCode {
	case 50000:
		response.Data = "Successful"
		break
	case 56000:
		response.Data = "Existed"
		response.Error = err.Error()
		break
	case 5700:
		response.Data = "Error"
		response.Error = err.Error()
		break
	}

	data, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/users/login"
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	var dataResource model.RegisterResource
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.ResponseError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	// Authenticate the login user
	user, err := userStore.Login(dataResource.Email, dataResource.Password)
	if err != nil {
		common.ResponseError(
			w,
			err,
			"Invalid login credentials",
			57001,
		)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(user.ID, user.Email, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	authUser := model.AuthUserModel{
		User:  user,
		Token: token,
	}

	data := model.ResponseModel{
		StatusCode: 50000,
		Data:       model.AuthUserResource{Data: authUser},
	}

	j, err := json.Marshal(data)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
