package xenon

import (
	"context"
	
	"github.com/cisordeng/beego/orm"
)

func GetOrmFromContext(ctx context.Context) orm.Ormer {
	o := ctx.Value("orm")
	if o == nil {
		return nil
	}
	return o.(orm.Ormer)
}