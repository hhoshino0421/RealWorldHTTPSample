package HttpGet

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func CookieGet(inUrl string) int {

	cookieObj, err := cookiejar.New(nil)

	if err != nil {
		//エラー処理
		log.Println("cookiejar new error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	clientObj := http.Client{
		Jar: cookieObj,
	}

	for i := 0; i < 2; i++ {
		
		resp, err := clientObj.Get(inUrl)

		if err != nil {
			//エラー処理
			log.Println("clientObj.Get error!")
			log.Println(err)
			//異常終了を通知
			return -1
		}

		dumpObj, err := httputil.DumpResponse(resp, true)

		if err != nil {
			//エラー処理
			log.Println("DumpResponse error!")
			log.Println(err)
			//異常終了を通知
			return -1
		}

		log.Println("dump: ", string(dumpObj))

	}

	//正常終了
	return 0
}
