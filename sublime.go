package main

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

func Marshal(setting SublSetting) []byte {
	b, _ := json.MarshalIndent(setting, "", "  ")
	return b
}

func UnMarshal(b []byte) SublSetting {
	var sublSetting SublSetting
	json.Unmarshal(b, &sublSetting)
	return sublSetting
}
