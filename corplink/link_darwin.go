package corplink

import (
	"fmt"
	"os/exec"
)

func ifconfig(name string, args ...any) error {
	strArgs := []string{name}
	for _, arg := range args {
		strArgs = append(strArgs, fmt.Sprint(arg))
	}
	cmd := exec.Command("ifconfig", strArgs...)
	return cmd.Run()
}

func routeAdd(name string, network string) error {
	cmd := exec.Command("route", "add", "-net", network, "dev", name)
	return cmd.Run()
}

func SetInterfaceUp(name string, up bool) error {
	if up {
		return ifconfig(name, "up")
	}
	return ifconfig(name, "down")
}

func SetInterfaceMTU(name string, mtu int) error {
	return ifconfig(name, "mtu", mtu)
}

func SetInterfaceAddress(name, addr string) error {
	return ifconfig(name, "add", addr)
}

func AddInterfaceRoute(name, network string) error {
	return routeAdd(name, network)
}
