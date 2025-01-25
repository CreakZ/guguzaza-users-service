package validation

import (
	"fmt"
	"guguzaza-users/domain/entities"
	"strings"
)

func isLatin(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isCyrillic(r rune) bool {
	isSmall := r >= 'а' && r <= 'я'
	isCapital := r >= 'А' && r <= 'Я'
	isWeirdE := r == 'ё' || r == 'Ё' // 'ё' ('Ё') letter check :)

	return isSmall || isCapital || isWeirdE
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// NICKNAME REQUIREMENTS:
// Nickname should be 2 to 50 symbols long
// and contain only latin, cyrillic and '+-=_:;().,' special symbols

const nicknameSpecSymbols = "+-=_:;().,"

func isNicknameSpecSymbol(r rune) bool {
	return strings.ContainsRune(nicknameSpecSymbols, r)
}

// nicknameRequirementsFunc returns true if r breaks nickname requirements
func nicknameRequirementsFunc(r rune) bool {
	validRune := isLatin(r) || isCyrillic(r) || isDigit(r) || isNicknameSpecSymbol(r)
	return !validRune
}

func checkNicknameValidity(nickname string) (valid bool, errMsg string) {
	if len(nickname) < 2 || len(nickname) > 50 {
		return false, fmt.Sprintf("длина никнейма должна составлять от 2 до 50 символов, длина текущего: %d", len(nickname))
	}

	invalid := strings.ContainsFunc(nickname, nicknameRequirementsFunc)
	if invalid {
		return false, "никнейм должен содержать только символы латиницы, кириллицы и символы из следующего набора: +-_:;()"
	}

	return true, ""
}

// PASSWORD REQUIREMENTS:
// Password should be 8 to 50 symbols long
// and contain only latin symbols, at least 1 digit
// and at least 1 symbol from '!@_-='

const passSpecSymbols = "!@_-="

func isPassSpecSymbol(r rune) bool {
	return strings.ContainsRune(passSpecSymbols, r)
}

func passwordRequirementsFunc(r rune) bool {
	validRune := isLatin(r) || isDigit(r) || isPassSpecSymbol(r)
	return !validRune
}

func checkPasswordValidity(password string) (valid bool, errMsg string) {
	if len(password) < 8 || len(password) > 50 {
		return false, fmt.Sprintf("длина пароля должна составлять от 8 до 50 символов, длина текущего: %d", len(password))
	}

	var (
		hasLatin      = false
		hasDigit      = false
		hasSpecSymbol = false
	)

	invalid := strings.ContainsFunc(password, passwordRequirementsFunc)
	if invalid {
		return false, "пароль должен содержать только латинские буквы, цифры и символы из набора: !@_-="
	}

	for _, r := range password {
		if isLatin(r) {
			hasLatin = true
		}
		if isDigit(r) {
			hasDigit = true
		}
		if isPassSpecSymbol(r) {
			hasSpecSymbol = true
		}
	}

	if !hasLatin {
		return false, "пароль не содержит латинских символов"
	}

	if !hasDigit {
		return false, "пароль не содержит цифры"
	}

	if !hasSpecSymbol {
		return false, "пароль не содержит специальные символы из набора: !@_-="
	}

	return true, ""
}

// SEX REQUIREMENTS:
// sex should be "a" to "f" (lexicographically) string

func checkSexValidity(sex string) (valid bool, errMsg string) {
	switch sex {
	case "a", "b", "c", "d", "e", "f":
		return true, ""
	default:
		return false, fmt.Sprintf("неверное значение пола: %s", sex)
	}
}

// ABOUT REQUIREMENTS:
// about length should be less or equal than 100 characters

func checkAboutValidity(about string) (valid bool, errMsg string) {
	if len(about) > 100 {
		return false, fmt.Sprintf("длина поля 'О себе' не должна превышать 100 символов, длина текущего: %d", len(about))
	}

	return true, ""
}

func validateCredentials(nickname, password string) (valid bool, errMsg string) {
	valid, errMsg = checkNicknameValidity(nickname)
	if !valid {
		return false, errMsg
	}

	return checkPasswordValidity(password)
}

func validateSexAbout(sex, about string) (valid bool, errMsg string) {
	valid, errMsg = checkSexValidity(sex)
	if !valid {
		return false, errMsg
	}

	return checkAboutValidity(about)
}

func ValidateMemberCreate(member entities.MemberCreate) (valid bool, errMsg string) {
	valid, errMsg = validateCredentials(member.Nickname, member.Password)
	if !valid {
		return false, errMsg
	}

	return validateSexAbout(member.Sex, member.About)
}

func ValidateAdminCreate(admin entities.AdminCreate) (valid bool, errMsg string) {
	return validateCredentials(admin.Nickname, admin.Password)
}
