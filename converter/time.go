package converter

import (
	"fmt"
	"reflect"
	"time"
)

func init() {
	registerConverter(reflect.TypeOf(time.Time{}), StringToTime)
}

func StringToTime(value string) (interface{}, error) {
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05",
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822Z,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, value); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", value)
}
