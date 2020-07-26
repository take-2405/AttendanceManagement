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
ターミナルでプログラムを実行する前に以下のコマンドを入力する
 ```
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=rootpassword \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=time_management_app
```

## APIローカル起動方法
```
$ go run main.go
```