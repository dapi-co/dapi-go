package constants

type goExtraHeadersType struct {
	LibraryVersion  string
	LibraryPlatform string
}

var GoExtraHeaders = goExtraHeadersType{
	LibraryVersion:  "1.4.0",
	LibraryPlatform: "dapi-go",
}
