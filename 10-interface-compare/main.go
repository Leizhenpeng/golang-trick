package main

func main() {
	pool := []interface{}{true}
	_ = pool[0]
	/*
		if tag {
			fmt.Println("This is a boolean value")
		}
	*/

	/*
		//type assertion , get dynamic type of interface
			if tag.(bool) {
				fmt.Println("This is a boolean value")
			}
	*/

	/*
		//https://go.dev/ref/spec#Comparison_operators
			if tag == true {
				fmt.Println("This is a boolean value")
			}
	*/
}
