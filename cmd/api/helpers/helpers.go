package helpers

import (
	"fmt"
	"os"
)

const (
	JWT_SECRET = "JWT_SECRET"
	JWT_ALGO   = "JWT_ALGO"
)

func LookEnv(name string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Errorf("environment variable not found : [%s]", name))
	}
	return v
}

func GetJwtSecret() string {
	return LookEnv(JWT_SECRET)
}

func GetJwtAlgorithm() string {
	return LookEnv(JWT_ALGO)
}
