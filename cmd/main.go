package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ShintaNakama/gosolr"
)

// passArgs accepts multiple arguments and returns their values.
func passArgs() (host, c, a, r, q, fq, sort, st, row, fl, wt, indent string, mock bool, err error) {
	envSolrHost := os.Getenv("SOLRHOST")
	if envSolrHost == "" {
		err = errors.New("SOLRHOSTが設定されていません")
		return
	}
	flag.StringVar(&host, "host", envSolrHost, "hostのURL")
	flag.StringVar(&c, "c", "group", "core: 対象のsolr-core")
	flag.StringVar(&a, "a", "select", "action: 実行するアクション。原則 select")
	flag.StringVar(&r, "recommend", "", "recommend-type")
	flag.StringVar(&q, "q", "*:*", "q: 検索クエリ ex) -q=client_code:r")
	flag.StringVar(&fq, "fq", "", "fq: サブ検索クエリ ex) -fq=client_code:n")
	flag.StringVar(&sort, "sort", "", "sort: ソート順 ex) -sort=asc")
	flag.StringVar(&st, "st", "0", "start: 開始位置 ex) -start=10")
	flag.StringVar(&row, "rows", "1", "rows: 取得数 ex) -rows=10")
	flag.StringVar(&fl, "fl", "", "fl: 取得するフィールド ex) -fl=dwelling_name")
	flag.StringVar(&wt, "wt", "json", "wt: レスポンス形式(json、xml、python、ruby、php、csv) ex) -wt=json")
	flag.StringVar(&indent, "indent", "true", "indent: インデント ex) -indent=true")

	// mock
	flag.BoolVar(&mock, "mock", false, "mockを使ってresponseを生成するかどうか (default false)")

	if flag.Parse(); flag.Parsed() {
		return
	}
	err = errors.New("引数のparseに失敗しました。")
	return
}

func main() {
	host, c, a, r, q, fq, sort, st, row, fl, wt, indent, mock, err := passArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli := gosolr.NewClient(host, c, a, r, q, fq, sort, st, row, fl, wt, indent, mock)
	ctx := context.Background()

	result, err := defaultSearchExec(ctx, cli)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	re, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, string(re))
}

func defaultSearchExec(ctx context.Context, cli *gosolr.Client) (*gosolr.Result, error) {
	rs, err := cli.DefaultSearch.GetData(ctx)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
