/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package conf

import (
	"bytes"
	"fmt"
	"github.com/EchoJamie/ejx/tools/filepath"
	"github.com/EchoJamie/ejx/tools/request"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
)

// InitConfig
// @desc	使用viper初始化配置的工具函数
//
//	  		使用时需保证在 ${project.home}/configs/ 目录下
//	  		创建 ejx.default.yaml 文件
//	  		例如：
//				配置文件名为 abc.xyz ,则需要创建 ${project.home}/configs/abc.default.xyz 文件
//
// @param 	confName	配置文件名
// @param 	confPaths 配置文件遍历路径
func InitConfig(confName string, confPaths ...string) {
	// 配置文件名
	viper.SetConfigFile(confName)
	// 配置文件遍历路径
	for _, confPath := range confPaths {
		viper.AddConfigPath(confPath)
	}
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		confPath := confPaths[0]
		// 路径是否存在
		if !filepath.Exist(confPath) {
			err := os.MkdirAll(confPath, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		// 文件是否存在
		confFile := confPath + "/" + confName
		if !filepath.Exist(confFile) {
			// 初始化配置
			viper.SetConfigType("yaml")
			template := gainConfigTemplate(confName)
			if template != nil {
				// 读取模板
				err := viper.ReadConfig(bytes.NewReader(template))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				// 保存配置
				err = viper.WriteConfigAs(confFile)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

// gainConfigTemplate
// @desc	获取配置模板
// @param	configName	配置文件名
// @return	[]byte		配置模板
func gainConfigTemplate(configName string) []byte {
	// 替换为 默认配置文件名
	replaceName := strings.Replace(configName, ".", ".default.", -1)
	return requestConfigTemplate(replaceName)
}

// requestConfigTemplate
// @desc	请求配置模板
// @param	filename	文件名
// @return	[]byte		配置模板
func requestConfigTemplate(filename string) []byte {

	resp := request.Get("https://raw.githubusercontent.com/EchoJamie/ejx/main/configs/"+filename, nil, request.JsonHeader())
	defer request.CloseRespBody(resp.Body)

	if resp.StatusCode == 200 {
		bytes, _ := io.ReadAll(resp.Body)
		return bytes
	}

	fmt.Println("请重试: ", resp.StatusCode)
	os.Exit(1)
	return nil
}
