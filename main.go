package main

import (
	"fmt"
	"math"
)

type converter struct{}

type centimeter float64
type feet float64
type minutes float64
type seconds float64
type milliseconds float64
type celsius float64
type fahrenheit float64
type radian float64
type degree float64
type kilogram float64
type pounds float64

func (cvr converter) CentimeterToFeet(c centimeter) feet {
	result := feet(c * 0.0328084)
	return result
}

func (cvr converter) FeetToCentimeter(f feet) centimeter {
	result := centimeter(f * 30.48)
	return result
}

func (cvr converter) MinutesToSeconds(m minutes) seconds {
	result := seconds(m * 60)
	return result
}

func (cvr converter) SecondsToMinutes(s seconds) minutes {
	result := minutes(s * 0.0166667)
	return result
}

func (cvr converter) SecondsToMilliseconds(s seconds) milliseconds {
	result := milliseconds(s * 1000)
	return result
}

func (cvr converter) MillisecondsToSeconds(mi milliseconds) seconds {
	result := seconds(mi * 0.001)
	return result
}

func (cvr converter) CelsiusToFahrenheit(c celsius) fahrenheit {
	result := fahrenheit((c*(1.8) + 32))
	return result
}

func (cvr converter) FahrenheitToCelsius(f fahrenheit) celsius {
	result := celsius((f - 32) * (0.5555555555555556))
	return result
}

func (cvr converter) RadianToDegree(r radian) degree {
	result := degree(r * (180 / math.Pi))
	return result
}

func (cvr converter) DegreeToRadian(d degree) radian {
	result := radian(d * (math.Pi / 180))
	return result
}

func (cvr converter) KilogramToPounds(k kilogram) pounds {
	result := pounds(k * 2.20462)
	return result
}

func (cvr converter) PoundsToKilogram(p pounds) kilogram {
	result := kilogram(p * 0.453592)
	return result
}

func main() {
	// to access the methods
	// we need an instance to the Converter sruct
	cvr := converter{}
	// calling our method and printing the result at the same time
	fmt.Println(cvr.CentimeterToFeet(10))
	fmt.Println(cvr.FeetToCentimeter(10))
	fmt.Println(cvr.MinutesToSeconds(10))
	fmt.Println(cvr.SecondsToMinutes(10))
	fmt.Println(cvr.SecondsToMilliseconds(10))
	fmt.Println(cvr.MillisecondsToSeconds(10))
	fmt.Println(cvr.CelsiusToFahrenheit(10))
	fmt.Println(cvr.FahrenheitToCelsius(10))
	fmt.Println(cvr.RadianToDegree(10))
	fmt.Println(cvr.DegreeToRadian(10))
	fmt.Println(cvr.KilogramToPounds(10))
	fmt.Println(cvr.PoundsToKilogram(10))

}
