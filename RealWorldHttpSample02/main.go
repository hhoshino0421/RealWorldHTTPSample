package main

import (
	"RealWorldHttpSample02/HttpGet"
	"RealWorldHttpSample02/HttpPost"
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

// HeadMethod HEADメソッド実行関数
func HeadMethod() {

	url := "https://www.google.co.jp/"

	ret := HttpGet.HeadMethodMain(url)

	if ret < 0 {
		//関数内でエラー発生時の処理
		println("Error!")
		return
	}

	println("Success.")

}

// FormUrlencodedPostMain FORM-URLENCODED形式のPOST処理関数
func FormUrlencodedPostMain() {

	url := "http://localhost:18888/"

	ret := HttpPost.FormUrlencodedPost(url)

	if ret < 0 {
		//関数内でエラー発生時の処理
		println("Error!")
		return
	}

	println("Success.")

}

// FilePostMain ファイル送信のPOST処理
func FilePostMain() {

	url := "http://localhost:18888/"

	ret := HttpPost.FilePost(url)

	if ret < 0 {
		//関数内でエラー発生時の処理
		println("Error!")
		return
	}

	println("Success.")

}

// MultipartPostMain Multipart/form-data形式でのPOST処理
func MultipartPostMain() {

	url := "http://localhost:18888/"

	ret := HttpPost.MultipartPost(url)

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
		f = flag.Int("f", -1, "Process Flag: 1=SimpleHTTPGet, 2:QueryHTTPGet, 3:HttpHead, 4:FormUrlencodedPost, 5:FilePost, 6:MultipartPost")
	)

	//コマンドライン引数を解析
	flag.Parse()

	//for debug
	fmt.Printf("param -f : %d\n", *f)

	//引数によって処理を分岐
	switch *f {
	case 1:
		//シンプルなHTTP GET処理
		SimpleHttp()
	case 2:
		//クエリ付きHTTP GET処理
		QueryHttp()
	case 3:
		//HTTP HEAD処理
		HeadMethod()
	case 4:
		//FORM-URLENCODED形式のPOST処理
		FormUrlencodedPostMain()
	case 5:
		//FileのPOST処理
		FilePostMain()
	case 6:
		//Multipart-form-dataのPOST処理
		MultipartPostMain()
	default:
		println("Parameter error. InParam:%d", *f)
	}

}

//エントリポイント関数
func main() {

	startMain()

}
