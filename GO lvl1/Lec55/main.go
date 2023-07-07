package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type OTP struct {
	Code      string
	CodeSpace string
}

func (o *OTP) GenerateCode(min, max int64) OTP {

	bg := big.NewInt(max - min)
	bigMin := big.NewInt(min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}
	n = n.Add(n, bigMin)

	code := n.String()

	codeSpace := code
	if len(code) > 5 {
		codeSpace = code[:3] + " " + code[3:]
	}

	return OTP{
		Code:      code,
		CodeSpace: codeSpace,
	}
}

func main() {
	var otpCode OTP
	otpCode = otpCode.GenerateCode(1000000000, 9999999999)
	fmt.Println(otpCode.Code)
}
