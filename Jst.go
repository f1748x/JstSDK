package Jst

import (
	"fmt"

	util "./util"
)

func PrintlnNow() {
	fmt.Println("聚水潭Sdk版本1--------1")
}
func GetOrder() {
	util.GetOrderList()
	util.SendOrder()
}
