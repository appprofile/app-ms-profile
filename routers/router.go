package routers

import (
	"app-ms-profile/controllers"

	"github.com/astaxie/beego"
)

// @APIVersion 1.0.0
// @Title app-ms-profile.
// @Description Profiles microservice.
// @Contact app.profile.kruger@gmail.com
// @TermsOfServiceUrl
// @License
// @LicenseUrl
func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/profile",
			beego.NSInclude(
				&controllers.ProfileController{},
				&controllers.ExperienceController{},
				&controllers.EducationController{},
			),
		),
	)
	// Register namespace.
	beego.AddNamespace(ns)
}
