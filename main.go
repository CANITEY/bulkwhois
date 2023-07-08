package main

import (
	"fmt"
    "flag"
	"github.com/likexian/whois"
	"github.com/likexian/whois-parser"
)

var name int
func init() {
    flag.IntVar(&name, "name", 4, "specify name")
    flag.IntVar(&name, "n", 4, "specify name")
}

func main() {
    fmt.Print("enter domain >> ")
    var domain string
    fmt.Scan(&domain)
    result, err := whois.Whois(domain)
    if err != nil {
        fmt.Println(err)
        return
    }
    rslt, err := whoisparser.Parse(result)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(rslt.Registrant.Organization)
    fmt.Println(name)
}
