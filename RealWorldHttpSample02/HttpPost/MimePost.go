package HttpPost

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

// MimePost MIME付きでMultipart/form-dataでPOST処理
func MimePost(inUrl string) int {

	var byteBuffer bytes.Buffer
	writerObj := multipart.NewWriter(&byteBuffer)
	writerObj.WriteField("name", "Mewmew")

	partObj := make(textproto.MIMEHeader)
	partObj.Set("Content-type", "image/jpeg")
	partObj.Set("Content-Disposition", `form-data; name="thumbnail"; filename="DSCN3077.JPG"`)

	fileWriterObj, err := writerObj.CreatePart(partObj)

	if err != nil {
		//エラー処理
		log.Println("CreatePart error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	readFileObj, err := os.Open("/Users/hhoshino/develop/Goland_workspace/RealWorldHttpSample02/DSCN3077.JPG")

	if err != nil {
		//エラー処理
		log.Println("File open error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	defer readFileObj.Close()

	io.Copy(fileWriterObj, readFileObj)

	writerObj.Close()

	resp, err := http.Post(inUrl, writerObj.FormDataContentType(), &byteBuffer)

	if err != nil {
		//エラー処理
		log.Println("Post(multipart/form-data[Mime]) error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	log.Println("Status: ", resp.Status)

	/*
		こんな感じで送信される
		POST / HTTP/1.1
		Host: localhost:18888
		Accept-Encoding: gzip
		Content-Length: 357923
		Content-Type: multipart/form-data; boundary=c89735ce918c7b09c7271250d0f17d7b682f52052685940087bcec13efdc
		User-Agent: Go-http-client/1.1

		--c89735ce918c7b09c7271250d0f17d7b682f52052685940087bcec13efdc
		Content-Disposition: form-data; name="name"

		Mewmew
		--c89735ce918c7b09c7271250d0f17d7b682f52052685940087bcec13efdc
		Content-Disposition: form-data; name="thumbnail"; filename="DSCN3077.JPG"
		Content-Type: image/jpeg

		????]ExifII*
		??(1?2?;%??74i?l%???          NIKONCOOLPIX A900,,COOLPIX A900   V1.12020:01:17 13:15:04$??"??*"?'?P0??0230?2?F??Z?
		b?j?	?
		?r|??,<??,z?0100??@?h
	*/

	//正常終了
	return 0

}
