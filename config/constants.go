package config

import "os"

const (
	DeployConfigFile string = "config/deployment/deployment.yaml"
	LocalConfigFile  string = "config/local/local.yaml"
)

var Env = os.Getenv("ENV_TYPE")
