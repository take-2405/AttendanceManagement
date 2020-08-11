# AttendanceManagement
退勤管理システムを作りたい

### MySQLWorkbenchの設定
MySQLへの接続設定をします。
1. MySQL Connections の + を選択
2. 以下のように接続設定を行う
 ```
 Connection Name: 任意
 Connection Method: Standard (TCP/IP)
 Hostname: 配布されたサーバのグローバルIP
 Port: 3306
 Username: 
 Password: 
 Default Schema: 
 ```

### API用のデータベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。<br>
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

Macの場合
```
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=ca-tech-dojo-go-2020 \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=dojo_api
```

Windowsの場合
```
$ SET MYSQL_USER=root
$ SET MYSQL_PASSWORD=ca-tech-dojo-go-2020
$ SET MYSQL_HOST=127.0.0.1
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=dojo_api
```

## APIローカル起動方法
```
上記の「ターミナルでプログラムを実行する前に以下のコマンドを入力する」の部分のコマンドを端末にて実行  
DBの作成 docker-compose up -d (Dockerのinstallが必要)
$ go run main.go  
```
