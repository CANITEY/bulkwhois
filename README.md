# GORG
Is a script that you can feed a domains list and it will return the organization name for every domain or filter domains for a specific organization name.


## Installation

To install the binary in $GOPATH use this command:

```
$ go install github.com/canitey/gorg@1.2
```

## Usage
```
$ gorg -h                                                                                                                                    

 ██████╗  ██████╗ ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗██╔══██╗██╔════╝ 
██║  ███╗██║   ██║██████╔╝██║  ███╗
██║   ██║██║   ██║██╔══██╗██║   ██║
╚██████╔╝╚██████╔╝██║  ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝      BY: CANITEY 
Usage:
  main [OPTIONS]

Application Options:
  -f, --file=FILE              specify file to "gorg" it
  -o, --output=FILE            specify file to save the output
  -a, --all                    when this flag is used the script return every query with Organization name associated with it
  -r, --org-name="ORG NAME"    specify Organization name to return domains that match it

Help Options:
  -h, --help                   Show this help message
```
**EXAMPLES**

`cat domains.txt | gorg -a -o output.txt`

`cat damains.txt | gorg -r "dummy org" -o output.txt`

`gorg -f domains.txt -r "dummy org" -o output.txt`
