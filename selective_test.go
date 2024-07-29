package selective_json_go

import (
	"strings"
	"testing"
)

func TestSelectiveMarshal(t *testing.T) {
	s1 := NewSelectiveMarshaller("s1")
	s2 := NewSelectiveMarshaller("s2")
	s3 := NewSelectiveMarshaller("s3")
	s4 := NewSelectiveMarshaller("")
	// normal case
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

	b, err := s1.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	str := string(b)
	if !strings.Contains(str, "f1") || // f1 should be always present
		strings.Contains(str, "f2") || // f2 should never be present
		!strings.Contains(str, "f3") || // f3 only present in s1
		!strings.Contains(str, "f4") || // f4 present in s1 and s2
		strings.Contains(str, "f5") || // f5 only present in s2
		!strings.Contains(str, "f6") || // f6 should not be omitted
		!strings.Contains(str, "f7") || // f7 only present in s1
		strings.Contains(str, "f8") { // f8 should be omitted
		t.Errorf("unexpected output: %s", str)
	}

	b, err = s2.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	str = string(b)
	if !strings.Contains(str, "f1") || // f1 should be always present
		strings.Contains(str, "f2") || // f2 should never be present
		strings.Contains(str, "f3") || // f3 only present in s1
		!strings.Contains(str, "f4") || // f4 present in s1 and s2
		!strings.Contains(str, "f5") || // f5 only present in s2
		!strings.Contains(str, "f6") || // f6 should not be omitted
		strings.Contains(str, "f7") || // f7 only present in s1
		strings.Contains(str, "f8") { // f8 should be omitted
		t.Errorf("unexpected output: %s", str)
	}

	b, err = s3.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	str = string(b)
	if !strings.Contains(str, "f1") || // f1 should be always present
		strings.Contains(str, "f2") || // f2 should never be present
		strings.Contains(str, "f3") || // f3 only present in s1
		strings.Contains(str, "f4") || // f4 present in s1 and s2
		strings.Contains(str, "f5") || // f5 only present in s2
		!strings.Contains(str, "f6") || // f6 should not be omitted
		strings.Contains(str, "f7") || // f7 only present in s1
		strings.Contains(str, "f8") { // f8 should be omitted
		t.Errorf("unexpected output: %s", str)
	}

	b, err = s4.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	str = string(b)
	if !strings.Contains(str, "f1") || // f1 should be always present
		strings.Contains(str, "f2") || // f2 should never be present
		!strings.Contains(str, "f3") || // f3 should be present since scenraio is not enabled
		!strings.Contains(str, "f4") || // f4 should be present since scenraio is not enabled
		!strings.Contains(str, "f5") || // f5 should be present since scenraio is not enabled
		!strings.Contains(str, "f6") || // f6 should not be omitted
		!strings.Contains(str, "f7") || // f7 should be present since scenraio is not enabled
		strings.Contains(str, "f8") { // f8 should be omitted
		t.Errorf("unexpected output: %s", str)
	}
}
