# ** nature **
> 博客api服务

## Get
`cd $GOPATH/src`
`git clone git@github.com:cisordeng/nature.git`

## Environment
### database

1. 在mysql中创建`nature`数据库: `CREATE DATABASE nature DEFAULT CHARSET UTF8MB4;`；
2. 将`nature`数据库授权给`nature`用户：`GRANT ALL ON nature.* TO 'nature'@localhost IDENTIFIED BY 's:66668888';`；
3. 项目目录下执行 `go run main.go syncdb -v`


## How use this ?

`bee run`

## Document
```

```
