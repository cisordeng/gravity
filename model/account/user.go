package account

import (
	"github.com/cisordeng/beego/orm"
	"time"
)

type User struct {
	Id int
	Name string
	Password string
	Avatar string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *User) TableName() string {
	return "account_user"
}

func init() {
	orm.RegisterModel(new(User))
}