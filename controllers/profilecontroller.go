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

// Profile API
type ProfileController struct {
	BaseController
}

// URLMapping Url mapping register.
func (c *ProfileController) URLMapping() {
	c.Mapping("CreateProfile", c.CreateProfile)
	c.Mapping("GetProfiles", c.GetProfiles)
}

// @Title CreateProfile
// @Description Create profile.
// @Accept json
// @Success 200  {object} models.Profile
// @router / [post]
func (c *ProfileController) CreateProfile() {
	// Unmarshall request.
	profile := new(models.Profile)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate request.
	err = models.Validator.Struct(profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Default values.
	profile.ID = bson.NewObjectId()
	profile.Created = time.Now()
	profile.Updated = time.Now()

	// Default values for nested experiences and educations.
	if profile.Experiences == nil {
		profile.Experiences = make([]*models.Experience, 0)
	}
	for _, experience := range profile.Experiences {
		experience.ID = bson.NewObjectId()
		experience.Created = time.Now()
		experience.Updated = time.Now()
	}
	if profile.Educations == nil {
		profile.Educations = make([]*models.Education, 0)
	}
	for _, education := range profile.Educations {
		education.ID = bson.NewObjectId()
		education.Created = time.Now()
		education.Updated = time.Now()
	}

	// Insert.
	err = models.Insert(models.ProfileCollectionName, profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = profile
	c.ServeJSON()
}

// @Title GetProfile
// @Description Get profile.
// @Param	profile_id	path	string	true	"Profile Id."
// @Success 200  {object} models.Profile
// @router /:profile_id [get]
func (c *ProfileController) GetProfile(profile_id *string) {
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

	// Serve JSON.
	c.Data["json"] = profile
	c.ServeJSON()
}

// @Title GetProfiles
// @Description Get profiles.
// @router / [get]
func (c *ProfileController) GetProfiles() {
	// Prepare query.
	profiles := make([]*models.Profile, 0)

	// Read all.
	err := models.ReadAll(models.ProfileCollectionName, &profiles)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(profiles)
	response["profiles"] = profiles

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateProfile
// @Description Update profile.
// @Param	profile_id	path	string	true	"Profile Id."
// @Success 200  {object} models.Profile
// @router /:profile_id [patch]
func (c *ProfileController) UpdateProfile(profile_id *string) {
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	profile := new(models.Profile)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Validate request.
	err = models.Validator.Struct(profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Update default fields.
	profile.ID = bson.ObjectIdHex(*profile_id)
	profile.Updated = time.Now()

	// Update.
	err = models.Update(models.ProfileCollectionName, *profile_id, profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = profile
	c.ServeJSON()
}

// @Title DeleteProfile
// @Description Delete profile.
// @Param	profile_id	path	string	true	"Profile Id."
// @router /:profile_id [delete]
func (c *ProfileController) DeleteProfile(profile_id *string) {
	if profile_id == nil {
		err := fmt.Errorf("profile_id can not be empty")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Delete.
	err := models.Delete(models.ProfileCollectionName, *profile_id)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}
}
