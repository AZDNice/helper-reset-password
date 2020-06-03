package password

import "github.com/sethvargo/go-password/password"

func GeneratePassword() (string, error) {
	pass, err := password.Generate(32, 8, 8, false, false)
	if err != nil {
		return "", err
	}

	return pass, nil
}
