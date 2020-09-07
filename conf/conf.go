package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/Guiyunweb/go-web-template/library/database/orm"
	"github.com/Guiyunweb/go-web-template/library/log"
	"io/ioutil"
)

const (
	_configKey = "cmd/application.toml"
)

var (
	confPath string
	Conf     = &Config{}
)

type Config struct {
	Datasource *orm.Datasource
}

func Init() (err error) {
	if err := local(); err != nil {
		return err
	}
	database()

	return
}

// 读取解析文件
func local() (err error) {
	// 读取文件
	confByte, err := ioutil.ReadFile(_configKey)
	if err != nil {
		log.Error("配置文件读取错误")
		return err
	}
	confPath := string(confByte)

	// 解析文件
	if _, err := toml.Decode(confPath, &Conf); err != nil {
		log.Error("配置文件解析错误")
		return err
	}
	return
}

func database() {
	orm.NewMySQL(Conf.Datasource)
}
