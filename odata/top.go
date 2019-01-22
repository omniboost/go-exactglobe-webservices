package odata

import "strconv"

func NewTop() *Top {
	return &Top{}
}

type Top struct {
	i int
}

func (t *Top) Set(i int) {
	t.i = i
}

func (t *Top) MarshalSchema() string {
	i := int64(t.i)
	if i == 0 {
		return ""
	}

	return strconv.FormatInt(i, 10)
}
