package modal

// "encoding/json"
// "fmt"
// "net/http"

// Cart Struct to reflect Database structure
type Cart struct {
	ID        string `json:"code" db:"id"`
	UserID    string `json:"data" db:"user_id"`
	Status    int    `json:"status" db:"status"`
	CreatedAt string `json:"createdAt" db:"created_at"`
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
