package main

import (
	"fmt"
	"net/http"

	_ "webapp/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           WebApp API
// @version         1.0
// @description     This is a sample webapp server.
// @host            localhost:8080
// @BasePath        /

// pingHandler godoc
// @Summary      Ping endpoint
// @Description  Returns a simple ping message and request headers
// @Tags         ping
// @Produce      plain
// @Success      200  {string}  string  "Hello I am Ping"
// @Router       /api/ping [get]
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello I am Ping")
	fmt.Fprintln(w, "Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

// advLogoutHandler godoc
// @Summary      Advanced Logout endpoint
// @Description  Returns a logout message and request headers
// @Tags         auth
// @Param        extranetUrl  query     string  false  "Extranet URL"
// @Produce      plain
// @Success      200  {string}  string  "Hello I am advanced logged-out endpoint"
// @Router       /realms/ext/npwl-user/logout [get]
func advLogoutHandler(w http.ResponseWriter, r *http.Request) {
	extranetUrl := r.URL.Query().Get("extranetUrl")
	fmt.Fprintln(w, "Hello I am advanced logged-out endpoint")
	fmt.Fprintf(w, "extranetUrl param: %s\n", extranetUrl)
	fmt.Fprintln(w, "Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

// advConnectHandler godoc
// @Summary      Advanced Connect endpoint
// @Description  Returns a connect message and request headers
// @Tags         auth
// @Param        extranetId  query     string  false  "Extranet ID"
// @Param        p           query     string  false  "p param"
// @Produce      plain
// @Success      200  {string}  string  "Hello I am advanced ping"
// @Router       /realms/ext/protocol/cas/connect [get]
func advConnectHandler(w http.ResponseWriter, r *http.Request) {
	extranetId := r.URL.Query().Get("extranetId")
	p := r.URL.Query().Get("p")
	fmt.Fprintln(w, "Hello I am advanced ping")
	fmt.Fprintf(w, "extranetId param: %s\n", extranetId)
	fmt.Fprintf(w, "p param: %s\n", p)
	fmt.Fprintln(w, "Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func main() {
	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/realms/ext/npwl-user/logout", advLogoutHandler)
	http.HandleFunc("/realms/ext/protocol/cas/connect", advConnectHandler)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/swagger/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	fmt.Println("Server starting on :8080...")
	fmt.Println("Swagger UI available at http://localhost:8080/swagger/index.html")
	fmt.Println("OpenAPI JSON available at http://localhost:8080/swagger/openapi.json")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
