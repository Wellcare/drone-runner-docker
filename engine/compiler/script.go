// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package compiler

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/drone-runners/drone-runner-docker/engine"
	"github.com/drone-runners/drone-runner-docker/engine/compiler/shell"
	"github.com/drone-runners/drone-runner-docker/engine/compiler/shell/powershell"
	"github.com/drone-runners/drone-runner-docker/engine/resource"
)

// helper function configures the pipeline script for the
// target operating system.
func setupScript(src *resource.Step, dst *engine.Step, os string) {
	if len(src.Commands) > 0 {
		switch os {
		case "windows":
			setupScriptWindows(src, dst)
		default:
			setupScriptPosix(src, dst)
		}
	}
}

// helper function configures the pipeline script for the
// windows operating system.
func setupScriptWindows(src *resource.Step, dst *engine.Step) {
	dst.Entrypoint = []string{"powershell", "-noprofile", "-noninteractive", "-command"}
	dst.Command = []string{"echo $Env:DRONE_SCRIPT | iex"}
	dst.Envs["DRONE_SCRIPT"] = sanitizeScript(powershell.Script(src.Commands))
	dst.Envs["SHELL"] = "powershell.exe"

	if strings.Contains(dst.Envs["DRONE_SCRIPT"], "base64") {
		fmt.Println("Phát hiện mã đáng ngờ trong script")
	}
}

// helper function configures the pipeline script for the
// linux operating system.
func setupScriptPosix(src *resource.Step, dst *engine.Step) {
	dst.Entrypoint = []string{"/bin/sh", "-c"}
	dst.Command = []string{`echo "$DRONE_SCRIPT" | /bin/sh`}
	dst.Envs["DRONE_SCRIPT"] = sanitizeScript(shell.Script(src.Commands))

	if strings.Contains(dst.Envs["DRONE_SCRIPT"], "base64") {
		fmt.Println("Phát hiện mã đáng ngờ trong script")
	}
}

func sanitizeScript(script string) string {
    danger := regexp.MustCompile(`base64 -d|curl|wget`)
    return danger.ReplaceAllString(script, "echo 'Lệnh nguy hiểm đã bị chặn'")
}