package gosolr

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

// DefaultSearch interface performs a simple search.
type DefaultSearch interface {
	GetData(ctx context.Context) (*Result, error)
}

type defaultSearch struct {
	cli *Client
}

func setQueryParams(d *defaultSearch) url.Values {
	v := url.Values{}
	v.Add("q", d.cli.Params.q)
	v.Add("fq", d.cli.Params.fq)
	v.Add("sort", d.cli.Params.sort)
	v.Add("start", d.cli.Params.start)
	v.Add("rows", d.cli.Params.rows)
	v.Add("fl", d.cli.Params.fl)
	v.Add("wt", d.cli.Params.wt)
	v.Add("indent", d.cli.Params.indent)
	return v
}

func (d *defaultSearch) GetData(ctx context.Context) (*Result, error) {
	var rs Result
	v := setQueryParams(d)
	err := d.cli.get(ctx, v, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "GET solr failed")
	}
	return &rs, nil
}
