package models 
type URI struct {

	// absolute
	Absolute bool `json:"absolute,omitempty"`

	// authority
	Authority string `json:"authority,omitempty"`

	// fragment
	Fragment string `json:"fragment,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// opaque
	Opaque bool `json:"opaque,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// raw authority
	RawAuthority string `json:"rawAuthority,omitempty"`

	// raw fragment
	RawFragment string `json:"rawFragment,omitempty"`

	// raw path
	RawPath string `json:"rawPath,omitempty"`

	// raw query
	RawQuery string `json:"rawQuery,omitempty"`

	// raw scheme specific part
	RawSchemeSpecificPart string `json:"rawSchemeSpecificPart,omitempty"`

	// raw user info
	RawUserInfo string `json:"rawUserInfo,omitempty"`

	// scheme
	Scheme string `json:"scheme,omitempty"`

	// scheme specific part
	SchemeSpecificPart string `json:"schemeSpecificPart,omitempty"`

	// user info
	UserInfo string `json:"userInfo,omitempty"`
}

