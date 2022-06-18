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

var wg sync.WaitGroup

func randString(l int) string {
	str := "0123456789abcdefghigklmnopqrstuvwxyz"
	strList := []byte(str)
	result := []byte{}
	i := 0
	for i < l {
		new := strList[rand.Intn(len(strList))]
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
func a(c1 chan<- interface{}) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		str := randString(4)
		c1 <- str
	}
}
func b(c1 chan<- interface{}) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		in := randInt()
		c1 <- in
	}
}
func c(c1 chan<- interface{}) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fl := strconv.FormatFloat(randFloat(), 'f', 1, 64)
		c1 <- fl
	}
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
	fmt.Println("GetMapValueToString-m: ", dataString)
	/*keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	i := 0
	for k, v := range keys {
		if v == key {
			i = k
		}
	}*/
	var dataValue string
	keyIndex := strings.Index(dataString, key)
	if keyIndex == len(m)-1 {
		dataValue = dataString[keyIndex+len(key):]
		return dataValue
	}
	//endIndex := strings.Index(dataString, keys[i+1])
	dataValue = dataString[keyIndex+len(key) : keyIndex+2*len(key)]
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
	fmt.Println("MapToString-dataString: ", dataString)
	//保存key值，以便后续组装
	respMap := make(map[string]string)
	keys := make([]string, 0, len(m))
	for key, _ := range m {
		fmt.Printf("%s ", key)
		keys = append(keys, key)
	}
	fmt.Println()
	fmt.Println("MapToString-keys: ", keys)

	//组装map
	for i := 0; i < len(keys); i++ {
		start := strings.Index(dataString, keys[i])
		if i < len(keys)-1 {
			//endIndex := strings.Index(dataString, keys[i+1])
			respMap[keys[i]] = dataString[start+len(keys[i]) : start+2*len(keys[i])]
			continue
		}
		respMap[keys[i]] = dataString[start+len(keys[i]):]
	}
	fmt.Println("MapToString-respMap:", respMap)
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
func main() {
	cc := make(chan interface{}, 20)
	wg.Add(3)
	go a(cc)
	go b(cc)
	go c(cc)
	wg.Wait()
	close(cc)
	m := make(map[string]interface{})
	keys := make([]string, 0, 30)
	for v := range cc {
		str := interfaceToString(v)
		keys = append(keys, str)
		m[str] = v
	}
	fmt.Println("m::", m)
	ind := rand.Intn(15)
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
	mq := MapToString(m)
	fmt.Println("lan(V): ", mq[keystr])
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
