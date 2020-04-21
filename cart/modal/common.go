package modal

// MessageQueueObject is mapping the rabbiMQ object
type MessageQueueObject struct {
	RoutingKey string      `json:"routingKey"`
	Data       interface{} `json:"data"`
}
