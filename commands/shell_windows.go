//go:build windows
// +build windows

// Merlin is a post-exploitation command and control framework.
// This file is part of Merlin.
// Copyright (C) 2022  Russel Van Tuyl

// Merlin is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.

// Merlin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Merlin.  If not, see <http://www.gnu.org/licenses/>.

package commands

import (
	"os"
	"strings"
)

// shell is used to execute a command on a host using the operating system's default shell
func shell(args []string) (stdout string, stderr string) {
	var shell string
	var arguments []string
	if s, ok := os.LookupEnv("COMSPEC"); ok {
		shell = s
		if strings.Contains(s, "cmd.exe") {
			arguments = []string{"/c"}
		} else if strings.Contains(s, "powershell.exe") {
			arguments = []string{"-Command"}
		}
	} else {
		shell = "cmd.exe"
		arguments = []string{"/c"}
	}
	return executeCommand(shell, append(arguments, args...))
}
