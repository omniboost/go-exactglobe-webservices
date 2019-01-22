package utils

import (
	nativeurl "net/url"
)

type URL struct {
	nativeurl.URL
}

// UnmarshalText calls url.Parse
func (u *URL) UnmarshalText(p []byte) error {
	nu, err := nativeurl.Parse(string(p))
	if err != nil {
		return err
	}
	// (*u) = URL(*nu)
	u.URL = nativeurl.URL(*nu)
	return nil
}

// MarshalText just calls String()
func (u *URL) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}
