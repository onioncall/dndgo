package defaultjsonconfigs

import (
	"embed"
)

//go:embed *.json
var Content embed.FS
