package validation

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
)

var (
	latinSmallRt = &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: 'a', Hi: 'z', Stride: 1},
		},
	}

	latinCapitalRt = &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: 'A', Hi: 'Z', Stride: 1},
		},
	}

	latinFullRt = &unicode.RangeTable{
		R16: append(latinSmallRt.R16, latinCapitalRt.R16...),
	}

	cyrillicRt = &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: 0x0410, Hi: 0x042F, Stride: 1},
			{Lo: 0x0430, Hi: 0x044F, Stride: 1},
		},
	}

	specSymbolsRt = &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: '+', Hi: '+', Stride: 1},
			{Lo: '-', Hi: '-', Stride: 1},
			{Lo: '_', Hi: '_', Stride: 1},
			{Lo: ':', Hi: ':', Stride: 1},
			{Lo: ';', Hi: ';', Stride: 1},
			{Lo: '(', Hi: '(', Stride: 1},
			{Lo: ')', Hi: ')', Stride: 1},
		},
	}

	digitsRt = &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: '0', Hi: '9', Stride: 1},
		},
	}

	latinSmallSet  = runes.In(latinSmallRt)
	latinFullSet   = runes.In(latinFullRt)
	cyrillicSet    = runes.In(cyrillicRt)
	specSymbolsSet = runes.In(specSymbolsRt)
	digitsSet      = runes.In(digitsRt)
)

// NICKNAME REQUIREMENTS:
// Nickname should be 2 to 50 symbols long
// and contain only latin, cyrillic and '+-_:;()' special symbols

func nicknameRequirementsFunc(r rune) bool {
	return latinFullSet.Contains(r) || cyrillicSet.Contains(r) || specSymbolsSet.Contains(r)
}

func CheckNicknameValidity(nickname string) (valid bool, errMsg string) {
	if len(nickname) < 2 || len(nickname) > 50 {
		return false, fmt.Sprintf("длина никнейма должна составлять от 2 до 50 символов, длина текущего: %d", len(nickname))
	}

	valid = strings.ContainsFunc(nickname, nicknameRequirementsFunc)
	if !valid {
		return false, "никнейм должен содержать только символы латиницы, кириллицы и символы из следующего набора: +-_:;()"
	}

	return true, ""
}

// PASSWORD REQUIREMENTS:
// Password should be 8 to 50 symbols long
// and contain only latin symbols, at least 1 of which is uppercase, and at least 1 digit

func passwordRequirementsFunc(r rune) bool {
	return latinSmallSet.Contains(r) || digitsSet.Contains(r)
}

func CheckPasswordValidity(password string) (valid bool, errMsg string) {
	if len(password) < 8 || len(password) > 50 {
		return false, fmt.Sprintf("длина пароля должна составлять от 8 до 50 символов, длина текущего: %d", len(password))
	}

	if !strings.ContainsFunc(password, unicode.IsUpper) {
		return false, "пароль должен содержать хотя бы одну заглавную латинскую букву"
	}

	return strings.ContainsFunc(password, passwordRequirementsFunc), ""
}

// SEX REQUIREMENTS:
// sex should be 'm' or 'f' char

func CheckSexValidity(sex byte) bool {
	return sex == 'm' || sex == 'f'
}
