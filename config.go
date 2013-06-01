package yourpackage

import "os"

var (
	config = getConfig()
)

type Config struct {
	ServiceName string
	Port        int
	DatabaseUrl string
}

func getConfig() *Config {
	config := new(Config)
	env := os.Getenv("ENV")

	// Load in defaults here
	config.ServiceName = "myservice"
	config.Port = 8080

	// Environment-specific
	switch {
	case env == "", env == "test":
		config.DatabaseUrl = "sqlite://:memory:"

	case env == "dev", env == "development":
		config.DatabaseUrl = "postgresql://dev-db/something"

	case env == "stage", env == "staging":
		config.DatabaseUrl = "postgresql://stage-db/something"

	case env == "prod", env == "production":
		config.DatabaseUrl = "postgresql://prod-db/something"

	}

	return config
}
