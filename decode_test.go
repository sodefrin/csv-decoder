package csvdecoder

import (
	"bytes"
	"errors"
	"testing"
)

func TestDecode(t *testing.T) {
	type A struct {
		Test  string
		Test2 *string
	}

	t.Run("invalid not slice", func(t *testing.T) {
		data := []byte("")
		out := A{}
		r := bytes.NewReader(data)
		if err := Decode(out, r); !errors.Is(err, ErrInvalidInterface) {
			t.Fatalf("want err: %s, but have: %s", ErrInvalidInterface, err)
		}
	})

	t.Run("invliad not pointer", func(t *testing.T) {
		data := []byte("")
		out := []A{}
		r := bytes.NewReader(data)
		if err := Decode(out, r); !errors.Is(err, ErrInvalidInterface) {
			t.Fatalf("want err: %s, but have: %s", ErrInvalidInterface, err)
		}
	})

	t.Run("success slice struct", func(t *testing.T) {
		data := []byte("Test\na\n")
		out := []A{}
		r := bytes.NewReader(data)
		if err := Decode(&out, r); err != nil {
			t.Fatal(err)
		}

		if out[0].Test != "a" {
			t.Fatalf("out.Test must be a")
		}
	})

	t.Run("success slice pointer", func(t *testing.T) {
		data := []byte("Test,Test2\na,b\n")
		out := []*A{}
		r := bytes.NewReader(data)
		if err := Decode(&out, r); err != nil {
			t.Fatal(err)
		}

		if out[0].Test != "a" {
			t.Fatalf("out.Test must be a")
		}
		if *out[0].Test2 != "b" {
			t.Fatalf("out.Test must be a")
		}
	})
}
