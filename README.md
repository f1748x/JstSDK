### 托马斯牌聚水潭SDK 第一版
###### golang 版本 聚水潭SDK 请使用第一版
#### 下载包
```go
go get -u github.com/f1748x/JstSDK
```
#### 如果你的版本太旧了 请及时更新sdk到最新版本
* 1.找到你项目的go.mod文件
* 2.打开找到``github.com/f1748x/JstSDK v0.0.3`` 如要更新版本至 ``github.com/f1748x/JstSDK v0.0.4``
```go 
第一个版本简易版可直接使用
1.系统参数
import (
    ""github.com/f1748x/JstSDK""
)
    Jst := jst.NewClient(Config.Jpartnerid, Config.Jpartnerkey, Config.Jtoken)
2.请求参数
    tmap := make(map[string]interface{})
	tmap["page_size"] = 50
	tmap["page_index"] = 1
	tmap["status"] = "Sent"
	tmap["modified_begin"] = "2021-03-1 00:00:00"
	tmap["modified_end"] = "2021-03-7 00:00:00"
//Jst.Send("method",data)
	data := Jst.Send("orders.single.query",tmap)//string

```
------------------------------------------------------------
### 托马斯牌聚水潭SDK 第二版
###### golang 版本 聚水潭SDK 暂不可用  请使用第一版
* 例子:
调用聚水潭订单查询接口: orders.single.query
```go
import (
    ""github.com/f1748x/JstSDK""
)
Jst := jst.NewClient(Config.Jpartnerid, Config.Jpartnerkey, Config.Jtoken)
data :=Jst.Order_Signle_Query("2021-03-1 00:00:00","2021-03-7 00:00:00")
//订单数据
```

### 对照表:
###### 暂不可用
<table>
        <tr>
            <th>聚水潭</th>
            <th>JstSDK</th>
            <th>参数说明</th>
            <th>接口名称</th>
        </tr>
        <tr>
            <th>order.signle.query</th>
            <th>Jst.Order_Signle_Query()</th>
            <th>args ...string</th>
            <th>获取订单</th>
        </tr>
         <tr>
            <th>shops.query</th>
            <th>Jst.Shops_Query()</th>
            <th>args ...string</th>
            <th>获取店铺</th>
</tr>
    </table>
