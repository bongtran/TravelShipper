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
	"TravelShipper/constants"
	"time"
	"github.com/gorilla/mux"
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
			constants.InternalError.V(),
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
		Email:        dataResource.Email,
		ActivateCode: code,
		CreatedDate:  time.Now().UTC(),
		ModifiedDate: time.Now().UTC(),
		Role:         "member",
	}

	// Insert User document
	statusCode, err := userStore.Create(user, dataResource.Password)

	response := model.ResponseModel{
		StatusCode: statusCode.V(),
	}

	switch statusCode {
	case constants.Successful:
		emails.SendVerifyEmail(dataResource.Email, code)
		response.Data = ""
		break
	case constants.ExitedEmail:
		response.Error = statusCode.T()
		//if err != nil {
		//	response.Error = err.Error()
		//}
		break
	case constants.Error:
		response.Error = statusCode.T()
		//if err != nil {
		//	response.Error = err.Error()
		//}
		break
	}

	data, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func Activate(w http.ResponseWriter, r *http.Request) {
	log.Println("Activate")
	var dataResource model.ActivateResource
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.ResponseError(
			w,
			err,
			"Invalid activate data",
			constants.InternalError.V(),
		)
		return
	}
	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	// Authenticate the login user
	user, err, status := userStore.Activate(dataResource)
	data := model.ResponseModel{
		StatusCode: status.V(),
	}
	switch status {
	case constants.ActivateFail:
		data.Error = status.T()
		//if err != nil {
		//	data.Error = err.Error()
		//}
		break
	case constants.Error:
		data.Error = status.T()
		//if err != nil {
		//	data.Error = err.Error()
		//}
		break
	case constants.Successful:
		token, err = common.GenerateJWT(user.ID, user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Eror while generating the access token",
				constants.InternalError.V(),
			)
			return
		}
		// Clean-up the hashpassword to eliminate it from response JSON
		user.HashPassword = nil
		authUser := model.AuthUserModel{
			User:  user,
			Token: token,
		}
		data.Data = authUser
		break
	}
	// Generate JWT token

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
			constants.InternalError.V(),
		)
		return
	}
	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	// Authenticate the login result
	result, err, status := userStore.Login(dataResource.Email, dataResource.Password)

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.NotActivated:
		data.Error = status.T()
		//if err != nil {
		//	data.Data = constants.NotActivated.T()
		//	data.Error = err.Error()
		//}
		break
	case constants.LoginFail:
		data.Error = status.T()
		//if err != nil {
		//	data.Error = err.Error()
		//}
		break
	case constants.Successful:
		// Generate JWT token
		token, err = common.GenerateJWT(result.ID, result.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Eror while generating the access token",
				constants.InternalError.V(),
			)
			return
		}
		// Clean-up the hashpassword to eliminate it from response JSON
		result.HashPassword = nil
		result.ActivateCode = ""
		authUser := model.AuthUserModel{
			User:  result,
			Token: token,
		}
		data.Data = authUser
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

func GetMyProfile(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value("user")
	if session == nil {
		common.ResponseErrorString(
			w,
			constants.InternalError.T(),
			"Invalid Token Data",
			constants.InternalError.V(),
		)
		return
	}
	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}

	//log.Println(session.(model.User).ID.String())
	// Authenticate the login result
	result, err, status := userStore.GetUser(session.(model.User).ID.Hex())

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.Fail:
		data.Error = status.T()
		//if err != nil {
		//	data.Error = err.Error()
		//}
		break
	case constants.Successful:
		result.HashPassword = nil
		result.ActivateCode = ""
		data.Data = result
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

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Profile")
	var dataResource model.User
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.ResponseError(
			w,
			err,
			"Invalid Login data",
			constants.InternalError.V(),
		)
		return
	}
	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	// Authenticate the login user
	err, status := userStore.UpdateUser(dataResource)

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.Fail:
		data.Error = status.T()
		//if err != nil {
		//	data.Error = err.Error()
		//}
		break
	case constants.Successful:
		data.Data = constants.Successful.T()
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

func ResendActivateCode(w http.ResponseWriter, r *http.Request) {
	log.Println("Resend")
	var dataResource model.RegisterResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			constants.InternalError.V(),
		)
		return
	}

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}

	// Insert User document
	activateCode, err, statusCode := userStore.GetActivateCode(dataResource.Email)

	response := model.ResponseModel{
		StatusCode: statusCode.V(),
	}

	switch statusCode {
	case constants.Successful:
		emails.SendVerifyEmail(dataResource.Email, activateCode)
		response.Data = "Successful"
		break
	case constants.NotExitedEmail:
		response.Error = statusCode.T()
		//if err != nil {
		//	response.Error = err.Error()
		//}
		break
	}

	data, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("RequestResetPassword")
	var dataResource model.RegisterResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			constants.InternalError.V(),
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

	// Insert User document
	err, statusCode := userStore.RequestResetPassord(dataResource.Email, code)

	response := model.ResponseModel{
		StatusCode: statusCode.V(),
	}

	switch statusCode {
	case constants.Successful:
		emails.SendVerifyEmail(dataResource.Email, code)
		response.Data = "Successful"
		break
	case constants.NotExitedEmail:
		response.Error = statusCode.T()
		//if err != nil {
		//	response.Error = err.Error()
		//}
		break
	}

	data, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("Reset")
	var dataResource model.ResetPasswordResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			constants.InternalError.V(),
		)
		return
	}

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}

	// Insert User document
	user, err, statusCode := userStore.ResetPassword(dataResource.Email, dataResource.Password, dataResource.ActivateCode)

	response := model.ResponseModel{
		StatusCode: statusCode.V(),
	}

	switch statusCode {
	case constants.Successful:
		token, err := common.GenerateJWT(user.ID, user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Eror while generating the access token",
				constants.InternalError.V(),
			)
			return
		}
		// Clean-up the hashpassword to eliminate it from response JSON
		user.HashPassword = nil
		user.ActivateCode = ""
		authUser := model.AuthUserModel{
			User:  user,
			Token: token,
		}
		response.Data = authUser
		break
	case constants.NotExitedEmail:
		response.Error = statusCode.T()
		//if err != nil {
		//	response.Error = err.Error()
		//}
		break
	case constants.ResetPasswordFail:
		response.Error = statusCode.T()
		//if err != nil {
		//	response.Error = err.Error()
		//}
		break
	}

	data, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dataStore := common.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	// Authenticate the login result
	result, err, status := userStore.GetUser(id)

	data := model.ResponseModel{
		StatusCode: status.V(),
	}

	switch status {
	case constants.Fail:
		data.Error = status.T()
		if err != nil {
			data.Error = err.Error()
		}
		break
	case constants.Successful:
		result.HashPassword = nil
		result.ActivateCode = ""
		data.Data = result
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