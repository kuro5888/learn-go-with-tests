package main

type Rectangle struct {
	Width float64
	Height float64
}

// Calculate Perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height) 
}

// Calculate Area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
