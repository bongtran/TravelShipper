package store

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"TravelShipper/model"
	"TravelShipper/controllers"
	"go/constant"
	"TravelShipper/constants"
)

// UserStore provides persistence logic for "users" collection.
type UserStore struct {
	C *mgo.Collection
}

// Create insert new User
func (store UserStore) Create(user model.User, password string) (int, error ){
	var fund model.User
	err := store.C.Find(bson.M{"email": user.Email}).One(&fund)
	if err == nil {
		return 57000, err
	}

	user.ID = bson.NewObjectId()
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 56000, err
	}
	user.HashPassword = hpass
	err = store.C.Insert(user)
	return 50000, err
}

func (store UserStore) Activate(model controllers.ActivateResource) (model.User, error) {
	var user = model.User{};
	err := store.C.Update(bson.M{"_id": model.ID},
	bson.M{"$set": bson.M{"activated": true}})
	if err != nil{
		return user, err
	}

	err = store.C.Find(bson.M{"_id": model.ID}).One(&user)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Login authenticates the User
func (store UserStore) Login(email, password string) (model.User, error) {
	var user model.User
	err := store.C.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return model.User{}, err
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (store UserStore) GetUser(email, password string) (model.User, error) {
	var user model.User
	err := store.C.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
