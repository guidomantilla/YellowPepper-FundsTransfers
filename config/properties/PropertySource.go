package properties

type PropertySource interface {
	Get(property string) string
}

type DefaultPropertySource struct {
	name       string
	properties Properties
}

func NewDefaultPropertySource(name string, properties Properties) *DefaultPropertySource {
	return &DefaultPropertySource{
		name:       name,
		properties: properties,
	}
}

func (source *DefaultPropertySource) Get(property string) string {
	return source.properties.Get(property)
}
