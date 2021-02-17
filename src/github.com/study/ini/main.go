package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

/**
ini配置文件解析
*/
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

/**
加载ini配置
*/
func loadIni(fileName string, data interface{}) (err error) {

	//参数校验 传进来的参数必须是指针类型（因为需要在函数中赋值） 并且类型必须是结构体
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("必须是指针类型")
		return
	}
	//判断是否是结构体指针 上边判断的是否是结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("必须是结构体指针")
		return
	}
	//1.读取文件得到字节数据类型
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//2.一行一行读取数据 如果是注释就跳过 如果不是注释方括号开头的页跳过
	dataConfigs := string(file)
	lineSlice := strings.Split(dataConfigs, "\r\n")
	var structName string
	for idx, line := range lineSlice {
		//去掉每一个字符串首尾空格
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			//如果是空行就跳过
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("配置文件第:%d错误", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("配置文件第:%d错误", idx+1)
			}
			//根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了嵌套的结构
					structName = field.Name
					break
				}
			}
		} else {
			//以=分隔这一行 =左边是key  =右边是value
			if strings.Index(line, "=") == -1 || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("第%d有语法错误", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//根据structName去data里边发对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //获取嵌套结构体值信息
			sType := sValue.Type()                     //获取嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				//如果不是结构体
				err = fmt.Errorf("data中的字段必须是结构体")
			}
			var fieldName string
			var fieldType reflect.StructField
			//遍历嵌套结构体的每一个字段 判断Tag是不是等于这个key
			for i := 0; i < sValue.NumField(); i++ {
				//嵌套结构体字段
				field := sType.Field(i) //Tag信息是存储在类型信息中的
				fieldType = field
				if field.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字段
				continue
			}
			//如果key==tag 那么就给这个字段赋值
			fieldObj := sValue.FieldByName(fieldName)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				//字符串转int
				var parseInt int64
				parseInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("第%d行语法错误", idx+1)
					return
				}
				fieldObj.SetInt(parseInt)
			case reflect.Bool:
				var parseBool bool
				parseBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("第%d行语法错误", idx+1)
					return
				}
				fieldObj.SetBool(parseBool)
			case reflect.Float32, reflect.Float64:
				var parseFloat float64
				parseFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("第%d行语法错误", idx+1)
					return
				}
				fieldObj.SetFloat(parseFloat)
			}

		}
	}
	return nil
}

func main() {
	var cfg Config
	err := loadIni("D:\\GoWorkspace\\src\\github.com\\study\\ini\\conf.ini", &cfg)
	if err != nil {
		fmt.Printf("加载失败 err:%s", err)
	}
	fmt.Println(cfg)
}
