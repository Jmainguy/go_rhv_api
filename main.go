package main

import (
    "net/http"
    "crypto/tls"
    "fmt"
    "time"
)

func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func main() {


    c := config()
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    for {
        for _, rhevHost := range c.Rhev_hosts {
            fmt.Println(rhevHost)
            accessCode := getToken(c.Rhev_user, c.Rhev_password, c.Rhev_domain, rhevHost)
            fmt.Println("Hosts health")
            getHostsHealth(rhevHost, accessCode, c)
            fmt.Println("Storage domains")
            getStorageHealth(rhevHost, accessCode, c)
        }
        fmt.Println("Sleepy bye bye time")
        time.Sleep(60 * time.Second)
    }
}
