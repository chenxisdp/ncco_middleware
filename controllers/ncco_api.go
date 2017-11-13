package controllers

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/ncco_middleware/nccologic"
)

type NccoController struct {
	beego.Controller
}

func (this *NccoController) Post() {
	var body nccologic.AnswerMessage
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	log.Println(body)
	if err != nil {
		this.RenderError(err.Error())
	} else {
		nccologic.BuildAnswerURL(body)
		this.RenderJson(body)
	}
}

func (this *NccoController) RenderError(msg string) {
	resp := NewDto()
	resp.Code = "fail"
	resp.Msg = msg
	this.Data["json"] = resp
	this.Ctx.Output.SetStatus(400)
	this.ServeJSON()
}

func (this *NccoController) RenderJson(json interface{}) {
	this.Data["json"] = json
	this.ServeJSON()
}

type Dto struct {
	Code        string `json:"code"`
	Msg         string `json:"message"`
	Host_id     string `json:"host_id"`
	Server_time string `json:"server_time"`
}

func NewDto() *Dto {
	host_name, _ := os.Hostname()
	return &Dto{
		Host_id:     host_name,
		Server_time: time.Now().Format(time.RFC3339),
	}
}
