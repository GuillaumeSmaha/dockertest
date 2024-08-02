// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build (!windows && !linux && !freebsd) || (freebsd && !cgo)
// +build !windows,!linux,!freebsd freebsd,!cgo

package mount // import "github.com/GuillaumeSmaha/dockertest/docker/pkg/mount"

import (
	"fmt"
	"runtime"
)

func parseMountTable() ([]*Info, error) {
	return nil, fmt.Errorf("mount.parseMountTable is not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
