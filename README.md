# BULKWOIS
is a script that you can feed a domains list and it will return the organization name for every domain or filter domains for a specific organization name


## installation

to install the binary in $GOPATH use this command

```
$ go install github.com/canitey/bulkwhois
```

## usage
```
$ bulkwhois -h                                                                                                                                    

██████╗ ██╗   ██╗██╗     ██╗  ██╗██╗    ██╗██╗  ██╗ ██████╗ ██╗███████╗
██╔══██╗██║   ██║██║     ██║ ██╔╝██║    ██║██║  ██║██╔═══██╗██║██╔════╝
██████╔╝██║   ██║██║     █████╔╝ ██║ █╗ ██║███████║██║   ██║██║███████╗
██╔══██╗██║   ██║██║     ██╔═██╗ ██║███╗██║██╔══██║██║   ██║██║╚════██║
██████╔╝╚██████╔╝███████╗██║  ██╗╚███╔███╔╝██║  ██║╚██████╔╝██║███████║
╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝╚══════╝
BY: CANITEY 
Usage:
  bulkwhois [OPTIONS]

Application Options:
  -f, --file=FILE              specify file to "whois" it
  -o, --output=FILE            specify file to save the output
  -a, --all                    when this flag is used the script return every query with Organization name associated with it
  -r, --org-name="ORG NAME"    specify Organization name to return domains that match it

Help Options:
  -h, --help                   Show this help message
```
**EXMAPLES**
`cat domains.txt | bulkwhois -a -o output.txt`
`cat damains.txt | bulkwhois -r "dummy org" -o output.txt`
`bulkwhois -f domains.txt -r "dummy org" -o output.txt`
