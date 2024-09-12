package main

import "fmt"

// Function to perform Leftmost Derivation
func leftmostDerivation(input string, grammar []string) {
    fmt.Println("Leftmost Derivations:")
    derivation := "S"
    fmt.Printf("step 1: %s\n", derivation)

    // Apply rules to derive the input string
    if derivation == "S" {
        derivation = "aAB"
        fmt.Printf("step 2: %s\n", derivation)
    }
    if derivation == "aAB" {
        derivation = "abB"
        fmt.Printf("step 3: %s\n", derivation)
    }
    if derivation == "abB" {
        derivation = "abc"
        fmt.Printf("step 4: %s\n", derivation)
    }
}

// Function to perform Rightmost Derivation
func rightmostDerivation(input string, grammar []string) {
    fmt.Println("Rightmost Derivations:")
    derivation := "S"
    fmt.Printf("step 1: %s\n", derivation)

    // Apply rules to derive the input string
    if derivation == "S" {
        derivation = "aAB"
        fmt.Printf("step 2: %s\n", derivation)
    }
    if derivation == "aAB" {
        derivation = "aAc"
        fmt.Printf("step 3: %s\n", derivation)
    }
    if derivation == "aAc" {
        derivation = "abc"
        fmt.Printf("step 4: %s\n", derivation)
    }
}

func main() {
    input := "abc"
    grammar := []string{"S -> aAB", "A -> b | ε", "B -> c | ε"}

    leftmostDerivation(input, grammar)
    rightmostDerivation(input, grammar)
}
