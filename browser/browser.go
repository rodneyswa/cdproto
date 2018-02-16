// Package browser provides the Chrome Debugging Protocol
// commands, types, and events for the Browser domain.
//
// The Browser domain defines methods and events for browser managing.
//
// Generated by the cdproto-gen command.
package browser

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
)

// CloseParams close browser gracefully.
type CloseParams struct{}

// Close close browser gracefully.
func Close() *CloseParams {
	return &CloseParams{}
}

// Do executes Browser.close against the provided context.
func (p *CloseParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandClose, nil, nil)
}

// GetVersionParams returns version information.
type GetVersionParams struct{}

// GetVersion returns version information.
func GetVersion() *GetVersionParams {
	return &GetVersionParams{}
}

// GetVersionReturns return values.
type GetVersionReturns struct {
	ProtocolVersion string `json:"protocolVersion,omitempty"` // Protocol version.
	Product         string `json:"product,omitempty"`         // Product name.
	Revision        string `json:"revision,omitempty"`        // Product revision.
	UserAgent       string `json:"userAgent,omitempty"`       // User-Agent.
	JsVersion       string `json:"jsVersion,omitempty"`       // V8 version.
}

// Do executes Browser.getVersion against the provided context.
//
// returns:
//   protocolVersion - Protocol version.
//   product - Product name.
//   revision - Product revision.
//   userAgent - User-Agent.
//   jsVersion - V8 version.
func (p *GetVersionParams) Do(ctxt context.Context, h cdp.Executor) (protocolVersion string, product string, revision string, userAgent string, jsVersion string, err error) {
	// execute
	var res GetVersionReturns
	err = h.Execute(ctxt, CommandGetVersion, nil, &res)
	if err != nil {
		return "", "", "", "", "", err
	}

	return res.ProtocolVersion, res.Product, res.Revision, res.UserAgent, res.JsVersion, nil
}

// GetHistogramsParams get Chrome histograms.
type GetHistogramsParams struct {
	Query string `json:"query,omitempty"` // Requested substring in name. Only histograms which have query as a substring in their name are extracted. An empty or absent query returns all histograms.
}

// GetHistograms get Chrome histograms.
//
// parameters:
func GetHistograms() *GetHistogramsParams {
	return &GetHistogramsParams{}
}

// WithQuery requested substring in name. Only histograms which have query as
// a substring in their name are extracted. An empty or absent query returns all
// histograms.
func (p GetHistogramsParams) WithQuery(query string) *GetHistogramsParams {
	p.Query = query
	return &p
}

// GetHistogramsReturns return values.
type GetHistogramsReturns struct {
	Histograms []*Histogram `json:"histograms,omitempty"` // Histograms.
}

// Do executes Browser.getHistograms against the provided context.
//
// returns:
//   histograms - Histograms.
func (p *GetHistogramsParams) Do(ctxt context.Context, h cdp.Executor) (histograms []*Histogram, err error) {
	// execute
	var res GetHistogramsReturns
	err = h.Execute(ctxt, CommandGetHistograms, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Histograms, nil
}

// GetHistogramParams get a Chrome histogram by name.
type GetHistogramParams struct {
	Name string `json:"name"` // Requested histogram name.
}

// GetHistogram get a Chrome histogram by name.
//
// parameters:
//   name - Requested histogram name.
func GetHistogram(name string) *GetHistogramParams {
	return &GetHistogramParams{
		Name: name,
	}
}

// GetHistogramReturns return values.
type GetHistogramReturns struct {
	Histogram *Histogram `json:"histogram,omitempty"` // Histogram.
}

// Do executes Browser.getHistogram against the provided context.
//
// returns:
//   histogram - Histogram.
func (p *GetHistogramParams) Do(ctxt context.Context, h cdp.Executor) (histogram *Histogram, err error) {
	// execute
	var res GetHistogramReturns
	err = h.Execute(ctxt, CommandGetHistogram, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Histogram, nil
}

// GetWindowBoundsParams get position and size of the browser window.
type GetWindowBoundsParams struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
}

// GetWindowBounds get position and size of the browser window.
//
// parameters:
//   windowID - Browser window id.
func GetWindowBounds(windowID WindowID) *GetWindowBoundsParams {
	return &GetWindowBoundsParams{
		WindowID: windowID,
	}
}

// GetWindowBoundsReturns return values.
type GetWindowBoundsReturns struct {
	Bounds *Bounds `json:"bounds,omitempty"` // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// Do executes Browser.getWindowBounds against the provided context.
//
// returns:
//   bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (p *GetWindowBoundsParams) Do(ctxt context.Context, h cdp.Executor) (bounds *Bounds, err error) {
	// execute
	var res GetWindowBoundsReturns
	err = h.Execute(ctxt, CommandGetWindowBounds, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Bounds, nil
}

// GetWindowForTargetParams get the browser window that contains the devtools
// target.
type GetWindowForTargetParams struct {
	TargetID target.ID `json:"targetId"` // Devtools agent host id.
}

// GetWindowForTarget get the browser window that contains the devtools
// target.
//
// parameters:
//   targetID - Devtools agent host id.
func GetWindowForTarget(targetID target.ID) *GetWindowForTargetParams {
	return &GetWindowForTargetParams{
		TargetID: targetID,
	}
}

// GetWindowForTargetReturns return values.
type GetWindowForTargetReturns struct {
	WindowID WindowID `json:"windowId,omitempty"` // Browser window id.
	Bounds   *Bounds  `json:"bounds,omitempty"`   // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// Do executes Browser.getWindowForTarget against the provided context.
//
// returns:
//   windowID - Browser window id.
//   bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (p *GetWindowForTargetParams) Do(ctxt context.Context, h cdp.Executor) (windowID WindowID, bounds *Bounds, err error) {
	// execute
	var res GetWindowForTargetReturns
	err = h.Execute(ctxt, CommandGetWindowForTarget, p, &res)
	if err != nil {
		return 0, nil, err
	}

	return res.WindowID, res.Bounds, nil
}

// SetWindowBoundsParams set position and/or size of the browser window.
type SetWindowBoundsParams struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
	Bounds   *Bounds  `json:"bounds"`   // New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
}

// SetWindowBounds set position and/or size of the browser window.
//
// parameters:
//   windowID - Browser window id.
//   bounds - New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
func SetWindowBounds(windowID WindowID, bounds *Bounds) *SetWindowBoundsParams {
	return &SetWindowBoundsParams{
		WindowID: windowID,
		Bounds:   bounds,
	}
}

// Do executes Browser.setWindowBounds against the provided context.
func (p *SetWindowBoundsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetWindowBounds, p, nil)
}

// Command names.
const (
	CommandClose              = "Browser.close"
	CommandGetVersion         = "Browser.getVersion"
	CommandGetHistograms      = "Browser.getHistograms"
	CommandGetHistogram       = "Browser.getHistogram"
	CommandGetWindowBounds    = "Browser.getWindowBounds"
	CommandGetWindowForTarget = "Browser.getWindowForTarget"
	CommandSetWindowBounds    = "Browser.setWindowBounds"
)
