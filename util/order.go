package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetOrderList() {
	fmt.Println("获取订单列表------------")
}
func SendOrder() {
	fmt.Println("上传订单------------")
}

// md5加密
func Md532(jmap map[string]string) string {
	// timer := time.Now().Unix
	// str := fmt.Sprintf("%s%s%s%s%s", jmap.Method, jmap.Jpartnerid, "token"+jmap.Token, fmt.Sprintf("ts%d", timer), jmap.Jpartnerkey)
	h := md5.New()
	// h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
