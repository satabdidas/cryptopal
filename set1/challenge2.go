package main

import (
	"fmt"
	"encoding/hex"
	"bytes"
)

func xor(b1 []byte, b2 []byte) (b3 []byte) {
	b3 = make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		b3[i] = b1[i] ^ b2[i]
	}
	return
}

func main() {
	b1, err1 := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err1 != nil {
		fmt.Println("String decoding failed", err1)
		return
	}
	b2, err2 := hex.DecodeString("686974207468652062756c6c277320657965")
	if err2 != nil {
		fmt.Println("String decoding failed", err2)
		return
	}

	b3 := xor(b1, b2)
	fmt.Println(hex.EncodeToString(b3)) 
	expected_buf := bytes.NewBufferString("746865206b696420646f6e277420706c6179")
	if bytes.Equal(b3, expected_buf.Bytes()) {
		fmt.Println("Wrong!")
	}
}
