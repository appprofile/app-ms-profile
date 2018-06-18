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
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	education := new(models.Education)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, education)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate request.
	err = models.Validator.Struct(education)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Default values.
	education.ID = bson.NewObjectId()
	education.Created = time.Now()
	education.Updated = time.Now()

	// Insert.
	dao := models.NewProfileDao()
	err = dao.AddEducation(*profile_id, education)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = education
	c.ServeJSON()

}

// @Title GetEducation
// @Description Get education.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	education_id	path	string	true	"Education Id."
// @Success 200  {object} models.Education
// @router /:profile_id/education/:education_id [get]
func (c *EducationController) GetEducation(profile_id, education_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate education ID.
	if education_id == nil {
		err := fmt.Errorf("education_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Read.
	education := new(models.Education)
	dao := models.NewProfileDao()
	err := dao.GetEducation(*profile_id, *education_id, education)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = education
	c.ServeJSON()
}

// @Title GetEducations
// @Description Get educations.
// @Param	profile_id	path	string	true	"Profile Id."
// @router /:profile_id/education [get]
func (c *EducationController) GetEducations(profile_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty.")
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

	if len(profile.Educations) == 0 {
		profile.Educations = make([]*models.Education, 0)
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(profile.Educations)
	response["educations"] = profile.Educations

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateEducations
// @Description Update educations.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	education_id	path	string	true	"Education Id."
// @Success 200  {object} models.Education
// @router /:profile_id/education/:education_id [patch]
func (c *EducationController) UpdateEducation(profile_id, education_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate education ID.
	if education_id == nil {
		err := fmt.Errorf("education_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	education := new(models.Education)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, education)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Force education ID.
	education.ID = bson.ObjectIdHex(*education_id)

	// Delete.
	dao := models.NewProfileDao()
	err = dao.UpdateEducation(*profile_id, education)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = education
	c.ServeJSON()
}

// @Title DeleteEducation
// @Description Delete education.
// @Param	profile_id	path	string	true	"Profile Id."
// @Param	education_id	path	string	true	"Education Id."
// @router /:profile_id/education/:education_id [delete]
func (c *EducationController) DeleteEducation(profile_id, education_id *string) {
	// Validate profile ID.
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate education ID.
	if education_id == nil {
		err := fmt.Errorf("education_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Delete.
	dao := models.NewProfileDao()
	err := dao.RemoveEducation(*profile_id, *education_id)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}
}
