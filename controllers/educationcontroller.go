package controllers

// Education API
type EducationController struct {
	BaseController
}

func (c *EducationController) URLMapping() {
}

// @Title AddEducation
// @Description Add education.
// @Param	profile_id	path	string	true	"Profile Id."
// @Accept json
// @Success 200  {object} models.Education
// @router /:profile_id/education [post]
func (c *EducationController) AddEducation(profile_id *string) {

}

// @Title GetEducation
// @Description Get education.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	education_id	path	string	true	"Education Id."
// @Success 200  {object} models.Education
// @router /:profile_id/education/:education_id [get]
func (c *EducationController) GetEducation(profile_id, education_id *string) {

}

// @Title GetEducations
// @Description Get educations.
// @Param	profile_id	path	string	true	"Profile Id."
// @router /:profile_id/education [get]
func (c *EducationController) GetEducations(profile_id *string) {

}

// @Title UpdateEducations
// @Description Update educations.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	education_id	path	string	true	"Education Id."
// @Success 200  {object} models.Education
// @router /:profile_id/education/:education_id [patch]
func (c *EducationController) UpdateEducation(profile_id, education_id *string) {

}

// @Title DeleteEducation
// @Description Delete education.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	education_id	path	string	true	"Education Id."
// @router /:profile_id/education/:education_id [delete]
func (c *EducationController) DeleteEducation(profile_id, education_id *string) {

}
