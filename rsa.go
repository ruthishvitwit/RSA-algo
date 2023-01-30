package main

import "fmt"

var public_key int
var private_key int
var n int

func Ncode(input string) []int64 {
	elist := []int64{}
	for i := 0; i < len(input); i++ {
		elist = append(elist, Ncrypt(int64(input[i])))
	}
	return elist
}

func Dcode(Ncode []int64) string {
	s := ""
	for i := 0; i < len(Ncode); i++ {
		s += fmt.Sprintf("%c", Dcrypt(Ncode[i]))
	}

	return s
}

func Ncrypt(txt int64) int64 {
	var etxt int64 = 1
	for e := public_key; e > 0; e-- {
		etxt *= int64(txt)
	}
	etxt %= int64(n)

	return etxt
}

func Dcrypt(etxt int64) int64 {
	var Dcrptd int64 = 1
	for d := private_key; d > 0; d-- {
		Dcrptd *= int64(etxt)
		Dcrptd %= int64(n)

	}
	return int64(Dcrptd)
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func init() {
	p := 7
	q := 19
	n = p * q
	fi := (p - 1) * (q - 1)
	e := 2
	for {
		if gcd(e, fi) == 1 {
			break
		}
		e++
	}
	public_key = e

	d := 2
	for {
		if (d*e)%fi == 1 {
			break
		}
		d++
	}
	private_key = d

}
