package main

import (
    "log"
    "net"
    "fmt"
    "time"
    "net/http"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func handler(w http.ResponseWriter, r *http.Request) {
    logIt(r)
    fmt.Fprintf(w, GetLocalIP())
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    logIt(r)
    fmt.Fprintf(w, time.Now().String()) 
}

func logIt(r *http.Request) {
    log.Println(r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/time", timeHandler)

    log.Println("Started Webby!")
    log.Println("Listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
