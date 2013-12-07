package yoga

import (
    "fmt"
    "net/http"

		"appengine"
		"appengine/user"
)

func init() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
}

const embedPage = `
<html>
<head>
<style>
body {
	background-color: black;
}
div#main {
  width: 750px;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
</head>
<body>
<div id="main">
<blockquote class="versal-embedded" data-id="nmnk6x" viewport="tall">
<noscript>
  <a href="https://versal.com/c/nmnk6x/yoga-for-beginners">
    <p>Yoga for Beginners on Versal.com</p>
    <img src="https://stack.versal.com:443/api2/assets/a2aa04df-9ce2-4116-8c52-0be2a2d0ad09">
  </a>
</noscript>
</blockquote>
<script src="https://versal.com/embed.js"></script>
</div>

</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
		 fmt.Fprint(w, embedPage)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
		 c := appengine.NewContext(r)
		 u := user.Current(c)
		 if u == nil {
		 		url, err := user.LoginURL(c, r.URL.String())
				if err != nil {
					 http.Error(w, err.Error(), http.StatusInternalServerError)
					 return
				}
				w.Header().Set("Location", url)
				w.WriteHeader(http.StatusFound)
				return
		 }
		 fmt.Fprintf(w, "Hello, %v!", u)
}