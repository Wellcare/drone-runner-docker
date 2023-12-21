// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package compiler

import (
	"fmt"
	"github.com/wellcare/drone-runner-docker/engine"
	"github.com/wellcare/drone-runner-docker/engine/compiler/shell"
	"github.com/wellcare/drone-runner-docker/engine/compiler/shell/powershell"
	"github.com/wellcare/drone-runner-docker/engine/resource"
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
	fmt.Println(src.Commands)
	dst.Envs["DRONE_SCRIPT"] = powershell.Script(src.Commands)
	dst.Envs["SHELL"] = "powershell.exe"
	fmt.Println(dst.Envs["DRONE_SCRIPT"])
}

// helper function configures the pipeline script for the
// linux operating system.
func setupScriptPosix(src *resource.Step, dst *engine.Step) {
	dst.Entrypoint = []string{"/bin/sh", "-c"}
	dst.Command = []string{`echo "$DRONE_SCRIPT" | /bin/sh`}
	fmt.Println("src.Commands")
	fmt.Println(src.Commands)
	re := regexp.MustCompile(`#END.*`)
	dst.Envs["DRONE_SCRIPT"] = re.ReplaceAllString(dst.Envs["DRONE_SCRIPT"], "")
	fmt.Println("DRONE_SCRIPT")
	fmt.Println(dst.Envs["DRONE_SCRIPT"])

}
