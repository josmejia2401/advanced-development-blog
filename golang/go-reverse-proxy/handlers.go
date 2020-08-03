package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*Realiza el envío de la petición al backend destino*/
func ServeReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", url.Host)
	r.Host = url.Host
	proxy.ServeHTTP(w, r)
	return
}

/*Busca el HOST dentro de los backend permitidos*/
func GetNewProxyUrl(urlSource *url.URL, host string) string {
	var newServer, _ = TARGET_HOST_ALLOW[host]
	if newServer != "" {
		return newServer
	} else {
		return ""
	}
}

/* Redirecciona la página 404 - Página no encontrada*/
func RedirectNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://soursop-dev.blogspot.com/404.html", http.StatusSeeOther)
	return
}

func HandleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	Log.Println("Llega la petición", r.Host, r.URL.String())
	url := GetNewProxyUrl(r.URL, r.Host)
	if url == "" {
		Log.Println("Host no existe " + r.URL.String())
		RedirectNotFound(w, r)
		return
	} else {
		Log.Println("Se redirecciona a ", url)
		ServeReverseProxy(url, w, r)
		return
	}
}
