# csv-decoder

parse csv like below.

```
    data := []byte("Test\na\n")
    out := []A{}
    r := bytes.NewReader(data)
    if err := Decode(&out, r); err != nil {
        t.Fatal(err)
    }

    if out[0].Test != "a" {
        t.Fatalf("out.Test must be a")
    }
```
