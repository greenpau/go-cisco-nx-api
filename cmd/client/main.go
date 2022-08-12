// Copyright 2018 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/greenpau/go-cisco-nx-api/pkg/client"
	"github.com/greenpau/versioned"
	log "github.com/sirupsen/logrus"
)

var (
	app        *versioned.PackageManager
	appVersion string
	gitBranch  string
	gitCommit  string
	buildUser  string
	buildDate  string
)

func init() {
	app = versioned.NewPackageManager("go-cisco-nx-api-client")
	app.Description = "Cisco NX-OS API client."
	app.Documentation = "https://github.com/greenpau/go-cisco-nx-api/"
	app.SetVersion(appVersion, "1.0.8")
	app.SetGitBranch(gitBranch, "main")
	app.SetGitCommit(gitCommit, "v1.0.7-5-ga4851ba")
	app.SetBuildUser(buildUser, "")
	app.SetBuildDate(buildDate, "")
}

func main() {
	var logLevel string
	var isShowVersion bool

	var host, proto, authUser, authPass, cliCommand string
	var port int
	var secure bool

	flag.StringVar(&host, "host", "", "target hostname or ip address")
	flag.IntVar(&port, "port", 443, "target port")
	flag.StringVar(&proto, "proto", "https", "protocol: https (default) or http")
	flag.BoolVar(&secure, "secure", false, "validate certificates, default: false")
	flag.StringVar(&authUser, "user", "", "username")
	flag.StringVar(&authPass, "pass", "", "password")
	flag.StringVar(&cliCommand, "cli", "", "cli command")
	flag.StringVar(&logLevel, "log.level", "info", "logging severity level")
	flag.BoolVar(&isShowVersion, "version", false, "version information")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n%s - %s\n\n", app.Name, app.Description)
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments]\n\n", app.Name)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nDocumentation: %s\n\n", app.Documentation)
	}
	flag.Parse()
	if isShowVersion {
		fmt.Fprintf(os.Stdout, "%s\n", app.Banner())
		os.Exit(0)
	}
	if level, err := log.ParseLevel(logLevel); err == nil {
		log.SetLevel(level)
	} else {
		log.Errorf(err.Error())
		os.Exit(1)
	}

	cli := client.NewClient()
	if err := cli.SetHost(host); err != nil {
		log.Fatalf("argument '-host': %s", err)
	}
	if err := cli.SetPort(port); err != nil {
		log.Fatalf("argument '-port': %s", err)
	}
	if err := cli.SetProtocol(proto); err != nil {
		log.Fatalf("argument '-proto': %s", err)
	}
	if err := cli.SetUsername(authUser); err != nil {
		log.Fatalf("argument '-user': %s", err)
	}
	if err := cli.SetPassword(authPass); err != nil {
		log.Fatalf("argument '-pass': %s", err)
	}
	if secure {
		if err := cli.SetSecure(); err != nil {
			log.Fatalf("argument '-secure': %s", err)
		}
	}
	log.Debugf("host: %s, port: %d, secure: %t,  user: %s, cli command: %s", host, port, secure, authUser, cliCommand)

	switch cliCommand {
	case "":
		log.Fatalf("argument '-cli' is empty")
	case "show version":
		start := time.Now()
		sysinfo, err := cli.GetSystemInfo()
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Fprintf(os.Stdout, "Hostname: %s\n", sysinfo.Hostname)
		fmt.Fprintf(os.Stdout, "Processor Board ID: %s\n", sysinfo.ProcessorBoardID)
		fmt.Fprintf(os.Stdout, "Kickstart Image Version: %s\n", sysinfo.KickstartImage.Version)
		fmt.Fprintf(os.Stdout, "Uptime: %d\n", sysinfo.Uptime)
		log.Debugf("took %s", time.Since(start))
	case "show vlan":
		start := time.Now()
		vlans, err := cli.GetVlans()
		if err != nil {
			log.Fatalf("%s", err)
		}
		for _, vlan := range vlans {
			fmt.Fprintf(os.Stdout, "Vlan ID %s, Name: %s\n", vlan.ID, vlan.Name)
		}
		log.Debugf("took %s", time.Since(start))
	case "show vlan counters":
		start := time.Now()
		vlanCounters, err := cli.GetVlanCounters()
		if err != nil {
			log.Fatalf("%s", err)
		}
		for _, vlanCounter := range vlanCounters {
			fmt.Fprintf(os.Stdout, "Vlan ID %d, InUcastBytes: %d, OutUcastBytes: %d\n", vlanCounter.ID, vlanCounter.InputUnicastBytes, vlanCounter.OutputUnicastBytes)
		}
		log.Debugf("took %s", time.Since(start))
	case "show interface":
		start := time.Now()
		ifaces, err := cli.GetInterfaces()
		if err != nil {
			log.Fatalf("%s", err)
		}
		for _, iface := range ifaces {
			var out strings.Builder
			out.WriteString(fmt.Sprintf("Interface Name: %s", iface.Name))
			out.WriteString(fmt.Sprintf(", State: %s/%s", iface.Props.State, iface.Props.AdminState))
			if iface.Props.BiaHwAddr != "" {
				out.WriteString(fmt.Sprintf(", MAC: %s", iface.Props.BiaHwAddr))
			}
			if iface.Props.IPAddress != "" {
				out.WriteString(fmt.Sprintf(", IP: %s/%d", iface.Props.IPAddress, iface.Props.IPMask))
			}
			fmt.Fprintf(os.Stdout, "%s\n", out.String())
		}
		log.Debugf("took %s", time.Since(start))
	case "show running-config":
		config, err := cli.GetRunningConfiguration()
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", config.Text)
	case "show startup-config":
		config, err := cli.GetStartupConfiguration()
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", config.Text)
	case "show ip bgp summary":
		bgpSummary, err := cli.GetBgpSummary()
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", bgpSummary.Text)
	case "show interface transceiver details":
		transceivers, err := cli.GetTransceivers()
		if err != nil {
			log.Fatalf("%s", err)
		}
		for _, t := range transceivers {
			fmt.Fprintf(os.Stdout, "%s\n", t.String())
		}
	default:
		start := time.Now()

		// show interface <name>
		if strings.HasPrefix(cliCommand, "show interface") {
			intfName := cliCommand[len("show interface "):]
			intfInfo, err := cli.GetInterface(intfName)
			if err != nil {
				log.Fatalf("%s", err)
			}
			var out strings.Builder
			out.WriteString(fmt.Sprintf("State: %s/%s", intfInfo.Props.State, intfInfo.Props.AdminState))
			if intfInfo.Props.BiaHwAddr != "" {
				out.WriteString(fmt.Sprintf(", MAC: %s", intfInfo.Props.BiaHwAddr))
			}
			if intfInfo.Props.IPAddress != "" {
				out.WriteString(fmt.Sprintf(", IP: %s/%d", intfInfo.Props.IPAddress, intfInfo.Props.IPMask))
			}
			fmt.Fprintf(os.Stdout, "%s\n", out.String())
			return
		}

		// show running-config interface <name>
		if strings.HasPrefix(cliCommand, "show running-config interface") {
			intfName := cliCommand[len("show running-config interface "):]
			intfInfo, err := cli.GetInterfaceRunningConfiguration(intfName)
			if err != nil {
				log.Fatalf("%s", err)
			}
			fmt.Fprintf(os.Stdout, "%s\n", intfInfo.Text)
			return
		}

		// interface configuration, use ; to break commands
		if strings.HasPrefix(cliCommand, "interface") {
			cmds := strings.Split(cliCommand, ";")
			resp, err := cli.Configure(cmds)
			if err != nil {
				log.Fatalf("%v", err)
			}
			for i, r := range resp {
				if r.Error != nil {
					log.Fatalf("failed to execute command %v:\n%v\n", cmds[i], r.Error)
				}
			}
			return
		}

		output, err := cli.GetGeneric(cliCommand)
		if err != nil {
			log.Fatalf("%s", err)
		}

		var respJSON client.JSONRPCResponse
		err = json.Unmarshal(output, &respJSON)
		if err != nil {
			log.Fatalf("JSON parsing failed:%v\n Input: %s", err, string(output))
		}

		if respJSON.Error != nil {
			log.Fatalf("Command returned failure: %v", respJSON.Error)
		}

		var body client.JSONRPCResponseBody
		err = json.Unmarshal(respJSON.Result, &body)
		if err != nil {
			log.Fatalf("JSON parsing failed:%v\n Input: %s", err, string(respJSON.Result))
		}
		fmt.Println(string(body.Body))

		log.Debugf("took %s", time.Since(start))
	}
}
