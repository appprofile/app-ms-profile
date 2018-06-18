package controllers

import (
	"app-ms-profile/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
)

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
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	experience := new(models.Experience)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, experience)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate request.
	err = models.Validator.Struct(experience)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Default values.
	experience.ID = bson.NewObjectId()
	experience.Created = time.Now()
	experience.Updated = time.Now()

	// Insert.
	dao := models.NewProfileDao()
	err = dao.AddExperience(*profile_id, experience)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = experience
	c.ServeJSON()
}

// @Title GetExperience
// @Description Get experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	experience_id	path	string	true	"Experience Id."
// @Success 200  {object} models.Experience
// @router /:profile_id/experience/:experience_id [get]
func (c *ExperienceController) GetExperience(profile_id, experience_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate experience ID.
	if experience_id == nil {
		err := fmt.Errorf("experience_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Read.
	experience := new(models.Experience)
	dao := models.NewProfileDao()
	err := dao.GetExperience(*profile_id, *experience_id, experience)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = experience
	c.ServeJSON()
}

// @Title GetExperiences
// @Description Get experiences.
// @Param	profile_id	path	string	true	"Profile Id."
// @router /:profile_id/experience [get]
func (c *ExperienceController) GetExperiences(profile_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	profile := new(models.Profile)

	// Read.
	err := models.Read(models.ProfileCollectionName, *profile_id, profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	if len(profile.Experiences) == 0 {
		profile.Experiences = make([]*models.Experience, 0)
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(profile.Experiences)
	response["experiences"] = profile.Experiences

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateExperience
// @Description Update experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	experience_id	path	string	true	"Experience Id."
// @Success 200  {object} models.Experience
// @router /:profile_id/experience/:experience_id [patch]
func (c *ExperienceController) UpdateExperience(profile_id, experience_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate experience ID.
	if experience_id == nil {
		err := fmt.Errorf("experience_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	experience := new(models.Experience)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, experience)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Force experience ID.
	experience.ID = bson.ObjectIdHex(*experience_id)

	// Delete.
	dao := models.NewProfileDao()
	err = dao.UpdateExperience(*profile_id, experience)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = experience
	c.ServeJSON()
}

// @Title DeleteExperience
// @Description Delete experience.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	experience_id	path	string	true	"Experience Id."
// @router /:profile_id/experience/:experience_id [delete]
func (c *ExperienceController) DeleteExperience(profile_id, experience_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate experience ID.
	if experience_id == nil {
		err := fmt.Errorf("experience_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Delete.
	dao := models.NewProfileDao()
	err := dao.RemoveExperience(*profile_id, *experience_id)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}
}
