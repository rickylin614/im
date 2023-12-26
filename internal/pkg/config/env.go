package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// updateConfigFromEnv env label fields would check os env variable instead of config setting.
func updateConfigFromEnv(config any) error {
	val := reflect.ValueOf(config).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Struct {
			err := updateConfigFromEnv(field.Addr().Interface())
			if err != nil {
				return err
			}
		} else {
			tag, ok := val.Type().Field(i).Tag.Lookup("env")
			if ok && len(tag) > 0 {
				envValue := os.Getenv(tag)
				if envValue != "" {
					err := setValue(field, envValue)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func setValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.Int:
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intVal))
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolVal)
	case reflect.String:
		field.SetString(value)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
	return nil
}
