package models

import (
	"app-upgrade-service/config"
	. "app-upgrade-service/logger"

	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

const (
	_ = iota
	// TIP 提示更新
	TIP
	// FORCE 强制更新
	FORCE
	// NEWEST 已经最新
	NEWEST
)

var (
	// VersionConfig 版本更新策略
	VersionConfig  = &Configure{}
	lastModifyTime = time.Now()
	// CloseCheck 关闭文件监控
	CloseCheck = make(chan Null, 1)
	// CurretPlatforms 当前用户客户端信息
	CurretPlatforms = make(map[string]platform, 2)
)

// ListenConfiguration 监听配置文件的修改
func ListenConfiguration() {
	path := config.GetString("rulesPath", true, "../conf/configure.json")
	err := format(path)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 2):
				err := checkFile(path)
				if err != nil {
					Logger.Error(err)
				}
			case <-CloseCheck:
				Logger.Info("file check closed")
				break
			}
		}
	}()
}

func format(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, VersionConfig)
	CurretPlatforms["ios"] = VersionConfig.IOS
	CurretPlatforms["android"] = VersionConfig.Android
	return err
}

func checkFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if lastModifyTime.Before(fileInfo.ModTime()) {
		Logger.Info("更新配置文件")
		lastModifyTime = fileInfo.ModTime()
		err := format(path)
		if err != nil {
			return err
		}
	}
	return nil
}

type (
	// platform 操作系统
	platform struct {
		Version string `json:"version"`
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		URL     string `json:"url"`
	}
	// Packages platforms
	Packages struct {
		IOS     platform `json:"ios"`
		Android platform `json:"android"`
	}
	// Rule 规则
	Rule struct {
		Platform   string `json:"platform"`
		MinVersion string `json:"minVersion"`
		MaxVersion string `json:"maxVersion"`
		Action     int    `json:"action"`
	}
	// Configure read from file
	Configure struct {
		Packages `json:"packages"`
		Rules    []Rule `json:"rules"`
	}
)

// Null 空结构体
type Null struct{}
