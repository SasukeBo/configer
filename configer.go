package configer

import (
	"errors"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	configFileDir             = "./config"
	entryConfigFileName       = "app.yaml"
	developmentConfigFileName = "dev.yaml"
	productionConfigFileName  = "prod.yaml"
	c                         config
)

// getRealConfigDir get real path where config files located.
func getRealConfigDir() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(path, configFileDir)
}

type config struct {
	Configs map[string]interface{}
}

func (c *config) getEnv(key string) interface{} {
	v, ok := c.Configs[key]
	if !ok {
		panic(errors.New(fmt.Sprintf("Missing [%s] in your configuration", key)))
	}

	return v
}

// GetEnv get string config value
func GetEnv(key string) interface{} {
	return c.getEnv(key)
}

func loadConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	nodes := strings.Split(path, "/")

	var entry []byte
	var directory string
	for i:= len(nodes); i > 0 ; i-- {
		path := strings.Join(nodes[0:i], "/")
		binary, err := ioutil.ReadFile(filepath.Join(path, configFileDir, entryConfigFileName))
		if err == nil {
			entry = binary
			directory = path
			break
		}
	}

	if len(entry) == 0 {
		panic(errors.New(fmt.Sprintf("read configuration from %s failed, file missing.", entryConfigFileName)))
	}

	if err := yaml.Unmarshal(entry, &c.Configs); err != nil {
		panic(err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = fmt.Sprint(c.getEnv("env"))
	}

	var filePath string
	switch env {
	case "prod", "PROD":
		filePath = filepath.Join(directory, configFileDir, productionConfigFileName)
	default: // 默认取dev环境
		filePath = filepath.Join(directory, configFileDir, developmentConfigFileName)
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fileName := filepath.Base(filePath)
		panic(errors.New(fmt.Sprintf("read configuration from %s failed: %v", fileName, err)))
	}

	var envConfigs map[string]interface{}
	if err := yaml.Unmarshal([]byte(content), &envConfigs); err != nil {
		panic(err)
	}

	for k, v := range envConfigs {
		c.Configs[k] = v
	}
}

func init() {
	loadConfig()
}
