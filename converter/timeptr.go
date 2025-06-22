package converter

import (
	"reflect"
	"time"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	var timeVar time.Time
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&timeVar), StringToTimePtr)
}

func StringToTimePtr(value string) (interface{}, error) {
	timeValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, err
	}
	return &timeValue, nil
}
