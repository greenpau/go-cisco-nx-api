<a href="https://github.com/greenpau/go-cisco-nx-api/actions/" target="_blank"><img src="https://github.com/greenpau/go-cisco-nx-api/workflows/build/badge.svg?branch=main"></a>
<a href="https://pkg.go.dev/github.com/greenpau/go-cisco-nx-api" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>

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
information, vlans and interfaces, etc.

* `GetSystemInfo()`: **show version**
* `GetVlans()`: **show vlan**
* `GetVlanCounters()`: **show vlan counters**
* `GetInterfaces()` **show interface**
* `GetSystemResources()` **show system resources** (CPU, Memory)
* `GetSystemEnvironment()` **show environment** (Fans, Power Supplies, Sensors)
* `GetRunningConfiguration()` **show running-config** (running configuration)
* `GetStartupConfiguration()` **show startup-config** (startup configuration)
* `GetBgpSummary()` **show ip bgp summary** (BGP routing summary)
* `GetTransceivers()` **show interface transceiver details** (fiber transceivers)
* `GetMacAddressTable()` **show mac address-table [interface name]** (MAC address table)
* `GetCDPNeighbors()` **show cdp neighbors** (CDP neighbors)
* `GetGeneric()`: runs any arbitrary command and produces JSON output

Additionally, the library allows "batch" execution of configuration commands,
e.g, change interface or vlan configurations.

* `Configure()`: execute a batch of configuration commands

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

The following snippet shuts down interface e1/1:

```golang
// shutdown interface e1/1
resp, err := cli.Configure([]string{"interface e1/1", "shutdown"})
if err != nil {
    log.Fatalf("client: %s", err)
}
for _, r := range resp {
    if r.Error != nil {
        log.Fatalf("failed to execute command %v:\n%v\n", r.ID, r.Error)
    }
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
bin/go-cisco-nx-api-client -cli "show vlan counters" -host 10.1.1.1 -user admin -pass cisco
bin/go-cisco-nx-api-client -cli "show running-config" -host 10.1.1.1 -user admin -pass cisco
bin/go-cisco-nx-api-client -cli "show startup-config" -host 10.1.1.1 -user admin -pass cisco
bin/go-cisco-nx-api-client -cli "show ip bgp summary" -host 10.1.1.1 -user admin -pass cisco
bin/go-cisco-nx-api-client -cli "show interface transceiver details" -host 10.1.1.1 -user admin -pass cisco
```

However, the same client can be used to run any command, e.g. **show ip arp**.
Here, the output is in JSON format.

```
bin/go-cisco-nx-api-client -cli "show ip arp" -host 10.1.1.1 -user admin -pass cisco
```

The client also supports interface configuration commands, e.g.:
```
bin/go-cisco-nx-api-client -cli "interface e1/1;shutdown" -host 10.1.1.1 -user admin -pass cisco
```

```;``` is used to pass multiple commands.
It can be extended to support other configuration commands with simple change.

## References

* [Cisco Nexus 9000 Series NX-OS Programmability Guide, Release 6.x](https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus9000/sw/6-x/programmability/guide/b_Cisco_Nexus_9000_Series_NX-OS_Programmability_Guide/b_Cisco_Nexus_9000_Series_NX-OS_Programmability_Guide_chapter_011.html)

Useful `curl` commands:

```
cat << EOF > show.version.json
[{ "jsonrpc": "2.0", "method": "cli", "params": { "cmd": "show version", "version": 1 }, "id": 1 }]
EOF
curl -v -u admin:cisco -H "Content-Type: application/json-rpc" -H "Cache-Control: no-cache" -d @show.version.json -X POST http://127.0.0.1:80/ins
```
