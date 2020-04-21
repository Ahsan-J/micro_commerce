package modal

// "encoding/json"
// "fmt"
// "net/http"

// Order Struct to handle Database structure
type Order struct {
	CartID    string `json:"cartId" db:"cart_id"`
	ProductID string `json:"productId" db:"product_id"`
	ItemAdded string `json:"itemAdded" db:"item_added"`
	Status    int    `json:"status" db:"status"`
}

// // SendAPI sends resposne with given data
// func (res *Response) SendAPI(w http.ResponseWriter, data interface{}) {
// 	// Writing data to the response method
// 	res.Data = data
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(res.Status)

// 	err := json.NewEncoder(w).Encode(res)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
