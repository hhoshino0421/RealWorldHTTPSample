package main

import (
	"RealWorldHttpSample02/HttpGet"
)

func main() {

	//接続先URL
	url := "https://www.yahoo.co.jp/"

	//シンプルHTTP GET実行
	ret := HttpGet.SimpleHttpGetMain(url)

	if ret < 0 {
		//関数内でエラー発生時の処理
		println("Error!")
		return
	}

	println("Success.")

}
