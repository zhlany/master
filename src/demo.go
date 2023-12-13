package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//盐值加密：
//通过随机生成数和密码进行组合
//数据库同时存储MD5值和salt值，验证使用正确的salt的值进行MD5即可

// randString 生成一个长度为 l 的随机字符串
func randString(l int) string {
	// 定义随机字符串的字符集
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
func randInt() int {
	return rand.Intn(100)
}
func randFloat() float64 {
	return rand.Float64() + float64(rand.Intn(99))
}

type data struct {
	k *int
	v *string
}

func setF(c interface{}) {

}

func main() {
	/*	options := &password.Options{
			SaltLen:      10,      //salt的长度
			Iterations:   100,     //迭代次数
			KeyLen:       32,      //生成长度
			HashFunction: md5.New, //哈希函数
		}
		salt, enCoderPwd := password.Encode("root@123", options)
		fmt.Println(&options.HashFunction)
		fmt.Println(salt)
		fmt.Println(enCoderPwd)
		check := password.Verify("root@123", salt, enCoderPwd, options)
		fmt.Println(check)*/
	m := make(map[string]interface{})
	channel := make(chan interface{}, 31)
	i := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(c1 chan interface{}) {
		for i < 10 {
			str := randString(4)
			fmt.Println("str:", str)
			c1 <- str
			i++
		}
		close(c1)
		wg.Done()
	}(channel)

	/*	j := 0
		for j < 5 {
			fl := randFloat()
			channel <- fl
			j++
		}
		k := 0
		for k < 5 {
			in := randInt()
			channel <- in
			k++
		}*/
	wg.Wait()
	keys := make([]string, 0, 30)
	for {
		v, ok := <-channel
		if !ok {
			break
		}
		str := interfaceToString(v)
		fmt.Println("vv:", str)
		keys = append(keys, str)
		m[str] = v
	}

	fmt.Println("map: ", m)
	ind := rand.Intn(10)
	fmt.Println("key:", keys[ind], "value:", m[keys[ind]])
	keystr := keys[ind]
	//方法一
	fmt.Println("--------------------------11111111111111111111111111")
	start := time.Now()
	fmt.Println("lan(V): ", interfaceToString(m[keystr]))
	end := time.Now()
	t := end.Sub(start)
	fmt.Println("t1:", t)

	//方法二
	fmt.Println("--------------------------222222222222222222222222")
	start2 := time.Now()
	fmt.Println("lan(V): ", MapToString(m)[keystr])
	end2 := time.Now()
	t2 := end2.Sub(start2)
	fmt.Println("t2:", t2)

	//方法3
	fmt.Println("--------------------------333333333333333333333333")
	start3 := time.Now()
	fmt.Println("lan(V): ", GetMapValueToString(m, keystr))
	end3 := time.Now()
	t3 := end3.Sub(start3)
	fmt.Println("t3:", t3)
}

func GetMapValueToString(m map[string]interface{}, key string) string {
	var dataValue string
	for _, v := range m {
		str := fmt.Sprintf("%v", v)
		dataValue += fmt.Sprintf("%s;", str)
	}
	return dataValue
}

func MapToString(m map[string]interface{}) map[string]string {
	dataType, _ := json.Marshal(m)
	fmt.Println("json:: ", dataType)
	dataString := string(dataType)

	fmt.Println("dataString: ", dataString)
	//保存key值，以便后续组装
	respMap := make(map[string]string)
	for k, v := range m {
		str := fmt.Sprintf("%s", v)
		respMap[k] = str
	}

	return respMap
}

func interfaceToString(value interface{}) string {
	// interface 转 string
	return fmt.Sprintf("%v", value)
}
