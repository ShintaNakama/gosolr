package gosolr

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/pkg/errors"
)

// Client is a struct for accessing solr
type Client struct {
	DefaultSearch DefaultSearch
	HTTPClient    *http.Client
	BaseURL       string
	Params        *Params
	Mock          bool
}

// NewClient creates a client struct
func NewClient(core, action, r, q, fq, sort, st, row, fl, wt, indent string, mock bool) *Client {
	var cli Client
	cli.DefaultSearch = &defaultSearch{cli: &cli}
	cli.BaseURL = "/" + core + "/" + action
	cli.Params = &Params{q: q, fq: fq, sort: sort, start: st, rows: row, fl: fl, wt: wt, indent: indent}
	cli.Mock = mock
	return &cli
}

func (cli *Client) httpClient() *http.Client {
	if cli.HTTPClient != nil {
		return cli.HTTPClient
	}
	return http.DefaultClient
}

func (cli *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	httpClient := cli.httpClient()
	// mock on の場合
	if cli.Mock {
		httpClient.Transport = newMockTransport()
	}
	// return cli.httpClient().Do(req)
	return httpClient.Do(req)
}

// 空インターフェースを使って、複数の型にsolrのresponseを当てはめれるようにする
func (cli *Client) get(ctx context.Context, params url.Values, v interface{}) error {
	reqURL := cli.BaseURL
	if params != nil {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return errors.Wrap(err, "cannot create HTTP request")
	}

	resp, err := cli.do(ctx, req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// 300 > StatusCode >= 200
	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return cli.error(resp.StatusCode, resp.Body)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return errors.Wrap(err, "cannot parse HTTP body")
	}

	return nil
}

func (cli *Client) error(statusCode int, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil || len(buf) == 0 {
		return errors.Errorf("request failed with status code %d", statusCode)
	}
	return errors.Errorf("StatusCode: %d, Error: %s", statusCode, string(buf))
}

type mockTransport struct{}

func newMockTransport() http.RoundTripper {
	return &mockTransport{}
}

// Implement http.RoundTripper
func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")

	// mockファイルが増えたら、ReadDirにする
	data, err := ioutil.ReadFile(filepath.FromSlash("../testdata/group.json"))
	if err != nil {
		return nil, err
	}
	responseBody := *(*string)(unsafe.Pointer(&data))
	// fmt.Println("debug")
	// fmt.Println(responseBody)
	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}
