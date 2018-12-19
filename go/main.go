package main

import (
	"fmt"
	"./ElectronicsShop"
)

//func main1() {
//	fmt.Println(DayOfTheProgrammer.Solve(1918))
//}

func main() {
	//fmt.Println(CountingValleys.CountingValleys(8, "UDDDUDUU"))
	//fmt.Println(CountingValleys.CountingValleys(12, "DDUUDDUDUUUD"))
	//fmt.Println(CountingValleys.CountingValleys(10, "UDUUUDUDDD"))
	keyboards := []int32{183477, 732159, 779867, 598794, 596985, 156054, 4}
	drives := []int32{45934, 156030, 99998, 58097, 459353, 866372, 3337}
	var b int32 = 374625
	fmt.Println(ElectronicsShop.GetMoneySpent(keyboards, drives, b))

}
