package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type appEnvType string

// IsDev returns true if we are in dev mode.
func (a *appEnvType) IsDev() bool {
	return *a == Development
}

const (
	Production  appEnvType = "PROD"
	Development appEnvType = "DEV"
)

var (
	InternalIngressPort string
	InternalIngressUrl  string
)

func init() {
	InternalIngressPort = GetOptionalEnv("INTERNAL_INGRESS_PORT", "51377")
	InternalIngressUrl = GetOptionalEnv("INTERNAL_INGRESS_URL", "http://localhost:51377")
}

func GetOptionalEnv(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}

func RequireEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logrus.Panicf("Required environment variable '%s' is not set. Please create a .env file at the root of the project with the appropriate configuration set.", key)
	}

	return val
}

func RequireSensitiveEnv(key string, appEnv appEnvType) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logrus.Panicf("Required environment variable '%s' is not set. Please create a .env file at the root of the project with the appropriate configuration set.", key)
	}

	if (strings.HasPrefix(val, "${{") && strings.HasSuffix(val, "}}")) || val == "" {
		logrus.Panicf("Environment variable '%s' is invalid and set to '%s'. Please make sure your environment variables are set to safe secret UUIDs.", val, key)
	}

	return val
}

func requireAppEnv(key string) appEnvType {
	val := RequireEnv(key)
	switch appEnvType(val) {
	case Development:
		return Development
	case Production:
		return Production
	default:
		logrus.Errorf("Invalid %s provided, falling back to %s", key, Production)
		return Production
	}
}
