package main

import (
	"fmt"
	"time"
)

var (
	cst *time.Location
)

// Go 
//Go 语言中的时间格式化和解析是基于一个特定的参考时间来定义的。这个参考时间是：
//Mon Jan 2 15:04:05 MST 2006
//这个时间表示的是 2006 年 1 月 2 日下午 3 点 4 分 5 秒（注意，这是美国山地标准时间，即 MST 时区）。这个时间点被用作时间格式化和解析中的一个基准，因为它提供了一个稳定的、可预测的、并且容易记忆的参考。
//原理上，Go 语言的时间格式化和解析系统使用这个参考时间来定义每个格式化字符的含义。当你提供一个格式化字符串时，Go 会将这个字符串中的每个字符与参考时间中的相应部分进行匹配。例如：
//2006 对应于年份（Year）
//01 对应于月份（Month）
//02 对应于日期（Day）
//15 对应于小时（Hour），24 小时制
//04 对应于分钟（Minute）
//05 对应于秒（Second）
//这种设计使得 Go 语言的时间格式化和解析变得非常灵活和强大。你可以使用各种不同的格式化字符串来表示时间，而 Go 会自动根据参考时间来解析和格式化。
//此外，由于这个参考时间是固定的，它还确保了时间格式化和解析的一致性。无论你在世界的哪个地方，无论你的系统时区如何，Go 语言都会按照相同的方式处理时间格式化和解析。
//这种设计的一个副作用是，如果你尝试使用一个与参考时间不匹配的格式化字符串，Go 可能无法正确解析时间。这就是为什么在使用 Go 语言进行时间格式化和解析时，你需要确保你的格式化字符串与参考时间相匹配。如果你的格式化字符串与参考时间不一致，Go 可能无法正确地识别各个时间组件，从而导致格式化或解析错误。
const CSTLayout = "2006-01-02 15:04:05"

// init 初始化方法
func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
}

func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

func main(){
	RFC3339Str := "2020-11-08T08:18:46+08:00"
	fmt.Println(RFC3339ToCSTLayout(RFC3339Str))
}
