package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"time"
)

func generateTOTP(email string) string {
	secret := email + "HENNGECHALLENGE004"

	counter := time.Now().Unix() / 30

	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(counter))

	h := hmac.New(sha512.New, []byte(secret))
	h.Write(buf[:])
	hash := h.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	binaryCode :=
		(int(hash[offset])&0x7F)<<24 |
			(int(hash[offset+1])&0xFF)<<16 |
			(int(hash[offset+2])&0xFF)<<8 |
			(int(hash[offset+3]) & 0xFF)

	otp := binaryCode % 10000000000

	return fmt.Sprintf("%010d", otp)
}

func main() {
	email := "hamza.siddiqui4747@gmail.com"
	res := generateTOTP(email)

	fmt.Println(res)
}
