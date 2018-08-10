package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/url"
)

func decodeAuth (jsonblob []byte) (accessToken string) {
    var ra RhevAuth
    json.Unmarshal(jsonblob, &ra)
    accessToken = ra.AccessToken
    return
}

func getToken(username, password, domain, apiurl string) (accessCode string) {
    userdomain := fmt.Sprintf("%s@%s", username, domain)

    var Url *url.URL
    Url, _ = url.Parse(apiurl)
    Url.Path += "/ovirt-engine/sso/oauth/token"
    param := url.Values{}
    param.Add("grant_type", "password")
    param.Add("scope", "ovirt-app-api")
    param.Add("username", userdomain)
    param.Add("password", password)
    Url.RawQuery = param.Encode()

    req, err := http.NewRequest("POST", Url.String(), nil)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    jsonblob, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    accessCode = decodeAuth(jsonblob)
    return accessCode
}
