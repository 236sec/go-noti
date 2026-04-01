package config

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"goboilerplate.com/src/pkg/utils"
)

// App config struct
type Config struct {
	YMLConfig YMLConfig
	EnvConfig EnvConfig
}

type YMLConfig struct {
	Server    ServerConfig
	Swagger   SwaggerConfig
	Database  DatabaseConfig
	Telemetry TelemetryConfig
}

type EnvConfig struct {
	Redis    RedisConfig
	Postgres PostgresConfig
	Email    EmailConfig
	MongoDB  MongoDBConfig
}

type DatabaseConfig struct {
	Driver string
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Email config
type EmailConfig struct {
	EmailHost     string `env:"EMAIL_HOST,default=smtp.gmail.com"`
	EmailUsername string `env:"EMAIL_USERNAME,default="`
	EmailPassword string `env:"EMAIL_PASSWORD,default="`
	EmailFrom     string `env:"EMAIL_FROM,default="`
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string `env:"POSTGRES_HOST,default=localhost"`
	PostgresqlPort     string `env:"POSTGRES_PORT,default=5432"`
	PostgresqlUser     string `env:"POSTGRES_USER,default=postgres"`
	PostgresqlPassword string `env:"POSTGRES_PASSWORD,default=password"`
	PostgresqlDbname   string `env:"POSTGRES_DB,default=postgres"`
}

// MongoDB config
type MongoDBConfig struct {
	MongodbHost     string `env:"MONGO_HOST,default=localhost"`
	MongodbPort     string `env:"MONGO_PORT,default=27017"`
	MongodbUser     string `env:"MONGO_USER,default=root"`
	MongodbPassword string `env:"MONGO_PASSWORD,default=password"`
	MongodbDbname   string `env:"MONGO_DB,default=local"`
}

// Redis config
type RedisConfig struct {
	RedisAddr      string `env:"REDIS_ADDR,default=localhost:6379"`
	RedisPassword  string `env:"REDIS_PASSWORD,default=password"`
	RedisDB        int    `env:"REDIS_DB,default=0"`
	RedisDefaultdb int    `env:"REDIS_DEFAULTDB,default=0"`
	MinIdleConns   int    `env:"MIN_IDLE_CONNS,default=0"`
	PoolSize       int    `env:"REDIS_POOL_SIZE,default=10"`
	PoolTimeout    int    `env:"REDIS_POOL_TIMEOUT,default=0"`
	Protocol       int    `env:"REDIS_PROTOCOL,default=2"`
}

// Swagger config
type SwaggerConfig struct {
	BasePath string
	FilePath string
	Path     string
	Title    string
}

// Telemetry config
type TelemetryConfig struct {
	Enabled     bool   `env:"TELEMETRY_ENABLED,default=false"`
	ServiceName string `env:"TELEMETRY_SERVICE_NAME,default=boilerplate-service"`
}

var GetConfig = sync.OnceValue(func() *Config {
	return NewConfig()
})

func NewConfig() *Config {
	configPath := utils.GetConfigPath(os.Getenv("config"))

	v, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	ymlConfig, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	envConfig := LoadEnvConfig()

	return &Config{
		YMLConfig: *ymlConfig,
		EnvConfig: *envConfig,
	}
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*YMLConfig, error) {
	var c YMLConfig

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

func LoadEnvConfig() *EnvConfig {
	// Load .env file if present; env vars already set in the OS take precedence.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to OS environment variables")
	}

	var environment EnvConfig
	if _, err := env.UnmarshalFromEnviron(&environment); err != nil {
		log.Fatal(err)
	}
	return &environment
}
