// parse
package main

import (
	"fmt"
	"strconv"
)

const (
	TOKEN_NUM = iota
	TOKEN_D
	TOKEN_X
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_END
	TOKEN_BESTOF
	TOKEN_INVALID
)

type token int

const (
	PARSE_BEGIN = iota
	PARSE_COMPLETE
	PARSE_CONTINUE
	PARSE_INVALID
)

type state int

type parseData struct {
	request  string
	curPos   int
	curState state
}

func charToToken(char string) token {
	var aToken token = TOKEN_INVALID

	if len(char) != 1 {
		fmt.Printf("Error: char has length of %d", len(char))
		return aToken
	}

	isNumeric := false
	if _, err := strconv.Atoi(char); err == nil {
		isNumeric = true
	}

	switch {
	case char == "+":
		aToken = TOKEN_PLUS
	case char == "-":
		aToken = TOKEN_MINUS
	case char == "d" || char == "D":
		aToken = TOKEN_D
	case char == "x" || char == "X":
		aToken = TOKEN_X
	case char == ",":
		aToken = TOKEN_BESTOF
	case isNumeric:
		aToken = TOKEN_NUM
	}

	return aToken
}

func (parse *parseData) takeToken() token {
	var curToken token = TOKEN_INVALID

	if parse.curPos < len(parse.request) {
		curToken = charToToken(parse.request[parse.curPos : parse.curPos+1])
		parse.curPos++
	} else {
		curToken = TOKEN_END
	}

	return curToken
}

func (parse *parseData) rewind() {
	parse.curPos--
}

func (parse parseData) peekToken() token {
	var nextToken token = TOKEN_INVALID

	if parse.curPos+1 <= len(parse.request) {
		nextToken = charToToken(parse.request[parse.curPos : parse.curPos+1])
	} else {
		nextToken = TOKEN_END
	}

	return nextToken
}

func (parse parseData) peekNextToken(index int) token {
	var nextToken token = TOKEN_INVALID

	if parse.curPos+index < len(parse.request) {
		nextToken = charToToken(parse.request[parse.curPos+index : parse.curPos+index+1])
	} else {
		nextToken = TOKEN_END
	}

	return nextToken
}

func (parse *parseData) takeNum() int64 {
	numCount := 0
	for parse.peekNextToken(numCount) == TOKEN_NUM {
		numCount++
	}

	foundNumString := parse.request[parse.curPos : parse.curPos+numCount]

	foundNum, _ := strconv.ParseInt(foundNumString, 10, 64)
	parse.curPos += numCount

	return foundNum
}

func Parse(request string) (spec *RollSpec, err error) {
	var parsed RollSpec
	parsed.DieCount = 1
	parsed.Times = 1

	var data parseData
	data.request = request
	data.curPos = 0
	data.curState = PARSE_BEGIN

	for data.curState != PARSE_COMPLETE && data.curState != PARSE_INVALID {
		switch {
		case data.curState == PARSE_BEGIN:
			curToken := data.takeToken()
			switch {
			case curToken == TOKEN_NUM:
				data.rewind()

				nextToken := data.peekNextToken(1)

				if nextToken == TOKEN_D {
					parsed.DieCount = data.takeNum()
					if parsed.BestOf >= parsed.DieCount {
						spec = nil
						err = fmt.Errorf("the best-of value must be smaller than the dice count")
						return
					}
					data.curState = PARSE_CONTINUE
				} else if nextToken == TOKEN_BESTOF {
					if data.peekNextToken(2) == TOKEN_NUM {
						parsed.BestOf = data.takeNum()
						data.takeToken()
					} else {
						spec = nil
						err = fmt.Errorf("when a best-of number is provided, a number must be present after the comma")
						return
					}
				} else {
					spec = nil
					err = fmt.Errorf("when a die count is provided, either the character d must follow the count number or the character , must be present to indicate a best-of number")
					return
				}
			case curToken == TOKEN_D:
				if data.peekToken() != TOKEN_NUM {
					spec = nil
					err = fmt.Errorf("a number must always follow the character d")
					return
				}
				parsed.Sides = data.takeNum()
				data.curState = PARSE_CONTINUE
			default:
				spec = nil
				err = fmt.Errorf("unsupported format: %s", request)
				return
			}
		case data.curState == PARSE_CONTINUE:
			curToken := data.takeToken()
			switch {
			case curToken == TOKEN_D:
				if data.peekToken() != TOKEN_NUM {
					spec = nil
					err = fmt.Errorf("a number must always follow the character d")
					return
				}
				parsed.Sides = data.takeNum()
				data.curState = PARSE_CONTINUE
			case curToken == TOKEN_PLUS || curToken == TOKEN_MINUS:
				if data.peekToken() != TOKEN_NUM {
					spec = nil
					err = fmt.Errorf("plus or minus must always be followed by a number")
					return
				}
				isPlus := curToken == TOKEN_PLUS
				modQuant := data.takeNum()
				if !isPlus {
					modQuant = modQuant * -1
				}
				parsed.Modifier = modQuant
			case curToken == TOKEN_X:
				if data.peekToken() != TOKEN_NUM {
					spec = nil
					err = fmt.Errorf("x must always be followed by a number")
					return
				}
				parsed.Times = data.takeNum()
			case curToken == TOKEN_BESTOF:
				spec = nil
				err = fmt.Errorf("the best-of character is not valid in that position")
				return
			case curToken == TOKEN_END:
				data.curState = PARSE_COMPLETE
			}
		default:

		}
	}

	spec = &parsed
	err = nil
	return
}
