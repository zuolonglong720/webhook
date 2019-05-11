package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
        this.Ctx.WriteString("hello")
}
func (this *MainController)Post(){
    //一个文件一个结构体，为避免麻烦，就创建了个map
    resp := make(map[string]interface{})

    //从前端获得数据，并解包放入resp中
    json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
//	data := this.Ctx.Input.RequestBody
//	fmt.Println(&resp)
//	fmt.Println(string(resp["ref"]))
	fmt.Println(resp["ref"].(string))

	this.Ctx.WriteString(string(resp['ref']))
	fmt.Println("这是一个webhook的测试")
}

