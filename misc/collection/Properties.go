package collection

type Properties struct {
	internalMap map[string]string
}

func NewProperties() *Properties {
	return &Properties{
		internalMap: make(map[string]string, 0),
	}
}

func (p *Properties) Add(property string, value string) {
	if p.internalMap[property] == "" {
		p.internalMap[property] = value
	}
}

func (p *Properties) Get(property string) string {
	return p.internalMap[property]
}
