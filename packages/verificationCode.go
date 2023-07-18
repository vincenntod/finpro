package packages

import (
	"golang/module/account/entities"
	"math/rand"
	"strconv"
	"time"
)

func GenerateVerificationCode(email string) string {

	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(9000) + 1000)

	entities.VerificationCodes[email] = code
	return code
}
