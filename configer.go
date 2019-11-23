package configer

import (
	"fmt"
	"github.com/SasukeBo/log"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var (
	configFileDir             = "./config"
	entryConfigFileName       = "app.yaml"
	developmentConfigFileName = "dev.yaml"
	productionConfigFileName  = "prod.yaml"
	c                         config
)

type configError struct {
	Message string
}

func (ce configError) Error() string {
	return ce.Message
}

// SetConfigFileDir set config files directory
func SetConfigFileDir(dir string) {
	configFileDir = dir
}

// SetEntryConfigFileName set entry config files name
func SetEntryConfigFileName(name string) {
	entryConfigFileName = name
}

// SetDevelopmentConfigFileName set entry config files name
func SetDevelopmentConfigFileName(name string) {
	developmentConfigFileName = name
}

// SetProductionConfigFileName set entry config files name
func SetProductionConfigFileName(name string) {
	productionConfigFileName = name
}

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

func (c *config) getstring(key string) (string, error) {
	switch v := c.Configs[key].(type) {
	case string:
		return v, nil
	default:
		if v == nil {
			return "", configError{Message: fmt.Sprintf("%s not found in config", key)}
		}

		return "", configError{Message: "config value is not a string value"}
	}
}

func (c *config) getint(key string) (int, error) {
	switch v := c.Configs[key].(type) {
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case int:
		return v, nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		vi, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}

		return vi, nil
	default:
		if v == nil {
			return 0, configError{Message: fmt.Sprintf("%s not found in config", key)}
		}

		return 0, configError{Message: "config value is not an int value"}
	}
}

func (c *config) getbool(key string) (bool, error) {
	switch v := c.Configs[key].(type) {
	case string:
		vb, err := strconv.ParseBool(v)
		if err != nil {
			return false, err
		}
		return vb, nil
	case bool:
		return v, nil
	default:
		if v == nil {
			return false, configError{Message: fmt.Sprintf("%s not found in config", key)}
		}

		return false, configError{Message: "config value is not a bool value"}
	}
}

func (c *config) getfloat(key string) (float64, error) {
	switch v := c.Configs[key].(type) {
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		vf, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}

		return vf, nil
	default:
		if v == nil {
			return 0, configError{Message: fmt.Sprintf("%s not found in config", key)}
		}

		return 0, configError{Message: "config value is not a float value"}
	}
}

// GetString get string config value
func GetString(key string) string {
	v, e := c.getstring(key)
	if e != nil {
		log.Errorln(e)
	}
	return v
}

// GetInt get int config value
func GetInt(key string) int {
	v, e := c.getint(key)
	if e != nil {
		log.Errorln(e)
	}
	return v
}

// GetBool get bool config value
func GetBool(key string) bool {
	v, e := c.getbool(key)
	if e != nil {
		log.Errorln(e)
	}
	return v
}

// GetFloat get float64 config value
func GetFloat(key string) float64 {
	v, e := c.getfloat(key)
	if e != nil {
		log.Errorln(e)
	}
	return v
}

// ReloadConfig reload config after customer configuration,
// for example set a custom configFileDir value
func ReloadConfig() error {
	err := loadConfig()
	if err != nil {
		return err
	}

	return nil
}

func loadConfig() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	entry, err := ioutil.ReadFile(filepath.Join(path, configFileDir, entryConfigFileName))
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(entry), &c.Configs); err != nil {
		return err
	}

	env, err := c.getstring("env")
	if err != nil {
		return err
	}

	var filePath string
	switch env {
	case "dev":
		filePath = filepath.Join(path, configFileDir, developmentConfigFileName)
	case "prod":
		filePath = filepath.Join(path, configFileDir, productionConfigFileName)
	default:
		return configError{Message: "unknown env"}
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Warn("%s file not found.", filePath)
		return nil
	}

	var envConfigs map[string]interface{}
	if err := yaml.Unmarshal([]byte(content), &envConfigs); err != nil {
		return err
	}

	for k, v := range envConfigs {
		c.Configs[k] = v
	}

	return nil
}

func init() {
	err := loadConfig()
	if err != nil {
		log.Errorln(err)
	}
}
