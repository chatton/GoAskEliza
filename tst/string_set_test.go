package test

import (
	"testing"

	"../src/util"
)

func TestIsEmpty(t *testing.T) {
	set := util.NewStringSet()
	if !set.IsEmpty() {
		t.Error("Set was not empty before adding elements.")
	}

	set.Add("Hello")
	if set.IsEmpty() {
		t.Error("Set was empty after adding elements.")
	}

}
func TestAddStringSetAddsElement(t *testing.T) {
	set := util.NewStringSet()

	set.Add("Hello")
	if !set.Contains("Hello") {
		t.Error("Set did not contain added element.")
	}

	set.Add("World")
	if !set.Contains("World") {
		t.Error("Set did not contain added element.")
	}

}

func TestRemoveStringRemovesElement(t *testing.T) {
	set := util.NewStringSet("Hello", "World")

	set.Remove("Hello")
	if set.Contains("Hello") {
		t.Error("Element was not removed correctly")
	}

	set.Remove("World")
	if set.Contains("World") {
		t.Error("Element was not removed correctly")
	}
}
