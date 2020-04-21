package modal

import (
	"cart/helper"
	"encoding/json"
	"net/http"
)

// Response Struct to send
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
}

// SendAPI sends resposne with given data
func (res *Response) SendAPI(w http.ResponseWriter, data interface{}) {
	// Writing data to the response method
	res.Data = data
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.Status)

	err := json.NewEncoder(w).Encode(res)
	helper.FailOnError(err, "Error in unmarshelling")
}
