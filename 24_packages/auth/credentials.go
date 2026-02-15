package auth

import "fmt"

//Capital letter and small letter function
//Capital letter outside scope and vice  like LoginWithCredentials
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
