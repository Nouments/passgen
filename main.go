package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
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

func ValidateInput(input string) (int,error){
	val:= strings.TrimSpace(input)
	intval,err:=strconv.Atoi(val)
	if err!=nil{
		return 0,fmt.Errorf("make sure input is integer")
	}
	return intval,nil
}
func main() {

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Entered password lenth: ")
	input,_:= reader.ReadString('\n')

	length,err:=ValidateInput(input)
	if err!=nil{
		log.Fatal(err)
	}

	pwd, err:= GeneratePassword(length)
	if err!=nil{
		fmt.Println(err)
		return 
	}

	fmt.Println("Generated password:", pwd)
}