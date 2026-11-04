//διβάζει μια λίστα και αφαιρεί το στοιχείο που βρίσκεται στη θέση που δίνει ο χρήστης
package main

import "fmt"

func main() {
	s := []float32{1.5, 2.5, 3.5, 4.5, 5.5}
	var idx int
	fmt.Print("Enter the index: ")
	fmt.Scan(&idx)

	if idx < 0 || idx >= len(s) {
		fmt.Println("Invalid index")
		return
	}
	s = append(s[:idx], s[idx+1:]...) 
	fmt.Println(s)
}