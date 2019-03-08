# ** mango **
> 博客api服务

## GET
`cd $GOPATH/src`
`git clone git@github.com:cisordeng/mango.git`

## 环境搭建
### 数据库

1. 在mysql中创建`mango`数据库: `CREATE DATABASE mango DEFAULT CHARSET UTF8MB4;`；
2. 将`mango`数据库授权给`mango`用户：`GRANT ALL ON mango.* TO 'mango'@localhost IDENTIFIED BY 's:66668888';`；
3. 项目目录下执行 `go run main.go syncdb -v`


## 运行

`bee run`

## 文档
```
$ sudo apt-get install golang
$ go get -u github.com/cisordeng/beego
$ go get -u github.com/beego/bee
$ go get -u github.com/go-sql-driver/mysql
```
