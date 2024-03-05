package config

import (
	"log"
	"os"
	"strconv"
	"yugod-backend/app/util"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	// System define
	Port                  int  `yaml:"Port"`
	Debug                 bool `yaml:"Debug"`
	Mode                  string
	MaximumUploadFileSize int64  `yaml:"MaximumUploadFileSize"`
	FrontendURL           string `yaml:"FrontendURL"`
	BackendURL            string `yaml:"BackendURL"`
	DownloadCode          uint   `yaml:"DownloadCode"`
}

func (app *AppConfig) setAppMode() {
	if app.Debug {
		app.Mode = "debug"
	} else {
		app.Mode = "release"
	}
}

func (app *AppConfig) overwritePortIfNeeded(key string) error {
	port := os.Getenv(key)
	portNumber, err := strconv.Atoi(port)
	if err == nil && portNumber > 0 && portNumber < 65536 {
		app.Port = portNumber
	}
	return err
}

var App AppConfig

func init() {
	var setting AppConfig

	if util.GetEnvBooleanValue("APP_USE_ENV") {
		// 使用环境变量配置项目
		setting = AppConfig{
			Debug:       util.GetEnvBooleanValue("APP_DEBUG"),
			FrontendURL: os.Getenv("APP_FRONTEND_URL"),
			BackendURL:  os.Getenv("APP_BACKEND_URL"),
		}

		downloadCode, _ := strconv.Atoi(os.Getenv("APP_DOWNLOAD_CODE"))
		setting.DownloadCode = uint(downloadCode)

		if err := setting.overwritePortIfNeeded("APP_PORT"); err != nil {
			setting.Port = 8080
		}
		maxUploadFileSize, _ := strconv.Atoi(os.Getenv("APP_MAXIMUM_UPLOAD_FILE_SIZE"))
		setting.MaximumUploadFileSize = int64(maxUploadFileSize)
	} else {
		// 使用yml配置项目
		config, err := os.ReadFile("config/app.yml")
		if err != nil {
			log.Fatal("App config not set.")
		}
		yamlErr := yaml.Unmarshal(config, &setting)
		if yamlErr != nil {
			log.Fatal("App config read error.")
		}
	}

	App = setting

	App.setAppMode()
	App.overwritePortIfNeeded("PORT")
}
