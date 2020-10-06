package main

import (
	"fmt"

	"github.com/komish/vendor-demo-mod-2/pkg/vehicles"
	"github.com/komish/vendor-demo-mod-3/pkg/colors"
)

func main() {
	c := colors.Color{
		NameEnglish: "Blue",
		NameSpanish: "Azul",
	}
	/* Let's get a blue sedan, but this one fails */
	// myCar := vehicles.GetSedan(c)

	/* Let's get a blue sedan, successfully this time */
	myCar := vehicles.GetSedan(vehicles.GetColor(c.NameEnglish, c.NameSpanish))

	/* Print out some information about our car */
	fmt.Println("I drive a", myCar.Color.NameEnglish, myCar.Type)
}
