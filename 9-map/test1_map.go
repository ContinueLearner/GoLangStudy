package main

import "fmt"

func main() {
	//第一种声明方式，不初始化值的形式
	myMap := make(map[int]string)
	myMap[1] = "china"
	myMap[2] = "England"
	myMap[3] = "Japan"

	//第二种声明方式，带有初始化值的
	myMap1 := map[int]string{
		1: "china",
		2: "England",
		3: "Japan",
	}
	for key, value := range myMap1 {
		fmt.Println(key, value)
	}

	//添加
	myMap[4] = "London"
	myMap[5] = "USA"
	myMap[6] = "NewYork"
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//删除
	delete(myMap, 5)
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//修改
	myMap[6] = "fkjajefa"
	for key, value := range myMap {
		fmt.Println(key, value)
	}

}
