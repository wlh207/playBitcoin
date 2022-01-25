package main

import (
	"math/rand"
	"strconv"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	for a:=0;a<100;a++{

		str := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		res := ""
		for i := 0; i < 13; i++ {
			//产生0到9的整数
			num := rand.Intn(58)
			//将整数转为字符串
			str1 := string(str[num])
			res += str1
			//println(str1)
		}

		res1 := ""
		for i := 0; i < 13; i++ {
			//产生0到9的整数
			num := rand.Intn(58)
			//将整数转为字符串
			str1 := string(str[num])
			res1 += str1
			//println(str1)
		}

		res2 := ""
		for i := 0; i < 13; i++ {
			//产生0到9的整数
			num := rand.Intn(58)
			//将整数转为字符串
			str1 := string(str[num])
			res2 += str1
			//println(str1)
		}

		res3 := ""
		for i := 0; i < 13; i++ {
			//产生0到9的整数
			num := rand.Intn(58)
			//将整数转为字符串
			str1 := string(str[num])
			res3 += str1
			//println(str1)
		}

		xs := strconv.Itoa(a+1)+"." +res + " "+res1 +" "+res2 +" "+res3
		println(xs)
	}


}


