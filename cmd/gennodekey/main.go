package main

import (
	"encoding/hex"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"fmt"
)

func main() {
	prv, _ := crypto.GenerateKey()
	buf := crypto.FromECDSA(prv)
	s := hex.EncodeToString(buf)
	fmt.Println("privatekey:", s)
	fmt.Println("address:", crypto.PubkeyToAddress(prv.PublicKey).Hex())
}
