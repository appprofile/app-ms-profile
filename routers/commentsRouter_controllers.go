package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"],
		beego.ControllerComments{
			Method: "AddEducation",
			Router: `/:profile_id/education`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"],
		beego.ControllerComments{
			Method: "GetEducations",
			Router: `/:profile_id/education`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"],
		beego.ControllerComments{
			Method: "GetEducation",
			Router: `/:profile_id/education/:education_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
				param.New("education_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"],
		beego.ControllerComments{
			Method: "UpdateEducation",
			Router: `/:profile_id/education/:education_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
				param.New("education_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:EducationController"],
		beego.ControllerComments{
			Method: "DeleteEducation",
			Router: `/:profile_id/education/:education_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
				param.New("education_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"],
		beego.ControllerComments{
			Method: "AddExperience",
			Router: `/:profile_id/experience`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"],
		beego.ControllerComments{
			Method: "GetExperiences",
			Router: `/:profile_id/experience`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"],
		beego.ControllerComments{
			Method: "GetExperience",
			Router: `/:profile_id/experience/:experience_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
				param.New("experience_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"],
		beego.ControllerComments{
			Method: "UpdateExperience",
			Router: `/:profile_id/experience/:experience_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
				param.New("experience_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ExperienceController"],
		beego.ControllerComments{
			Method: "DeleteExperience",
			Router: `/:profile_id/experience/:experience_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
				param.New("experience_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "CreateProfile",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "GetProfiles",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "GetProfile",
			Router: `/:profile_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "UpdateProfile",
			Router: `/:profile_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"] = append(beego.GlobalControllerRouter["app-ms-profile/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "DeleteProfile",
			Router: `/:profile_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("profile_id", param.IsRequired, param.InPath),
			),
			Params: nil})

}
