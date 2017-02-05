package main

import "testing"

func Exampleoutput() {
	output("test")
	// Output:
	// test
}

var data = []struct {
	in  string
	out string
}{
	{"test.txt", "hoge\nfuga"},
}

func TestGetData(t *testing.T) {
	for _, tt := range data {
		got, err := getData(tt.in)
		if err != nil {
			t.Error(err)
		}
		want := tt.out
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
}
func TestReplaceTemplate(t *testing.T) {
	templ := "hoge:1 fuga:2"
	r := []string{"test", "test2"}
	got := replaceTemplate(templ, r)
	want := "hogetest fugatest2"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
