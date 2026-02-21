package appx

import (
	"github.com/Deirror/servette/env"
)

func IsProdMode(mode string) bool {
	return  mode == env.Prod
}

func IsDevMode(mode string) bool {
	return  mode == env.Dev
}

func IsStageMode(mode string) bool {
	return  mode == env.Staging
}
