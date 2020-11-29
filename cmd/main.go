package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type Prefix struct {
	IpPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	NetworkBorderGroup string `json:"network_border_group"`
	Service            string `json:"service"`
}

type Ipv6Prefix struct {
	Ipv6Prefix         string `json:"ipv6_prefix"`
	Region             string `json:"region"`
	NetworkBorderGroup string `json:"network_border_group"`
	Service            string `json:"service"`
}

type IpRange struct {
	SyncToken    string       `json:"syncToken"`
	CreateDate   string       `json:"createDate"`
	Prefixes     []Prefix     `json:"prefixes"`
	Ipv6Prefixes []Ipv6Prefix `json:"ipv6_prefixes"`
}

func main() {
	const ipRangeLink string = "https://ip-ranges.amazonaws.com/ip-ranges.json"
	fmt.Println("Downloading the ip range json file ...")
	resp, err := http.Get(ipRangeLink)
	if err != nil {
		panic("Could not download the aws ip-ranges.json file!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Could not parse the aws ip-ranges.json file!")
	}
	var ipRangeJSON IpRange
	json.Unmarshal([]byte(body), &ipRangeJSON)
	fmt.Printf("Found %d IP ranges (including duplicates)\n", len(ipRangeJSON.Prefixes))
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [<ip_addr> | - <domain_name>] \n", os.Args[0])
		return
	}

	var argIP string = os.Args[1]
	var parsedIPs []net.IP
	if argIP == "-" {
		if len(os.Args) < 3 {
			fmt.Printf("usage: %s [<ip_addr> | - <domain_name>] \n", os.Args[0])
			return
		}
		var domainName string = os.Args[2]
		parsedIPs, err = net.LookupIP(domainName)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		parsedIPs = make([]net.IP, 1)
		parsedIPs[0] = net.ParseIP(argIP)
	}
	var found uint = 0
	for _, prefix := range ipRangeJSON.Prefixes {
		_, ipnet, err := net.ParseCIDR(prefix.IpPrefix)
		if err != nil {
			continue
		}
		for _, parsedIP := range parsedIPs {
			if ipnet.Contains(parsedIP) {
				found++
				fmt.Printf("Found address: %s\n", parsedIP)
				json, err := json.MarshalIndent(prefix, "", "  ")
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(json))
				break
			}
		}
	}
	fmt.Printf("\n%d result(s)\n", found)
}
