package constants

type StatusCode int

const (
	InternalError StatusCode = 500
	Successful StatusCode = 500000
	Fail StatusCode = 50001
	Error StatusCode = 56000
	Exited StatusCode = 57000
	LoginFail StatusCode = 57001
)


