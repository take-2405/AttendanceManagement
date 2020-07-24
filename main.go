package main

import (
	"flag"
	"AttendanceManagement/pkg"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	pkg.Server.Run(addr)
}