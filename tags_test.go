// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package selective_json_go

import (
	"reflect"
	"testing"
)

func TestTagParsing(t *testing.T) {
	name, opts := parseTag("field,foobar,foo")
	if name != "field" {
		t.Fatalf("name = %q, want field", name)
	}
	for _, tt := range []struct {
		opt  string
		want bool
	}{
		{"foobar", true},
		{"foo", true},
		{"bar", false},
	} {
		if opts.Contains(tt.opt) != tt.want {
			t.Errorf("Contains(%q) = %v", tt.opt, !tt.want)
		}
	}
}

func TestScenarioTagParsing(t *testing.T) {
	name, opts := parseTag("field,foobar,foo,scenarios:apiA|apiB")
	if name != "field" {
		t.Fatalf("name = %q, want field", name)
	}
	for _, tt := range []struct {
		opt  string
		want bool
	}{
		{"foobar", true},
		{"foo", true},
		{"bar", false},
		{"scenarios:apiA|apiB", true},
	} {
		if opts.Contains(tt.opt) != tt.want {
			t.Errorf("Contains(%q) = %v", tt.opt, !tt.want)
		}
	}
}

func TestParseScenarios(t *testing.T) {
	_, opts := parseTag("field,foobar,foo,scenarios:apiA|apiB")
	opts.Scenarios()

	for _, tt := range []struct {
		opt  string
		want []string
	}{
		{"", []string{}},
		{"scenarios:", []string{}},
		{"scenarios:apiA|apiB", []string{"apiA", "apiB"}},
		{"scenarios:apiA|apiB|apiC", []string{"apiA", "apiB", "apiC"}},
	} {
		if reflect.DeepEqual(tagOptions(tt.opt).Scenarios(), tt.want) != true {
			t.Errorf("Contains(%q) = %v", tagOptions(tt.opt).Scenarios(), tt.want)
		}
	}
}
