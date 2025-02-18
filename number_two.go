package main

import (
	"fmt"
	"strings"
)

func NumberTwo(totalAmount, paymentAmount int) string {
	var (
		moneyFractions  = []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
		cashbackTotal   = paymentAmount - totalAmount
		cashbackRounded int
		cashbackMoneys  = []string{}
	)

	//validate payment amount
	if paymentAmount < totalAmount {
		return "false, kurang bayar"
	}

	for _, fraction := range moneyFractions {
		totalCountFraction := cashbackTotal / fraction
		fractionDetail := "LEMBAR"
		if fraction < 1000 {
			fractionDetail = "KOIN"
		}

		if totalCountFraction > 0 {
			cashbackMoneys = append(cashbackMoneys, fmt.Sprintf("%v %v %v", totalCountFraction, fractionDetail, fraction))
			cashbackTotal -= totalCountFraction * fraction
			cashbackRounded += totalCountFraction * fraction
		}
	}

	return fmt.Sprintf("Total kembalian : %v\n%v", cashbackRounded, strings.Join(cashbackMoneys, "\n"))
}
