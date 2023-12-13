package controller

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"zhl/src/gin/app/model"
	ser "zhl/src/gin/app/service"
	gro "zhl/src/gin/databases"
	"zhl/src/gin/until"
)

var ex ser.StudentService

//receiver 接收者 Requestor 请求者
//生成n个学生信息
func RandStudent(n int) []*model.Student {
	students := make([]*model.Student, 0, n)
	uuidAndPwd := make(map[string]string)

	for i := 0; i < 10; i++ {
		student := new(model.Student)
		student.Uuid = until.UUID()
		student.Name = until.RandString(5)
		student.Sex = model.Sex(rand.Intn(3))
		student.Phone = until.GenerateRandomPhoneNumber()
		// 生成随机密码
		randPwd := fmt.Sprintf("1%010d", rand.Intn(999999))
		//生成安全密码保存数据库
		codeDate := until.SavePwd(randPwd)
		student.CodePwd = codeDate.EnCoderPwd
		slat := codeDate.Slat
		student.Salt = slat
		//获取系统当前时间
		currentTime := time.Now()
		student.Uptime = currentTime.Format("2006-01-02 15:04:05")
		student.Text = randPwd
		uuidAndPwd[student.Uuid] = randPwd

		students = append(students, student)
	}
	result := gro.DB.CreateInBatches(students, len(students))
	fmt.Println("result:", result.RowsAffected)
	var stu []model.Student
	fmt.Println("ALL:", gro.DB.Find(&stu))
	// 打印学生数据
	for _, ss := range stu {
		fmt.Printf("ID: %d, 姓名: %s %s\n", ss.Uuid, ss.Name)
		// 可以根据需要打印其他字段信息
	}
	if result.Error != nil {
		fmt.Println("Failed to create batch:", result.Error)
		return nil
	}
	writeMapToFile(uuidAndPwd, "src/gin/resources/data/uuidAndPwd.txt")
	//fmt.Println("生成学生信息成功", students)

	return students
}

func writeMapToFile(data map[string]string, filePath string) error {
	// 检查并创建文件
	err := checkAndCreateTxtFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 转换 map 为字符串
	var builder strings.Builder
	for key, value := range data {
		builder.WriteString(fmt.Sprintf("%s: %s\n", key, value))
	}
	content := builder.String()

	// 打开文件进行追加写入
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return err
	}
	defer file.Close()

	// 写入字符串到文件中
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("写入文件出错:", err)
		return err
	}

	fmt.Println("信息已追加到文件")

	return nil
}

func checkAndCreateTxtFile(filePath string) error {
	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 文件不存在，创建新文件
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Println("File created:", filePath)
	} else if err != nil {
		return err
	} else {
		fmt.Println("File already exists:", filePath)
	}

	return nil
}
