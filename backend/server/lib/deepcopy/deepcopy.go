package deepcopy

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Struct(oldObj interface{}) interface{} {
	newObj := reflect.New(reflect.TypeOf(oldObj).Elem())
	oldVal := reflect.ValueOf(oldObj).Elem()
	newVal := newObj.Elem()
	for i := 0; i < oldVal.NumField(); i++ {
		newValField := newVal.Field(i)
		if newValField.CanSet() {
			newValField.Set(oldVal.Field(i))
		}
	}

	return newObj.Interface()
}

func Any(val interface{}) interface{} {
	return nil
}

func Map(src map[string]interface{}) map[string]interface{} {
	return nil
}

func Copy(dst interface{}, src interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal into dst: %s", err)
	}
	return nil
}
