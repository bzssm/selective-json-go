# Selective-json-go

Marshal json with selective fields.

## Examples
```go
a := struct {
	F1 string `json:"f1"`
	F2 string `json:"-"`
	F3 string `json:"f3,scenarios:s1"`
	F4 string `json:"f4,scenarios:s1|s2"`
	F5 string `json:"f5,scenarios:s2"`
	F6 string `json:",omitempty"`
	F7 string `json:",scenarios:s1"`
	F8 string `json:",omitempty"`
}{
	F1: "f1",
	F2: "f2",
	F3: "f3",
	F4: "f4",
	F5: "f5",
	F6: "f6",
	F7: "f7",
}
s1 := NewSelectiveMarshaller("s1")
s1.Marshal(a)
// output:
// {"f1":"f1","f3":"f3","f4":"f4","f6":"f6","f7":"f7"}
```