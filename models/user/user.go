package user

import (
	"github.com/cisordeng/beego/orm"
	"time"
)

type User struct {
	Id		  int
	Username 	  string
	Password  string
	Avatar    string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (o *User) TableName() string {
	return "user_user"
}

func init() {
	orm.RegisterModel(new(User))
}