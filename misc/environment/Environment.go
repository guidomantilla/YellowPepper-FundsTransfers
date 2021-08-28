package environment

import (
	"YellowPepper-FundsTransfers/misc/properties"
)

type Environment interface {
	Get(property string) EnvVar
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

func (environment *DefaultEnvironment) Get(property string) EnvVar {

	var value string
	for _, source := range environment.propertySources {
		internalValue := source.Get(property)
		if internalValue != "" {
			value = internalValue
			break
		}
	}
	return NewEnvVar(value)
}
