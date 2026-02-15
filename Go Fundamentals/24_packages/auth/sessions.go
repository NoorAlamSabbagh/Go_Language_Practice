package auth

func extractSession() string {
	return "logged In"
}

func GetSession() string {
	return extractSession()
}
