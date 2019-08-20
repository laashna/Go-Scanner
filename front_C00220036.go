package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

var charClass int
var LEFT_PAREN int = 25
var RIGHT_PAREN int = 26
var ADD_OP int = 21
var SUB_OP int = 22
var MULT_OP int = 23
var DIV_OP int = 24
var EOF int = -1
var nextToken int
var lexeme string
var IDENT int = 11
var INT_LIT int = 10
var LETTER int = 0
var DIGIT int = 1
var UNKNOWN int = 99
var nextChar rune
var (
	reader *bufio.Reader
)

func main() {
	f, err := os.Open("front.in")
	if err != nil {
		panic(err)
		
	}
	defer f.Close() 
	reader = bufio.NewReader(f)
	my_getChar()
	for nextToken != EOF {
		lex()
	}

}

func my_getChar() {
	nc, _, err := reader.ReadRune() 
	nextChar = nc
	if err == io.EOF {
		charClass = EOF
	} else if err != nil {
		charClass = EOF
	}

	if IsLetter(rune(nextChar)) == true {
		charClass = LETTER
	} else if IsDigit(rune(nextChar)) == true {
		charClass = DIGIT
	} else {
		charClass = UNKNOWN
	}

}


func IsLetter(s rune) bool {

	if !unicode.IsLetter(s) {
		return false

	}
	return true
}
func IsDigit(s rune) bool {

	if !unicode.IsDigit(s) {
		return false

	}
	return true
}
func IsSpace(s rune) bool {
	if !unicode.IsSpace(s) {
		return false
	}
	return true
}
func my_addChar() {
	lexeme = lexeme + string(nextChar)
}
func getNonBlank() {
	for IsSpace(nextChar) == true { 
		my_getChar()
	}

}
func lookup(ch rune) {
	switch ch {
	case '(':
		my_addChar()
		nextToken = LEFT_PAREN
	case ')':
		my_addChar()
		nextToken = RIGHT_PAREN
	case '+':
		my_addChar()
		nextToken = ADD_OP
	case '-':
		my_addChar()
		nextToken = SUB_OP
	case '*':
		my_addChar()
		nextToken = MULT_OP
	case '/':
		my_addChar()
		nextToken = DIV_OP

	default:
		my_addChar()
		nextToken = EOF
	}

}
func lex() {
	lexeme = ""
	getNonBlank()
	switch charClass {
	case LETTER:
		my_addChar()
		my_getChar()
		for charClass == LETTER || charClass == DIGIT {
			my_addChar()
			my_getChar()
		}
		nextToken = IDENT
	case DIGIT:
		my_addChar()
		my_getChar()
		for charClass == DIGIT {
			my_addChar()
			my_getChar()
		}
		nextToken = INT_LIT
	case UNKNOWN:
		lookup(nextChar)
		my_getChar()
	case EOF:
		nextToken = EOF
		lexeme = "EOF"
	}
	fmt.Println("Next token is: ", nextToken, "Next lexeme is: ", lexeme)  
}
