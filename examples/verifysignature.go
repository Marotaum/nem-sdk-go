package main

import (
	"fmt"

	"github.com/Marotaum/nem-sdk-go/model"
	"github.com/Marotaum/nem-sdk-go/utils"
)

func main() {
	// Create keypair
	kp := model.KeyPairCreate("")

	// Data to sign
	data := "NEM is awesome !"

	// Sign data
	sig := kp.Sign(data)

	// Review
	fmt.Println("Public key: ", kp.PublicString())
	fmt.Println("Original data: ", data)
	fmt.Println("Signature: ", utils.Bt2Hex(sig))

	// Result
	fmt.Println()
	if model.Verify(kp.Public, []byte(data), sig) {
		fmt.Println("Signature is valid")
	} else {
		fmt.Println("Signature is invalid")
	}
}
