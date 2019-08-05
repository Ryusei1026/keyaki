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
	"uemu-":    "03",
	"ozeki":    "04",
	"oda":      "05",
	"michan":   "06",
	"yuipon":   "07",
	"saitou":   "08",
	"satoshi":  "09",
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
	"inori":    "43",
	"seki":     "44",
	"takemoto": "45",
	"hono":     "46",
	"karin":    "47",
	"marina":   "48",
	"rikopi":   "49",
	"hikaru":   "50",
	"ten":      "51",
}

const U = " = "

var wait1 sync.WaitGroup

func Kanzi() {
	var chara = []string{"nizika", "uemu-", "ozeki", "oda", "michan", "yuipon", "saitou", "satoshi", "yukka", "suzumon", "na-ko", "habuchan", "aoi", "techi", "berika",
		"berisa", "neru", "inori", "seki", "takemoto", "hono", "karin", "marina", "rikopi", "hikaru", "ten"}

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
