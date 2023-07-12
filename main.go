package main

import (
	"bufio"
	"fmt"
	flag "github.com/jessevdk/go-flags"
	"github.com/likexian/whois"
	"github.com/likexian/whois-parser"
	"os"
	"strings"
)

const VERSION = "v1.2"

// these are the parameters
var opts struct {
	File      string `short:"f" long:"file" description:"specify file to \"gorg\" it" value-name:"FILE"`
	Out       string `short:"o" long:"output" description:"specify file to save the output" value-name:"FILE"`
	All       bool   `short:"a" long:"all" description:"when this flag is used the script return every query with Organization name associated with it"`
	Version   bool   `short:"v" long:"version" description:"print the script version"`
	Org_name  string `short:"r" long:"org-name" description:"specify Organization name to return domains that match it" value-name:"\"ORG NAME\""`
	Org_email string `short:"e" long:"org-email" description:"specify Registrant name to return domains that match it" value-name:"\"example@mail.com\""`
}

func all(domain string) string {
	domain = strings.TrimSpace(domain)
	query, err := whois.Whois(domain)
	if err != nil {
		return ""
	}
	parsed, err := whoisparser.Parse(query)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v\t\t%v\n", domain, parsed.Registrant.Organization)
}

func returnOrgName(domain string, orgName string) string {
	domain = strings.TrimSpace(domain)
	query, err := whois.Whois(domain)
	if err != nil {
		return ""
	}
	parsed, err := whoisparser.Parse(query)
	if err != nil {
		return ""
	}
	if parsed.Registrant == nil {
		return ""
	}
	if parsed.Registrant.Organization == orgName {
		return fmt.Sprintf("[] %v\n", domain)
	} else if parsed.Registrant.Organization != orgName || parsed.Registrant == nil {
		return ""
	} else {
		return ""
	}
}

func returnOrgEmail(domain string, orgEmail string) string {
	domain = strings.TrimSpace(domain)
	query, err := whois.Whois(domain)
	if err != nil {
		return ""
	}
	parsed, err := whoisparser.Parse(query)
	if err != nil {
		return ""
	}
	if parsed.Registrant == nil {
		return ""
	}
	if parsed.Registrant.Email == orgEmail {
		return fmt.Sprintf("[] %v\n", domain)
	} else if parsed.Registrant.Organization != orgEmail || parsed.Registrant == nil {
		return ""
	} else {
		return ""
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	fmt.Println(`
 ██████╗  ██████╗ ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗██╔══██╗██╔════╝ 
██║  ███╗██║   ██║██████╔╝██║  ███╗
██║   ██║██║   ██║██╔══██╗██║   ██║
╚██████╔╝╚██████╔╝██║  ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝      BY: CANITEY `)
	// this part parses the arguments
	_, err := flag.Parse(&opts)
	if err != nil {
		return
	}
	if !opts.All && opts.Org_name == "" && !opts.Version {
		fmt.Println("use -h/--h flag to print the help message")
		return
	}
    if opts.Version {
        fmt.Printf("Version is %v\n", VERSION)
        return
    }
	// this part for file/stdin reading
	var f *os.File
	if opts.File != "" {
		f, err = os.Open(opts.File)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer f.Close()
	} else {
		f = os.Stdin
		defer f.Close()
	}
	if opts.All && opts.Org_name != "" && opts.Org_email != "" {
		fmt.Println("ERROR: Arguments collision")
		return
	}

	var outFile *os.File
	if opts.Out != "" {
		outFile, err = os.Create(opts.Out)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer outFile.Close()
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if opts.All {
			query := all(scanner.Text())
			fmt.Print(query)
			if opts.Out != "" {
				outFile.WriteString(query)
			}
		} else if opts.Org_name != "" {
			query := returnOrgName(scanner.Text(), opts.Org_name)
			fmt.Print(query)
			if opts.Out != "" {
				outFile.WriteString(query)
			}
		} else if opts.Org_email != "" {
			query := returnOrgEmail(scanner.Text(), opts.Org_email)
			fmt.Print(query)
			if opts.Out != "" {
				outFile.WriteString(query)
			}
		}
	}
}
