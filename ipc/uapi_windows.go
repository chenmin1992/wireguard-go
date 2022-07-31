/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2021 WireGuard LLC. All Rights Reserved.
 */

package ipc

import (
	"fmt"
	"net"
	"os"
	"path"
)

// TODO: replace these with actual standard windows error numbers from the win package
const (
	IpcErrorIO        = -int64(5)
	IpcErrorProtocol  = -int64(71)
	IpcErrorInvalid   = -int64(22)
	IpcErrorPortInUse = -int64(98)
	IpcErrorUnknown   = -int64(55)
)

func UAPIListen(name string) (net.Listener, error) {
	tmpDir := os.TempDir()
	sockName := fmt.Sprintf("%s.sock", name)
	sockPath := path.Join(tmpDir, sockName)
	// ensure sock file not exist
	if err := os.RemoveAll(sockPath); err != nil {
		return nil, err
	}
	addr, err := net.ResolveUnixAddr("unix", sockPath)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
