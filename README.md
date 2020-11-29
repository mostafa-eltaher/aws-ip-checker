# AWS IP Checker
A small utility tool to check whether an IP Address belongs to AWS. The input IP Address can be specified directly or via a Domain Name

It is writting in Go, mainly to avoid the need to install any runtime environment (e.g. JVM, Python, or node.js).

## How it works
Simply, it downloads the [ip-ranges.json](https://ip-ranges.amazonaws.com/ip-ranges.json) file, parse it, then search in it for desired IP address.

For that, it needs access to the internet (also if a Domain Name is used, then the access to a DNS server is needed)

More details are [here](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html)
## How to Build (on windows)
```bash
go build -o bin\aws-ip-addr-windows-amd64.exe .\cmd
```
or (for a release version)
```bash
go build -o bin\aws-ip-addr-windows-amd64-release.exe -ldflags "-s -w" .\cmd
```

## How to use (on windows)
- To check an IP address:
```bash
bin\aws-ip-addr-windows-amd64.exe 35.179.8.33
```
- To check a Domian Name:
```bash
bin\aws-ip-addr-windows-amd64.exe - dnssec-name-and-shame.com
```