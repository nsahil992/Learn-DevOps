 package main
 import "fmt"

 func calcArea(r float64)float64 {
	return 3.14 * r * r
 }

 func calcPerimeter(r float64)float64 {
	return 2 * 3.14 * r
 }

 func calcDiameter(r float64)float64 {
	return 2 * r
 }

 func main() {

	var query int
	var radius float64

	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
	fmt.Scanf("%d", &query)

	if query == 1 {
		fmt.Println("Result: ", calcArea(radius))
	} else if query == 2{
		fmt.Println("Result: ", calcPerimeter(radius))
	} else if query == 3{
		fmt.Println("Result: ", calcDiameter(radius))
	} else {
		fmt.Println("Invalid query")
	}
	fmt.Println("Thankyou!")
 }
