package services

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ValidationNum(Phone string) error{
	
	if len(strings.ReplaceAll(Phone,"+",""))!=12{
		return errors.New("Error with Number!!")
	}
	_,err:=strconv.Atoi(strings.ReplaceAll(Phone,"+",""))
	if err!=nil{
		return errors.New("Error with Number!!")
	}
	return nil
}

func CreateTranId() string{

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö" +
    "0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // Например "ExcbsVQs"
	return str
}
