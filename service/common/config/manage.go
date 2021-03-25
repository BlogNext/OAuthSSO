package config

import (
	"errors"
	"fmt"
)

//单例模式
var manage *configManage

func init() {
	if manage == nil {
		manage = &configManage{
			configMap: make(map[string]*Config),
		}
	}
}

//configManage管理
type configManage struct {
	configMap map[string]*Config
}

//提过一个对外的函数加载配置文件进入
func LoadConfig(fileName, filePath, fileType string) error {

	//路径转化为相对项目的绝对路径
	config, err := NewConfig(
		SetFileNameOption(fileName),
		SetFilePathOption(filePath),
		SetFileType(fileType),
	)

	if err != nil {
		return err
	}

	manage.configMap[fileName] = config
	return nil
}

//获取配置文件
func GetConfig(fileName string) (*Config, error) {

	if config, ok := manage.configMap[fileName]; ok {
		return config, nil
	}

	return nil, errors.New(fmt.Sprintf("没有找到配置文件: %s", fileName))
}
