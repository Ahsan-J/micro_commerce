package cart

// SendHelloWorld an example of function that sends data
func SendHelloWorld(data interface{}) {
	publish("hello", data)
}
