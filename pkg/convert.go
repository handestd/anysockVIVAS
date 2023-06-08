package pkg

import (
	"fmt"
	"strconv"
)

func InterfaceToInt64(t interface{}) (int64, error) {
	switch t := t.(type) { // This is a type switch.
	case int64:
		return t, nil // All done if we got an int64.
	case int:
		return int64(t), nil // This uses a conversion from int to int64
	case string:
		return strconv.ParseInt(t, 10, 64)
	default:
		return 0, fmt.Errorf("type %T not supported", t)
	}
}
