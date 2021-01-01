// Global functions that can be used in diffrent sections of the project
package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// Simple error handler

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Function that will allow us to hash a password using depencides from golang.org/x/crypto/bcrypt
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}
