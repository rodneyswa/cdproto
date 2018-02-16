package page

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
)

// EventDomContentEventFired [no description].
type EventDomContentEventFired struct {
	Timestamp *cdp.MonotonicTime `json:"timestamp"`
}

// EventFrameAttached fired when frame has been attached to its parent.
type EventFrameAttached struct {
	FrameID       cdp.FrameID         `json:"frameId"`         // Id of the frame that has been attached.
	ParentFrameID cdp.FrameID         `json:"parentFrameId"`   // Parent frame identifier.
	Stack         *runtime.StackTrace `json:"stack,omitempty"` // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
}

// EventFrameClearedScheduledNavigation fired when frame no longer has a
// scheduled navigation.
type EventFrameClearedScheduledNavigation struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that has cleared its scheduled navigation.
}

// EventFrameDetached fired when frame has been detached from its parent.
type EventFrameDetached struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that has been detached.
}

// EventFrameNavigated fired once navigation of the frame has completed.
// Frame is now associated with the new loader.
type EventFrameNavigated struct {
	Frame *cdp.Frame `json:"frame"` // Frame object.
}

// EventFrameResized [no description].
type EventFrameResized struct{}

// EventFrameScheduledNavigation fired when frame schedules a potential
// navigation.
type EventFrameScheduledNavigation struct {
	FrameID cdp.FrameID                    `json:"frameId"` // Id of the frame that has scheduled a navigation.
	Delay   float64                        `json:"delay"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
	Reason  FrameScheduledNavigationReason `json:"reason"`  // The reason for the navigation.
	URL     string                         `json:"url"`     // The destination URL for the scheduled navigation.
}

// EventFrameStartedLoading fired when frame has started loading.
type EventFrameStartedLoading struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that has started loading.
}

// EventFrameStoppedLoading fired when frame has stopped loading.
type EventFrameStoppedLoading struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that has stopped loading.
}

// EventInterstitialHidden fired when interstitial page was hidden.
type EventInterstitialHidden struct{}

// EventInterstitialShown fired when interstitial page was shown.
type EventInterstitialShown struct{}

// EventJavascriptDialogClosed fired when a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload) has been closed.
type EventJavascriptDialogClosed struct {
	Result    bool   `json:"result"`    // Whether dialog was confirmed.
	UserInput string `json:"userInput"` // User input in case of prompt.
}

// EventJavascriptDialogOpening fired when a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload) is about to open.
type EventJavascriptDialogOpening struct {
	URL           string     `json:"url"`                     // Frame url.
	Message       string     `json:"message"`                 // Message that will be displayed by the dialog.
	Type          DialogType `json:"type"`                    // Dialog type.
	DefaultPrompt string     `json:"defaultPrompt,omitempty"` // Default dialog prompt.
}

// EventLifecycleEvent fired for top level page lifecycle events such as
// navigation, load, paint, etc.
type EventLifecycleEvent struct {
	FrameID   cdp.FrameID        `json:"frameId"`  // Id of the frame.
	LoaderID  cdp.LoaderID       `json:"loaderId"` // Loader identifier. Empty string if the request is fetched from worker.
	Name      string             `json:"name"`
	Timestamp *cdp.MonotonicTime `json:"timestamp"`
}

// EventLoadEventFired [no description].
type EventLoadEventFired struct {
	Timestamp *cdp.MonotonicTime `json:"timestamp"`
}

// EventScreencastFrame compressed image data requested by the
// startScreencast.
type EventScreencastFrame struct {
	Data      string                   `json:"data"`      // Base64-encoded compressed image.
	Metadata  *ScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
	SessionID int64                    `json:"sessionId"` // Frame number.
}

// EventScreencastVisibilityChanged fired when the page with currently
// enabled screencast was shown or hidden .
type EventScreencastVisibilityChanged struct {
	Visible bool `json:"visible"` // True if the page is visible.
}

// EventWindowOpen fired when a new window is going to be opened, via
// window.open(), link click, form submission, etc.
type EventWindowOpen struct {
	URL            string   `json:"url"`            // The URL for the new window.
	WindowName     string   `json:"windowName"`     // Window name.
	WindowFeatures []string `json:"windowFeatures"` // An array of enabled window features.
	UserGesture    bool     `json:"userGesture"`    // Whether or not it was triggered by user gesture.
}
