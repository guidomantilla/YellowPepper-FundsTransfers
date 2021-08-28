package environment

import (
	"YellowPepper-FundsTransfers/pkg/misc/properties"
)

type Environment interface {
	GetValue(property string) EnvVar
	GetValueOrDefault(property string, defaultValue string) EnvVar
	GetPropertySources() []properties.PropertySource
}

type DefaultEnvironment struct {
	propertySources []properties.PropertySource
}

func NewDefaultEnvironment(propertySources ...properties.PropertySource) *DefaultEnvironment {

	propertySourcesArray := make([]properties.PropertySource, len(propertySources))
	for index, source := range propertySources {
		propertySourcesArray[index] = source
	}

	return &DefaultEnvironment{
		propertySources: propertySourcesArray,
	}
}

func (environment *DefaultEnvironment) GetValue(property string) EnvVar {

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

func (environment *DefaultEnvironment) GetValueOrDefault(property string, defaultValue string) EnvVar {

	envVar := environment.GetValue(property)
	if envVar != "" {
		return envVar
	}
	return NewEnvVar(defaultValue)
}

func (environment *DefaultEnvironment) GetPropertySources() []properties.PropertySource {
	return environment.propertySources
}
