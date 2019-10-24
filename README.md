# gosolr
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
  # solrのhost-urlを環境変数で設定
  ex) export SOLRHOST=http://solr-host-url/solr
  # デフォルトのパラメータで検索
  ./gosolr
  # 検索クエリ、取得するフィールドを設定
  ./gosolr -q=example_field:hoge, -fl=example_field
  ```
  - help
  ```
  ./gosolr --help

  # 特定の1つのfieldを環境変数で登録し、検索できるようにする
  ex) export FINDFIELD=hoge
  ./gosolr -find=fuga
  # -fq=hoge:fuga で検索する

  # Numfoundの値だけ取得する
  ./gosolr --count
  ```
