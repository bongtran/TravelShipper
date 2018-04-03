package store

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"TravelShipper/model"
	"TravelShipper/constants"
	"time"
	"log"
)

// UserStore provides persistence logic for "users" collection.
type UserStore struct {
	C *mgo.Collection
}

// Create insert new User
func (store UserStore) Create(user model.User, password string) (constants.StatusCode, error) {
	var fund model.User
	err := store.C.Find(bson.M{"email": user.Email}).One(&fund)
	if err == nil {
		return constants.ExitedEmail, err
	}

	user.ID = bson.NewObjectId()
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return constants.Error, err
	}
	user.HashPassword = hpass
	err = store.C.Insert(user)
	return constants.Successful, err
}

func (store UserStore) Activate(model model.ActivateResource) (user model.User, err error, code constants.StatusCode) {
	//var user = model.User{};
	log.Println(model.ActivateCode)
	err = store.C.Update(bson.M{"email": model.Email, "activatecode": model.ActivateCode},
		bson.M{"$set": bson.M{"activated": true, "modifieddate": time.Now().UTC()}})
	if err != nil {
		return user, err, constants.ActivateFail
	}

	err = store.C.Find(bson.M{"email": model.Email}).One(&user)

	if err != nil {
		return user, err, constants.Error
	}

	return user, nil, constants.Successful
}

// Login authenticates the User
func (store UserStore) Login(email, password string) (model.User, error, constants.StatusCode) {
	var user model.User
	err := store.C.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return model.User{}, err, constants.NotExitedEmail
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return model.User{}, err, constants.LoginFail
	}

	if !user.Activated {
		return model.User{}, err, constants.NotActivated
	}

	return user, nil, constants.Successful
}

func (store UserStore) UpdateUser(user model.User) (error, constants.StatusCode) {
	err := store.C.Update(bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{
			"firstname":   user.FirstName,
			"lastname":    user.LastName,
			"description": user.Description,
			"myurl":       user.MyUrl,
			"phone_number": user.PhoneNumber,
			//"idcardurl": user.FirstName,
			//"firstname": user.FirstName,
			"modifieddate": time.Now().UTC()}})
	if err != nil {
		return err, constants.UpdateProfileFail
	}

	return nil, constants.Successful
}

func (store UserStore) GetUser(id string) (model.User, error, constants.StatusCode) {
	var user model.User
	err := store.C.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return model.User{}, err, constants.Fail
	}

	return user, nil, constants.Successful
}

func (store UserStore) GetActivateCode(email string) (string, error, constants.StatusCode) {
	var user model.User
	err := store.C.Find(bson.M{"email": email, "activated": false}).One(&user)
	if err != nil {
		return "", err, constants.NotExitedEmail
	}

	return user.ActivateCode, nil, constants.Successful
}

func (store UserStore) RequestResetPassord(email string, code string) (error, constants.StatusCode) {
	err := store.C.Update(bson.M{"email": email},
		bson.M{"$set": bson.M{"activatecode": code, "modifieddate": time.Now().UTC()}})
	if err != nil {
		return err, constants.NotExitedEmail
	}

	return nil, constants.Successful
}

func (store UserStore) ResetPassword(email string, password string, code string) (model.User, error, constants.StatusCode) {
	var user model.User
	err := store.C.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return model.User{}, err, constants.NotExitedEmail
	}

	err = store.C.Update(bson.M{"email": email, "activatecode": code},
		bson.M{"$set": bson.M{"hashpassword": password, "modifieddate": time.Now().UTC()}})
	if err != nil {
		return model.User{}, err, constants.ResetPasswordFail
	}
	return user, nil, constants.Successful
}
