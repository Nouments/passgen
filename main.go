package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func GeneratePassword(length int) (string, error) {
	if length < 8 {
		return "", fmt.Errorf("password length too short")
	}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789&#!:/;.,?*%%=+_-@$^([]{})"
	password := make([]byte, length)
	for i:=0;i<length;i++{
		index,_:=rand.Int(rand.Reader,big.NewInt(int64(len(chars))))
		password[i]=chars[index.Int64()]
	}
	return  string(password),nil
}

func main() {
	 length:=flag.Int("l",8,"length")
	 flag.Parse()
	pwd, err:= GeneratePassword(*length)
	if err!=nil{
		fmt.Println(err)
		return 
	}

	fmt.Println("Generated password:", pwd)
}
