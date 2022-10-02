package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"net/http"
	"os"
)

var clientID = ""
var clientSecret = ""
var redirectURL = "http://localhost:18888"
var state = "your state"

func apiTest() {

	//デバッグ用
	fmt.Println("Hello Mewmew!")

	//OAuth2の接続先などの情報
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"user:email", "gist"},
		Endpoint:     github.Endpoint,
	}
	//これをこれから初期化する
	var token *oauth2.Token

	//ローカルにて既に保存済？
	file, err := os.Open("access_token.json")
	if os.IsNotExist(err) {
		//初回アクセス
		//まず許可画面のURLを取得
		url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

		//コールバックを受け取るWebサーバをセットアップ

		code := make(chan string)
		var server *http.Server

		//ハンドラ関数を定義
		h1 := func(w http.ResponseWriter, r *http.Request) {
			//クエリパラメータからcodeを取得しブラウザを閉じる
			w.Header().Set("Content-Type", "text/html")
			io.WriteString(w, "<html><script>window.open('about:blank','_self').close()</script></html>")
			w.(http.Flusher).Flush()
			code <- r.URL.Query().Get("code")
			//サーバも閉じる
			server.Shutdown(context.Background())
		}

		server = &http.Server{
			Addr:    ":1888",
			Handler: http.HandlerFunc(h1),
		}
		go server.ListenAndServe()

		//ブラウザで認可画面を開く
		//GitHubの認可が完了すれば上記のサーバにリダイレクトされてHandlerが実行される
		open.Start(url)

		//同期したコードをアクセストークンに交換
		token, err = conf.Exchange(oauth2.NoContext, <-code)
		if err != nil {
			//異常終了
			panic(err)
		}

		//アクセストークンをファイルに保存
		file, err = os.Create("access_token.json")
		if err != nil {
			//異常終了
			panic(err)
		}
		json.NewEncoder(file).Encode(token)

	} else if err == nil {
		//一度認可をしてローカルに保存済
		token = &oauth2.Token{}
		json.NewDecoder(file).Decode(token)

	} else {
		//異常終了
		panic(err)
	}

	client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))

	//ここでさまざまな処理を実行する

	client.Get("https://github.com/")

}
