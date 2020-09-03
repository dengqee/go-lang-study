// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
)

//!+String

func (v Var) String(env Env) {
	fmt.Printf("%g", env[v])
}

func (l literal) String(env Env) {
	fmt.Printf("%g", float64(l))
}

func (u unary) String(env Env) {
	fmt.Printf("%s", string(u.op))
	u.x.String(env)
}

func (b binary) String(env Env) {
	b.x.String(env)
	fmt.Printf("%s", string(b.op))
	b.y.String(env)
}

func (c call) String(env Env) {
	switch c.fn {
	case "pow":
		fmt.Printf("pow(")
		c.args[0].String(env)
		fmt.Printf(",")
		c.args[1].String(env)
		fmt.Printf(")")
	case "sin":
		fmt.Printf("sin(")
		c.args[0].String(env)
		fmt.Printf(")")
	case "sqrt":
		fmt.Printf("sqrt(")
		c.args[0].String(env)
		fmt.Printf(")")

	}
}

//!-String
