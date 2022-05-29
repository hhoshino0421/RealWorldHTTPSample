package HttpGet

import (
	"io/ioutil"
	"log"
	"net/http"
)

// SimpleHttpGetMain シンプル HTTP GET関数
func SimpleHttpGetMain(url string) {

	//指定されたURLへのGETメソッド処理実行
	resp, err := http.Get(url)

	if err != nil {
		//エラー処理
		print("Get error!")
		return
	}

	//この関数を抜けた後の処理を定義する
	//ソケットからBODYデータを読み込んだ後にクローズ処理を実行する
	defer resp.Body.Close()

	//BODYデータをソケットからバイト列として読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//エラー処理
		print("Body data read error!")
		return
	}

	//バイト列をUTF-8文字列に変換して出力
	log.Println(string(body))

}
