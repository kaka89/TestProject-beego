package rest

import (
	//"TestProject/domain/entities"
	//"TestProject/application/dto"
	//"encoding/json"

	"github.com/astaxie/beego"
	"TestProject-beego/domain/entities"
	"TestProject-beego/application/dto"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	dto.Object	true		"The object content"
// @Success 200 {string} dto.Object.Id
// @Failure 403 body is empty
//@router / [post]
func (o *ObjectController) Post() {
	var ob entities.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := entities.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} dto.ObjectDto
// @Failure 403 :objectId is empty
//router /:objectId [get]
func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := entities.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			var objectDto dto.ObjectDto
			objectDto.ObjectId = ob.ObjectId
			objectDto.PlayerName = ob.PlayerName
			objectDto.Score = 1000
			o.Data["json"] = objectDto
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} dto.Object
// @Failure 403 :objectId is empty
// router / [get]
func (o *ObjectController) GetAll() {
	obs := entities.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	dto.Object	true		"The body"
// @Success 200 {object} dto.Object
// @Failure 403 :objectId is empty
// router /:objectId [put]
func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob entities.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := entities.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// router /:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	entities.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

