package complexnumbers

import "math"

// Number represents a complex number.
type Number struct {
	real      float64
	imaginary float64
}

// Real returns the real part of the complex number.
func (n Number) Real() float64 {
	return n.real
}

// Imaginary returns the imaginary part of the complex number.
func (n Number) Imaginary() float64 {
	return n.imaginary
}

// Add returns the sum of two complex numbers.
func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imaginary + n2.imaginary}
}

// Subtract returns the difference of two complex numbers.
func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imaginary - n2.imaginary}
}

// Multiply returns the product of two complex numbers.
func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.real*n2.real - n1.imaginary*n2.imaginary, n1.real*n2.imaginary + n1.imaginary*n2.real}
}

// Times returns the product of a complex number and a scalar.
func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.imaginary * factor}
}

// Divide returns the quotient of two complex numbers.
func (n1 Number) Divide(n2 Number) Number {
	return Number{(n1.real*n2.real + n1.imaginary*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary),
		(n1.imaginary*n2.real - n1.real*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary)}
}

// Conjugate returns the conjugate of a complex number.
func (n Number) Conjugate() Number {
	return Number{n.real, -n.imaginary}
}

// Abs returns the absolute value of a complex number.
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

// Exp returns the exponential of a complex number.
func (n Number) Exp() Number {
	return Number{math.Exp(n.real) * math.Cos(n.imaginary), math.Exp(n.real) * math.Sin(n.imaginary)}
}