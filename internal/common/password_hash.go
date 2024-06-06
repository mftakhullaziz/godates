package common

import "golang.org/x/crypto/bcrypt"

func HashingPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

func ComparedPassword(hashedPwd string, plainPassword []byte) (bool, error) {
	//log := LoggerParent()
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
