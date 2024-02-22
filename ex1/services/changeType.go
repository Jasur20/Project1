package services

import "strings"

func TypeChanger(resp string) {
	req:=RemoveNonXMLCharacters(resp)

}

func RemoveNonXMLCharacters(xmlSTR string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x13, 0x10, 0x09:
			return -1
		default:
			return r
		}
	}, xmlSTR)
}