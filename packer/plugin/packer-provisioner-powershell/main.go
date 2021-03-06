// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/MSOpenTech/packer-hyperv/packer/provisioner/powershell"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(powershell.Provisioner))
	server.Serve()
}
