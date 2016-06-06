### 主要针对配置文件解析

```
go get github.com/sosop/libconfig
```

#### 一、解析ini文件

app.ini

```
appname = sosop_app
mode = dev



[prod]
db = mysql@172.78.66.88:3306/dbname
port = 80

[dev]
db = mysql@localhost:3306/dbname
port = 8080

[test]
db = mysql@172.78.66.28:3306/dbname
port = 8088
```

```
package main

import (
	"fmt"
	"github.com/sosop/libconfig"
)

func main() {
	iniConfig := libconfig.NewIniConfig("app.ini")
	appname := iniConfig.GetString("appname")
	mode := iniConfig.GetString("mode")
	devDB := iniConfig.GetString("dev::db")
	testPort := iniConfig.GetInt("test::port")
	fmt.Println(appname, mode, devDB, testPort)
}
```

输出：sosop_app dev mysql@localhost:3306/dbname 8088


#### 二、解析json

config.json

```
{
	"redisCluster": [
		{"host": "192.168.1.100", "port": 6379}, 
		{"host": "192.168.1.101", "port": 6380},
		{"host": "192.168.1.102", "port": 6381}],
	"dbCluster": [
		{"host": "172.20.10.8", "port": 3306}, 
		{"host": "172.20.10.9", "port": 3308},
		{"host": "172.20.10.10", "port": 3310}
	]
}
```


```
type StoreConfig struct {
	RedisCluster []struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"redisCluster`
	DbCluster []struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"dbCluster"`
}

func main() {
	storeConf := &StoreConfig{}
	libconfig.NewJsonConfig("config.json", storeConf)
	fmt.Println(*storeConf)
}
```
输出：{[{192.168.1.100 6379} {192.168.1.101 6380} {192.168.1.102 6381}] [{172.20.10.8 3306} {172.20.10.9 3308} {172.20.10.10 3310}]}

#### 三、解析xml
config.xml

```
<?xml version="1.0"?>
<config>
	<appname>testXml</appname>
	<host>0.0.0.0</host>
  	<port>8888</port>
	
	<group>
		<value>log</value>
		<value>queue</value>
	</group>
	
	
	<es name="es1">
		<host>192.168.1.2</host>
		<port>7989</port>
		<shard>1</shard>
	</es>
	<es name="es2">
		<host>192.168.1.3</host>
		<port>7986</port>
		<shard>2</shard>
	</es>
	<es name="es3">
		<host>192.168.1.4</host>
		<port>7988</port>
		<shard>3</shard>
	</es>
</config>
```

```
type ES struct {
	Name  string `xml:"name,attr"`
	Host  string `xml:"host"`
	Port  int    `xml:"port"`
	Shard int    `xml:"shard"`
}

type XMLConfig struct {
	Appname string   `xml:"appname"`
	Host    string   `xml:"host"`
	Port    int      `xml:"port"`
	Group   []string `xml:"group>value"`
	ESS     []ES     `xml:"es"`
}

func main() {
	xmlConf := &XMLConfig{}
	libconfig.NewXmlConfig("config.xml", xmlConf)
	fmt.Println(*xmlConf)
}
```

输出：{testXml 0.0.0.0 8888 [log queue] [{es1 192.168.1.2 7989 1} {es2 192.168.1.3 7986 2} {es3 192.168.1.4 7988 3}]}

#### 四、yaml
[go-yaml](https://github.com/go-yaml/yaml)


