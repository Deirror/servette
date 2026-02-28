package languages

const (
	LangKey = "lang"
)

type RequestType int

// Enum constants for deciding which language resolver func to call.
const (
	FromCookie RequestType = iota
	FromURL 
)
