package responses

import (
	"cart/modal"
	"net/http"
)

// Success

// GetOrderSuccess return response success order
var GetOrderSuccess = modal.Response{Status: http.StatusOK, Message: "Order Success", Code: "SUCCCART001"}

// CreateOrderSuccess returns response creating order
var CreateOrderSuccess = modal.Response{Status: http.StatusOK, Message: "Order Created Successfully", Code: "SUCCCART001"}

// Errors

// OrderNotFound return response on no or empty order
var OrderNotFound = modal.Response{Status: http.StatusNotFound, Message: "Order not Found", Code: "ERRCART001"}

// OrderIsBlocked return response on blocked order
var OrderIsBlocked = modal.Response{Status: http.StatusForbidden, Message: "User is unauthorized to access", Code: "ERRUSR010"}
