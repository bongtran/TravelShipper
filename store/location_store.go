package store

import (
	"gopkg.in/mgo.v2"
	"TravelShipper/model"
	"TravelShipper/constants"
	"gopkg.in/mgo.v2/bson"
)

type LocationStore struct {
	C *mgo.Collection
}

func (store LocationStore) SetLocation(data model.Location) (error, constants.StatusCode) {
	err := store.C.Update(bson.M{"user_id":data.UserID},
	bson.M{"$set": bson.M{"current_location": false}})
	data.CurrentLocation = true
	err = store.C.Insert(data)

	if err != nil{
		return err, constants.SetLocationFail
	}

	return nil, constants.Successful
}

func (store LocationStore) GetLocation(userID string) (model.Location, error, constants.StatusCode) {
	var location model.Location

	err := store.C.Find(bson.M{"user_id": bson.ObjectIdHex(userID)}).One(&location)
	if err != nil{
		return model.Location{}, err, constants.GetLocationFail
	}

	return location, nil, constants.Successful
}
