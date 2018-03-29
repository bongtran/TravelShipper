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
		return model.User{}, err, constants.LoginFail
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return model.User{}, err, constants.LoginFail
	}
	return user, nil, constants.Successful
}

func (store UserStore) GetUser(email, password string) (model.User, error) {
	var user model.User
	err := store.C.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
