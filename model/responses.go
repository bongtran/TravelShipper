package model

type (
	Response struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}

	ResponseModel struct {
		StatusCode int         `json:"code"`
		Data       interface{} `json:"d"`
		Error      string      `json:"error"`
	}
)

