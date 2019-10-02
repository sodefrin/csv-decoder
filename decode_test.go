package csvdecoder

import (
	"bytes"
	"errors"
	"testing"

	"cloud.google.com/go/civil"
)

func TestDecode(t *testing.T) {
	type A struct {
		Test  string
		Test2 *string
		Test3 civil.Date `json:"test3"`
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
		data := []byte("Test,Test2,test3\na,b,2019-09-02\n")
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
		if out[0].Test3.String() != "2019-09-02" {
			t.Fatalf("out.Test must be c")
		}
	})
}
