package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var ivec rune = 'd'
	var ptxt string = "Tom"
	res := XorStr(StrToBinArrs(ptxt), ToBinSlice(ivec))
	rev := XorDec(res, ToBinSlice(ivec))

	fmt.Printf("\nString to be XOr'd: '%v'"+
		" \nInitialization vector (Ascii base 10): %v\n", ptxt, ivec)
	fmt.Printf("\nOperation: \n")
	fmt.Printf("XOr of '%s' in binary:     %v\n", ptxt, res)
	fmt.Printf("XOr of XOr'd('%s') result: %v\n", ptxt, rev)
	fmt.Printf("\nReturned back to ascii: '%s'\n", BinToAscii(rev))

}

func ToBinSlice(r rune) (b []int) {

	a := int64(r)
	sbin := "0" + strconv.FormatInt(a, 2)
	ibin := strings.Split(sbin, "")
	b = make([]int, len(ibin))
	for i, v := range ibin {
		b[i], _ = strconv.Atoi(v)
	}

	return
}

func StrToBinArrs(s string) (ba [][]int) {

	rs := make([]rune, len(s))
	for i, v := range s {
		rs[i] = rune(v)
	}
	ba = make([][]int, len(rs))
	for i, v := range rs {
		ba[i] = ToBinSlice(v)
	}

	return
}

func Xor(r, s []int) (t []int) {

	t = make([]int, len(r))
	for i := range r {
		if r[i]+s[i] == 1 {
			t[i] = 1
		} else {
			t[i] = 0
		}
	}

	return
}

func XorStr(clear [][]int, iv []int) (xr [][]int) {

	xr = make([][]int, len(clear))

	for i := range clear {
		xr[i] = Xor(clear[i], iv)
		iv = xr[i]
	}

	return
}

func XorDec(xd [][]int, iv []int) (xr [][]int) {

	xr = make([][]int, len(xd))

	for i := range xd {
		xr[i] = Xor(xd[i], iv)
		iv = xd[i]
	}

	return
}

func BinToAscii(b [][]int) (o string) {

	m := make([]string, len(b))
	for i := range b {
		n := make([]string, len(b[i]))
		for j, v := range b[i] {
			n[j] = strconv.Itoa(v)
		}
		s := strings.Join(n, "")
		m[i] = s
	}
	p := make([]string, len(m))
	for i := range m {
		if j, err := strconv.ParseInt(m[i], 2, 64); err != nil {
			fmt.Println(err)
		} else {
			p[i] = string(j)
		}
	}
	o = strings.Join(p, "")

	return
}
