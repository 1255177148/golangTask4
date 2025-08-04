# golangTask4---个人博客系统后端

## 运行命令
```shell
go run ./cmd/main.go
```

## 数据库初始化脚本
scripts目录下的initBlogDB.sql

## 初始化依赖
```shell
go mod tidy
```
```shell
go mod download
```

## 生成swagger文档
```shell
swag init -g cmd/main.go
```
swagger-ui地址：http://localhost:9080/swagger/index.html

在用户登录返回token后，swagger-ui右上角有个Authorize按钮，点一下，传入Bearer +返回的token拼接起来的字符串，
之后的接口，就会带着这个token请求头发送请求了
## 依赖的第三方工具
本项目登录时，依赖了redis，需要本地安装一个redis，密码可以自己设置，或者不设置，
设置了密码的话，需要在配置文件里加下，并配置到config.go和bootstrap目录下的redis.go中