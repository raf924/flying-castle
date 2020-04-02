package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"sync"
)

type Config struct {
	DbUrl         string `json:"db_url" required:"true"`
	DataPath      string `json:"data_path" required:"true" type:"url"`
	S3AccessId    string `json:"s3_access_id" type:"alphanum"`
	S3Secret      string `json:"s3_secret" type:"alphanum"`
	S3Profile     string `json:"s3_profile" type:"alphanum"`
	S3Credentials string `json:"s3_credentials" type:"real_path"`
}

type ConfigFlags struct {
	ConfigFile string `flag:"config" required:"true" default:"config.json" usage:"Path to the config file"`
}

func (c *ConfigFlags) Validate() {

}

var config *Config
var configOnce sync.Once

type validationType string

const (
	Path         validationType = "path"
	RealPath     validationType = "real_path"
	Url          validationType = "url"
	Mail         validationType = "mail"
	Alphanumeric validationType = "alphanum"
	Alpha        validationType = "alpha"
)

var validationMap map[validationType]func(value string) bool

func init() {
	validationMap = make(map[validationType]func(value string) bool)
	validationMap[Path] = func(value string) bool {
		return true
	}
	validationMap[RealPath] = func(value string) bool {
		_, err := os.Stat(value)
		return err == nil
	}
	validationMap[Alphanumeric] = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	validationMap[Alpha] = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	validationMap[Url] = func(value string) bool {
		_, err := url.Parse(value)
		return err == nil
	}
	validationMap[Mail] = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString
}

func readConfig() {
	var configFlags = ConfigFlags{}
	ReadFlags(&configFlags)
	file, err := os.Open(configFlags.ConfigFile)
	if err != nil {
		panic(err)
	}
	config = &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}
	var configValue = reflect.Indirect(reflect.ValueOf(config))
	var configType = configValue.Type()
	var numField = configValue.NumField()
	for i := 0; i < numField; i++ {
		var typeField = configType.Field(i)
		if tagValue, ok := typeField.Tag.Lookup("required"); ok {
			isRequired, err := strconv.ParseBool(tagValue)
			if err != nil {
				continue
			}
			if isRequired {
				valueField := configValue.Field(i)
				jsonField := typeField.Tag.Get("json")
				if !valueField.IsValid() {
					panic(fmt.Sprintf("invalid value for %s", jsonField))
				}
				if valueField.String() == "" {
					panic(fmt.Sprintf("missing value for %s", jsonField))
				}
				if fieldType, ok := typeField.Tag.Lookup("type"); ok {
					if validation, ok := validationMap[validationType(fieldType)]; ok && !validation(valueField.String()) {
						panic(fmt.Sprintf("invalid value for %s", jsonField))
					}
				}
			}
		}
	}
}

func GetConfig() *Config {
	configOnce.Do(func() {
		readConfig()
	})
	return config
}
