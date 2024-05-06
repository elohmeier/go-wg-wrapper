package main

import (
	"fmt"
	"net"
	"os/exec"

	"github.com/aschmidt75/go-wg-wrapper/pkg/wgwrapper"
)

// Example setup
// uses all functions of wgwraper. Caution: creates, modifies wireguard interface and
// system routes.

func main() {
	// get a new wrapper
	wg := wgwrapper.New()

	// set up a new wireguard interface struct w/ some defaults
	wgi := wgwrapper.NewWireguardInterface("wg-wrap-0", net.IPNet{
		IP:   net.IPv4(10, 99, 99, 99),
		Mask: net.CIDRMask(24, 32),
	})

	// add the interface.
	err := wg.AddInterface(wgi)
	if err != nil {
		panic(err)
	}

	// Calling ip link show yields an interface wg-wrap-0
	out, err := exec.Command("ip", "-d", "link", "show", "dev", "wg-wrap-0").Output()
	println(string(out))

	// we're able to locate it
	ex, err := wg.HasInterface(wgi)
	if err != nil {
		panic(err)
	}

	// Configure the interface by creating a keypair and make it
	// listen on a port
	wgi.ListenPort = 46534
	err = wg.Configure(&wgi)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%#v\n", wgi) // must have public key here and endpoint/listenport assigned

	err = wg.SetInterfaceUp(wgi)
	if err != nil {
		panic(err)
	}

	// we're able to show its address
	out, err = exec.Command("ip", "addr", "show", "wg-wrap-0").Output()
	println(string(out))

	// add a sample peering to nowhere
	_, ipv4Net, err := net.ParseCIDR("127.0.0.1/32")

	_, err = wg.AddPeer(wgi, wgwrapper.WireguardPeer{
		RemoteEndpointIP: "10.1.2.3",
		ListenPort:       43210,
		Pubkey:           "9g4Eec+u+wBuMF06+qnsYl3G81l2PNCnG7nvtss9O2I=",
		AllowedIPs: []net.IPNet{
			*ipv4Net,
		},
		Psk: nil,
	})

	if err != nil {
		panic(err)
	}

	// calling /usr/bin/wg yields some output, including the peer. Assumes wireguard-tools is installed
	out, err = exec.Command("wg").Output()
	println(string(out))

	// we can iterate peer here, also
	wg.IteratePeers(wgi, func(p wgwrapper.WireguardPeer) {
		fmt.Printf("Peer: %s %s:%d\n", p.Pubkey, p.RemoteEndpointIP, p.ListenPort)
	})

	// set the route
	err = wg.SetRoute(wgi, "10.99.99.99/32")
	if err != nil {
		panic(err)
	}

	out, err = exec.Command("ip", "ro", "show", "dev", "wg-wrap-0").Output()
	println(string(out))

	// show the default interface
	defaultInterface, err := wg.DefaultRouteInterface()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Default interface is: %s\n", defaultInterface)

	// delete the wireguard interface
	err = wg.DeleteInterface(wgi)
	if err != nil {
		panic(err)
	}

	// it's not around anymore
	ex, _ = wg.HasInterface(wgi)
	if ex {
		panic("Error: Interface exists but should have been deleted.")
	}
}
