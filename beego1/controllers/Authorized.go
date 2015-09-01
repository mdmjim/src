package controllers

import (
	"github.com/astaxie/beego"
	"XspiderModels"
    "gopkg.in/mgo.v2/bson"
)

type AuthorizedController struct { 
	Usr XspiderModels.Xusr
	Token string
	beego.Controller
}

func (this *AuthorizedController) Prepare(){
	
}

func (this *AuthorizedController) prepare_cert(){
	this.Token=this.GetSession("token").(string)
	this.Usr=XspiderModels.Xusr{}  
	this.Usr.Id=this.GetSession("Usr_Id").(bson.ObjectId)
	this.Usr.Role =this.GetSession("Usr_Role").(string)
}