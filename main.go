package main

import (
	"fmt"
	"time"
)

func main ()  {
	day1 := time.Date(2017, time.August, 30, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2018, time.August, 30, 0, 0, 0, 0, time.UTC)

	difference := day2.Sub(day1)

	days := int(difference.Hours() /24)

	fmt.Printf("banyaknya hari antara 30 Agustus 2017 dengan 30 Agustus 2018 adalah %d hari", days)
}