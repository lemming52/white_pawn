// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToK converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return (Celsius(k) + AbsoluteZeroC) }

// KToF converts Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

//!-
