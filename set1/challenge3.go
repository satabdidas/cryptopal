
package main

import (
	"encoding/hex"
	"fmt"
	"math"
)

func get_ascii_value(b byte) (val int) {
	val = -1

	if b >= 65 && b <= 90 {
		val = int(b - 65)
	} else if b >= 97 && b <= 122 {
		val = int(b - 97)
	}
	// fmt.Printf("%x : %d\n", b, val)
	return
}

func find_char_freq(b1 []byte) (char_freq [26]float64) {
	total := len(b1)
	for i := 0; i < total; i++ {
		val := get_ascii_value(b1[i])
		if val != -1 {
			char_freq[val]++
		}
	}
	for k := range char_freq {
		char_freq[k] = char_freq[k] / float64(total)
	}
	// fmt.Println("Char freq of ",  char_freq)
	return
}

func get_diff(b1 []byte) (diff float64) {
	expected_char_freq := [26]float64{8.12,	1.49, 2.71,
		4.32, 12.02, 2.30,
		2.03, 5.92, 7.31,
		0.10, 0.69, 3.98,
		2.61, 6.95, 7.68,
		1.82, 0.11, 6.02,
		6.28, 9.10, 2.88,
		1.11, 2.09, 0.17,
		2.11, 0.07}
	
	actual_char_freq := find_char_freq(b1)

	for i, _ := range expected_char_freq {
		diff += math.Abs(expected_char_freq[i] - actual_char_freq[i])
	}

	return
}

func decipher_xor_cipher(hexstr string) {
	b1, err := hex.DecodeString(hexstr);
	if (err != nil) {
		fmt.Println("String decoding failed", err);
		return
	}
	fmt.Println("Input String", string(b1[:len(b1)]))

	min_diff := math.MaxFloat64
	var key byte
	for k := 0; k < 256; k++ {
		b2 := make([]byte, len(b1))
		for i := 0; i < len(b1); i++ {
			b2[i] = (b1[i] ^ byte(k))
		}
		
		// fmt.Println(b2)

		diff := get_diff(b2)
		// fmt.Println(diff)
		if diff < min_diff {
			min_diff = diff
			key = byte(k)
		}
	}
	
	// fmt.Println("Key", key, "Diff", min_diff)
	fmt.Println(string(func () []byte{
		b2 := make([]byte, len(b1))
		for i := 0; i < len(b1); i++ {
			b2[i] = b1[i] ^ key
		}
		return b2
	} ()))
}

func main() {
	decipher_xor_cipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
}
