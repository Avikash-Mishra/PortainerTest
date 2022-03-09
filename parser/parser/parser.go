package parser

import (
	"strings"
)

// ParseStructure takes in a string and tries to parse the structure
// returns bool
func ParseStructure(input string) bool {
	// stores all open brackets
	var bracketStack []string
	chars := strings.Split(input, "")

	// loops through the characters adding the open brackets onto the stack
	// while also popping off valid closing brackets
	for _, char := range chars {
		// check if open bracket
		if "(" == char || "[" == char || "{" == char {
			bracketStack = append(bracketStack, char)
			continue
		}

		// there is nothing on the stack to check against so return false
		length := len(bracketStack)
		if bracketStack == nil || length == 0 {
			return false
		}

		// check if valid close bracket
		lastOpenBracket := bracketStack[length-1]
		switch char {
		case ")":
			if "(" != lastOpenBracket {
				return false
			}
			break
		case "]":
			if "[" != lastOpenBracket {
				return false
			}
			break
		case "}":
			if "{" != lastOpenBracket {
				return false
			}
			break
		}

		// pop off the stack
		bracketStack = bracketStack[:length-1]
	}

	// if there is nothing left on the stack we have successfully validated the structure
	if len(bracketStack) == 0 {
		return true
	}

	return false
}
