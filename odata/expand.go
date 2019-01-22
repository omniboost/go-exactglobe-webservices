package odata

import "strings"

func NewExpand(allowed []string) *Expand {
	return &Expand{
		Values:  []string{},
		allowed: allowed,
	}
}

type Expand struct {
	Values  []string
	allowed []string
}

func (e *Expand) Add(key string) bool {
	keys := strings.Split(key, ",")
	allowed := []string{}

	for _, k := range keys {
		if !e.IsAllowed(k) {
			return false
		}
		allowed = append(allowed, k)
	}

	for _, key := range allowed {
		e.Values = append(e.Values, key)
	}
	return true
}

func (e *Expand) IsAllowed(key string) bool {
	ok := false
	for _, a := range e.allowed {
		if a == key {
			ok = true
			continue
		}
	}
	return ok
}

func (e *Expand) Query() string {
	return strings.Join(e.Values, ",")
}

func (e *Expand) MarshalSchema() string {
	return e.Query()
}
