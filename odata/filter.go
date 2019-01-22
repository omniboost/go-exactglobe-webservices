package odata

func NewFilter() *Filter {
	return &Filter{}
}

type Filter struct {
	Query string
}

func (f *Filter) Set(q string) {
	f.Query = q
}

func (f *Filter) MarshalSchema() string {
	return f.Query
}
