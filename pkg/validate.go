package pkg

import "reflect"

func IsStructContainNil(s interface{}) bool {
	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).IsNil() {
			return true
		}
	}
	return false
}
