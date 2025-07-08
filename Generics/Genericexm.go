// implementation 
package main 
import(
	"fmt"
)

type gasEngine struct{
	gallons float32
	mpg float32
}

type electricEngine struct{
	kwh float32
	mpkwh float32

}

type  car [T any] struct{ //or we can write (gasEngine| electricEngine)
	carMake string
	carModel string
	engine T  // Engine type (either gasEngine or electricEngine)
}
func main (){
	var gasCar = car [gasEngine]{
		carMake: "honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 23,
			mpg: 50,
		},
	}
	var electricCar = car[electricEngine]{
		carMake: "mahindra",
		carModel: "XEV 9E",
		engine: electricEngine{
			kwh: 90,
			mpkwh: 100,
		},
	}
	fmt.Printf("Gas Car: %s %s\n", gasCar.carMake, gasCar.carModel)
	fmt.Printf("Gallons: %.2f, MPG: %.2f\n", gasCar.engine.gallons, gasCar.engine.mpg)

 	fmt.Printf("Electric Car: %s %s\n", electricCar.carMake, electricCar.carModel)
	fmt.Printf("KWH: %.2f, MPKWH: %.2f\n", electricCar.engine.kwh, electricCar.engine.mpkwh)
}
