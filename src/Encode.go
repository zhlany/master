package main

import (
	"crypto/md5"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
)

//盐值加密：
//通过随机生成数和密码进行组合
//数据库同时存储MD5值和salt值，验证使用正确的salt的值进行MD5即可
/*
var options = &password.Options{
	SaltLen:      16,      //salt的长度
	Iterations:   100,     //迭代次数
	KeyLen:       32,      //生成长度
	HashFunction: md5.New, //哈希函数
}*/

func MakeOptions(SaltLen, Iterations, KeyLen int) *password.Options {
	options := &password.Options{
		SaltLen:      SaltLen,    //salt的长度
		Iterations:   Iterations, //迭代次数
		KeyLen:       KeyLen,     //生成长度
		HashFunction: md5.New,    //哈希函数
	}
	return options
}

// Encode 加密密码
func Encode(pwd string, options *password.Options) (string, string) {
	salt, enCoderPwd := password.Encode(pwd, options)
	fmt.Println(&options.HashFunction)
	fmt.Println(salt)
	fmt.Println(enCoderPwd)
	return salt, enCoderPwd
}

// Verify 验证密码
func Verify(pwd, salt, enCoderPwd string, options *password.Options) bool {
	check := password.Verify(pwd, salt, enCoderPwd, options)
	return check
}
