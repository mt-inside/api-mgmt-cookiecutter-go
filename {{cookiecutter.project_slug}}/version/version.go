package version

import (
	"fmt"
	"runtime"
)

var Name string
var Version string
var GitCommit string
var BuildDate string
var GoVersion = runtime.Version()
var GoOsArch = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
