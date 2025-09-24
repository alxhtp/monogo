package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds all configuration for the application
type Config struct {
	AppConfig
	SwaggerAuth
	DatabaseConfig
	LogConfig
	CORSConfig
	JWTConfig
	RateLimitConfig
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	AppName      string `envconfig:"APP_NAME" default:"app"`
	Environment  string `envconfig:"APP_ENVIRONMENT" default:"development"`
	Debug        bool   `envconfig:"APP_DEBUG" default:"false"`
	AppPort      int    `envconfig:"APP_PORT" default:"8080"`
	AppHost      string `envconfig:"APP_HOST" default:"localhost"`
	Scheme       string `envconfig:"APP_SCHEME" default:"http"`
	ReadTimeout  int    `envconfig:"APP_READ_TIMEOUT" default:"10"`
	WriteTimeout int    `envconfig:"APP_WRITE_TIMEOUT" default:"120"`
	IdleTimeout  int    `envconfig:"APP_IDLE_TIMEOUT" default:"60"`
	BodyLimit    int    `envconfig:"APP_BODY_LIMIT" default:"4"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	DBHost          string `envconfig:"DB_HOST" default:"localhost"`
	DBPort          int    `envconfig:"DB_PORT" default:"5432"`
	DBName          string `envconfig:"DB_NAME" default:"app"`
	DBUser          string `envconfig:"DB_USER" default:"user"`
	DBPassword      string `envconfig:"DB_PASSWORD" default:"password"`
	DBSSLMode       string `envconfig:"DB_SSL_MODE" default:"disable"`
	MaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS" default:"10"`
	MaxOpenConns    int    `envconfig:"DB_MAX_OPEN_CONNS" default:"100"`
	ConnMaxLifetime string `envconfig:"DB_CONN_MAX_LIFETIME" default:"1h"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level  string `envconfig:"LOG_LEVEL" default:"info"`
	Format string `envconfig:"LOG_FORMAT" default:"json"`
	Output string `envconfig:"LOG_OUTPUT" default:"stdout"`
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins   []string `envconfig:"CORS_ALLOWED_ORIGINS" default:"*"`
	AllowedMethods   []string `envconfig:"CORS_ALLOWED_METHODS" default:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   []string `envconfig:"CORS_ALLOWED_HEADERS" default:"Origin,Content-Type,Accept,Authorization"`
	AllowCredentials bool     `envconfig:"CORS_ALLOW_CREDENTIALS" default:"true"`
	MaxAge           int      `envconfig:"CORS_MAX_AGE" default:"300"`
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey                string `envconfig:"JWT_SECRET_KEY" required:"true"`
	AccessTokenExpiryInHours int    `envconfig:"JWT_ACCESS_TOKEN_EXPIRY_IN_HOURS" default:"1"`
	RefreshTokenExpiryInDays int    `envconfig:"JWT_REFRESH_TOKEN_EXPIRY_IN_DAYS" default:"7"`
	TokenIssuer              string `envconfig:"JWT_TOKEN_ISSUER" default:"upskill-tb"`
	TokenAudience            string `envconfig:"JWT_TOKEN_AUDIENCE" default:"upskill-tb-api"`
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	MaxRequests int `envconfig:"RATE_LIMIT_MAX_REQUESTS" default:"100"`
}

// SwaggerAuth holds swagger authentication configuration
type SwaggerAuth struct {
	SwaggerUsername string `envconfig:"SWAGGER_USERNAME" required:"true"`
	SwaggerPassword string `envconfig:"SWAGGER_PASSWORD" required:"true"`
}

// Load loads configuration
func Load() (*Config, error) {
	var cfg Config

	_ = godotenv.Load()

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("envconfig: %w", err)
	}

	return &cfg, nil
}
