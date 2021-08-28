package config

import (
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/properties"
	"os"
)

var singletonEnvironment environment.Environment

func LoadEnvironment(args *[]string) environment.Environment {
	if singletonEnvironment == nil {
		osArgs := os.Environ()
		osSource := properties.NewDefaultPropertySource("OS", properties.NewPropertiesFromArray(&osArgs))
		argsSource := properties.NewDefaultPropertySource("CMD", properties.NewPropertiesFromArray(args))
		singletonEnvironment = environment.NewDefaultEnvironment(argsSource, osSource)

		//ioutil.ReadFile()
	}
	return singletonEnvironment
}
