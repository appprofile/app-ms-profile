package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	AppName string
	Profile string
	Server  string
	Branch  string

	CloudConfig *SpringCloudConfig
)

func init() {
	// Init vars.
	AppName = beego.AppConfig.String("appname")
	Profile = beego.AppConfig.String("runmode")
	Server = beego.AppConfig.String("configserver::server")
	Branch = beego.AppConfig.String("configserver::branch")

	// Load configuration.
	CloudConfig = new(SpringCloudConfig)
	loadConfigurationFrom(Server, AppName, Profile, Branch, CloudConfig)
}

// Loads configuration.
func loadConfigurationFrom(server, appName, profile, branch string, cloudConfig *SpringCloudConfig) {
	url := fmt.Sprintf("%s/%s/%s/%s", server, appName, profile, branch)
	logs.Info("Loading config from %s", url)

	// Make HTTP call.
	resp, err := http.Get(url)
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}

	// Unmarshall response.
	err = json.Unmarshal(body, cloudConfig)
	if err != nil {
		panic("Cannot parse configuration, message: " + err.Error())
	}

	// Use beego configuration context. Also viper could be used.
	for key, value := range cloudConfig.PropertySources[0].Source {
		beego.AppConfig.Set(key, value.(string))
	}

	logs.Info("Successfully loaded configuration for service %s", Server)
}

// SpringCloudConfig Response from Spring Cloud Config.
type SpringCloudConfig struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           string           `json:"label"`
	Version         string           `json:"version"`
	PropertySources []PropertySource `json:"propertySources"`
}

// PropertySource Property source.
type PropertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}
