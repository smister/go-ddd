package attr

type attributer interface {
	Set(string, interface{})
	Attributes() map[string]interface{}
	SetAttributes(attrs map[string]interface{})
	Refresh()
}

type Attribute struct {
	attributes map[string]interface{}
}

func (a *Attribute) Set(f string, v interface{}) {
	if a.attributes == nil {
		a.attributes = make(map[string]interface{})
	}

	a.attributes[f] = v
}

func (a *Attribute) SetAttributes(attrs map[string]interface{}) {
	a.attributes = attrs
}

func (a *Attribute) Refresh() {
	a.attributes = nil
}

func (a *Attribute) Attributes() map[string]interface{} {
	return a.attributes
}
