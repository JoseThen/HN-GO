package main

import (
	"reflect"
	"testing"

	get "github.com/JoseThen/HN-GO/util"
)

const testNum = 81
const testType = "string"

func TestIdGetter(t *testing.T) {
	got := get.Ids(testNum, "top")

	if len(got) != testNum {
		t.Errorf("got '%d' want '%d'", len(got), testNum)
	}
}

func TestDataGetter(t *testing.T) {
	got := get.Data(1)
	theType := reflect.TypeOf(got.Title)

	if theType != reflect.TypeOf(testType) {
		t.Errorf("got %s want $%s", theType, testType)
	}
}
