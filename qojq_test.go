package jsonql

import (
	"testing"
)

var jsonArray = `
[
  {
    "name": "elgs",
    "gender": "m",
    "skills": [
      "Golang",
      "Java",
      "C"
    ]
  },
  {
    "name": "enny",
    "gender": "f",
    "skills": [
      "IC",
      "Electric design",
      "Verification"
    ]
  },
  {
    "name": "sam",
    "gender": "m",
    "skills": [
      "Eating",
      "Sleeping",
      "Crawling"
    ],
	"hello": null,
	"hello_world":true
  }
]
`

func TestParseJsonArray(t *testing.T) {

	parserArray, err := NewJSONStringQuery(jsonArray)
	if err != nil {
		t.Error(err)
	}

	var pass = []struct {
		in string
		ex interface{}
	}{
		{"[0].name", "elgs"},
		{"[1].gender", "f"},
		{"[2].skills.[1]", "Sleeping"},
		{"[2].hello", nil},
		{"[2].hello_world", true},
	}
	var fail = []struct {
		in string
		ex interface{}
	}{}
	for _, v := range pass {
		result, err := parserArray.Query(v.in)
		if err != nil {
			t.Error(err)
		}
		if v.ex != result {
			t.Error("Expected:", v.ex, "actual:", result)
		}
	}
	for range fail {

	}
}

var jsonObj = `
{
  "name": "sam",
  "gender": "m",
  "skills": [
    "Eating",
    "Sleeping",
    "Crawling"
  ],
  "hello":null
}
`

func TestParseJsonObj(t *testing.T) {

	parserObj, err := NewJSONStringQuery(jsonObj)
	if err != nil {
		t.Error(err)
	}

	var pass = []struct {
		in string
		ex interface{}
	}{
		{"name", "sam"},
		{"gender", "m"},
		{"skills.[1]", "Sleeping"},
		{"hello", nil},
	}
	var fail = []struct {
		in string
		ex interface{}
	}{}
	for _, v := range pass {
		result, err := parserObj.Query(v.in)
		if err != nil {
			t.Error(err)
		}
		if v.ex != result {
			t.Error("Expected:", v.ex, "actual:", result)
		}
	}
	for range fail {

	}
}
