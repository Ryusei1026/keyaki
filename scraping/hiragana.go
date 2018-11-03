package scraping

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ryusei/scraping/keyaki/scraping/get"
)

var member2 = map[string]string{
	"iguchi":   "23",
	"sarina":   "24",
	"memi":     "25",
	"kagechan": "26",
	"katochi":  "27",
	"kyonko":   "28",
	"kumi":     "29",
	"mirei":    "30",
	"takase":   "31",
	"takamoto": "32",
	"mei":      "33",
	"kanemura": "34",
	"hina":     "35",
	"kosakana": "36",
	"tomita":   "37",
	"nibu":     "38",
	"hiyori":   "39",
	"konoka":   "40",
	"miyata":   "41",
	"bemiho":   "42",
}

const URL = "http://www.keyakizaka46.com/s/k46o/diary/member/list?ima=0000&ct="

var wait2 sync.WaitGroup

func Hiragana() {
	var chara = []string{"iguchi", "sarina", "memi", "kagechan", "katoshi", "kyonko", "kumi", "mirei", "takase", "takamoto", "mei", "kanemura", "hina", "kosakana", "tomita", "nibu", "hiyori", "konoka", "miyata", "bemiho"}

	start := time.Now()

	for _, v := range chara {
		wait2.Add(1)
		go func(v string) {
			k := member2[v]
			url := URL + k
			os.MkdirAll("member/"+v, 0775)
			get.GetPage(v, url) //jpgファイルを取得する関数にわたす
			wait2.Done()
		}(v)
	}
	wait2.Wait()

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
