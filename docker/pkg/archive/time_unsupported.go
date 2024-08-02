// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !linux
// +build !linux

package archive // import "github.com/GuillaumeSmaha/dockertest/docker/pkg/archive"

import (
	"syscall"
	"time"
)

func timeToTimespec(time time.Time) (ts syscall.Timespec) {
	nsec := int64(0)
	if !time.IsZero() {
		nsec = time.UnixNano()
	}
	return syscall.NsecToTimespec(nsec)
}
