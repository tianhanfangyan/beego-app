package controllers

import (
	"beego-app/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Student
type StudentController struct {
	beego.Controller
}

// @Title CreateStudent
// @Description create users
// @Param	body		body 	models.Student	true		"body for user content"
// @Success 200 {int} models.Student.Id
// @Failure 403 body is empty
// @router / [post]
func (u *StudentController) Post() {
	var user models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddStudent(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Student
// @Success 200 {object} models.Student
// @router / [get]
func (u *StudentController) GetAll() {
	users := models.GetAllStudent()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get student by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Student
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *StudentController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetStudent(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the student
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Student	true		"body for user content"
// @Success 200 {object} models.Student
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *StudentController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.Student
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateStudent(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the student
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *StudentController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteStudent(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
