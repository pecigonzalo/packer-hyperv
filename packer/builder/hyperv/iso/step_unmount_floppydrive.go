// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package iso

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	powershell "github.com/MSOpenTech/packer-hyperv/packer/powershell"
	common "github.com/MSOpenTech/packer-hyperv/packer/builder/hyperv/common"
)


type StepUnmountFloppyDrive struct {
}

func (s *StepUnmountFloppyDrive) Run(state multistep.StateBag) multistep.StepAction {
	//config := state.Get("config").(*config)
	//driver := state.Get("driver").(hypervcommon.Driver)
	ui := state.Get("ui").(packer.Ui)

	errorMsg := "Error Unmounting floppy drive: %s"
	vmName := state.Get("vmName").(string)

	powershell, _ := powershell.Command()

	ui.Say("Unmounting floppy drive...")

	var script common.ScriptBuilder
	script.WriteLine("param([string]$vmName)")
	script.WriteLine("Set-VMFloppyDiskDrive -VMName $vmName -Path $null")

	err := powershell.RunFile(script.Bytes(), vmName)

	if err != nil {
		err := fmt.Errorf(errorMsg, err)
		state.Put("error", err)
		ui.Error(err.Error())
	}

	return multistep.ActionContinue
}

func (s *StepUnmountFloppyDrive) Cleanup(state multistep.StateBag) {
	// do nothing
}
