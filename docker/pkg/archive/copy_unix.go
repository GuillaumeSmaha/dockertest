// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !windows
// +build !windows

package archive // import "github.com/GuillaumeSmaha/dockertest/docker/pkg/archive"

import (
	"path/filepath"
)

func normalizePath(path string) string {
	return filepath.ToSlash(path)
}
