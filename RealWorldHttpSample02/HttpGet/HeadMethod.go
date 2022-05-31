package HttpGet

import (
	"log"
	"net/http"
)

func HeadMethodMain(url string) int {

	resp, err := http.Head(url)

	if err != nil {
		//エラー処理
		print("HEAD error!")
		//異常終了を通知
		return -1
	}

	log.Println("Status: ", resp.Status)
	log.Println("StatusCode: ", resp.StatusCode)
	log.Println("Headers: ", resp.Header)

	//正常終了
	return 0

}
