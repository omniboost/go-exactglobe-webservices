package edm

import (
	"encoding/json"
	"regexp"
	"strconv"
	"time"
)

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time)
}

func (d DateTime) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// /Date(1488939627017)/
	re := regexp.MustCompile(`[0-9]+`)
	match := re.FindString(value)
	if match == "" {
		return nil
	}

	milis, err := strconv.Atoi(match)
	if err != nil {
		return err
	}

	// new Date(milis)
	d.Time = time.Unix(0, int64(milis)*int64(time.Millisecond))
	return err
}
