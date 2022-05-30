package HttpGet

import (
	"io/ioutil"
	"log"
	"net/http"
)

// QueryHttpGetMain クエリー付き HTTP Get関数
func QueryHttpGetMain(url string, query string) int {

	urlQuery := url + "?" + query

	//指定されたURLへのGETメソッド処理実行
	resp, err := http.Get(urlQuery)

	if err != nil {
		//エラー処理
		print("Get error!")
		//異常終了を通知
		return -1
	}

	//この関数を抜けた後の処理を定義する
	//ソケットからBODYデータを読み込んだ後にクローズ処理を実行する
	//Pythonのwithのようなコードを作れる
	defer resp.Body.Close()

	//BODYデータをソケットからバイト列として読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//エラー処理
		print("Body data read error!")
		//異常終了を通知
		return -1
	}

	//ステータスを出力
	log.Println("Status:", resp.Status)
	//ステータスコードを出力
	log.Println("StatusCode:", resp.StatusCode)

	//ヘッダを出力
	log.Println("Header:", resp.Header)

	//BODYのバイト列をUTF-8文字列に変換して出力
	log.Println("body:", string(body))

	//正常終了
	return 0

}
