package http

import (
	//"encoding/json"
	"fmt"
	"github.com/niean/opsadmin/g"
	"github.com/niean/opsadmin/proc"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func configWechatSenderApiRoutes() {
	http.HandleFunc("/wechat/sender", func(w http.ResponseWriter, req *http.Request) {
		// statistics
		proc.HttpRequestCnt.Incr()

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if req.Method != "POST" {
			RenderDataJson(w, fmt.Sprintf("%s not supported", req.Method))
			return
		}

		req.ParseForm()
		params := req.Form
		msg, exist := params["msg"]
		if !exist || len(msg[0]) < 1 {
			RenderDataJson(w, "bad msg")
			return
		}
		data := make(url.Values)
		data["ak"] = []string{g.GetConfig().Wechat.Ak}
		data["sk"] = []string{g.GetConfig().Wechat.Sk}
		data["content"] = msg

		res, err := http.PostForm(g.GetConfig().Wechat.Url, data)
		if err != nil {
			RenderDataJson(w, "send,error")
			proc.WechatSendErrCnt.Incr()
			return
		} else {
			proc.WechatSendCnt.Incr()
		}

		result, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		log.Println(string(result))

		proc.WechatSendOkCnt.Incr()
		RenderDataJson(w, "ok")
	})
}
