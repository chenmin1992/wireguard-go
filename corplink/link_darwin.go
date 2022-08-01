package corplink

import (
	"fmt"
	"net/netip"
	"os/exec"
	"unsafe"

	"golang.org/x/sys/unix"
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
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return err
	}
	ip.Is6()
	fd := int(tunDev.File().Fd())
	req := uint(unix.SIOCSIFADDR)
	var devName [unix.IFNAMSIZ]byte
	copy(devName[:], name)
	var ifReq any
	if ip.Is6() {
		ifReq = struct {
			Name [unix.IFNAMSIZ]byte
			Addr unix.RawSockaddrInet6
		}{
			Name: devName,
			Addr: unix.RawSockaddrInet6{
				Len:    1 + 1 + 2 + 4 + 16 + 4,
				Family: unix.AF_INET,
				Addr:   ip.As16(),
			},
		}

	} else {
		ifReq = struct {
			Name [unix.IFNAMSIZ]byte
			Addr unix.RawSockaddrInet4
		}{
			Name: devName,
			Addr: unix.RawSockaddrInet4{
				Len:    1 + 1 + 2 + 4 + 8,
				Family: unix.AF_INET,
				Addr:   ip.As4(),
			},
		}
	}
	ptr := unsafe.Pointer(&ifReq)
	return unix.IoctlSetPointerInt(fd, req, int(uintptr(ptr)))
}

func AddInterfaceRoute(name, network string) error {
	return routeAdd(name, network)
}
