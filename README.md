# go-wx-minigame-service
Build a wechat minigame service via gin(golang) + mongodb.

## 项目介绍
本项目是微信小游戏服务端，使用 Gin +MongoDB 等技术实现，主要功能有：
- 用户登录
- 用户信息更新
- 用户信息查询
- 排行榜

## 环境启动
### 1. 启动 MongoDB 服务 可以通过 docker 启动服务
```shell
docker run -d -p 27017:27017 -v /data/db:/data/db --name mongo-server -d mongo
```
### 2. 配置WX_APPID 和 WX_APPSECRET
```constants/wx_constants.go``` 文件修改自己小游戏的 ```WX_APPID``` 和 ```WX_APPSECRET```
> 通过 mp.weixin.qq.com 登陆小程序后台，点击左侧菜单栏的开发-开发设置，即可看到 AppSecret。

### 3. 启动服务
在根目录下执行
```shell
go run .
```
## TODO
- [ ] 配置证书，支持 https
- [ ] 调整user 路由的测试用例
