package app

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations and
// can be used in any place with app.Config.Variable
var Config appConfig

type appConfig struct {
	// Environment select the environment that will be used. Defaults to "production"
	Environment string `mapstructure:"environment"`
	// DatabaseFile yaml file with database connection data. Defaults to "config/database.yml"
	DatabaseFile string `mapstructure:"database_file"`
	// Migrate is a flag to rebuild the database (Migration). Defaults to false
	Migrate bool `mapstructure:"migrate"`
	// ShowSQL show all sql queries in cmd. Defaults to false
	ShowSQL bool `mapstructure:"show_sql"`
	// ServerPort is the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// JWTSigningMethod is the signing method for JWT (encryption). Defaults to "HS256"
	JWTSigningMethod string `mapstructure:"jwt_signing_method"`
	// JWTSigningKey is the JWT signing key. required.
	JWTSigningKey string `mapstructure:"jwt_signing_key"`
	// JWTVerificationKey is the JWT verification key. required.
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

// Validate check if the required config about the aplication is filled.
// Emmits a panic error if doesn't
func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
//
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically, like
// RESTFUL_SHOW_SQL=true or RESTFUL_MIGRATE=true (All characters with capslock)
//
// All load env variables area added into Config variable.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")
	v.SetDefault("environment", "production")
	v.SetDefault("database_file", "config/database.yml")
	v.SetDefault("migrate", false)
	v.SetDefault("show_sql", false)
	v.SetDefault("server_port", 8080)
	v.SetDefault("jwt_signing_method", "HS256")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
