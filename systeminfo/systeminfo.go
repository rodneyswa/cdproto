// Package systeminfo provides the Chrome Debugging Protocol
// commands, types, and events for the SystemInfo domain.
//
// The SystemInfo domain defines methods and events for querying low-level
// system information.
//
// Generated by the cdproto-gen command.
package systeminfo

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// GetInfoParams returns information about the system.
type GetInfoParams struct{}

// GetInfo returns information about the system.
func GetInfo() *GetInfoParams {
	return &GetInfoParams{}
}

// GetInfoReturns return values.
type GetInfoReturns struct {
	Gpu          *GPUInfo `json:"gpu,omitempty"`          // Information about the GPUs on the system.
	ModelName    string   `json:"modelName,omitempty"`    // A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
	ModelVersion string   `json:"modelVersion,omitempty"` // A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
	CommandLine  string   `json:"commandLine,omitempty"`  // The command line string used to launch the browser. Will be the empty string if not supported.
}

// Do executes SystemInfo.getInfo against the provided context.
//
// returns:
//   gpu - Information about the GPUs on the system.
//   modelName - A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
//   modelVersion - A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
//   commandLine - The command line string used to launch the browser. Will be the empty string if not supported.
func (p *GetInfoParams) Do(ctxt context.Context, h cdp.Executor) (gpu *GPUInfo, modelName string, modelVersion string, commandLine string, err error) {
	// execute
	var res GetInfoReturns
	err = h.Execute(ctxt, CommandGetInfo, nil, &res)
	if err != nil {
		return nil, "", "", "", err
	}

	return res.Gpu, res.ModelName, res.ModelVersion, res.CommandLine, nil
}

// Command names.
const (
	CommandGetInfo = "SystemInfo.getInfo"
)
