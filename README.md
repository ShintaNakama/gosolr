# gozsolr
- CLI
  - Apache Solr にCLIでアクセスし、検索結果を標準出力で表示する
  - Solr は index が設定されている前提。
  - build
  ```
  cd /cmd
  go build -o gozsolr
  ```
  - example
  ```
  # デフォルトのパラメータで検索
  ./gosolr
  # 検索クエリ、取得するフィールドを設定
  ./gosolr -q=example_field:hoge, -fl=example_field
  ```
  - help
  ```
  ./gosolr --help
  ```


  
# gosolr
