# golangTask4---个人博客系统后端

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
## 依赖的第三方工具
本项目登录时，依赖了redis，需要本地安装一个redis，密码可以自己设置，或者不设置，
设置了密码的话，需要在配置文件里加下，并配置到config.go和bootstrap目录下的redis.go中