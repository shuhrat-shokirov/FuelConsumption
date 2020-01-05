package main

import "fmt"

func main() {
	consuption := 8
	volume := 5
	fmt.Println(fuels(consuption,volume))
}
func fuels(consumption int, volume int)int {
	consumptionPer := volume * 100 / consumption
	return consumptionPer
}
