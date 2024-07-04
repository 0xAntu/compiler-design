package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Define patterns for different tokens
	const (
		KEYWORD_PATTERN           = `\b(if|else|byte|int|long|float|double|char|boolean|for|while|return|switch|case|final|do|goto|new|private|public|protected)\b`
		IDENTIFIER_PATTERN        = `\b[a-zA-Z_][a-zA-Z0-9_]*\b`
		CONSTANT_PATTERN          = `\b\d+\b`
		ARITHMETIC_OPERATOR_PATTERN = `[+\-*/%]`
		ASSIGNMENT_OPERATOR_PATTERN = `=`
		LOGICAL_OPERATOR_PATTERN  = `(==|!=|<=|>=|<|>|&&|\|\|)`
		PUNCTUATION_PATTERN       = `[;:,]`
		PARENTHESIS_PATTERN       = `[(){}\[\]]`
	)

	// Combine all patterns
	pattern := regexp.MustCompile(
		KEYWORD_PATTERN + "|" + IDENTIFIER_PATTERN + "|" + CONSTANT_PATTERN + "|" +
			ARITHMETIC_OPERATOR_PATTERN + "|" + LOGICAL_OPERATOR_PATTERN + "|" +
			ASSIGNMENT_OPERATOR_PATTERN + "|" + PUNCTUATION_PATTERN + "|" + PARENTHESIS_PATTERN,
	)

	// Read input from the console
	fmt.Println("Enter the input string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Initialize map to store tokens
	tokens := map[string][]string{
		"Keyword":             {},
		"Identifier":          {},
		"Constant":            {},
		"Arithmetic Operator": {},
		"Logical Operator":    {},
		"Assignment Operator": {},
		"Punctuation":         {},
		"Parenthesis":         {},
	}

	// Find all matches
	matches := pattern.FindAllString(input, -1)
	for _, token := range matches {
		switch {
		case regexp.MustCompile(KEYWORD_PATTERN).MatchString(token):
			tokens["Keyword"] = append(tokens["Keyword"], token)
		case regexp.MustCompile(IDENTIFIER_PATTERN).MatchString(token):
			tokens["Identifier"] = append(tokens["Identifier"], token)
		case regexp.MustCompile(CONSTANT_PATTERN).MatchString(token):
			tokens["Constant"] = append(tokens["Constant"], token)
		case regexp.MustCompile(ARITHMETIC_OPERATOR_PATTERN).MatchString(token):
			tokens["Arithmetic Operator"] = append(tokens["Arithmetic Operator"], token)
		case regexp.MustCompile(LOGICAL_OPERATOR_PATTERN).MatchString(token):
			tokens["Logical Operator"] = append(tokens["Logical Operator"], token)
		case regexp.MustCompile(ASSIGNMENT_OPERATOR_PATTERN).MatchString(token):
			tokens["Assignment Operator"] = append(tokens["Assignment Operator"], token)
		case regexp.MustCompile(PUNCTUATION_PATTERN).MatchString(token):
			tokens["Punctuation"] = append(tokens["Punctuation"], token)
		case regexp.MustCompile(PARENTHESIS_PATTERN).MatchString(token):
			tokens["Parenthesis"] = append(tokens["Parenthesis"], token)
		}
	}

	// Print results
	for tokenType, tokenList := range tokens {
		if len(tokenList) > 0 {
			fmt.Printf("%s (%d): %s\n", tokenType, len(tokenList), strings.Join(tokenList, ", "))
		}
	}
}
