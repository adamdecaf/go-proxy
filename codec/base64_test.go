package codec

import (
	"testing"
)

func TestToAndFromBase64(t *testing.T) {
	examples := []string{"", " ", "a", "foo", "FOO", "/url/foo",}
	for i := range examples {
		r1 := examples[i]
		r2 := FromBase64(ToBase64(r1))
		if r1 != r2 {
			t.Fatalf("unable to decode our encoding for '%s'", r1)
		}
	}
}
