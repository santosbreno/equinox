/*
package proto defines a set of structures used to negotiate an update between an
an application (the client) and an equinox update service.
*/
package proto

import "time"

type Request struct {
	AppID         string `json:"app_id"`
	Channel       string `json:"channel"`
	OS            string `json:"os"`
	Arch          string `json:"arch"`
	GoARM         string `json:"goarm"`
	TargetVersion string `json:"target_version"`

	CurrentVersion string `json:"current_version"`
	CurrentSHA256  string `json:"current_sha256"`
}

type Response struct {
	Available   bool    `json:"available"`
	DownloadURL string  `json:"download_url"`
	Checksum    string  `json:"checksum"`
	Signature   string  `json:"signature"`
	PatchType   string  `json:"patch_type"`
	Version     string  `json:"version"`
	Release     Release `json:"release"`
}

type Release struct {
	Title       string    `json:"title"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	CreateDate  time.Time `json:"create_date"`
}
