package mode

type JstParams struct {
	Method      string `json:"method" form:"method"`
	Jpartnerid  string `json:"jpartnerid" form:"jpartnerid"`
	Token       string `json:"token" form:"token"`
	Ts          string `json:"ts" form:"ts"`
	Jpartnerkey string `json:"jpartnerkey" form:"jpartnerkey"`
	Sign        string `json:"sign" form:"sign"`

	OrderSingleQuery func(method string)
}
