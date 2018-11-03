package get

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Errorhandle(chara string, s *goquery.Selection) {
	s.Find("div").Each(func(_ int, t *goquery.Selection) {
		s2, _ := t.Attr("style")
		log.Print(s2)
		if s2 == "text-align: start;" {
			t.Find("img").Each(func(_ int, u *goquery.Selection) {
				url, _ := u.Attr("src")
				GetImage(chara, url) //画像をダウンロードするための関数へ渡す
			})
		}
	})
}

func GetImage(chara, url string) { //画像ダウンロード
	response, err := http.Get(url)
	if err != nil {
		panic(chara)
	}
	defer response.Body.Close()
	filename := strings.LastIndex(url, "/") + 1                                //保存するファイル名決定
	file, err := os.Create(fmt.Sprintf("member/%s/%s", chara, url[filename:])) //ディレクトリおよびファイル作成
	if err != nil {
		panic(chara)
	}
	defer file.Close()
	io.Copy(file, response.Body) //作成したファイルに画像の情報をいれる
}

func GetPage(chara, url string) { //画像URL取得
	doc, _ := goquery.NewDocument(url)
	doc.Find("div").Each(func(_ int, s *goquery.Selection) {
		s2, _ := s.Attr("class")
		if s2 == "box-article" {
			s.Find("img").Each(func(_ int, k *goquery.Selection) {
				url, _ := k.Attr("src")
				if url == "" {
					Errorhandle(chara, s)
				} else {
					GetImage(chara, url)
				}
			})
		}
	})
}
