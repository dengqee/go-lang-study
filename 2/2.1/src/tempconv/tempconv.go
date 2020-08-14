// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
    AbsoluteZeroC Celsius = -273.15
    AbsoluteZeroK Kelvin = 0
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g'C",c) }
func (f Fahrenheit) String() string {return fmt.Sprintf("%g'F",f)}
func (k Kelvin) String() string { return fmt.Sprintf("%gk", k) }



