package main

import (
	"testing"
)

var iv rune = 'D'
var pt string = "Tom"

// control variables = expected results
var c0 = []int{0, 0, 0, 1, 0, 0, 0, 0} // T xor'd with D manually
var c1 = []int{0, 1, 1, 1, 1, 1, 1, 1} // o xor'd with c0 manually
var c2 = []int{0, 0, 0, 1, 0, 0, 1, 0} // m xor'd with c1 manually

// logging out easy stuff/basic syntax tests:
func Test_tom(t *testing.T) {

	// must convert string to rune slice/array
	var s string = "Tom"
	var r []rune = []rune(s)
	t.Logf("Rune slice: \n%v", r)

}

// must translate ascii to binary codes
// must load binary codes do slice array
func Test_toBinSlice(t *testing.T) {

	s := ToBinSlice('T')
	if s[1] != 1 || s[6] != 0 {
		t.Errorf("Expected 1 in s[1] position, got %d", s[1])
	}
}

// must perform xor calculation with iv
func Test_Xor(t *testing.T) {

	rs := make([]rune, len(pt))
	for i, v := range pt {
		rs[i] = rune(v)
	}

	// log out starting outputs
	t.Logf("Starting output for T:  %v", ToBinSlice(rs[0]))
	t.Logf("Starting output for iv: %v", ToBinSlice(iv))

	// testing xor of single values
	if !testEq(Xor(ToBinSlice(rs[0]), ToBinSlice(iv)), c0) {
		t.Errorf("Expected %v, got %v", c0, Xor(ToBinSlice(rs[0]), ToBinSlice(iv)))
	} else {
		t.Log("Output of xor rune tests")
		t.Logf("Expected outcome xor 1: %v", Xor(ToBinSlice(rs[0]), ToBinSlice(iv)))
	}

	if !testEq(Xor(ToBinSlice(rs[1]), c0), c1) {
		t.Errorf("Expected %v, got %v", c1, Xor(ToBinSlice(rs[1]), c0))
	} else {
		t.Logf("Expected outcome xor 2: %v", Xor(ToBinSlice(rs[1]), c0))
	}

	if !testEq(Xor(ToBinSlice(rs[2]), c1), c2) {
		t.Errorf("Expected %v, got %v", c2, Xor(ToBinSlice(rs[2]), c1))
	} else {
		t.Logf("Expected outcome xor 1: %v", Xor(ToBinSlice(rs[2]), c1))
	}

}

// must loop thru slice and perform xor'ing successfully
func Test_XorStr(t *testing.T) {
	// result := make([][]int, len(rs))
	// x := ToBinSlice(iv)
	// for i := range rs {
	// 	result[i] = Xor(ToBinSlice(rs[i]), x)
	// 	x = result[i]
	// }

	result := XorStr(StrToBinArrs(pt), ToBinSlice(iv))

	if !testEq(result[0], c0) || !testEq(result[1], c1) || !testEq(result[2], c2) {
		t.Errorf("Expected %v: got %v", c0, result[0])
		t.Errorf("Expected %v: got %v", c1, result[1])
		t.Errorf("Expected %v: got %v", c2, result[2])
	} else {
		t.Log("Output of looping thru xor tests")
		t.Logf("Expected %v: got %v", c0, result[0])
		t.Logf("Expected %v: got %v", c1, result[1])
		t.Logf("Expected %v: got %v", c2, result[2])
	}
}

// must convert back to ascii
// test works but returns non readable characters.
func Test_BinToAscii(t *testing.T) {

	result := BinToAscii(XorStr(StrToBinArrs(pt), ToBinSlice(iv)))
	// test set results in non-readable characters.
	t.Logf("Checking ascii to see if readable: %v<-- no visible chars?", result)
}

// must reverse calculation with iv successfully
func Test_reversal(t *testing.T) {

	rev := XorDec(XorStr(StrToBinArrs(pt), ToBinSlice(iv)), ToBinSlice(iv))
	t.Logf("Should say Tom in binary: %v", rev)
	t.Logf(string(84))
}

// testing if a slice is equal
func testEq(a, b []int) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
