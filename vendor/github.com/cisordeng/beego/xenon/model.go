package xenon

import (
	"fmt"
	"github.com/cisordeng/beego"
	"github.com/cisordeng/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

func RegisterModels() {
	// set default database
	maxIdle := 30
	maxConn := 100

	host := beego.AppConfig.String("db::DB_HOST")
	port := beego.AppConfig.String("db::DB_PORT")
	db := beego.AppConfig.String("db::DB_NAME")
	user := beego.AppConfig.String("db::DB_USER")
	password := beego.AppConfig.String("db::DB_PASSWORD")
	charset := beego.AppConfig.String("db::DB_CHARSET")
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Asia%%2FShanghai", user, password, host, port, db, charset)

	beego.Notice("connect mysql: ", mysqlURL)
	orm.RegisterDataBase("default", "mysql", mysqlURL, maxIdle, maxConn)
	orm.RunSyncdb("default", false, true)
}
