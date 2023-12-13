package until

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
	_ "zhl/src/gin/default"
	_default "zhl/src/gin/default"

	"github.com/anaskhan96/go-password-encoder"

	"github.com/google/uuid"
)

// RandString 生成一个长度为 l 的随机字符串
func RandString(l int) string {
	// 定义随机字符串的字符集
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 创建一个新的随机数生成器，使用当前时间作为随机数种子
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 初始化结果切片为指定长度
	result := make([]byte, l)
	// 遍历结果切片中的每个字符位置
	for i := range result {
		// 在 charset 中生成一个随机索引，并将对应的字符添加到结果切片中
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	// 将结果切片转换为字符串并返回
	return string(result)
}

// GenerateRandomPhoneNumber 生成随机手机号码
func GenerateRandomPhoneNumber() string {
	rand.Seed(time.Now().UnixNano())

	// 生成随机的手机号码
	phone := fmt.Sprintf("1%010d", rand.Intn(9999999999))

	return phone
}

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
	Slat       string
	EnCoderPwd string
}

// SavePwd 保存密码
func SavePwd(pwd string) *CodeCheck {

	salt, enCoderPwd := password.Encode(pwd, options)
	c := new(CodeCheck)
	code := &CodeCheck{pwd, salt, enCoderPwd}
	c.pwd = pwd
	return code
}

func EncryptPassword(str string) string {
	code := SavePwd(str)
	return code.EnCoderPwd
}

// CheckPwd 检查密码
func CheckPwd(checkValue *CodeCheck) bool {
	check := password.Verify(checkValue.pwd, checkValue.Slat, checkValue.EnCoderPwd, options)
	return check
}
