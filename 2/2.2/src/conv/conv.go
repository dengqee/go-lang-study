// Package conv performs foot and meter, pound and kilogram.
package conv

import "fmt"


type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

func (f Foot) String() string { return fmt.Sprintf("%gfeet", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string { return fmt.Sprintf("%gpounds",p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg",k) }

func FToM(f Foot) Meter { return Meter(f*0.3048) }
func MToF(m Meter) Foot { return Foot(m/0.3048) }
func PToK(p Pound) Kilogram { return Kilogram(p*0.4535924) }
func KToP(k Kilogram) Pound { return Pound(k/0.4535924) }


