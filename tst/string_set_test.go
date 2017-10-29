package test

import (
	"testing"

	"../src/util"
)

func contains(slice []string, val string) bool {
	for _, elem := range slice {
		if elem == val {
			return true
		}
	}
	return false
}

func containsAll(slice []string, vals ...string) bool {
	for _, val := range vals {
		if !contains(slice, val) {
			return false
		}
	}
	return true
}

func TestValuesReturnsAllValues(t *testing.T) {
	set := util.NewStringSet("Hello", "World", "My", "Friend")
	values := set.Values()
	if !containsAll(values, "Hello", "World", "My", "Friend") {
		t.Error("Values did not give back the correct elements.")
	}
}
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
	set.Add("World")

	if !set.Contains("Hello") || !set.Contains("World") {
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

func TestRemovingElementThatDoesntExistReturnsError(t *testing.T) {
	set := util.NewStringSet("Hello", "World")
	err := set.Remove("Element that doesn't exist in the set.")
	if err == nil { // no error.
		t.Error("Did not return error on removal of element that didn't exist.", err)
	}
}
