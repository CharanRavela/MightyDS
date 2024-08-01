package main

import "fmt"

func main() {
	fmt.Println("******************* Simple Stack Example - START ******************")
	fmt.Println()
	simple_stack()
	fmt.Println("******************* Simple Stack Example - END ******************")

	fmt.Println()

	fmt.Println("******************* Simple Stack From Data Example - START ******************")
	fmt.Println()
	simple_stack_from_data()
	fmt.Println("******************* Simple Stack From Data Example - END ******************")

	fmt.Println()

	fmt.Println("******************* Bounded Stack Example - START ******************")
	fmt.Println()
	bounded_stack()
	fmt.Println("******************* Bounded Stack Example - END ******************")

	fmt.Println()

	fmt.Println("******************* Bounded Stack From Data Example - START ******************")
	fmt.Println()
	bounded_stack_from_data()
	fmt.Println("******************* Bounded Stack From Data Example - END ******************")
}
