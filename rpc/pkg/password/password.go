package password

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

//生成密码
func GeneratePassword(str string, salt string) (password string) {
	// md5
	m := md5.New()
	m.Write([]byte(str))
	mByte := m.Sum(nil)

	// hmac
	h := hmac.New(sha256.New, []byte(salt))
	h.Write(mByte)
	password = hex.EncodeToString(h.Sum(nil))

	return
}

//比较密码是否相同
func ComparePassword(str string, salt string, dbPassword string) bool {
	return GeneratePassword(str, salt) == dbPassword
}
