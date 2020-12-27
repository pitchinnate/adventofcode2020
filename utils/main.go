package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func ReadInputFile(path string) string {
	contents, _ := ioutil.ReadFile(fmt.Sprintf("c:/www/adventofcode2020/%s", path))
	return string(contents)
}

func SplitByLine(content string) []string {
	return strings.Split(content, "\n")
}

func StringsToInts(content []string) []int64 {
	var vals []int64
	for _, val := range content {
		trimmed := strings.TrimSpace(val)
		if trimmed != "" {
			intVal, _ := strconv.Atoi(trimmed)
			vals = append(vals, int64(intVal))
		}
	}
	return vals
}

func SetField(v interface{}, name string, value string) error {
	// v must be a pointer to a struct
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}

	// Dereference pointer
	rv = rv.Elem()

	// Lookup field by name
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	// Field must be exported
	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	// We expect a string field
	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}

	// Set the value
	fv.SetString(value)
	return nil
}
