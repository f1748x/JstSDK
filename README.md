### 托马斯牌聚水潭SDK
#### 下载包
```go
go get -u github.com/f1748x/JstSDK
```
#### 如果你的版本太旧了 请及时更新sdk到最新版本
* 1.找到你项目的go.mod文件
* 2.打开找到``github.com/f1748x/JstSDK v0.0.3`` 如要更新版本至 ``github.com/f1748x/JstSDK v0.0.4``
```go 
1.系统参数
    Jst := jst.NewClient(Config.Jpartnerid, Config.Jpartnerkey, Config.Jtoken)
2.请求参数
    tmap := make(map[string]interface{})
	tmap["page_size"] = 50
	tmap["page_index"] = 1
	tmap["status"] = "Sent"
	tmap["modified_begin"] = "2021-03-1 00:00:00"
	tmap["modified_end"] = "2021-03-7 00:00:00"

	data := Jst.Order_Signle_Query(tmap)//string

```
<table>
        <tr>
            <th>设备</th>
            <th>设备文件名</th>
            <th>文件描述符</th>
            <th>类型</th>
        </tr>
        <tr>
            <th>键盘</th>
            <th>/dev/stdin</th>
            <th>0</th>
            <th>标准输入</th>
        </tr>
        <tr>
            <th>显示器</th>
            <th>/dev/stdout</th>
            <th>1</th>
            <th>标准输出</th>
        </tr>
        <tr>
            <th>显示器</th>
            <th>/dev/stderr</th>
            <th>2</th>
            <th>标准错误输出</th>
        </tr>
    </table>