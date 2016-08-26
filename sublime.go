package subpro

import (
	"encoding/json"
)

type SublSetting struct {
	Folders  []Folder    `json:"folders"`
	Settings interface{} `json:"settings"`
}

type Folder struct {
	FollowSymlinks        bool        `json:"follow_symlinks"`
	Path                  string      `json:"path"`
	FileExcludePatterns   interface{} `json:"file_exclude_patterns"`
	FolderExcludePatterns interface{} `json:"folder_exclude_patterns"`
}

func MarshalSetting(setting SublSetting) []byte {
	b, err := json.MarshalIndent(setting, "", "  ")
	if err != nil {
		panic(err)
	}

	return b
}

func UnMarshalSetting(b []byte) SublSetting {
	var sublSetting SublSetting
	err := json.Unmarshal(b, &sublSetting)
	if err != nil {
		panic(err)
	}
	return sublSetting
}
