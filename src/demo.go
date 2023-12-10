package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

//盐值加密：
//通过随机生成数和密码进行组合
//数据库同时存储MD5值和salt值，验证使用正确的salt的值进行MD5即可

//随机字符串
func randString(l int) string {
	str := "0123456789abcdefghigklmnopqrstuvwxyz"
	strList := []byte(str)

	result := []byte{}
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
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
	dataType, _ := json.Marshal(m)
	dataString := string(dataType)
	//json转string
	//去掉 { } , : " 等符号，只保留键值连在一起的字符串
	stringReplace := []string{"{", "}", ",", ":", "\""}
	for i := 0; i < len(stringReplace); i++ {
		dataString = strings.Replace(dataString, stringReplace[i], "", -1)
	}

	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	i := 0
	for k, v := range keys {
		if v == key {
			i = k
		}
	}
	var dataValue string
	keyIndex := strings.Index(dataString, keys[i])
	if i == len(keys)-1 {
		dataValue = dataString[keyIndex+len(key):]
		return dataValue
	}
	endIndex := strings.Index(dataString, keys[i+1])
	dataValue = dataString[keyIndex+len(key) : endIndex]
	return dataValue
}

func MapToString(m map[string]interface{}) map[string]string {
	dataType, _ := json.Marshal(m)
	dataString := string(dataType)
	//json转string
	//去掉 { } , : " 等符号，只保留键值连在一起的字符串
	stringReplace := []string{"{", "}", ",", ":", "\""}
	for i := 0; i < len(stringReplace); i++ {
		dataString = strings.Replace(dataString, stringReplace[i], "", -1)
	}
	fmt.Println("srrtttttt: ", dataString)
	//保存key值，以便后续组装
	respMap := make(map[string]string)
	keys := make([]string, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	fmt.Println("{{{{{{key: ", keys)

	//组装map
	for i := 0; i < len(keys); i++ {
		start := strings.Index(dataString, keys[i])
		if i < len(keys)-1 {
			endIndex := strings.Index(dataString, keys[i+1])
			respMap[keys[i]] = dataString[start+len(keys[i]) : endIndex]
			continue
		}
		respMap[keys[i]] = dataString[start+len(keys[i]):]
	}
	return respMap
}

func interfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
