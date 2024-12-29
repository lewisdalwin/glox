// This package defines the different categories for our tokens
package token

type TokenType int 

const (
	// special tokens
	ILLEGAL TokenType = iota
	EOF
	COMMENT
	// single-character token
	LEFT_PAREN  	// (
	RIGHT_PAREN		// )
	LEFT_BRACE		// {
	RIGHT_BRACE		// }
	COMMA			// ,
	DOT 			// .
	MINUS 			// -
	PLUS 			// +
	SEMICOLON		// ;
	FORWARD_SLASH	// /
	STAR
	// one or more characters
	BANG 			// !
	BANG_EQUAL		// !=
	EQUAL 			// =
	EQUAL_EQUAL 	// ==
	GREATER_THAN	// > 
	GREATER_THAN_EQUAL 	// >=
	LESS_THAN		// <
	LESS_THAN_EQUAL 	// <=
	// literals
	IDENTIFIER 			// main
	STRING 				// "Belize"
	NUMBER 				// 44, 44.56
	// keywords
	AND 
	CLASS 
	ELSE 
	FALSE 
	FUN 
	FOR 
	IF 
	NIL 
	OR 
	PRINT 
	RETURN 
	SUPER 
	THIS 
	TRUE 
	VAR 
	WHILE 
)

