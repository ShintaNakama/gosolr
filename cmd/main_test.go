package main_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/ShintaNakama/gosolr"
	main "github.com/ShintaNakama/gosolr/cmd"
)

func TestPassArgs(t *testing.T) {
	expected := map[string]string{"c": "group", "a": "select", "r": "none"}
	expectedMock := map[string]bool{"m": false}
	c, a, r, q, fq, sort, st, row, fl, wt, indent, mock, err := main.PassArgs()

	if err != nil {
		t.Errorf("expected:%v, err:%v", nil, err)
	}
	if c != expected["c"] {
		t.Errorf("expected:%v, core:%v", expected["c"], c)
	}
	if a != expected["a"] {
		t.Errorf("expected:%v, action:%v", expected["a"], a)
	}
	if r != expected["r"] {
		t.Errorf("expected:%v, recommend:%v", expected["r"], r)
	}
	if mock != expectedMock["m"] {
		t.Errorf("expected:%v, recommend:%v", expectedMock["m"], mock)
	}
}

type exampleArgs struct {
	core      string
	action    string
	recommend string
	mock      bool
}

func TestDefaultSearchExec(t *testing.T) {
	var rs gosolr.Result
	// mockファイルが増えたら、ReadDirにする
	raw, err := ioutil.ReadFile(filepath.FromSlash("../testdata/group.json"))
	if err != nil {
		t.Errorf("expected:%v, err:%v", nil, err)
	}
	json.Unmarshal(raw, &rs)
	expected := &rs

	examples := []exampleArgs{
		{core: "group", action: "select", recommend: "none", mock: true},
	}

	ctx := context.Background()
	for _, e := range examples {
		cli := gosolr.NewClient(e.core, e.action, e.recommend, e.mock)
		result, err := main.DefaultSearchExec(ctx, cli)
		if err != nil {
			t.Errorf("expected:%v, err:%v", nil, err)
		}
		if result.Response.Docs[0] != expected.Response.Docs[0] {
			t.Errorf("expected:%v\n, result:%v", expected.Response.Docs[0], result.Response.Docs[0])
		}
	}

}
