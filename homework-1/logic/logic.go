package logic

import "math"

const course = 65

// ConvertFromRubToUsd - convert from Rub to Usd
func ConvertFromRubToUsd(amount float64) float64 {
	return amount * course
}

// CalculateHypotenuse - calc lenghth of hypotenuse
func CalculateHypotenuse(cat1, cat2 float64) float64 {
	return math.Sqrt(cat1*cat1 + cat2*cat2)
}

// CalculatePerimeter - calc perimeter of triangle
func CalculatePerimeter(cat1, cat2 float64) float64 {
	return cat1 + cat2 + CalculateHypotenuse(cat1, cat2)
}

// CalculateSquare - calc square of triangle
func CalculateSquare(cat1, cat2 float64) float64 {
	return (cat1 * cat2) / 2
}

// CalculatePerspectiveAmount - calc amount in years with interest rate
func CalculatePerspectiveAmount(amount, rate, years float64) float64 {
	return amount + (amount/rate)*years
}
