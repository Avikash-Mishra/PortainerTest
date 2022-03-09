package parser

import (
	"fmt"
	"testing"
)

// TestValid tests a set of valid inputs
func TestValid(t *testing.T) {
	validCases := []string{"", "()", "[]", "{}", "({[]})", "()()()(){}{}{}{}[][][][]"}
	for _, val := range validCases {
		valid := ParseStructure(val)
		if !valid {
			t.Error(fmt.Sprintf("`%s` expected to be valid structure", val))
		}
	}
}

// TestInvalid tests a set of invalid inputs
func TestInvalid(t *testing.T)  {
	validCases := []string{"[","]","())", "[[]", "()}{}", "({[])", "()(()()(){}{}{}{}[][}][][}]"}
	for _, val := range validCases {
		valid := ParseStructure(val)
		if valid {
			t.Error(fmt.Sprintf("`%s` expected to be invalid structure", val))
		}
	}
}