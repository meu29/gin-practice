プロジェクト作成の手順<br>
①go mod init go_app<br>
②go get github.com/gin-gonic/gin github.com/gin-contrib/sessions github.com/gin-contrib/sessions/cookie github.com/jinzhu/gorm github.com/mattn/go-sqlite3<br>
③go run main.go<br>
自分が作成したパッケージをインポートする場合は(go.modに書かれたmodule名)/(パッケージ名)<br>