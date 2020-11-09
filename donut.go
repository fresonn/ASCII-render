/*
	I was inspired by this: https://www.youtube.com/watch?v=DEqXNfs_HhY

	"floatMemset" and "byteMemset" - analogue of memset from C
	https://stackoverflow.com/questions/30614165/is-there-analog-of-memset-in-go

	Original C code: https://www.dropbox.com/s/79ga2m7p2bnj1ga/donut_deobfuscated.c?dl=0
*/

package main

import (
	"fmt"
	"math"
	"time"
)

const (
	delay      = 16 * time.Millisecond // 60 fps hahaha
	coreString = ".,-~:;=!*#$@"
)

type sliceType interface {
	len() int
}

func floatMemset(arr []float64, v float64) {
	for i := range arr {
		arr[i] = v
	}
}

func byteMemset(arr []string, v string) {
	for i := range arr {
		arr[i] = v
	}
}

func main() {
	A := float64(0)
	B := float64(0)

	var i, j float64
	var k int

	z := make([]float64, 1760)
	b := make([]string, 1760)

	fmt.Print("\033[H\033[2J") // clear previous stdout

	for {
		byteMemset(b, " ")
		floatMemset(z, 0)

		for j = 0; j < 6.28; j += 0.07 {
			for i = 0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e

				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))

				o := int(x + 80*y)

				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))

				if y < 22 && y > 0 && x > 0 && x < 80 && D > z[o] {
					z[o] = D

					// golang doesn't have ternary operator
					point := 0
					if N > 0 {
						point = N
					}

					b[o] = string(coreString[point])
				}

			}
		}

		print("\x1b[H")

		for k = 0; k < 1761; k++ {

			v := "\n"

			if k%80 > 0 {
				v = string(b[k])
			}

			fmt.Printf(v)

			A += 0.00004
			B += 0.00002
		}

		time.Sleep(delay)

	}

}
