import re

# Define patterns for different tokens
KEYWORD_PATTERN = r'\b(if|else|byte|int|long|float|double|char|boolean|for|while|return|switch|case|final|do|goto|new|private|public|protected)\b'
IDENTIFIER_PATTERN = r'\b[a-zA-Z_][a-zA-Z0-9_]*\b'
CONSTANT_PATTERN = r'\b\d+\b'
ARITHMETIC_OPERATOR_PATTERN = r'[+\-*/%]'
ASSIGNMENT_OPERATOR_PATTERN = r'='
LOGICAL_OPERATOR_PATTERN = r'(==|!=|<=|>=|<|>|&&|\|\|)'
PUNCTUATION_PATTERN = r'[;:,]'
PARENTHESIS_PATTERN = r'[(){}\[\]]'

# Combine all patterns
combined_pattern = re.compile(
    f"({KEYWORD_PATTERN})|({IDENTIFIER_PATTERN})|({CONSTANT_PATTERN})|"
    f"({ARITHMETIC_OPERATOR_PATTERN})|({LOGICAL_OPERATOR_PATTERN})|"
    f"({ASSIGNMENT_OPERATOR_PATTERN})|({PUNCTUATION_PATTERN})|({PARENTHESIS_PATTERN})"
)

def main():
    # Read input from the console
    input_string = input("Enter the input string: ")

    # Initialize dictionary to store tokens
    tokens = {
        "Keyword": [],
        "Identifier": [],
        "Constant": [],
        "Arithmetic Operator": [],
        "Logical Operator": [],
        "Assignment Operator": [],
        "Punctuation": [],
        "Parenthesis": [],
    }

    # Find all matches
    matches = combined_pattern.findall(input_string)

    # Flatten the list of tuples to a list of strings
    matches = [token for match in matches for token in match if token]

    for token in matches:
        if re.fullmatch(KEYWORD_PATTERN, token):
            tokens["Keyword"].append(token)
        elif re.fullmatch(IDENTIFIER_PATTERN, token):
            tokens["Identifier"].append(token)
        elif re.fullmatch(CONSTANT_PATTERN, token):
            tokens["Constant"].append(token)
        elif re.fullmatch(ARITHMETIC_OPERATOR_PATTERN, token):
            tokens["Arithmetic Operator"].append(token)
        elif re.fullmatch(LOGICAL_OPERATOR_PATTERN, token):
            tokens["Logical Operator"].append(token)
        elif re.fullmatch(ASSIGNMENT_OPERATOR_PATTERN, token):
            tokens["Assignment Operator"].append(token)
        elif re.fullmatch(PUNCTUATION_PATTERN, token):
            tokens["Punctuation"].append(token)
        elif re.fullmatch(PARENTHESIS_PATTERN, token):
            tokens["Parenthesis"].append(token)

    # Print results
    for token_type, token_list in tokens.items():
        if token_list:
            print(f"{token_type} ({len(token_list)}): {', '.join(token_list)}")

if __name__ == "__main__":
    main()
