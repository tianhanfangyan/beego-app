// @APIVersion 1.0.0
// @Title beego app API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact dystargate@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
