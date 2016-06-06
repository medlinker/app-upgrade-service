package models

import "strings"

type (
	// RequestParam 请求参数
	RequestParam struct {
		Platform      string `json:"platform"`
		ClientVersion string `json:"clientVersion"`
		ClientChannel string `json:"clientChannel"`
	}

	// Response 返回
	Response struct {
		Platform       string `json:"platform"`
		ClientVersion  string `json:"clientVersion"`
		ClientChannel  string `json:"clientChannel"`
		CurrentVersion string `json:"currentVersion"`
		Data           `json:"data"`
	}

	// Data 返回具体数据
	Data struct {
		Title  string `json:"title"`
		Desc   string `json:"desc"`
		Action int    `json:"action"`
		URL    string `json:"url"`
	}
)

// Illegal 判断请求参数合法性
func (req RequestParam) Illegal() bool {
	if strings.TrimSpace(req.ClientChannel) == "" ||
		strings.TrimSpace(req.ClientVersion) == "" ||
		strings.TrimSpace(req.Platform) == "" {
		return true
	}
	return false
}
