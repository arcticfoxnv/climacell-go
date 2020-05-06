package climacell

type DataField string
type DataFieldList []DataField

type DataPoint struct {
	Units string      `json:"units,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type UnitSystem string

func (d DataField) String() string {
	return string(d)
}

func (d DataFieldList) Strings() []string {
	l := make([]string, len(d))
	for i, v := range d {
		l[i] = v.String()
	}

	return l
}

func (d DataPoint) Present() bool {
	return d.Value != nil
}

func (d DataPoint) Unset() bool {
	return !d.Present()
}

func (u UnitSystem) String() string {
	return string(u)
}
