package constants

type StatusCode int

const (
	InternalError  StatusCode = 500
	Successful     StatusCode = 50000
	Fail           StatusCode = 51000
	Error          StatusCode = 56000
	LoginFail      StatusCode = 51001
	ExitedEmail    StatusCode = 51002
	NotExitedEmail StatusCode = 51003
	ActivateFail   StatusCode = 51004
	NotActivated   StatusCode = 51005
)

var statusValue = map[StatusCode]int{
	InternalError:  500,
	Successful:     50000,
	Fail:           51000,
	Error:          56000,
	LoginFail:      51001,
	ExitedEmail:    51002,
	NotExitedEmail: 51003,
	ActivateFail:   51004,
	NotActivated:   51005,
}

var statusString = map[StatusCode]string{
	InternalError:  "InternalError",
	Successful:     "Successful",
	Fail:           "Fail",
	Error:          "Error",
	LoginFail:      "LoginFail",
	ExitedEmail:    "ExitedEmail",
	NotExitedEmail: "NotExitedEmail",
	ActivateFail:   "ActivateFail",
	NotActivated:   "NotActivated",
}

func (status StatusCode) V() int {
	return statusValue[status]
}

func (status StatusCode) T() string {
	return statusString[status]
}
