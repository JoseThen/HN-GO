package main

import (
	"reflect"
	"testing"
)

const testNum = 81
const testType = "string"

func TestIdGetter(t *testing.T) {
	got := getIds(testNum, "top")

	if len(got) != testNum {
		t.Errorf("got '%d' want '%d'", len(got), testNum)
	}
}

func TestDataGetter(t *testing.T) {
	got := getData(1)
	theType := reflect.TypeOf(got.Title)

	if theType != reflect.TypeOf(testType) {
		t.Errorf("got %s want $%s", theType, testType)
	}
}
