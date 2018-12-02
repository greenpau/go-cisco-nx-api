[![Go Report Card](https://goreportcard.com/badge/github.com/greenpau/go-cisco-nx-api)](https://goreportcard.com/report/github.com/greenpau/go-cisco-nx-api) [![Build Status](https://travis-ci.org/greenpau/go-cisco-nx-api.svg?branch=master)](https://travis-ci.org/greenpau/go-cisco-nx-api) [![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/greenpau/go-cisco-nx-api)

# go-cisco-nx-api

Cisco NX-OS API client library written in Go.

## Table of Contents

* [Getting Started](#getting-started)
* [Cisco NX-API Configuration](#cisco-nx-api-configuration)
* [Command-Line Client](#command-line-client)
* [References](#references)

## Getting Started

This library is well documented and comes with a command-line utility for code
implementation and usage purposes. The library helps its users querying system
information, vlans, and interfaces, etc.

* `GetSystemInfo()`: **show version**
* `GetVlans()`: **show vlan**
* `GetInterfaces()` **show interface**
* `GetGeneric()`: runs any arbitrary command and produces JSON output

For example, the following snippet queries system information:

```golang
import (
    "github.com/greenpau/go-cisco-nx-api/pkg/client"
)

cli := client.NewClient()
cli.SetHost("nysw01")
cli.SetPort(443)
cli.SetProtocol("https")
cli.SetUsername("admin")
cli.SetPassword("cisco")
if data, err := cli.GetSystemInfo(); err != nil {
    log.Fatalf("%s", err)
} else {
    // do something with  the data
}
```

## Cisco NX-API Configuration

Use the following command to check the status of NX-API server:

```
switch# show nxapi
nxapi enabled
HTTPS Listen on port 443
Certificate Information:
    Issuer:   /C=US/ST=CA/L=San Jose/O=Cisco Systems Inc./OU=dcnxos/CN=nxos
    Expires:  Nov 24 22:00:40 2018 GMT
switch#
```

If it is not enabled, enable it:

```
conf t
  feature nxapi
  nxapi http port 80
```

## Command-Line Client

The `make` command builds `bin/go-cisco-nx-api-client` binary.

```
go-cisco-nx-api-client - Cisco NX-OS API client

Usage: go-cisco-nx-api-client [arguments]

  -cli string
        cli command
  -host string
        target hostname or ip address
  -log.level string
        logging severity level (default "info")
  -pass string
        password
  -port int
        target port (default 443)
  -proto string
        protocol: https (default) or http (default "https")
  -secure
        validate certificates, default: false
  -user string
        username
  -version
        version information

Documentation: https://github.com/greenpau/go-cisco-nx-api/
```

Here are a number of examples how to invoke the clients. In these example,
the output is actually a structure or a list of structures.

```
bin/go-cisco-nx-api-client -cli "show version" -host 10.1.1.1 -user admin -port 443 -proto https -log.level debug -pass cisco
bin/go-cisco-nx-api-client -cli "show version" -host 10.1.1.1 -user admin -port 80 -proto http -log.level info -pass cisco
bin/go-cisco-nx-api-client -cli "show version" -host 10.1.1.1 -user admin -pass cisco
bin/go-cisco-nx-api-client -cli "show interface" -host 10.1.1.1 -user admin -pass cisco
bin/go-cisco-nx-api-client -cli "show vlan" -host 10.1.1.1 -user admin -pass cisco
```

However, the same client can be used to run any command, e.g. **show ip arp**.
Here, the output is in JSON format.

```
bin/go-cisco-nx-api-client -cli "show ip arp" -host 10.1.1.1 -user admin -pass cisco
```

## References

* [Cisco Nexus 9000 Series NX-OS Programmability Guide, Release 6.x](https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus9000/sw/6-x/programmability/guide/b_Cisco_Nexus_9000_Series_NX-OS_Programmability_Guide/b_Cisco_Nexus_9000_Series_NX-OS_Programmability_Guide_chapter_011.html)

Useful `curl` commands:

```
cat << EOF > show.version.json
[{ "jsonrpc": "2.0", "method": "cli", "params": { "cmd": "show version", "version": 1 }, "id": 1 }]
EOF
curl -v -u admin:cisco -H "Content-Type: application/json-rpc" -H "Cache-Control: no-cache" -d @show.version.json -X POST http://127.0.0.1:80/ins
```
