# AWS IP Checker
A small utility tool to check whether an IP Address belongs to AWS. The input IP Address can be specified directly or via a Domain Name

It is writting in Go, mainly to avoid the need to install any runtime environment (e.g. JVM, Python, or node.js).

## How it works
Simply, it downloads the [ip-ranges.json](https://ip-ranges.amazonaws.com/ip-ranges.json) file, parse it, then search in it for desired IP address.

For that, it needs access to the internet (also if a Domain Name is used, then the access to a DNS server is needed)

More details are [here](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html)
## How to Build

### Build on windows
```cmd
go build -o bin\windows_amd64\aws-ip-checker.exe .\cmd\aws-ip-checker
```
or (for a release version)
```cmd
go build -o bin\windows_amd64\aws-ip-checker.exe -ldflags "-s -w" .\cmd\aws-ip-checker
```
### Build on Linux
```bash
go build -o bin/linux_amd64/aws-ip-checker ./cmd/aws-ip-checker
```
or (for a release version)
```bash
go build -o bin/linux_amd64/aws-ip-checker -ldflags '-s -w' ./cmd/aws-ip-checker
```

### Build by `Make`
```bash
make
```
The output will be in `bin/<version>/<os_arch>/aws-ip-checker`

## How to use
### To check an IP address

```bash
aws-ip-checker 35.179.8.33
```
the output should look like:
```text
Downloading the ip range json file ...
Found 6908 IP ranges (including duplicates)
Found address: 35.179.8.33
{
  "ip_prefix": "35.178.0.0/15",
  "region": "eu-west-2",
  "network_border_group": "eu-west-2",
  "service": "AMAZON"
}
Found address: 35.179.8.33
{
  "ip_prefix": "35.178.0.0/15",
  "region": "eu-west-2",
  "network_border_group": "eu-west-2",
  "service": "EC2"
}

2 result(s)
```

### To check a Domian Name:

```bash
aws-ip-checker - www.netflix.com
```