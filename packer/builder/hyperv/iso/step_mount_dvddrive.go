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


type StepMountDvdDrive struct {
	path string
}

func (s *StepMountDvdDrive) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	//driver := state.Get("driver").(hypervcommon.Driver)
	ui := state.Get("ui").(packer.Ui)

	powershell, _ := powershell.Command()
	errorMsg := "Error mounting dvd drive: %s"
	vmName := state.Get("vmName").(string)
	isoPath := config.RawSingleISOUrl

	ui.Say("Mounting dvd drive...")

	var script common.ScriptBuilder
	script.WriteLine("param([string]$vmName,[string]$path)")
	script.WriteLine("Set-VMDvdDrive -VMName $vmName -Path $path")

	err := powershell.RunFile(script.Bytes(), vmName, isoPath)

	if err != nil {
		err := fmt.Errorf(errorMsg, err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	s.path = isoPath

	return multistep.ActionContinue
}

func (s *StepMountDvdDrive) Cleanup(state multistep.StateBag) {
	if s.path == "" {
		return
	}

	powershell, _ := powershell.Command()
	errorMsg := "Error unmounting dvd drive: %s"

	vmName := state.Get("vmName").(string)
	//driver := state.Get("driver").(hypervcommon.Driver)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Unmounting dvd drive...")

	var err error = nil

	var script common.ScriptBuilder
	script.WriteLine("param([string]$vmName)")
	script.WriteLine("Set-VMDvdDrive -VMName $vmName -Path $null")

	err = powershell.RunFile(script.Bytes(), vmName)

	if err != nil {
		ui.Error(fmt.Sprintf(errorMsg, err))
	}
}
