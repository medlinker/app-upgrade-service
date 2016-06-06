package libconfig

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type IniConfig struct {
	filename string
	entry    map[string]interface{}
}

func NewIniConfig(filename string) *IniConfig {
	iniConfig := &IniConfig{filename, make(map[string]interface{}, 32)}
	iniConfig.parse()
	return iniConfig
}

func (c *IniConfig) parse() {
	file, err := os.Open(c.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	tagName := ""
	for scan.Scan() {
		line := scan.Text()
		n := len(line)
		if n > 0 && (line[0] == '[' && line[n-1] == ']') {
			tagName = line[1 : n-1]
			continue
		}
		if line = strings.TrimSpace(line); strings.Contains(line, "=") {
			kv := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(kv[0])
			if tagName != "" {
				key = tagName + "::" + key
			}
			c.entry[key] = strings.TrimSpace(kv[1])
		}
	}
	if err = scan.Err(); err != nil {
		panic(err)
	}
}

func (c *IniConfig) GetString(key string, defaultValue ...string) string {
	if val, ok := c.entry[key]; ok {
		return val.(string)
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func (c *IniConfig) GetBool(key string, defaultValue ...bool) bool {
	if val, ok := c.entry[key]; ok {
		ret, err := strconv.ParseBool(val.(string))
		if err != nil {
			panic(err)
		}
		return ret
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

func (c *IniConfig) GetInt(key string, defaultValue ...int) int {
	if val, ok := c.entry[key]; ok {
		ret, err := strconv.Atoi(val.(string))
		if err != nil {
			panic(err)
		}
		return ret
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

func (c *IniConfig) Set(key string, value interface{}) {
	c.entry[strings.TrimSpace(key)] = value
}
