# ** gravity **
> 博客api服务

## Get
`cd $GOPATH/src`
`git clone git@github.com:cisordeng/gravity.git`

## Environment
### database

1. 在mysql中创建`gravity`数据库: `CREATE DATABASE gravity DEFAULT CHARSET UTF8MB4;`；
2. 将`gravity`数据库授权给`gravity`用户：`GRANT ALL ON gravity.* TO 'gravity'@localhost IDENTIFIED BY 's:66668888';`；
3. 项目目录下执行 `go run main.go syncdb -v`


## How use this ?

`bee run`

## Document
```

```
