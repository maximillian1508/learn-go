package hello

import (
	"testing"
)

func TestSay(t *testing.T) { // to run the test, use the command: go test ./..., this will run all the tests in the current directory and all subdirectories
	want := "Hello, World!" // expected result
	got := Say("World") // actual result to be tested

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
} // the command result will be in the terminal, if the test fails, the test will be marked as failed, if the test passes, nothing will be printed

func TestSayMany(t *testing.T){
	want := "Hello, world!" // expected result
	got := SayMany([]string{"world"})

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestSaySubtest(t *testing.T) {
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Max"},
			result: "Hello, Max!",
		},
	}

	for _, st := range subtests {
		if s := SayMany(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s);
		}
	}
}