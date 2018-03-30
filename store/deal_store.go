package store

import (
	"gopkg.in/mgo.v2"
	"TravelShipper/model"
)

type DealStore struct {
	C *mgo.Collection
}

func (store DealStore) CreateDealResponse(item model.Item) () {

}

func (store DealStore) AcceptDeal()  {

}

func (store DealStore) FinishDeal()  {

}

func (store DealStore) GetDealResponses()  {

}

func (store DealStore) GetDealResponse()  {

}

func (store DealStore) GetDeals()  {

}

func (store DealStore) GetDeal()  {

}

func (store DealStore) GetCancelDeal()  {

}

func (store DealStore) CancelDeal()  {

}
