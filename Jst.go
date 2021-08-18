package jst

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/15902124763/go-base/http"
)

type JstParams struct {
	Host        string `json:"host" form:"host"`
	Method      string `json:"method" form:"method"`
	Jpartnerid  string `json:"jpartnerid" form:"jpartnerid"`
	Token       string `json:"token" form:"token"`
	Ts          string `json:"ts" form:"ts"`
	Jpartnerkey string `json:"jpartnerkey" form:"jpartnerkey"`
	Sign        string `json:"sign" form:"sign"`
	Url         string `json:"url"`

	OrderSingleQuery func(method string)
}

func NewClient(Jpartnerid, Jpartnerkey, Token string) (jmap *JstParams) {

	jmap = &JstParams{}
	jmap.Jpartnerid = Jpartnerid
	jmap.Token = Token
	jmap.Jpartnerkey = Jpartnerkey

	jmap.Host = "https://open.erp321.com/api/open/query.aspx"

	jmap.Url = fmt.Sprintf("&partnerid=%s&token=%s&ts=", jmap.Jpartnerid, jmap.Token)

	return

}

// orders.single.query
// func (jmap *JstParams) Order_Signle_Query(tmap map[string]interface{}) string {
// 	jmap.Method = "orders.single.query"
// 	ts := time.Now().Unix()
// 	timeTs := fmt.Sprintf("%d", ts)
// 	sign := jmap.Md532(timeTs)
// 	jmap.Sign = sign //%s?method=%s
// 	nUrl := fmt.Sprintf("%s?method=%s%s%s%s%s", jmap.Host, jmap.Method, jmap.Url, timeTs, "&sign=", sign)
// 	res := http.Post(nUrl, tmap, "application/json")
// 	return res
// }
func (jmap *JstParams) Send(method string, tmap map[string]interface{}) string {
	jmap.Method = method
	ts := time.Now().Unix()
	timeTs := fmt.Sprintf("%d", ts)
	sign := jmap.Md532(timeTs)
	jmap.Sign = sign //%s?method=%s
	nUrl := fmt.Sprintf("%s?method=%s%s%s%s%s", jmap.Host, jmap.Method, jmap.Url, timeTs, "&sign=", sign)
	res := http.Post(nUrl, tmap, "application/json")
	return res
}

// md5加密
func (jsting *JstParams) Md532(ts string) string {
	// str := fmt.Sprintf("%s%s%s%s%s%s%s", "orders.single.query", j.Jpartnerid, "token", j.Jtoken, "ts", fmt.Sprintf("%d", timer), j.Jpartnerkey)
	str := fmt.Sprintf("%s%s%s%s%s", jsting.Method, jsting.Jpartnerid, "token"+jsting.Token, "ts"+ts, jsting.Jpartnerkey)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
