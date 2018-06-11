package app

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// escolha qual ambiente será utilizado
	Environment string `mapstructure:"environment"`
	// local do arquivo yaml com os dados de conexão
	DatabaseFile string `mapstructure:"database_file"`
	// faz a migration do banco de dados
	Migrate bool `mapstructure:"migrate"`
	// faz o log de todas as queries executadas
	ShowSQL bool `mapstructure:"show_sql"`
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the signing method for JWT. Defaults to "HS256"
	JWTSigningMethod string `mapstructure:"jwt_signing_method"`
	// JWT signing key. required.
	JWTSigningKey string `mapstructure:"jwt_signing_key"`
	// JWT verification key. required.
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically.
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
