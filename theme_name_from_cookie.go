package dashboard

import (
	"log"
	"net/http"

	"github.com/dracory/dashboard/shared"
)

func ThemeNameRetrieveFromCookie(r *http.Request) string {
	themeNameFromCookie, err := r.Cookie(shared.THEME_COOKIE_KEY)

	if err != nil {
		if err != http.ErrNoCookie {
			log.Println(err.Error())
		}

		return ""

	}

	return themeNameFromCookie.Value
}
