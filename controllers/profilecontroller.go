package controllers

import (
	"app-ms-profile/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
)

// Profile API
type ProfileController struct {
	baseController
}

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
		err := fmt.Errorf("profile_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	var profile models.Profile

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
		err := fmt.Errorf("profile_id can not be empty.")
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
		err := fmt.Errorf("profile_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	profile := &models.Profile{Id: bson.ObjectIdHex(*profile_id)}

	// Delete.
	err := models.Delete(models.ProfileCollectionName, profile)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}
}
