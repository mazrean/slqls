// Code generated by "stringer -type Kind kind.go"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SQLKeyword-0]
	_ = x[Number-1]
	_ = x[Char-2]
	_ = x[SingleQuotedString-3]
	_ = x[NationalStringLiteral-4]
	_ = x[Comma-5]
	_ = x[Whitespace-6]
	_ = x[Comment-7]
	_ = x[Eq-8]
	_ = x[Neq-9]
	_ = x[Lt-10]
	_ = x[Gt-11]
	_ = x[LtEq-12]
	_ = x[GtEq-13]
	_ = x[Plus-14]
	_ = x[Minus-15]
	_ = x[Mult-16]
	_ = x[Div-17]
	_ = x[Caret-18]
	_ = x[Mod-19]
	_ = x[LParen-20]
	_ = x[RParen-21]
	_ = x[Period-22]
	_ = x[Colon-23]
	_ = x[DoubleColon-24]
	_ = x[Semicolon-25]
	_ = x[Backslash-26]
	_ = x[LBracket-27]
	_ = x[RBracket-28]
	_ = x[Ampersand-29]
	_ = x[LBrace-30]
	_ = x[RBrace-31]
	_ = x[ILLEGAL-32]
}

const _Kind_name = "SQLKeywordNumberCharSingleQuotedStringNationalStringLiteralCommaWhitespaceCommentEqNeqLtGtLtEqGtEqPlusMinusMultDivCaretModLParenRParenPeriodColonDoubleColonSemicolonBackslashLBracketRBracketAmpersandLBraceRBraceILLEGAL"

var _Kind_index = [...]uint8{0, 10, 16, 20, 38, 59, 64, 74, 81, 83, 86, 88, 90, 94, 98, 102, 107, 111, 114, 119, 122, 128, 134, 140, 145, 156, 165, 174, 182, 190, 199, 205, 211, 218}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
