package until

import (
	"crypto/md5"
	"fmt"
	_ "zhl/src/gin/default"
	_default "zhl/src/gin/default"

	"github.com/anaskhan96/go-password-encoder"

	"github.com/google/uuid"
)

func UUID() string {
	U, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("get uuid failed, err: ", err)
	}
	return U.String()
}

var options = &password.Options{
	SaltLen:      _default.SaltLen,    //salt的长度
	Iterations:   _default.Iterations, //迭代次数
	KeyLen:       _default.KeyLen,     //生成长度
	HashFunction: md5.New,             //哈希函数
}

type CodeCheck struct {
	pwd        string
	slat       string
	enCoderPwd string
}

func CheckPwd(pwd string) *CodeCheck {

	salt, enCoderPwd := password.Encode(pwd, options)
	c := new(CodeCheck)
	code := &CodeCheck{pwd, salt, enCoderPwd}
	c.pwd = pwd
	return code
}

func savePwd(checkValue *CodeCheck) bool {
	check := password.Verify(checkValue.pwd, checkValue.slat, checkValue.enCoderPwd, options)
	return check
}
