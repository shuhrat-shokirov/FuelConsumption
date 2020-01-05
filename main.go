package main

func main() {

}
func fuels(consumption int, volume int)int {
	const MaxPercent = 100
	consumptionPer := volume * MaxPercent / consumption
	return consumptionPer
}
