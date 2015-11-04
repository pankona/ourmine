# これは何をするかと言うと

* 指定したプロジェクト内で閉じられていないチケットを一覧表示します。 

# Install

* `go get github.com/pankona/ourmine`

# Usage

* 以下の環境変数を設定します。

  * REDMINE_URL
    * redmine へのURL

  * REDMINE_API_KEY
    * redmine の REST API を利用するための API KEY

  * 以下のコマンドを実行すると、チケット一覧が表示されます。

    * `$ ourmine`

  * `-o {ticket num}` のオプションをつけると、指定されたチケットをブラウザで開きます。

    * `$ ourmine -o 12345`

# License

* MIT

# Contribution

* Any contribution is welcome!
