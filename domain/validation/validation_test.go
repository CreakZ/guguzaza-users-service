package validation_test

import (
	"guguzaza-users/domain/validation"
	"testing"
)

type testStr struct {
	value string
	valid bool
}

func TestCheckNicknameValidity(t *testing.T) {
	tests := []testStr{
		{
			value: "фимоз_1337",
			valid: true,
		},
		{
			value: "victor",
			valid: true,
		},
		{
			value: "a",
			valid: false,
		},
		{
			value: "",
			valid: false,
		},
		{
			value: "-_-",
			valid: true,
		},
		{
			value: "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef",
			valid: false,
		},
		{
			value: "Creakz",
			valid: true,
		},
		{
			value: "Look, I have whitespace",
			valid: false,
		},
	}

	for _, tt := range tests {
		valid, _ := validation.CheckNicknameValidity(tt.value)
		if valid != tt.valid {
			t.Errorf("ошибка на никнейме \"%s\": ожидалось %t, на деле %t \n", tt.value, tt.valid, valid)
		}
	}
}

func TestCheckPasswordValidity(t *testing.T) {
	tests := []testStr{
		{
			value: "",
			valid: false,
		},
		{
			value: "very_good_password",
			valid: false,
		},
		{
			value: "Very_good_password",
			valid: false,
		},
		{
			value: "Very_goodpassw0rd",
			valid: true,
		},
	}

	for _, tt := range tests {
		valid, errMsg := validation.CheckPasswordValidity(tt.value)
		if valid != tt.valid {
			t.Errorf("ошибка на пароле \"%s\": ожидалось %t, на деле %t \n", tt.value, tt.valid, valid)
			if !valid {
				t.Logf("пароль неверен: %s\n", errMsg)
			}
		}
	}
}

func TestCheckSexValidity(t *testing.T) {
	tests := []testStr{
		{
			value: "a",
			valid: true,
		},
		{
			value: "b",
			valid: true,
		},
		{
			value: "f",
			valid: true,
		},
		{
			value: "g",
			valid: false,
		},
	}

	for _, tt := range tests {
		valid, errMsg := validation.CheckSexValidity(tt.value)
		if valid != tt.valid {
			t.Errorf("ошибка на \"%s\": ожидалось %t, на деле %t \n", tt.value, tt.valid, valid)
			if !valid {
				t.Logf("значение пола неверно: %s\n", errMsg)
			}
		}
	}
}

func TestCheckAboutValidity(t *testing.T) {
	tests := []testStr{
		{
			value: "Я РЕАЛЬНО КРУТОЙ Я ПРЯМ ОТВЕЧАЮ 100%",
			valid: true,
		},
		{
			value: "",
			valid: true,
		},
		{
			value: "аааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааа",
			valid: false,
		},
	}

	for _, tt := range tests {
		valid, errMsg := validation.CheckAboutValidity(tt.value)
		if valid != tt.valid {
			t.Errorf("ошибка на \"%s\": ожидалось %t, на деле %t \n", tt.value, tt.valid, valid)
			if !valid {
				t.Logf("значение \"О себе\" неверно: %s\n", errMsg)
			}
		}
	}
}
