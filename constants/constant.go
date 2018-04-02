package constants

type StatusCode int

const (
	InternalError     StatusCode = 500
	Successful        StatusCode = 50000
	Fail              StatusCode = 51000
	Error             StatusCode = 56000
	LoginFail         StatusCode = 51001
	ExitedEmail       StatusCode = 51002
	NotExitedEmail    StatusCode = 51003
	ActivateFail      StatusCode = 51004
	NotActivated      StatusCode = 51005
	ResetPasswordFail StatusCode = 51006
	SetLocationFail   StatusCode = 51007
	InsertItemFail    StatusCode = 51008
	GetItemDetailFail StatusCode = 51009
	GetLocationFail   StatusCode = 51010
	EditItemFail      StatusCode = 51011
	DeleteItemFail    StatusCode = 51012
	UpdateProfileFail StatusCode = 51013
	GetProfileFail    StatusCode = 51014
)

var statusValue = map[StatusCode]int{
	InternalError:     500,
	Successful:        50000,
	Fail:              51000,
	Error:             56000,
	LoginFail:         51001,
	ExitedEmail:       51002,
	NotExitedEmail:    51003,
	ActivateFail:      51004,
	NotActivated:      51005,
	ResetPasswordFail: 51006,
	SetLocationFail:   51007,
	InsertItemFail:    51008,
	GetItemDetailFail: 51009,
	GetLocationFail:   51010,
	EditItemFail:      51011,
	DeleteItemFail:    51012,
	UpdateProfileFail: 51013,
	GetProfileFail:    51014,
}

var statusString = map[StatusCode]string{
	InternalError:     "InternalError",
	Successful:        "Successful",
	Fail:              "Fail",
	Error:             "Error",
	LoginFail:         "LoginFail",
	ExitedEmail:       "ExitedEmail",
	NotExitedEmail:    "NotExitedEmail",
	ActivateFail:      "ActivateFail",
	NotActivated:      "NotActivated",
	ResetPasswordFail: "ResetPasswordFail",
	SetLocationFail:   "SetLocationFail",
	InsertItemFail:    "InsertItemFail",
	GetItemDetailFail: "GetItemDetailFail",
	GetLocationFail:   "GetLocationFail",
	EditItemFail:      "EditItemFail",
	DeleteItemFail:    "DeleteItemFail",
	UpdateProfileFail: "UpdateProfileFail",
	GetProfileFail:    "GetProfileFail",
}

func (status StatusCode) V() int {
	return statusValue[status]
}

func (status StatusCode) T() string {
	return statusString[status]
}
