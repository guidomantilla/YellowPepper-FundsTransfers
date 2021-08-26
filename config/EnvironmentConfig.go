package config

import (
	"YellowPepper-FundsTransfers/config/properties"
	"os"
)

type Environment interface {
	Get(property string) string
}

type DefaultEnvironment struct {
	propertySources []properties.PropertySource
}

func NewDefaultEnvironment(propertySources ...properties.PropertySource) *DefaultEnvironment {

	var propertySourcesArray []properties.PropertySource
	for _, source := range propertySources {
		propertySourcesArray = append(propertySources, source)
	}

	return &DefaultEnvironment{
		propertySources: propertySourcesArray,
	}
}

func (environment *DefaultEnvironment) Get(property string) string {

	var value string
	for _, source := range environment.propertySources {
		internalValue := source.Get(property)
		if internalValue != "" {
			value = internalValue
			break
		}
	}
	return value
}

var environment Environment
func LoadEnvironment(args *[]string) Environment {
	if environment == nil {
		osArgs := os.Environ()
		osSource := properties.NewDefaultPropertySource("OS", properties.NewPropertiesFromArray(&osArgs))
		argsSource := properties.NewDefaultPropertySource("CMD", properties.NewPropertiesFromArray(args))
		environment = NewDefaultEnvironment(argsSource, osSource)
	}
	return environment
}

