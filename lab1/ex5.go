//γράψιμο λίστας με συγκεκριμένη σειρά 
package main

import "fmt"

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Printf("%q\n", slice)
}