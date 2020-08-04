package csvdecoder

import (
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
		out := A{}
		if err := Decode("testdata/empty.csv", &out); !errors.Is(err, ErrInvalidInterface) {
			t.Fatalf("want err: %s, but have: %s", ErrInvalidInterface, err)
		}
	})

	t.Run("invliad not pointer", func(t *testing.T) {
		out := []A{}
		if err := Decode("testdata/empty.csv", out); !errors.Is(err, ErrInvalidInterface) {
			t.Fatalf("want err: %s, but have: %s", ErrInvalidInterface, err)
		}
	})

	t.Run("success slice struct", func(t *testing.T) {
		out := []A{}
		if err := Decode("testdata/test.csv", &out); err != nil {
			t.Fatal(err)
		}

		if out[0].Test != "a" {
			t.Fatalf("out.Test must be a but have %s", out[0].Test)
		}
	})

	t.Run("success slice pointer", func(t *testing.T) {
		out := []*A{}
		if err := Decode("testdata/test2.csv", &out); err != nil {
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
