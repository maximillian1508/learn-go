package hello

import (
	"testing"
)

func TestSay(t *testing.T) { // to run the test, use the command: go test ./..., this will run all the tests in the current directory and all subdirectories
	// to run a specific test, use the command: go test -run TestSay ./...
	want := "Hello, World!" // expected result
	got := Say("World")     // actual result to be tested

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
} // the command result will be in the terminal, if the test fails, the test will be marked as failed, if the test passes, nothing will be printed

func TestSayMany(t *testing.T) {
	want := "Hello, world!" // expected result
	got := SayMany([]string{"world"})

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

// Subtests are a way to test multiple scenarios in a single test function
// Subtests in this case is using the Table-Driven Test pattern
// The pattern is to define a table of test cases, and then iterate over the table and run the test for each case
func TestSaySubtest(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Max"},
			result: "Hello, Max!",
		},
	}

	for _, st := range subtests {
		if s := SayMany(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}

// One way to improve this pattern is to use the t.Run function to run each subtest as a separate test
// This way, we can see the output of each subtest in the terminal
// We can also run a specific subtest: go test -run TestSayNameSubtest/Max ./...
func TestSayNameSubtest(t *testing.T) {
	subtests := []struct {
		name   string
		items  []string
		result string
	}{
		{
			name:   "World",
			items:  []string{"World"},
			result: "Hello, World!",
		},
		{
			name:   "Max",
			items:  []string{"Max"},
			result: "Hello, Max!",
		},
	}

	for _, st := range subtests {
		t.Run(st.name, func(t *testing.T) {
			if s := SayMany(st.items); s != st.result {
				t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
			}
		})
	}
}
