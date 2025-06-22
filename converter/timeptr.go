package converter

import (
	"reflect"
	"time"
)

func init() {
	var timeVar time.Time
	registerConverter(reflect.TypeOf(&timeVar), StringToTimePtr)
}

func StringToTimePtr(value string) (interface{}, error) {
	timeValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, err
	}
	return &timeValue, nil
}
