package environment

import (
	"YellowPepper-FundsTransfers/pkg/misc/properties"
)

type Environment interface {
	GetValue(property string) EnvVar
	GetValueOrDefault(property string, defaultValue string) EnvVar
	GetPropertySources() []properties.PropertySource
	AppendPropertySources(propertySources ...properties.PropertySource)
}

type DefaultEnvironment struct {
	propertySources []properties.PropertySource
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

func (environment *DefaultEnvironment) AppendPropertySources(propertySources ...properties.PropertySource) {
	environment.propertySources = append(environment.propertySources, propertySources...)
}

// DefaultEnvironmentBuilder

type DefaultEnvironmentBuilder struct {
	defaultEnvironment *DefaultEnvironment
}

func NewDefaultEnvironment() *DefaultEnvironmentBuilder {
	return &DefaultEnvironmentBuilder{
		defaultEnvironment: &DefaultEnvironment{},
	}
}
func (builder *DefaultEnvironmentBuilder) WithPropertySources(propertySources ...properties.PropertySource) *DefaultEnvironmentBuilder {
	builder.defaultEnvironment.propertySources = propertySources
	return builder
}

func (builder *DefaultEnvironmentBuilder) Build() *DefaultEnvironment {
	return builder.defaultEnvironment
}
