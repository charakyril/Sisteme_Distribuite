//ταξινόμηση πίνακα με τον αλγόριθμο της φυσαλίδας 
package main
 
import "fmt"
 
func main() {
	arr := [6]int{5, 2, 9, 1, 7, 3}
 
	// Bubble sort
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
 
