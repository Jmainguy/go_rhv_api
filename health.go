package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strings"
    "github.com/Jmainguy/zenoss"
)

func decodeHostsHealth (jsonblob []byte, c Config) (hostHealth map[string]string) {
    var rh RhevHosts
    hostHealth = make(map[string]string)
    json.Unmarshal(jsonblob, &rh)
    for _, host := range rh.Host {
        hostHealth[host.Name] = host.ExternalStatus
    }
    return hostHealth
}

func decodeStorageHealth (jsonblob []byte, c Config) {
    var rsd StorageDomains
    json.Unmarshal(jsonblob, &rsd)
    for _, sd := range rsd.StorageDomain {
        if strings.Compare("ok", sd.ExternalStatus) != 0 {
            msg := fmt.Sprintf("ERROR: %s %s %s", sd.Name, sd.Status, sd.ExternalStatus)
            zmsg := fmt.Sprintf("Storage %s: is %s",sd.Name, sd.ExternalStatus)
            UUID, success := zenoss.CreateAlarm(c.Zenoss_url, c.Zenoss_user, c.Zenoss_password, zmsg, c.Zenoss_rhev_hostname)
            fmt.Printf("%s, %t\n", UUID, success)
            fmt.Println(msg)
        }
        msg := fmt.Sprintf("%s %s %s", sd.Name, sd.Status, sd.ExternalStatus)
        fmt.Println(msg)
    }
}

func decodeClusters (jsonblob []byte) {
    var rc Clusters
    json.Unmarshal(jsonblob, &rc)
    for _, cluster := range rc.Cluster {
        msg := fmt.Sprintf("%s", cluster.Name)
        fmt.Println(msg)
    }
}

func getHostsHealth(domain, token string, c Config) {

    url := fmt.Sprintf("%s/ovirt-engine/api/hosts", domain)
    bear := fmt.Sprintf("Bearer %s", token)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Set("Authorization", bear)
    req.Header.Set("Accept", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    jsonblob, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    hostsHealth := decodeHostsHealth(jsonblob, c)
    for host, health := range hostsHealth {
        if strings.Compare(health, "ok") != 0 {
            msg := fmt.Sprintf("ERROR: %s has health of %s", host, health)
            UUID, success := zenoss.CreateAlarm(c.Zenoss_url, c.Zenoss_user, c.Zenoss_password, msg, c.Zenoss_rhev_hostname)
            fmt.Printf("%s, %t\n", UUID, success)
        }
        msg := fmt.Sprintf("%s %s", host, health)
        fmt.Println(msg)
    }
}

func getStorageHealth(domain, token string, c Config) {

    url := fmt.Sprintf("%s/ovirt-engine/api/storagedomains", domain)
    bear := fmt.Sprintf("Bearer %s", token)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Set("Authorization", bear)
    req.Header.Set("Accept", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    jsonblob, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    decodeStorageHealth(jsonblob, c)
}

func getClusters(domain, token string) {

    url := fmt.Sprintf("%s/ovirt-engine/api/clusters", domain)
    bear := fmt.Sprintf("Bearer %s", token)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Set("Authorization", bear)
    req.Header.Set("Accept", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    jsonblob, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    decodeClusters(jsonblob)
}
