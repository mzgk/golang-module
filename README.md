# Go Modules
- これまでのパッケージ管理は$GOPATHを利用した仕組み
  - $GOPATH/srcにインポートしたパッケージが格納される
- Moduleは、$GOPATHに代わってパッケージ管理する仕組み
  - ver1.13からのデフォルト機能
  - 必要なパッケージを自動でDLしてくれる
  - デフォルトでDLされるバージョンは同一バージョン内の最新版
  - 指定する場合は、importに/v2のようにバージョンを記述する
- 作成するモジュールの直下で、以下を実施
  - `$ go mod init <モジュール名>`
  - 2ファイルが作成される
    - go.mod: 依存するパッケージが追加される
    - go.sum: ロックファイル
- あとは、通常同様にコーディングして保存（もしくはBuild）
  - importに記述したパッケージが自動でDLされる
  - DL先は、$GOPATH/pkg/mod
  - GOPATHの設定がなければ、HOME/goがGOPATHになる
- go getは不要
- 参考
  - https://qiita.com/pokeh/items/139d0f1fe56e358ba597
  - http://www.songmu.jp/riji/entry/2019-03-28-go-modules.html
  - https://techblog.asahi-net.co.jp/entry/2019/03/11/122341


## 手順
```bash
$ go version
go version go1.12.4 windows/amd64

$ cd <任意の場所>/newmod
$ go mod init newmod

# あとはコーディング
# 自作のlibパッケージはnewmod/libに作成
# lib内でimportしている"golang.org/x/text/width"は、$GOPATH/src/pkg/mod/にDLされる

$ go run main.go
```