package main

// Config is the base configuration
type Config struct {
	PostgresURL string // `json: "postgresURL"`
	Port        int    // `json: "port"`
	RabbitURL   string
}

// GetConfiguration exports the configuration of the application
func GetConfiguration() Config {
	return Config{
		PostgresURL: "",
		Port:        5000,
		RabbitURL:   "",
	}
}
