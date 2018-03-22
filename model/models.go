package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User type represents the registered user.
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	// Bookmark type represents the metadata of a bookmark.
	Bookmark struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Location    string        `json:"location"`
		Priority    int           `json:"priority"` // Priority (1 -5)
		CreatedBy   string        `json:"createdby"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}

	UserKeyPair struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Seed      string        `json:"seed"`
		Address   string        `json:"address"`
		UserID    bson.ObjectId `json:"userid"`
		CreatedOn time.Time     `json:"createdon,omitempty"`
	}

	TransferModel struct {
		From      string `json:"from"`
		To        string `json:"to"`
		Issuer    string `json:"issuer"`
		Amount    string `json:"amount"`
		AssetCode string `json:"code"`
	}

	OfferModel struct {
		From          string `json:"from"`
		BuyingCode    string `json:"buying_code"`
		BuyingIssuer  string `json:"buying_issuer"`
		SellingCode   string `json:"selling_code"`
		SellingIssuer string `json:"selling_issuer"`
		Amount        string `json:"amount"`
		Price         string `json:"price"`
	}
)
