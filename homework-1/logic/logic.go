package logic

import "math"

const course = 65

// ConvertFromRubToUsd - convert from Rub to Usd
func ConvertFromRubToUsd(amount float64) float64 {
	return amount * course
}

// GetHypotenuse - calc lenghth of hypotenuse
func GetHypotenuse(cat1, cat2 float64) float64 {
	return math.Sqrt(cat1*cat1 + cat2*cat2)
}

// GetPerimeter - calc perimeter of triangle
func GetPerimeter(cat1, cat2 float64) float64 {
	return cat1 + cat2 + GetHypotenuse(cat1, cat2)
}

// GetSquare - calc square of triangle
func GetSquare(cat1, cat2 float64) float64 {
	return (cat1 * cat2) / 2
}

// GetPerspectiveAmount - calc amount in years with interest rate
func GetPerspectiveAmount(amount, rate, years float64) float64 {
	return amount + (amount/rate)*years
}
