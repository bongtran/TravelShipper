package store

import (
	"gopkg.in/mgo.v2"
	"TravelShipper/model"
	"TravelShipper/constants"
	"gopkg.in/mgo.v2/bson"
)

type ItemStore struct {
	C *mgo.Collection
}

func (store ItemStore) CreateItem(item model.Item) (error, constants.StatusCode) {
	err := store.C.Insert(item)
	if err != nil {
		return err, constants.InsertItemFail
	}

	return nil, constants.Successful
}

func (store ItemStore) EditItem(item model.Item) () {

}

func (store ItemStore) DeleteItem(id string) (error, constants.StatusCode) {
	err := store.C.Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{"item_status": 0}})
	if err != nil {
		return err, constants.DeleteItemFail
	}

	return nil, constants.Successful
}

func (store ItemStore) GetItemDetail(id string) (model.Item, error, constants.StatusCode) {
	var item model.Item
	err := store.C.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(item)
	if err != nil {
		return model.Item{}, err, constants.GetItemDetailFail
	}

	return item, nil, constants.Successful
}

func (store ItemStore) MyItem(userID string) []model.Item {
	var result []model.Item
	itor := store.C.Find(bson.M{"user_id": bson.ObjectIdHex(userID)}).Iter()

	item := model.Item{}
	for itor.Next(item) {
		result = append(result, item)
	}

	return result
}

func (store ItemStore) SuggestItem(userID string) {

}
