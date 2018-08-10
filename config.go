package main

import (
    "github.com/ghodss/yaml"
    "io/ioutil"
)

func config() (credents Config){
    config_file, err := ioutil.ReadFile("/opt/rhev/config.yaml")
    check(err)
    yaml.Unmarshal(config_file, &credents)
    return
}

type Config struct {
    Rhev_hosts []string `json:"rhev_hosts"`
    Rhev_user string `json:"rhev_user"`
    Rhev_password string `json:"rhev_password"`
    Rhev_domain string `json:"rhev_domain"`
    Zenoss_url string `json:"zenoss_url"`
    Zenoss_user string `json:"zenoss_user"`
    Zenoss_password string `json:"zenoss_password"`
    Zenoss_rhev_hostname string `json:"zenoss_rhev_hostname"`
}
