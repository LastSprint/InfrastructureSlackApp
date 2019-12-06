package utils

import "testing"

func TestJoinByCharacterNoAddDelimToOneElement(t *testing.T) {
	// Arrange

	arr := []string{"tmp"}

	// Act

	res := JoinByCharacter(arr, ",", "")

	// Assert

	if res != "tmp" {
		t.Fail()
	}
}

func TestJoinByCharacterAddDelim(t *testing.T) {
	// Arrange

	arr := []string{"foo", "bar"}

	// Act

	res := JoinByCharacter(arr, ",", "")

	// Assert

	if res != "foo,bar" {
		t.Fail()
	}
}

func TestJoinByCharacterAddSurrounds(t *testing.T) {
	// Arrange

	arr := []string{"foo", "bar"}

	// Act

	res := JoinByCharacter(arr, ",", "@")

	// Assert

	if res != "@foo@,@bar@" {
		t.Fail()
	}
}

func TestJoinByCharacterAddSurroundsToOneSymbol(t *testing.T) {
	// Arrange

	arr := []string{"foo"}

	// Act

	res := JoinByCharacter(arr, ",", "@")

	// Assert

	if res != "@foo@" {
		t.Fail()
	}
}
