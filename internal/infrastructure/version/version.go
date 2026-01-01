package version

import (
	"encoding/json"
	"os"
)

type Version struct {
	writer *json.Encoder
}

func New() *Version {
	writer := json.NewEncoder(os.Stdout)
	writer.SetIndent("", "  ")
	return &Version{
		writer: writer,
	}
}

func (version *Version) Selected() bool {
	return len(os.Args) > 1 && os.Args[1] == "version"
}

func (version *Version) Print(data map[string]string) {
	version.writer.Encode(data)
}
