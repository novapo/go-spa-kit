package admin

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/novapo/go-spa-kit/config"
)

var store = sessions.NewCookieStore([]byte("very-secure"))

// AdminHandler serves the admin area for spa-kit
func AdminHandler(w http.ResponseWriter, req *http.Request) {
	adminPage, err := ioutil.ReadFile("admin/admin.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	session, _ := store.Get(req, "session-name")
	// Set some session values.
	if session.Values["foo"] == "bar" && session.Values[42] == 43 {
		w.Write(adminPage)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// LoginHandler handles admin login for spa-kit
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		loginPage, err := ioutil.ReadFile("admin/login.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(loginPage)
	}
	if req.Method == "POST" {
		req.ParseForm()

		login := req.PostForm.Get("login")
		password := req.PostForm.Get("password")

		if login == config.AdnConf.Login && password == config.AdnConf.Password {
			session, _ := store.Get(req, "session-name")
			// Set some session values.
			session.Values["foo"] = "bar"
			session.Values[42] = 43
			// Save it.
			session.Save(req, w)
			http.Redirect(w, req, "/admin", 302)
		}
	}
}
