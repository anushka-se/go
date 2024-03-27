package main

import "fmt"

func variable() {
	// initializing varaibles here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	adminName := "Anushka"
	salary := 35000
	spending := 19194.678

	// changing from float to int
	spendingInt := int(spending)

	fmt.Printf(
		"%v %f %v %q \nName = %q \nSalary = %v \nSpending = %f \nSpending (int) = %v\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
		adminName,
		salary,
		spending,
		spendingInt,
	)
}