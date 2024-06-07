package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should repeat character given (5 times)", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("should return empty string if negative count is provided", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""

		if repeated != expected {
			t.Errorf("negative count value should return default string value (empty)")
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)
	// Output: aaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
