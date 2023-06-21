package main

import "fmt"

const (
	daysInMonth       = 30
	workingDayInHours = 8
)

func CalculateIfThereAreEnoughBarbers(countOfBarbersInBarbershop int, countOfMenInCity int) (bool, string) {
	enough := true

	numberOfMenWhoComeEveryDay := countOfMenInCity / daysInMonth
	menWhoCanCutBarriesInADay := countOfBarbersInBarbershop * 8

	barriesCantCutNumberOfMen := numberOfMenWhoComeEveryDay - menWhoCanCutBarriesInADay

	if barriesCantCutNumberOfMen != 0 {
		numberOfBarbersNeeded := barriesCantCutNumberOfMen / countOfBarbersInBarbershop

		enough = false

		return enough, fmt.Sprintf("не хвататет барберов: %d", numberOfBarbersNeeded)
	}

	return enough, "барберы не нужны"
}

func main() {
	var (
		countOfMenInCity           int
		countOfBarbersInBarbershop int
	)

	fmt.Println("введите сколько мужчин проживает в городе:")
	fmt.Scan(&countOfMenInCity)

	fmt.Println("введите сколько барберов работает в барбершопе:")
	fmt.Scan(&countOfBarbersInBarbershop)

	enoughBarbers, reason := CalculateIfThereAreEnoughBarbers(countOfBarbersInBarbershop, countOfMenInCity)

	fmt.Println(enoughBarbers, reason)
}
