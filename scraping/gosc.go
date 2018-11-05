package scraping

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ryusei/scraping/keyaki/scraping/get"
)

var member1 = map[string]string{
	"nizika":   "01",
	"zumiko":   "02",
	"uemu-":    "03",
	"ozeki":    "04",
	"oda":      "05",
	"michan":   "06",
	"yuipon":   "07",
	"saitou":   "08",
	"satoshi":  "09",
	"manaka":   "10",
	"yukka":    "11",
	"suzumon":  "12",
	"na-ko":    "13",
	"habuchan": "14",
	"aoi":      "15",
	"techi":    "17",
	"akanen":   "18",
	"yone":     "19",
	"berika":   "20",
	"berisa":   "21",
	"neru":     "22",
}

const U = "http://www.keyakizaka46.com/s/k46o/diary/member/list?ima=0000&ct="

var wait1 sync.WaitGroup

func Kanzi() {
	var chara = []string{"nizika", "zumiko", "uemu-", "ozeki", "oda", "michan", "yuipon", "saitou", "satoshi", "manaka", "yukka", "suzumon", "na-ko", "habuchan", "aoi", "techi", "yone", "berika",
		"berisa", "neru"}

	start := time.Now()

	for _, v := range chara {
		wait1.Add(1)
		go func(v string) {
			k := member1[v]
			url := U + k
			os.MkdirAll("member/"+v, 0775)
			get.GetPage(v, url) //jpgファイルを取得する関数にわたす
			wait1.Done()
		}(v)
	}
	wait1.Wait()

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
