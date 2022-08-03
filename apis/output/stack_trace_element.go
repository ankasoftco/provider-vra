package models 
type StackTraceElement struct {

	// class loader name
	ClassLoaderName string `json:"classLoaderName,omitempty"`

	// file name
	FileName string `json:"fileName,omitempty"`

	// line number
	LineNumber int32 `json:"lineNumber,omitempty"`

	// method name
	MethodName string `json:"methodName,omitempty"`

	// module name
	ModuleName string `json:"moduleName,omitempty"`

	// module version
	ModuleVersion string `json:"moduleVersion,omitempty"`
}

