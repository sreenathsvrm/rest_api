package helpers

import (
	"fmt"
	
	"loginPage/db"
	"net/http"
)

func CheckSession(r *http.Request) bool {

	cookie, err := r.Cookie("session")

	if !CheckError(err) {

		if session, ok := db.Sessions[cookie.Value]; ok {

			if !session.Sessionexpired() {
				return true
			}
		}
	}

	return false
}

func CheckError(err error) bool {
	if err != nil {
		fmt.Println("Error found")
		return true
	}
	return false

}


