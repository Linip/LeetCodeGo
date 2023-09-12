package RansomNote

import "testing"

func TestCanConstruct1(t *testing.T) {
	ransomNote := "a"
	magazine := "b"

	expected := false

	output := canConstruct(ransomNote, magazine)

	if output != expected {
		t.Error(
			"For", ransomNote, magazine,
			"expected", expected,
			"got", output,
		)
	}
}

func TestCanConstruct2(t *testing.T) {
	ransomNote := "aa"
	magazine := "ab"

	expected := false

	output := canConstruct(ransomNote, magazine)

	if output != expected {
		t.Error(
			"For", ransomNote, magazine,
			"expected", expected,
			"got", output,
		)
	}
}

func TestCanConstruct3(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"

	expected := true

	output := canConstruct(ransomNote, magazine)

	if output != expected {
		t.Error(
			"For", ransomNote, magazine,
			"expected", expected,
			"got", output,
		)
	}
}

func TestCanConstruct4(t *testing.T) {
	ransomNote := "aab"
	magazine := "baa"

	expected := true

	output := canConstruct(ransomNote, magazine)

	if output != expected {
		t.Error(
			"For", ransomNote, magazine,
			"expected", expected,
			"got", output,
		)
	}
}
