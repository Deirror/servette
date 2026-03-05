package httpx

func IsSuccessStatus(status int) bool {
	return status/100 == 2
}
