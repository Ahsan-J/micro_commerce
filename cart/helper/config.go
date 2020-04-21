package helper

import (
	"os"
	"strconv"
)

type mappingInfo struct {
	Cart     string `json:"cart"`
	User     string `json:"user"`
	Common   string `json:"common"`
	Product  string `json:"product"`
	Logistic string `json:"logistic"`
}

// Config is the base configuration
type Config struct {
	PostgresURL string      `json:"postgresURL"`
	Port        int         `json:"port"`
	RabbitURL   string      `json:"rabbitMQ_URL"`
	Exchange    mappingInfo `json:"exchange"`
	Queue       mappingInfo `json:"queue"`
}

func getEnv(field string, defaultValue string) string {
	fieldValue := os.Getenv(field)
	if fieldValue == "" || fieldValue == "0" {
		return defaultValue
	}
	return fieldValue
}

var defaultConfig Config

func parseInt(value string) int {
	v, _ := strconv.ParseInt(value, 10, 64)
	return int(v)
}

// GetConfiguration exports the configuration of the application
func GetConfiguration() Config {
	if (Config{}) == defaultConfig {
		ImportJSON("config.json", &defaultConfig) // import JSON into defaultConfig only if it is undefined
	}

	return Config{
		PostgresURL: getEnv("POSTGRES_URL", defaultConfig.PostgresURL),
		Port:        parseInt(getEnv("PORT", strconv.Itoa(defaultConfig.Port))),
		RabbitURL:   getEnv("RABBITMQ_URL", defaultConfig.RabbitURL),
		Exchange: mappingInfo{
			Cart:     getEnv("EXCHANGE_CART", defaultConfig.Exchange.Cart),
			Common:   getEnv("EXCHANGE_COMMON", defaultConfig.Exchange.Common),
			User:     getEnv("EXCHANGE_USER", defaultConfig.Exchange.User),
			Logistic: getEnv("EXCHANGE_LOGISTICS", defaultConfig.Exchange.Logistic),
			Product:  getEnv("EXCHANGE_PRODUCT", defaultConfig.Exchange.Product),
		},
		Queue: mappingInfo{
			Cart:     getEnv("QUEUE_CART", defaultConfig.Queue.Cart),
			Common:   getEnv("QUEUE_COMMON", defaultConfig.Queue.Common),
			User:     getEnv("QUEUE_USER", defaultConfig.Queue.User),
			Logistic: getEnv("QUEUE_LOGISTICS", defaultConfig.Queue.Logistic),
			Product:  getEnv("QUEUE_PRODUCT", defaultConfig.Queue.Product),
		},
	}
}
