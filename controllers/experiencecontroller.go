package controllers

// Experience API
type ExperienceController struct {
	BaseController
}

func (c *ExperienceController) URLMapping() {
}

// @Title AddExperience
// @Description Add experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Accept json
// @Success 200  {object} models.Experience
// @router /:profile_id/experience [post]
func (c *ExperienceController) AddExperience(profile_id *string) {

}

// @Title GetExperience
// @Description Get experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	experience_id	path	string	true	"Experience Id."
// @Success 200  {object} models.Experience
// @router /:profile_id/experience/:experience_id [get]
func (c *ExperienceController) GetExperience(profile_id, experience_id *string) {

}

// @Title GetExperiences
// @Description Get experiences.
// @Param	profile_id	path	string	true	"Profile Id."
// @router /:profile_id/experience [get]
func (c *ExperienceController) GetExperiences(profile_id *string) {

}

// @Title UpdateExperience
// @Description Update experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	experience_id	path	string	true	"Experience Id."
// @Success 200  {object} models.Experience
// @router /:profile_id/experience/:experience_id [patch]
func (c *ExperienceController) UpdateExperience(profile_id, experience_id *string) {

}

// @Title DeleteExperience
// @Description Delete experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	experience_id	path	string	true	"Experience Id."
// @router /:profile_id/experience/:experience_id [delete]
func (c *ExperienceController) DeleteExperience(profile_id, experience_id *string) {

}
