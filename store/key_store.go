package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"TravelShipper/model"
	"time"
)

// UserStore provides persistence logic for "users" collection.
type KeyStore struct {
	C *mgo.Collection
}

// Create insert new User
func (store KeyStore) Create(user model.UserKeyPair) error {
	user.CreatedOn = time.Now()
	user.ID = bson.NewObjectId()
	err := store.C.Insert(user)
	return err
}
