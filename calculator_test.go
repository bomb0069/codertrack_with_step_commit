package codetrack

import "testing"

func Test_1_add_2_is_3(t *testing.T) {
	expected := "3"
	actual := add("1", "2")
	if expected != actual {
		t.Error(expected + " != " + actual)
	}

}

func Test_2_add_3_is_5(t *testing.T) {
	expected := "5"
	actual := add("2", "3")
	if expected != actual {
		t.Error(expected + " != " + actual)
	}

}

func Test_max_length_of_int(t *testing.T) {
	expected := "19999999999999999998"
	actual := add("9999999999999999999", "9999999999999999999")
	if expected != actual {
		t.Error(expected + " != " + actual)
	}

}
