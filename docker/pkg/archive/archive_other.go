// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !linux
// +build !linux

package archive // import "github.com/GuillaumeSmaha/dockertest/docker/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat) tarWhiteoutConverter {
	return nil
}
