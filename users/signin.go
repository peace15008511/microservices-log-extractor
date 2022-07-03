package users

func Login(username string, password string) string {
	usn := "peace"
	pass := "1234"
	var response string
	if usn == username && pass == password {
		response = "valid pass and username"
	} else {

	}
	return response

}
