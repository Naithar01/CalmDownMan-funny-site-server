package action

import "golang.org/x/crypto/bcrypt"

func UserRegisterHashPassWord(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), err

}
