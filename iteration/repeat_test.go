package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}

func TestCompare(t *testing.T) {
	compared := strings.Compare("aaa", "aaa")
	expected := 0

	if compared != expected {
		t.Errorf("expected %v but got %v", expected, compared)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b")
	fmt.Println(repeated)
	// output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
