package store

import (
	"gopkg.in/mgo.v2"
	"TravelShipper/model"
	"TravelShipper/constants"
	"gopkg.in/mgo.v2/bson"
	"time"
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

func (store ItemStore) EditItem(item model.Item) (error, constants.StatusCode) {
	err := store.C.Update(bson.M{"_id": item.ID},
		bson.M{"$set": bson.M{
			"name":          item.Name,
			"description":   item.Price,
			"price":         item.Price,
			"quantity":      item.Quantity,
			"item_url":      item.ItemUrl,
			"country":       item.Country,
			"province":      item.Province,
			"latitude":      item.Latitude,
			"longitude":     item.Longitude,
			"modified_date": time.Now().UTC(),
		}})
	if err != nil {
		return err, constants.EditItemFail
	}

	return nil, constants.Successful
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
	//query := []bson.M{{
	//	"$lookup": bson.M{ // lookup the documents table here
	//		"from":         "user",
	//		"localField":   "_id",
	//		"foreignField": "folderID",
	//		"as":           "documents",
	//	}},
	//	{"$match": bson.M{
	//		"level": bson.M{"$lte": user.Level},
	//		"userIDs": user.ID,
	//	}}}
	//
	//pipe := store.C.Pipe(query)
	//err := pipe.All(&folders)
}
