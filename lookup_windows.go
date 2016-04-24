package lookup

import (
	"net"
	"os"
	"syscall"
	"unsafe"
)

var (
	modws2_32        = syscall.NewLazyDLL("ws2_32.dll")
	procGetNameInfoW = modws2_32.NewProc("GetNameInfoW")
)

const (
	NI_NAMEREQD = 0x04
	NI_MAXHOST  = 1025
)

func getNameInfoW(sockAddr uintptr, sockAddrLength uintptr) (string, error) {
	err := procGetNameInfoW.Find()
	if err != nil {
		return "", err
	}

	buf := make([]uint16, NI_MAXHOST)

	errcode, _, _ := procGetNameInfoW.Call(
		sockAddr,
		sockAddrLength,
		uintptr(unsafe.Pointer(&buf[0])),
		NI_MAXHOST,
		0,
		0,
		NI_NAMEREQD)

	if errcode != 0 {
		return "", syscall.Errno(errcode)
	}

	host := syscall.UTF16ToString(buf)
	return host, nil
}

func lookupAddr(ip string) (string, error) {
	var data syscall.WSAData
	err := syscall.WSAStartup(uint32(0x202), &data)
	if err != nil {
		return "", err
	}
	defer syscall.WSACleanup()

	hints := syscall.AddrinfoW{
		Family:   syscall.AF_UNSPEC,
		Socktype: syscall.SOCK_STREAM,
		Protocol: syscall.IPPROTO_IP,
	}
	var result *syscall.AddrinfoW
	err = syscall.GetAddrInfoW(syscall.StringToUTF16Ptr(ip), nil, &hints, &result)
	if err != nil {
		return "", &net.DNSError{Err: os.NewSyscallError("getaddrinfow", err).Error(), Name: ip}
	}
	defer syscall.FreeAddrInfoW(result)

	return getNameInfoW(result.Addr, result.Addrlen)
}
