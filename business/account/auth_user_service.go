package account

import (
	"encoding/json"
	"github.com/cisordeng/beego"
	"github.com/cisordeng/beego/xenon"
)

func AuthUser(name string, password string) string {
	user := GetUserByName(name)
	userMap := EncodeUser(user)
	if user.Password == xenon.EncodeMD5(password) {
		decodedByteToken, err := json.Marshal(userMap)
		xenon.PanicNotNilError(err)
		decodedToken := string(decodedByteToken)

		commonKey := beego.AppConfig.String("api::aesCommonKey")
		sid, err := xenon.EncodeAesWithCommonKey(decodedToken, commonKey)
		xenon.PanicNotNilError(err)
		return sid
	}
	return ""
}