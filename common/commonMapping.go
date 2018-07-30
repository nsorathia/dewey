package common

import (
	"fmt"
	"time"
)

//GetInteger reads a given key from an object map, parses its value and returns an integer
func GetInteger(data map[string]interface{}, key string) (int, error) {
	val, ok := data[key]
	if ok {
		floatVal, ok := val.(float64)
		if ok {
			return int(floatVal), nil
		}
		return 0, fmt.Errorf("%v value could not be parsed to integer:  %v", key, val)
	}
	return 0, fmt.Errorf("'%v' key not found in object", key)
}

//GetString reads a given key from an object map, parses its value and returns a string
func GetString(data map[string]interface{}, key string) (string, error) {
	val, ok := data[key]
	if !ok {
		return "", fmt.Errorf("'%v' key not found in object", key)
	}
	return val.(string), nil
}

//GetDate reads a given key from an object map, parses its value and returns a Time
func GetDate(data map[string]interface{}, key string) (time.Time, error) {
	val, ok := data[key]
	if !ok {
		return time.Now(), fmt.Errorf("'%v' key not found in object", key)
	}
	dateObj, err := time.Parse(time.RFC3339, val.(string))
	if err != nil {
		return time.Now(), fmt.Errorf("'%v' val could not be parsed to Time format", key)
	}
	return dateObj, nil
}
