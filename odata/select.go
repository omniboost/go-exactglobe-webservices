package odata

import "strings"

func NewSelect(allowed []string) *Select {
	return &Select{
		Values:  []string{},
		allowed: allowed,
	}
}

type Select struct {
	Values  []string
	allowed []string
}

func (s *Select) Add(key string) bool {
	keys := strings.Split(key, ",")
	allowed := []string{}

	for _, k := range keys {
		if !s.IsAllowed(k) {
			return false
		}
		allowed = append(allowed, k)
	}

	for _, key := range allowed {
		s.Values = append(s.Values, key)
	}
	return true
}

func (s *Select) IsAllowed(key string) bool {
	ok := false
	for _, a := range s.allowed {
		if a == key {
			ok = true
			continue
		}
	}
	return ok
}

func (t *Select) MarshalSchema() string {
	return strings.Join(t.Values, ",")
}
