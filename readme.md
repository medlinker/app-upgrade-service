# 更新提示接口

- 版本强制更新配置,根据客户端当前安装版本和渠道来配置更新策略和相关信息
- 客户端对接受到的更新提示信息进行展示
- 强制更新要求用户不进行更新则无法使用APP
- 接口需要在App每次打开后进行请求

## 安装
*已经安装go1.6+*  

```
git clone git@github.com:medlinker/app-upgrade-service.git

// 执行安装脚本
./install.sh

// 启动服务
./start.sh

// 测试, 同时支持GET和POST两种请求
curl http://127.0.0.1:8866/version?platform=ios&clientVersion=2.0.0&clientChannel=xiaomistore
```

*环境配置在：conf/env.conf，可以根据实际环境修改；conf/configure.json可以动态修改更新规则，不用重启服务；conf/log.xml日志配置文件*

## 接口

### 接口地址


```
http://domain/version?platform=xxx&clientVersion=2.0.0&clientChannel=xiaomistore
```

接口根据客户端上报的渠道号和安装版本返回版本更新提示,更新链接以及更新的方式

### 请求数据说明 ###

| 名称 | 类型 |  说明 | 示例 |
|---|---|---|---|
| platform | string | 客户端平台 | android,ios |
| clientVersion | string | 客户端版本信息 | 2.0.0,3.2.0|
| clientChannel | string |  渠道ID  | appstore,xiaomistore |


### 返回数据说明 ###

| 名称 | 类型 |  说明 | 可选值 |
|---|---|---|---|
| title | string | 提示标题 | "4.0新版本发布了,快来更新吧!" |
| desc | string | 版本特性描述(支持html语法) | "4.0版本增加了<strong>自动印钱</strong>功能xxxx" |
| action | int |  更新类型  | 1 仅显示提示,2 强制更新,3 已经是最新版本|
| url| string | 更新包URL |  http://store.mi.com/download.php?id=87654321|



### 数据结构

##### Req

```json
{
    "platform":"android",
    "clientVersion":"3.1.0",
    "clientChannel":""
}
```

##### Response

```json
{
    "platform":"android",
    "clientVersion":"3.1.0",
    "clientChannel":"",
    "currentVersion":"4.0.0",
    "data":{
        "title":"4.0新版本发布了,快来更新吧!",
        "desc":"4.0版本增加了自动印钱功能xxxx",
        "action":1,
        "url":"http://medlinker.com/medlinker_4.0.apk"
    }
}
```

### 配置数据建议

>
>  根据客户端的渠道号替换$channel变量,生成相应渠道的安装包地址.
>

```json
{
    "packages" :{
        "android":{
            "version":"4.0.0",
            "title":"xxx",
            "desc":"xxx",
            "url":"http://medlinker.com/apk/4.0.0_${channel}.apk"
        },
        "ios":{
            "version":"3.0.0"
            "title":"xxx",
            "desc":"xxx",
            "url":"http://store.apple.com/download?id=2877365"
            }

    },
    "rules" :[
        {
            "platform":"android",          
            "minVersion":"2.0.0" ,
            "maxVersion":"3.0.0" ,
            "action":2 ,
        },
        {
            "platform":"android",
            "minVersion":"3.0.0" ,
            "maxVersion":"4.0.0" ,
            "action":1 ,
        }
        ,
        {
            "platform":"ios",
            "minVersion":"2.0.0" ,
            "maxVersion":"3.0.0" ,
            "action":1 ,
        }
    ]
}

```


