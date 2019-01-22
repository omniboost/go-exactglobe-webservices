package odata

import (
	"fmt"
	"strings"
)

func NewOrderBy(allowed []string) *OrderBy {
	return &OrderBy{
		Values:  map[string]string{},
		allowed: allowed,
	}
}

type OrderBy struct {
	Values  map[string]string
	allowed []string
}

func (o *OrderBy) Add(field string, order string) bool {
	if !o.IsAllowed(field) {
		return false
	}

	o.Values[field] = order
	return true
}

func (o *OrderBy) IsAllowed(key string) bool {
	ok := false
	for _, a := range o.allowed {
		if a == key {
			ok = true
			continue
		}
	}
	return ok
}

func (o *OrderBy) MarshalSchema() string {
	pairs := []string{}
	for k, v := range o.Values {
		v = strings.ToLower(v)
		pair := fmt.Sprintf("%s %s", k, v)
		pairs = append(pairs, pair)
	}
	return strings.Join(pairs, ",")
}
