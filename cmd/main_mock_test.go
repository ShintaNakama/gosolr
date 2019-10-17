package main_test

import (
	"context"

	"github.com/ShintaNakama/gosolr"
)

type MockDefaultSearch struct {
	gosolr.DefaultSearch
	MockGetData func(ctx context.Context) (*gosolr.Result, error)
}

func (d *MockDefaultSearch) GetData(ctx context.Context) (*gosolr.Result, error) {
	return d.MockGetData(ctx)
}
