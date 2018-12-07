package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("-----------------------")
	fmt.Println("年：", start.Year())  //年
	fmt.Println("月：", start.Month()) //月
	fmt.Println("日：", start.Day())
	fmt.Println("星期几：", start.Weekday())   //星期几
	fmt.Println("时：", start.Hour())        //时
	fmt.Println("分：", start.Minute())      //分
	fmt.Println("秒：", start.Second())      //秒
	fmt.Println("纳秒：", start.Nanosecond()) //纳秒[0, 999999999]
	fmt.Print("年 月 日：")                    //返回年 月 日
	fmt.Println(start.Date())
	fmt.Println("----------------------------")
	fmt.Println(start.ISOWeek()) //返回年份 今年的第几周
	fmt.Println(start.Clock())   //时 分 秒
	fmt.Println(start.YearDay()) //今年的第几天
	fmt.Println("----------------------------")
	var d time.Duration = 3600000000000 //初始化的救赎纳秒数
	fmt.Println(d.Minutes())            //把d转换成分钟
	fmt.Println(d.Hours())              //把d转换成小时
	fmt.Println(d.Seconds())            //把d转换成秒钟
	fmt.Println(d.String())             //把d转换成：t时m分s秒
	fmt.Println(d.Nanoseconds())        //纳秒数
	fmt.Println(d.Round(120))
	fmt.Println("-----------------------------")
	fmt.Println(time.Until(start))     //返回的类型是Duration，内部返回的：start.Sub(time.Now())
	fmt.Println(time.Now().Sub(start)) //time.Now()-start返回微妙
	fmt.Println(start.Add(d))          //start+d
	fmt.Println(time.Since(start))     //与Until相反
	fmt.Println(start.UTC())           //返回UTC时间，与我们这儿的时间间隔八个小时，也就是格林威治时间
	fmt.Println(start.Local())         //返回CST时间，与我们这里的时间相同
	fmt.Println()
}
