package HttpPost

import (
	"log"
	"net/http"
	"os"
)

func FilePost(inUrl string) int {

	//送信するファイル定義
	fileName := "/Users/hhoshino/develop/Goland_workspace/RealWorldHttpSample02/main.go"

	//ファイルオープン処理
	fileObj, err := os.Open(fileName)

	if err != nil {
		log.Println("File open error!")
		log.Println(err)
		//異常終了
		return -1
	}

	//POST送信処理
	resp, err := http.Post(inUrl, "text/plane", fileObj)

	if err != nil {
		log.Println("Post error!")
		log.Println(err)
		//異常終了
		return -1
	}

	log.Println("Status: ", resp.Status)

	/*
		こんな感じのヘッダーで送信される
		POST / HTTP/1.1
		Host: localhost:18888
		Transfer-Encoding: chunked
		Accept-Encoding: gzip
		Content-Type: text/plane
		User-Agent: Go-http-client/1.1
	*/

	//正常終了
	return 0
}
