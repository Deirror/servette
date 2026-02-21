package translation

import (
	"github.com/Deirror/servette/translation/languages"
)

// Wrapper struct, containing all needed data for translation.
type Translator struct {
	Bundle  *Bundle
	Resolver *languages.Resolver
}

func New(b *Bundle, r *languages.Resolver) *Translator {
	return &Translator{
		Bundle:  b,
		Resolver: r,
	}
}
