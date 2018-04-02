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
		Description  string        `json:"description"`
		Avatar       string        `json:"avatar"`
		MyUrl        string        `json:"my_url"`
		IDCardUrl    string        `json:"id_card_url"`
		CreatedDate  time.Time     `json:"created_date"`
		ModifiedDate time.Time     `json:"modified_date"`
		Activated    bool          `json:"activated"`
		ActivateCode string        `json:"activate_code"`
		Role         string        `json:"role"`
		//Locations    []Location    `json:"locations"`
	}

	Device struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserID      bson.ObjectId `json:"user_id"`
		DeviceOS    string        `json:"device_os"`
		DeviceModel string        `json:"device_model"`
		DeviceToken string        `json:"device_token"`
	}

	Location struct {
		ID              bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserID          bson.ObjectId `json:"user_id"`
		Country         string        `json:"country"`
		Province        string        `json:"province"`
		Latitude        float32       `json:"latitude"`
		Longitude       float32       `json:"longitude"`
		Hometown        bool          `json:"hometown"`
		BeginTime       time.Time     `json:"begin_time"`
		EndTime         time.Time     `json:"end_time"`
		CurrentLocation bool          `json:"current_location"`
	}

	Item struct {
		ID           bson.ObjectId  `bson:"_id,omitempty" json:"id"`
		UserID       bson.ObjectId  `json:"user_id"`
		Name         string         `json:"name"`
		Description  string         `json:"description"`
		Price        float32        `json:"price"`
		Quantity     int            `json:"quantity"`
		ItemUrl      string         `json:"item_url"`
		Country      string         `json:"country"`
		Province     string         `json:"province"`
		Latitude     float32        `json:"latitude"`
		Longitude    float32        `json:"longitude"`
		CreatedDate  time.Time      `json:"created_date"`
		ModifiedDate time.Time      `json:"modified_date"`
		ItemStatus   int            `json:"item_status"`
		Offers       []ShipperOffer `json:"offers"`
		Deal         Deal           `json:"deal"`
	}

	ShipperOffer struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ItemID      bson.ObjectId `json:"item_id"`
		BuyerID     bson.ObjectId `json:"buyer_id"`
		ShipperID   bson.ObjectId `json:"shipper_id"`
		Fee         float32       `json:"fee"`
		Description string        `json:"description"`
		Status      int           `json:"status"`
		CreatedDate time.Time     `json:"created_date"`
	}

	Deal struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ItemID      bson.ObjectId `json:"item_id"`
		BuyerID     bson.ObjectId `json:"buyer_id"`
		ShipperID   bson.ObjectId `json:"shipper_id"`
		Price       string        `json:"price"`
		Fee         float32       `json:"fee"`
		Description string        `json:"description"`
		Status      int           `json:"status"`
		CreatedDate time.Time     `json:"created_date"`
		FinishDate  time.Time     `json:"finish_date"`
	}

	Conversation struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ItemID    bson.ObjectId `json:"item_id"`
		BuyerID   bson.ObjectId `json:"buyer_id"`
		ShipperID bson.ObjectId `json:"shipper_id"`
		Chats     []Chat        `json:"chats"`
	}

	Chat struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ChatterID   bson.ObjectId `json:"item_id"`
		Content     string        `json:"content"`
		Status      int           `json:"status"`
		CreatedDate time.Time     `json:"created_date"`
	}

	Rate struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ItemID    bson.ObjectId `json:"item_id"`
		BuyerID   bson.ObjectId `json:"buyer_id"`
		ShipperID bson.ObjectId `json:"shipper_id"`
		Content   string        `json:"content"`
		Value     int           `json:"value"`
	}

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
