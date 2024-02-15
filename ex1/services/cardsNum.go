package services

import (
	"errors"
	"strconv"
	"strings"
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