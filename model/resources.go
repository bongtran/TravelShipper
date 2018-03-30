package model

import (
	"gopkg.in/mgo.v2/bson"
)

//Models for JSON resources
type (
	// UserResource For Post - /user/register
	RegisterResource struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	ResetPasswordResource struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		ActivateCode string `json:"activate_code"`
	}

	UserResource struct {
		Data UserModel `json:"data"`
	}

	ActivateResource struct {
		ID bson.ObjectId `json:"id"`
		Email string `json:"email"`
		ActivateCode string `json:"activate_code"`
	}
	// AuthUserResource Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// BookmarkResource For Post/Put - /bookmarks
	// For Get - /bookmarks/id
	BookmarkResource struct {
		Data Bookmark `json:"data"`
	}
	// BookmarksResource for Get - /bookmarks
	BookmarksResource struct {
		Data []Bookmark `json:"data"`
	}

	TransferResource struct {
		Data TransferModel `json:"data"`
	}

	OfferResource struct {
		Data OfferModel `json:"data"`
	}

	// UserModel reperesents a user
	UserModel struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	// AuthUserModel for authorized user with access token
	AuthUserModel struct {
		User  User `json:"user"`
		Token string     `json:"token"`
	}


)
