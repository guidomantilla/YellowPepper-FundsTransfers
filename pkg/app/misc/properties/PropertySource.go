package properties

type PropertySource interface {
	Get(property string) string
	AsMap() map[string]interface{}
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

func (source *DefaultPropertySource) AsMap() map[string]interface{} {

	internalMap := make(map[string]interface{})
	internalMap["name"], internalMap["value"] = source.name, source.properties.AsMap()
	return internalMap
}
