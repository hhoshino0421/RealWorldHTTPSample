package HttpPost

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// MultipartPost Multipart/form-data形式でのPOST処理
func MultipartPost(inUrl string) int {

	var byteBuffer bytes.Buffer
	writerObj := multipart.NewWriter(&byteBuffer)
	writerObj.WriteField("name", "Mewmew")

	fileWriter, err := writerObj.CreateFormFile("thumbnail", "DSCN3077.JPG")

	if err != nil {
		//エラー処理
		log.Println("CreateFormFile error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	readFile, err := os.Open("/Users/hhoshino/develop/Goland_workspace/RealWorldHttpSample02/DSCN3077.JPG")

	if err != nil {
		//エラー処理
		log.Println("File open error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	defer readFile.Close()

	io.Copy(fileWriter, readFile)

	writerObj.Close()

	resp, err := http.Post(inUrl, writerObj.FormDataContentType(), &byteBuffer)

	if err != nil {
		//エラー処理
		log.Println("Post(multipart/form-data) error!")
		log.Println(err)
		//異常終了を通知
		return -1
	}

	log.Println("Status: ", resp.Status)

	/*
		こんな感じで送信される
		2022/05/31 12:49:58 start http listening : 18888

		POST / HTTP/1.1
		Host: localhost:18888
		Accept-Encoding: gzip
		Content-Length: 357937
		Content-Type: multipart/form-data; boundary=501d706eed9a6c01217c477b0203addd11d19eae5a023d0fc06b49d2d555
		User-Agent: Go-http-client/1.1

		--501d706eed9a6c01217c477b0203addd11d19eae5a023d0fc06b49d2d555
		Content-Disposition: form-data; name="name"

		Mewmew
		--501d706eed9a6c01217c477b0203addd11d19eae5a023d0fc06b49d2d555
		Content-Disposition: form-data; name="thumbnail"; filename="DSCN3077.JPG"
		Content-Type: application/octet-stream

		????]ExifII*
		??(1?2?;%??74i?l%???          NIKONCOOLPIX A900,,COOLPIX A900   V1.12020:01:17 13:15:04$??"??*"?'?P0??0230?2?F??Z?
		b?j?	?
		?r|??,<??,z?0100??@?h
		                     ?????????#??	?
		?
		 ?
		(
		2020:01:17 13:15:042020:01:17 13:15:04
		#
		>
		?SCIId46,,R980100NikonII'0200??

	*/

	//正常終了
	return 0
}
