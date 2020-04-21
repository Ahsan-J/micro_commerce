package responses

import (
	"cart/modal"
	"net/http"
)

// Success

// GetCartSuccess return response success cart
var GetCartSuccess = modal.Response{Status: http.StatusOK, Message: "Cart Success", Code: "SUCCCART001"}

// CreateCartSuccess returns response creating cart
var CreateCartSuccess = modal.Response{Status: http.StatusOK, Message: "Cart Created Successfully", Code: "SUCCCART001"}

// Errors

// CartNotFound return response on no or empty cart
var CartNotFound = modal.Response{Status: http.StatusNotFound, Message: "Cart not Found", Code: "ERRCART001"}

// CartIsBlocked return response on blocked cart
var CartIsBlocked = modal.Response{Status: http.StatusForbidden, Message: "User is unauthorized to access", Code: "ERRUSR010"}
