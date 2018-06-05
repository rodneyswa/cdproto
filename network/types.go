package network

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/cdproto/security"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// RequestID unique request identifier.
type RequestID string

// String returns the RequestID as string value.
func (t RequestID) String() string {
	return string(t)
}

// InterceptionID unique intercepted request identifier.
type InterceptionID string

// String returns the InterceptionID as string value.
func (t InterceptionID) String() string {
	return string(t)
}

// ErrorReason network level fetch failure reason.
type ErrorReason string

// String returns the ErrorReason as string value.
func (t ErrorReason) String() string {
	return string(t)
}

// ErrorReason values.
const (
	ErrorReasonFailed               ErrorReason = "Failed"
	ErrorReasonAborted              ErrorReason = "Aborted"
	ErrorReasonTimedOut             ErrorReason = "TimedOut"
	ErrorReasonAccessDenied         ErrorReason = "AccessDenied"
	ErrorReasonConnectionClosed     ErrorReason = "ConnectionClosed"
	ErrorReasonConnectionReset      ErrorReason = "ConnectionReset"
	ErrorReasonConnectionRefused    ErrorReason = "ConnectionRefused"
	ErrorReasonConnectionAborted    ErrorReason = "ConnectionAborted"
	ErrorReasonConnectionFailed     ErrorReason = "ConnectionFailed"
	ErrorReasonNameNotResolved      ErrorReason = "NameNotResolved"
	ErrorReasonInternetDisconnected ErrorReason = "InternetDisconnected"
	ErrorReasonAddressUnreachable   ErrorReason = "AddressUnreachable"
	ErrorReasonBlockedByClient      ErrorReason = "BlockedByClient"
	ErrorReasonBlockedByResponse    ErrorReason = "BlockedByResponse"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ErrorReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ErrorReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ErrorReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ErrorReason(in.String()) {
	case ErrorReasonFailed:
		*t = ErrorReasonFailed
	case ErrorReasonAborted:
		*t = ErrorReasonAborted
	case ErrorReasonTimedOut:
		*t = ErrorReasonTimedOut
	case ErrorReasonAccessDenied:
		*t = ErrorReasonAccessDenied
	case ErrorReasonConnectionClosed:
		*t = ErrorReasonConnectionClosed
	case ErrorReasonConnectionReset:
		*t = ErrorReasonConnectionReset
	case ErrorReasonConnectionRefused:
		*t = ErrorReasonConnectionRefused
	case ErrorReasonConnectionAborted:
		*t = ErrorReasonConnectionAborted
	case ErrorReasonConnectionFailed:
		*t = ErrorReasonConnectionFailed
	case ErrorReasonNameNotResolved:
		*t = ErrorReasonNameNotResolved
	case ErrorReasonInternetDisconnected:
		*t = ErrorReasonInternetDisconnected
	case ErrorReasonAddressUnreachable:
		*t = ErrorReasonAddressUnreachable
	case ErrorReasonBlockedByClient:
		*t = ErrorReasonBlockedByClient
	case ErrorReasonBlockedByResponse:
		*t = ErrorReasonBlockedByResponse

	default:
		in.AddError(errors.New("unknown ErrorReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ErrorReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Headers request / response headers as keys / values of JSON object.
type Headers map[string]interface{}

// ConnectionType the underlying connection technology that the browser is
// supposedly using.
type ConnectionType string

// String returns the ConnectionType as string value.
func (t ConnectionType) String() string {
	return string(t)
}

// ConnectionType values.
const (
	ConnectionTypeNone       ConnectionType = "none"
	ConnectionTypeCellular2g ConnectionType = "cellular2g"
	ConnectionTypeCellular3g ConnectionType = "cellular3g"
	ConnectionTypeCellular4g ConnectionType = "cellular4g"
	ConnectionTypeBluetooth  ConnectionType = "bluetooth"
	ConnectionTypeEthernet   ConnectionType = "ethernet"
	ConnectionTypeWifi       ConnectionType = "wifi"
	ConnectionTypeWimax      ConnectionType = "wimax"
	ConnectionTypeOther      ConnectionType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ConnectionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ConnectionType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ConnectionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ConnectionType(in.String()) {
	case ConnectionTypeNone:
		*t = ConnectionTypeNone
	case ConnectionTypeCellular2g:
		*t = ConnectionTypeCellular2g
	case ConnectionTypeCellular3g:
		*t = ConnectionTypeCellular3g
	case ConnectionTypeCellular4g:
		*t = ConnectionTypeCellular4g
	case ConnectionTypeBluetooth:
		*t = ConnectionTypeBluetooth
	case ConnectionTypeEthernet:
		*t = ConnectionTypeEthernet
	case ConnectionTypeWifi:
		*t = ConnectionTypeWifi
	case ConnectionTypeWimax:
		*t = ConnectionTypeWimax
	case ConnectionTypeOther:
		*t = ConnectionTypeOther

	default:
		in.AddError(errors.New("unknown ConnectionType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ConnectionType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// CookieSameSite represents the cookie's 'SameSite' status:
// https://tools.ietf.org/html/draft-west-first-party-cookies.
type CookieSameSite string

// String returns the CookieSameSite as string value.
func (t CookieSameSite) String() string {
	return string(t)
}

// CookieSameSite values.
const (
	CookieSameSiteStrict CookieSameSite = "Strict"
	CookieSameSiteLax    CookieSameSite = "Lax"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CookieSameSite) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CookieSameSite) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CookieSameSite) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CookieSameSite(in.String()) {
	case CookieSameSiteStrict:
		*t = CookieSameSiteStrict
	case CookieSameSiteLax:
		*t = CookieSameSiteLax

	default:
		in.AddError(errors.New("unknown CookieSameSite value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CookieSameSite) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ResourceTiming timing information for the request.
type ResourceTiming struct {
	RequestTime       float64 `json:"requestTime"`       // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart        float64 `json:"proxyStart"`        // Started resolving proxy.
	ProxyEnd          float64 `json:"proxyEnd"`          // Finished resolving proxy.
	DNSStart          float64 `json:"dnsStart"`          // Started DNS address resolve.
	DNSEnd            float64 `json:"dnsEnd"`            // Finished DNS address resolve.
	ConnectStart      float64 `json:"connectStart"`      // Started connecting to the remote host.
	ConnectEnd        float64 `json:"connectEnd"`        // Connected to the remote host.
	SslStart          float64 `json:"sslStart"`          // Started SSL handshake.
	SslEnd            float64 `json:"sslEnd"`            // Finished SSL handshake.
	WorkerStart       float64 `json:"workerStart"`       // Started running ServiceWorker.
	WorkerReady       float64 `json:"workerReady"`       // Finished Starting ServiceWorker.
	SendStart         float64 `json:"sendStart"`         // Started sending request.
	SendEnd           float64 `json:"sendEnd"`           // Finished sending request.
	PushStart         float64 `json:"pushStart"`         // Time the server started pushing request.
	PushEnd           float64 `json:"pushEnd"`           // Time the server finished pushing request.
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"` // Finished receiving response headers.
}

// ResourcePriority loading priority of a resource request.
type ResourcePriority string

// String returns the ResourcePriority as string value.
func (t ResourcePriority) String() string {
	return string(t)
}

// ResourcePriority values.
const (
	ResourcePriorityVeryLow  ResourcePriority = "VeryLow"
	ResourcePriorityLow      ResourcePriority = "Low"
	ResourcePriorityMedium   ResourcePriority = "Medium"
	ResourcePriorityHigh     ResourcePriority = "High"
	ResourcePriorityVeryHigh ResourcePriority = "VeryHigh"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ResourcePriority) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ResourcePriority) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ResourcePriority) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ResourcePriority(in.String()) {
	case ResourcePriorityVeryLow:
		*t = ResourcePriorityVeryLow
	case ResourcePriorityLow:
		*t = ResourcePriorityLow
	case ResourcePriorityMedium:
		*t = ResourcePriorityMedium
	case ResourcePriorityHigh:
		*t = ResourcePriorityHigh
	case ResourcePriorityVeryHigh:
		*t = ResourcePriorityVeryHigh

	default:
		in.AddError(errors.New("unknown ResourcePriority value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ResourcePriority) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Request HTTP request data.
type Request struct {
	URL              string                    `json:"url"`                        // Request URL.
	Method           string                    `json:"method"`                     // HTTP request method.
	Headers          Headers                   `json:"headers"`                    // HTTP request headers.
	PostData         string                    `json:"postData,omitempty"`         // HTTP POST request data.
	HasPostData      bool                      `json:"hasPostData,omitempty"`      // True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
	MixedContentType security.MixedContentType `json:"mixedContentType,omitempty"` // The mixed content type of the request.
	InitialPriority  ResourcePriority          `json:"initialPriority"`            // Priority of the resource request at the time request is sent.
	ReferrerPolicy   ReferrerPolicy            `json:"referrerPolicy"`             // The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	IsLinkPreload    bool                      `json:"isLinkPreload,omitempty"`    // Whether is loaded via link preload.
}

// SignedCertificateTimestamp details of a signed certificate timestamp
// (SCT).
type SignedCertificateTimestamp struct {
	Status             string              `json:"status"`             // Validation status.
	Origin             string              `json:"origin"`             // Origin.
	LogDescription     string              `json:"logDescription"`     // Log name / description.
	LogID              string              `json:"logId"`              // Log ID.
	Timestamp          *cdp.TimeSinceEpoch `json:"timestamp"`          // Issuance date.
	HashAlgorithm      string              `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string              `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string              `json:"signatureData"`      // Signature data.
}

// SecurityDetails security details about a request.
type SecurityDetails struct {
	Protocol                          string                            `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                       string                            `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup                  string                            `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                            string                            `json:"cipher"`                            // Cipher name.
	Mac                               string                            `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateID                     security.CertificateID            `json:"certificateId"`                     // Certificate ID value.
	SubjectName                       string                            `json:"subjectName"`                       // Certificate subject name.
	SanList                           []string                          `json:"sanList"`                           // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                            string                            `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                         *cdp.TimeSinceEpoch               `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                           *cdp.TimeSinceEpoch               `json:"validTo"`                           // Certificate valid to (expiration) date
	SignedCertificateTimestampList    []*SignedCertificateTimestamp     `json:"signedCertificateTimestampList"`    // List of signed certificate timestamps (SCTs).
	CertificateTransparencyCompliance CertificateTransparencyCompliance `json:"certificateTransparencyCompliance"` // Whether the request complied with Certificate Transparency policy
}

// CertificateTransparencyCompliance whether the request complied with
// Certificate Transparency policy.
type CertificateTransparencyCompliance string

// String returns the CertificateTransparencyCompliance as string value.
func (t CertificateTransparencyCompliance) String() string {
	return string(t)
}

// CertificateTransparencyCompliance values.
const (
	CertificateTransparencyComplianceUnknown      CertificateTransparencyCompliance = "unknown"
	CertificateTransparencyComplianceNotCompliant CertificateTransparencyCompliance = "not-compliant"
	CertificateTransparencyComplianceCompliant    CertificateTransparencyCompliance = "compliant"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CertificateTransparencyCompliance) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CertificateTransparencyCompliance) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CertificateTransparencyCompliance) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CertificateTransparencyCompliance(in.String()) {
	case CertificateTransparencyComplianceUnknown:
		*t = CertificateTransparencyComplianceUnknown
	case CertificateTransparencyComplianceNotCompliant:
		*t = CertificateTransparencyComplianceNotCompliant
	case CertificateTransparencyComplianceCompliant:
		*t = CertificateTransparencyComplianceCompliant

	default:
		in.AddError(errors.New("unknown CertificateTransparencyCompliance value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CertificateTransparencyCompliance) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// BlockedReason the reason why request was blocked.
type BlockedReason string

// String returns the BlockedReason as string value.
func (t BlockedReason) String() string {
	return string(t)
}

// BlockedReason values.
const (
	BlockedReasonOther             BlockedReason = "other"
	BlockedReasonCsp               BlockedReason = "csp"
	BlockedReasonMixedContent      BlockedReason = "mixed-content"
	BlockedReasonOrigin            BlockedReason = "origin"
	BlockedReasonInspector         BlockedReason = "inspector"
	BlockedReasonSubresourceFilter BlockedReason = "subresource-filter"
	BlockedReasonContentType       BlockedReason = "content-type"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t BlockedReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t BlockedReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *BlockedReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch BlockedReason(in.String()) {
	case BlockedReasonOther:
		*t = BlockedReasonOther
	case BlockedReasonCsp:
		*t = BlockedReasonCsp
	case BlockedReasonMixedContent:
		*t = BlockedReasonMixedContent
	case BlockedReasonOrigin:
		*t = BlockedReasonOrigin
	case BlockedReasonInspector:
		*t = BlockedReasonInspector
	case BlockedReasonSubresourceFilter:
		*t = BlockedReasonSubresourceFilter
	case BlockedReasonContentType:
		*t = BlockedReasonContentType

	default:
		in.AddError(errors.New("unknown BlockedReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *BlockedReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Response HTTP response data.
type Response struct {
	URL                string           `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             int64            `json:"status"`                       // HTTP response status code.
	StatusText         string           `json:"statusText"`                   // HTTP response status text.
	Headers            Headers          `json:"headers"`                      // HTTP response headers.
	HeadersText        string           `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string           `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     Headers          `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText string           `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool             `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionID       float64          `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    string           `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         int64            `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      bool             `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  bool             `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	EncodedDataLength  float64          `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *ResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           string           `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      security.State   `json:"securityState"`                // Security state of the request resource.
	SecurityDetails    *SecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocketRequest webSocket request data.
type WebSocketRequest struct {
	Headers Headers `json:"headers"` // HTTP request headers.
}

// WebSocketResponse webSocket response data.
type WebSocketResponse struct {
	Status             int64   `json:"status"`                       // HTTP response status code.
	StatusText         string  `json:"statusText"`                   // HTTP response status text.
	Headers            Headers `json:"headers"`                      // HTTP response headers.
	HeadersText        string  `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     Headers `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText string  `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocketFrame webSocket frame data.
type WebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket frame opcode.
	Mask        bool    `json:"mask"`        // WebSocke frame mask.
	PayloadData string  `json:"payloadData"` // WebSocke frame payload data.
}

// CachedResource information about the cached resource.
type CachedResource struct {
	URL      string            `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     page.ResourceType `json:"type"`               // Type of this resource.
	Response *Response         `json:"response,omitempty"` // Cached response data.
	BodySize float64           `json:"bodySize"`           // Cached response body size.
}

// Initiator information about the request initiator.
type Initiator struct {
	Type       InitiatorType       `json:"type"`                 // Type of this initiator.
	Stack      *runtime.StackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	URL        string              `json:"url,omitempty"`        // Initiator URL, set for Parser type or for Script type (when script is importing module) or for SignedExchange type.
	LineNumber float64             `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
}

// Cookie cookie object.
type Cookie struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	Domain   string         `json:"domain"`             // Cookie domain.
	Path     string         `json:"path"`               // Cookie path.
	Expires  float64        `json:"expires"`            // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size     int64          `json:"size"`               // Cookie size.
	HTTPOnly bool           `json:"httpOnly"`           // True if cookie is http-only.
	Secure   bool           `json:"secure"`             // True if cookie is secure.
	Session  bool           `json:"session"`            // True in case of session cookie.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
}

// CookieParam cookie parameter object.
type CookieParam struct {
	Name     string              `json:"name"`               // Cookie name.
	Value    string              `json:"value"`              // Cookie value.
	URL      string              `json:"url,omitempty"`      // The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain   string              `json:"domain,omitempty"`   // Cookie domain.
	Path     string              `json:"path,omitempty"`     // Cookie path.
	Secure   bool                `json:"secure,omitempty"`   // True if cookie is secure.
	HTTPOnly bool                `json:"httpOnly,omitempty"` // True if cookie is http-only.
	SameSite CookieSameSite      `json:"sameSite,omitempty"` // Cookie SameSite type.
	Expires  *cdp.TimeSinceEpoch `json:"expires,omitempty"`  // Cookie expiration date, session cookie if not set
}

// AuthChallenge authorization challenge for HTTP status code 401 or 407.
type AuthChallenge struct {
	Source AuthChallengeSource `json:"source,omitempty"` // Source of the authentication challenge.
	Origin string              `json:"origin"`           // Origin of the challenger.
	Scheme string              `json:"scheme"`           // The authentication scheme used, such as basic or digest
	Realm  string              `json:"realm"`            // The realm of the challenge. May be empty.
}

// AuthChallengeResponse response to an AuthChallenge.
type AuthChallengeResponse struct {
	Response AuthChallengeResponseResponse `json:"response"`           // The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username string                        `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password string                        `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// InterceptionStage stages of the interception to begin intercepting.
// Request will intercept before the request is sent. Response will intercept
// after the response is received.
type InterceptionStage string

// String returns the InterceptionStage as string value.
func (t InterceptionStage) String() string {
	return string(t)
}

// InterceptionStage values.
const (
	InterceptionStageRequest         InterceptionStage = "Request"
	InterceptionStageHeadersReceived InterceptionStage = "HeadersReceived"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InterceptionStage) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InterceptionStage) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InterceptionStage) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch InterceptionStage(in.String()) {
	case InterceptionStageRequest:
		*t = InterceptionStageRequest
	case InterceptionStageHeadersReceived:
		*t = InterceptionStageHeadersReceived

	default:
		in.AddError(errors.New("unknown InterceptionStage value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InterceptionStage) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// RequestPattern request pattern for interception.
type RequestPattern struct {
	URLPattern        string            `json:"urlPattern,omitempty"`        // Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	ResourceType      page.ResourceType `json:"resourceType,omitempty"`      // If set, only requests for matching resource types will be intercepted.
	InterceptionStage InterceptionStage `json:"interceptionStage,omitempty"` // Stage at which to begin intercepting requests. Default is Request.
}

// SignedExchangeSignature information about a signed exchange signature.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1.
type SignedExchangeSignature struct {
	Label        string   `json:"label"`                  // Signed exchange signature label.
	Signature    string   `json:"signature"`              // The hex string of signed exchange signature.
	Integrity    string   `json:"integrity"`              // Signed exchange signature integrity.
	CertURL      string   `json:"certUrl,omitempty"`      // Signed exchange signature cert Url.
	CertSha256   string   `json:"certSha256,omitempty"`   // The hex string of signed exchange signature cert sha256.
	ValidityURL  string   `json:"validityUrl"`            // Signed exchange signature validity Url.
	Date         int64    `json:"date"`                   // Signed exchange signature date.
	Expires      int64    `json:"expires"`                // Signed exchange signature expires.
	Certificates []string `json:"certificates,omitempty"` // The encoded certificates.
}

// SignedExchangeHeader information about a signed exchange header.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation.
type SignedExchangeHeader struct {
	RequestURL      string                     `json:"requestUrl"`      // Signed exchange request URL.
	RequestMethod   string                     `json:"requestMethod"`   // Signed exchange request method.
	ResponseCode    int64                      `json:"responseCode"`    // Signed exchange response code.
	ResponseHeaders Headers                    `json:"responseHeaders"` // Signed exchange response headers.
	Signatures      []*SignedExchangeSignature `json:"signatures"`      // Signed exchange response signature.
}

// SignedExchangeErrorField field type for a signed exchange related error.
type SignedExchangeErrorField string

// String returns the SignedExchangeErrorField as string value.
func (t SignedExchangeErrorField) String() string {
	return string(t)
}

// SignedExchangeErrorField values.
const (
	SignedExchangeErrorFieldSignatureSig         SignedExchangeErrorField = "signatureSig"
	SignedExchangeErrorFieldSignatureIntegrity   SignedExchangeErrorField = "signatureIntegrity"
	SignedExchangeErrorFieldSignatureCertURL     SignedExchangeErrorField = "signatureCertUrl"
	SignedExchangeErrorFieldSignatureCertSha256  SignedExchangeErrorField = "signatureCertSha256"
	SignedExchangeErrorFieldSignatureValidityURL SignedExchangeErrorField = "signatureValidityUrl"
	SignedExchangeErrorFieldSignatureTimestamps  SignedExchangeErrorField = "signatureTimestamps"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SignedExchangeErrorField) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SignedExchangeErrorField) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SignedExchangeErrorField) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SignedExchangeErrorField(in.String()) {
	case SignedExchangeErrorFieldSignatureSig:
		*t = SignedExchangeErrorFieldSignatureSig
	case SignedExchangeErrorFieldSignatureIntegrity:
		*t = SignedExchangeErrorFieldSignatureIntegrity
	case SignedExchangeErrorFieldSignatureCertURL:
		*t = SignedExchangeErrorFieldSignatureCertURL
	case SignedExchangeErrorFieldSignatureCertSha256:
		*t = SignedExchangeErrorFieldSignatureCertSha256
	case SignedExchangeErrorFieldSignatureValidityURL:
		*t = SignedExchangeErrorFieldSignatureValidityURL
	case SignedExchangeErrorFieldSignatureTimestamps:
		*t = SignedExchangeErrorFieldSignatureTimestamps

	default:
		in.AddError(errors.New("unknown SignedExchangeErrorField value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SignedExchangeErrorField) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SignedExchangeError information about a signed exchange response.
type SignedExchangeError struct {
	Message        string                   `json:"message"`                  // Error message.
	SignatureIndex int64                    `json:"signatureIndex,omitempty"` // The index of the signature which caused the error.
	ErrorField     SignedExchangeErrorField `json:"errorField,omitempty"`     // The field which caused the error.
}

// SignedExchangeInfo information about a signed exchange response.
type SignedExchangeInfo struct {
	OuterResponse   *Response              `json:"outerResponse"`             // The outer response of signed HTTP exchange which was received from network.
	Header          *SignedExchangeHeader  `json:"header,omitempty"`          // Information about the signed exchange header.
	SecurityDetails *SecurityDetails       `json:"securityDetails,omitempty"` // Security details for the signed exchange header.
	Errors          []*SignedExchangeError `json:"errors,omitempty"`          // Errors occurred while handling the signed exchagne.
}

// ReferrerPolicy the referrer policy of the request, as defined in
// https://www.w3.org/TR/referrer-policy/.
type ReferrerPolicy string

// String returns the ReferrerPolicy as string value.
func (t ReferrerPolicy) String() string {
	return string(t)
}

// ReferrerPolicy values.
const (
	ReferrerPolicyUnsafeURL                   ReferrerPolicy = "unsafe-url"
	ReferrerPolicyNoReferrerWhenDowngrade     ReferrerPolicy = "no-referrer-when-downgrade"
	ReferrerPolicyNoReferrer                  ReferrerPolicy = "no-referrer"
	ReferrerPolicyOrigin                      ReferrerPolicy = "origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicy = "origin-when-cross-origin"
	ReferrerPolicySameOrigin                  ReferrerPolicy = "same-origin"
	ReferrerPolicyStrictOrigin                ReferrerPolicy = "strict-origin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strict-origin-when-cross-origin"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ReferrerPolicy) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ReferrerPolicy) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ReferrerPolicy) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ReferrerPolicy(in.String()) {
	case ReferrerPolicyUnsafeURL:
		*t = ReferrerPolicyUnsafeURL
	case ReferrerPolicyNoReferrerWhenDowngrade:
		*t = ReferrerPolicyNoReferrerWhenDowngrade
	case ReferrerPolicyNoReferrer:
		*t = ReferrerPolicyNoReferrer
	case ReferrerPolicyOrigin:
		*t = ReferrerPolicyOrigin
	case ReferrerPolicyOriginWhenCrossOrigin:
		*t = ReferrerPolicyOriginWhenCrossOrigin
	case ReferrerPolicySameOrigin:
		*t = ReferrerPolicySameOrigin
	case ReferrerPolicyStrictOrigin:
		*t = ReferrerPolicyStrictOrigin
	case ReferrerPolicyStrictOriginWhenCrossOrigin:
		*t = ReferrerPolicyStrictOriginWhenCrossOrigin

	default:
		in.AddError(errors.New("unknown ReferrerPolicy value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ReferrerPolicy) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// InitiatorType type of this initiator.
type InitiatorType string

// String returns the InitiatorType as string value.
func (t InitiatorType) String() string {
	return string(t)
}

// InitiatorType values.
const (
	InitiatorTypeParser         InitiatorType = "parser"
	InitiatorTypeScript         InitiatorType = "script"
	InitiatorTypePreload        InitiatorType = "preload"
	InitiatorTypeSignedExchange InitiatorType = "SignedExchange"
	InitiatorTypeOther          InitiatorType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InitiatorType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InitiatorType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InitiatorType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch InitiatorType(in.String()) {
	case InitiatorTypeParser:
		*t = InitiatorTypeParser
	case InitiatorTypeScript:
		*t = InitiatorTypeScript
	case InitiatorTypePreload:
		*t = InitiatorTypePreload
	case InitiatorTypeSignedExchange:
		*t = InitiatorTypeSignedExchange
	case InitiatorTypeOther:
		*t = InitiatorTypeOther

	default:
		in.AddError(errors.New("unknown InitiatorType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InitiatorType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AuthChallengeSource source of the authentication challenge.
type AuthChallengeSource string

// String returns the AuthChallengeSource as string value.
func (t AuthChallengeSource) String() string {
	return string(t)
}

// AuthChallengeSource values.
const (
	AuthChallengeSourceServer AuthChallengeSource = "Server"
	AuthChallengeSourceProxy  AuthChallengeSource = "Proxy"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AuthChallengeSource) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AuthChallengeSource) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AuthChallengeSource) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AuthChallengeSource(in.String()) {
	case AuthChallengeSourceServer:
		*t = AuthChallengeSourceServer
	case AuthChallengeSourceProxy:
		*t = AuthChallengeSourceProxy

	default:
		in.AddError(errors.New("unknown AuthChallengeSource value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AuthChallengeSource) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AuthChallengeResponseResponse the decision on what to do in response to
// the authorization challenge. Default means deferring to the default behavior
// of the net stack, which will likely either the Cancel authentication or
// display a popup dialog box.
type AuthChallengeResponseResponse string

// String returns the AuthChallengeResponseResponse as string value.
func (t AuthChallengeResponseResponse) String() string {
	return string(t)
}

// AuthChallengeResponseResponse values.
const (
	AuthChallengeResponseResponseDefault            AuthChallengeResponseResponse = "Default"
	AuthChallengeResponseResponseCancelAuth         AuthChallengeResponseResponse = "CancelAuth"
	AuthChallengeResponseResponseProvideCredentials AuthChallengeResponseResponse = "ProvideCredentials"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AuthChallengeResponseResponse) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AuthChallengeResponseResponse) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AuthChallengeResponseResponse) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AuthChallengeResponseResponse(in.String()) {
	case AuthChallengeResponseResponseDefault:
		*t = AuthChallengeResponseResponseDefault
	case AuthChallengeResponseResponseCancelAuth:
		*t = AuthChallengeResponseResponseCancelAuth
	case AuthChallengeResponseResponseProvideCredentials:
		*t = AuthChallengeResponseResponseProvideCredentials

	default:
		in.AddError(errors.New("unknown AuthChallengeResponseResponse value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AuthChallengeResponseResponse) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
