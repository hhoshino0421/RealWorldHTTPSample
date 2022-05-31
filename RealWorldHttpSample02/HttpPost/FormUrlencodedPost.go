package HttpPost

import (
	"log"
	"net/http"
	"net/url"
)

// FormUrlencodedPost FORM URLEncoded形式のPOST送信処理
func FormUrlencodedPost(inUrl string) int {

	//POST用のパラメータ定義
	values := url.Values{
		"test": {"mewmew"},
		"Kuma": {"Yuna", "Bear"},
	}

	/*
		以下のような形で受信される

		POST / HTTP/1.1
		Host: localhost:18888
		Accept-Encoding: gzip
		Content-Length: 31
		Content-Type: application/x-www-form-urlencoded
		User-Agent: Go-http-client/1.1

		Kuma=Yuna&Kuma=Bear&test=mewmew
	*/

	//POST送信処理
	resp, err := http.PostForm(inUrl, values)

	if err != nil {
		//エラー処理
		log.Println("Post(FormUrlencoded) error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	log.Println("Status: ", resp.Status)

	//正常終了
	return 0

}
