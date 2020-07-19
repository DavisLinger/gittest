package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type davis struct {
	a int
}

// Generated by https://quicktype.io

type Ctrfb struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	Ctr       []Ctr  `json:"ctr"`
	Nobid     bool   `json:"nobid"`
	Price     int64  `json:"price"`
	RequestID string `json:"request_id"`
	Ver       string `json:"ver"`
}

type Ctr struct {
	AdID  string  `json:"ad_id"`
	Score float64 `json:"score"`
}

// CTRConversionResponseMsg is total msg which is from ctr trace back api.
type CTRConversionResponseMsg struct {
	Code int
	Msg  string
	Data CtrResponseData
}

// CtrResponseData is adjective data from ctr trace back response.
type CtrResponseData struct {
	RequestID      string `json:"request_id"`
	Timestamp      string `json:"timestamp"`
	Os             string `json:"os"`
	Osv            string `json:"osv"`
	Model          string `json:"model"`
	Make           string `json:"make"`
	Carrier        string `json:"carrier"`
	ConnectionType string `json:"connectiontype"`
	PosID          string `json:"pos_id"`
	ADID           string `json:"ad_id"`
	UA             string `json:"ua"`
	ExtData        string `json:"ext_data"`
}

var lastClick *CTRConversionResponseMsg

// this project is use to mock api interface service.
func main() {
	g := gin.Default()
	g.LoadHTMLFiles("./index.html")
	g.Any("/interactive",func(ctx *gin.Context){
		ctx.HTML(200,"index.html",gin.H{
			"title":"yoboo_callback_test",
		})
	})
	g.Any("/api/v1/fb", func(ctx *gin.Context) {
		// log.Println(ctx.Request.URL.String())
		var clickfb CTRConversionResponseMsg = CTRConversionResponseMsg{
			Code: 200,
			Msg:  "ok",
			Data: CtrResponseData{
				RequestID:      ctx.Query("request_id"),
				Timestamp:      strconv.FormatInt(time.Now().Unix(), 10),
				Os:             ctx.Query("media_os"),
				Osv:            ctx.Query("media_osv"),
				Model:          ctx.Query("device_model"),
				Make:           ctx.Query("device_brand"),
				Carrier:        ctx.Query("device_carrier"),
				ConnectionType: ctx.Query("device_network"),
				PosID:          ctx.Query("pos_id"),
				ADID:           ctx.Query("ad_id"),
				UA:             ctx.Query("ua"),
				ExtData:        ctx.Query("ext_data"),
			},
		}
		lastClick = &clickfb
		ctx.Status(200)
	})
	g.Any("/consume", func(ctx *gin.Context) {
		time.Sleep(3 * time.Second)
		// log.Println(ctx.Request.URL.String())
		ctx.Status(200)
	})
	g.Any("/api/v1/ctr", func(ctx *gin.Context) {
		// log.Println(ctx.Request.URL.String())
		// time.Sleep(3 * time.Second)
		// ctx.Status(200)
		data := Ctrfb{
			Code: 200,
			Msg:  "问题不大",
			Data: Data{
				Nobid: false,
				Ver:   "9999",
			},
		}
		d, _ := json.Marshal(data)
		ctx.Data(200, "json", d)
	})
	g.Any("/api/v1/track", func(ctx *gin.Context) {
		resp := CTRConversionResponseMsg{
			Code: 200,
			Msg:  "StatusOK",
			Data: CtrResponseData{
				RequestID:      "test_req_id_20200603",
				Timestamp:      "20200603",
				Os:             "ios",
				Osv:            "9.3",
				Model:          "iphone",
				Make:           "apple",
				Carrier:        "70120",
				ConnectionType: "4g",
				PosID:          "117279",
				ADID:           "10315",
				UA:             "Go-http-client/1.1",
				ExtData:        "CgM3MDcSCDEwMDA2MTMwIgUxMDI3NzoFMTAyNjlCBTEwMTQ1YgEwagEx",
			},
		}
		if lastClick != nil {
			resp = *lastClick
		}
		ctx.JSON(200, resp)
	})
	g.Any("/mock_conversion", func(ctx *gin.Context) {
		ctx.Status(200)
	})
	g.Any("/mock_conversion2", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	g.Any("/mock_conversion3", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	g.Any("/mock_conversion4", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	g.Any("/conversion-test", func(ctx *gin.Context) {
		callbackurl := ctx.Query("callback")
		log.Println("call back url:", callbackurl)
		if callbackurl == "" {
			// log.Println(callbackurl)
			return
		}
		// ctx.Data(200, "text", []byte("ojbk"))
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("开始请求转化追踪回调")
			req, _ := http.NewRequest("GET", callbackurl, nil)
			http.DefaultClient.Do(req)
		}()
		ctx.Data(200, "text", []byte("ojbk"))
	})
	g.Run(":17279")
}

// func main() {
// 	for i := 1; i < 100; i++ {
// 		consume()
// 	}
// }

func consume() {
	url := "http://127.0.0.1:8089/consume"
	method := "POST"
	jsonText := `
	{
		"accountId":"10006130",
		"adId":"10412",
		"campaignId":"10184",
		"groupId":"10355",
		"money":401
		}
	`
	payload := strings.NewReader(jsonText)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}