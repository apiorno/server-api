package view

//Response is an http response
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

//PostRequest is an http post request
type PostRequest struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}
