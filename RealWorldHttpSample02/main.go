package main

import (
	"RealWorldHttpSample02/HttpGet"
	"flag"
	"fmt"
)

// SimpleHttp シンプル HTTP GET関数
func SimpleHttp() {

	//接続先URL
	url := "https://www.yahoo.co.jp/"

	//シンプルHTTP GET実行
	ret := HttpGet.SimpleHttpGetMain(url)

	if ret <= 0 {
		//関数内でエラー発生時の処理
		println("Error!")
		return
	}

	println("Success.")

}

// QueryHttp クエリー付き HTTP GET関数
func QueryHttp() {

	//接続先URL
	url := "https://www.google.co.jp/maps"

	//クエリー定義
	query := "hl=ja"

	ret := HttpGet.QueryHttpGetMain(url, query)

	if ret < 0 {
		//関数内でエラー発生時の処理
		println("Error!")
		return
	}

	println("Success.")

}

// startMain 処理開始関数
func startMain() {

	var (
		f = flag.Int("f", -1, "Process Flag: 1=SimpleHTTPGet, 2:QueryHTTPGet")
	)

	//コマンドライン引数を解析
	flag.Parse()

	//for debug
	fmt.Printf("param -f : %d\n", *f)

	//引数によって処理を分岐
	switch *f {
	case 1:
		SimpleHttp()
	case 2:
		QueryHttp()
	default:
		println("Parameter error. InParam:%d", *f)
	}

}

//エントリポイント関数
func main() {

	startMain()

}
