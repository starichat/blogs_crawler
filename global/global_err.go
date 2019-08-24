package global

import "errors"

// 全局错误处理

var (
	ErrNoRecord = errors.New("没有找到记录")
)

func CheckErr(err error){
	if err != nil{
		panic(err)
	}
}
