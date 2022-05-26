package mutils

import (
	"reflect"
	"strings"
)

/*
SliceContains - checks if any of the slice values contains specified content.
Panics if input and content type do not match.
Allowed types: string | int | float64

*/
func SliceContains(input interface{}, content interface{}) bool {

	inType := reflect.TypeOf(input).String()
	conType := reflect.TypeOf(content).String()

	// panic if types are not matching
	if !strings.Contains(inType, conType) {
		panic("type missmatch in 'SliceContains', expected []" + conType + ", got " + inType)
	}

	switch t := input.(type) {

	// slice of strings
	case []string:
		for _, value := range t {
			if value == content.(string) {
				return true
			}
		}

	// slice of ints
	case []int:
		for _, value := range t {
			if value == content.(int) {
				return true
			}
		}

	// slice of floats
	case []float64:
		for _, value := range t {
			if value == content.(float64) {
				return true
			}
		}

	default:
		return false

	}

	return false

}
