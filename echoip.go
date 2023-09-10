package main 

import (
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

    if len(auth) != 2 || auth[0] != "Bearer" {
        fmt.Fprintln(w, "Hello, world")
        return
    }

    bearerToken := os.Getenv("BEARER_TOKEN")

    if auth[1] != bearerToken {
        fmt.Fprintln(w, "Hello, world")
        return
    }

    ip := getIP(r)

    fmt.Fprint(w, ip)
}

// getIP retrieves a correct IP from the request.
// It checks X-Forwarded-For header at first and if it is empty,
// retrieves an IP from RemoteAddr.
func getIP(r *http.Request) string {
    forwarded := r.Header.Get("X-Forwarded-For")
    if forwarded != "" {
        // If X-Forwarded-For Header is present (usually through proxies),
        // use it's first value (client original IP)
        return strings.Split(forwarded, ",")[0]
    }
    return strings.Split(r.RemoteAddr, ":")[0]
}

func logHandler(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        h.ServeHTTP(w, r)

        end := time.Now()
        duration := end.Sub(start)

        ip := strings.Split(r.RemoteAddr, ":")[0]
        
        fmt.Printf("Access log: IP %s - %s %s - %v\n", ip, r.Method, r.URL.Path, duration)
    }
}

func main() {
    http.HandleFunc("/", logHandler(handler))

    port := os.Getenv("PORT")
    if port == "" {
        // Set a default port if there is nothing in the environment
        port = "8080"
    }

    if err := http.ListenAndServe(":" + port, nil); err != nil {
        fmt.Println(err.Error())
    }
}
