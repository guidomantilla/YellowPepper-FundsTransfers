package config

import (
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/properties"
	"os"
)

const (
	OS_PROPERTY_SOURCE_NAME  = "OS"
	CMD_PROPERTY_SOURCE_NAME = "CMD"
)

func LoadEnvironment() environment.Environment {

	//ioutil.ReadFile()

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties.NewPropertiesFromArray(&osArgs))

	cmdArgs := os.Args[1:]
	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties.NewPropertiesFromArray(&cmdArgs))

	env := environment.NewDefaultEnvironment(cmdSource, osSource)
	return env
}
