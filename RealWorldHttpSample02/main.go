package main

import (
	"RealWorldHttpSample02/HttpGet"
)

func main() {

	//接続先URL
	url := "https://www.yahoo.co.jp/"

	//シンプルHTTP GET実行
	HttpGet.SimpleHttpGetMain(url)

}
