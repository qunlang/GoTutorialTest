package main

import (
	"fmt"
)

var countryCapitalMap map[string]string /*创建集合 */
func main() {
	println("hello,world!!")

	data := make(chan map[string]string)
	go pump(data)

	for country := range data {
		fmt.Println(country)
		//go func(v map[string]string) {
		//	println(v)
		//}(v)
	}

}

func pump(ch chan map[string]string) {

	countryCapitalMap := map[string]string{"France": "巴黎", "Italy": "罗马", "Japan": "东京"}
	countryCapitalMap["India "] = "新德里"
	ch <- countryCapitalMap

	//for i := 0; i < 10; i++ {
	//	countryCapitalMap["France"] = "巴黎"
	//}
}

func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
