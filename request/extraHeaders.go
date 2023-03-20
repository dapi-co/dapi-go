package request

type goExtraHeadersType struct {
	LibraryVersion  string
	LibraryPlatform string
}

var goExtraHeaders = goExtraHeadersType{
	LibraryVersion:  "1.4.0",
	LibraryPlatform: "dapi-go",
}
