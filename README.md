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
上記の「ターミナルでプログラムを実行する前に以下のコマンドを入力する」の部分のコマンドを端末にて実行  
DBの作成 docker-compose up -d (Dockerのinstallが必要)
$ go run main.go  
```