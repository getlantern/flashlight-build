package server

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/getlantern/golog"
)

var log = golog.LoggerFor("autoupdate-server")

// Initiative type.
type Initiative string

const (
	INITIATIVE_NEVER  Initiative = "never"
	INITIATIVE_AUTO              = "auto"
	INITIATIVE_MANUAL            = "manual"
)

// PatchType represents the type of a binary patch, if any. Only bsdiff is supported
type PatchType string

const (
	PATCHTYPE_BSDIFF PatchType = "bsdiff"
	PATCHTYPE_NONE             = ""
)

// Params represent parameters sent by the go-update client.
type Params struct {
	// protocol version
	Version int `json:"version"`
	// identifier of the application to update
	//AppId string `json:"app_id"`

	// version of the application updating itself
	AppVersion string `json:"app_version"`
	// operating system of target platform
	OS string `json:"-"`
	// hardware architecture of target platform
	Arch string `json:"-"`
	// application-level user identifier
	//UserId string `json:"user_id"`
	// checksum of the binary to replace (used for returning diff patches)
	Checksum string `json:"checksum"`
	// release channel (empty string means 'stable')
	//Channel string `json:"-"`
	// tags for custom update channels
	Tags map[string]string `json:"tags"`
}

// Result represents the answer to be sent to the client.
type Result struct {
	// should the update be applied automatically/manually
	Initiative Initiative `json:"initiative"`
	// url where to download the updated application
	URL string `json:"url"`
	// a URL to a patch to apply
	PatchURL string `json:"patch_url"`
	// the patch format (only bsdiff supported at the moment)
	PatchType PatchType `json:"patch_type"`
	// version of the new application
	Version string `json:"version"`
	// expected checksum of the new application
	Checksum string `json:"checksum"`
	// signature for verifying update authenticity
	Signature string `json:"signature"`
}

// CheckForUpdate receives a *Params message and emits a *Result. If both res
// and err are nil it means no update is available.
func (g *ReleaseManager) CheckForUpdate(p *Params) (res *Result, err error) {

	// Keep for the future.
	if p.Version < 1 {
		p.Version = 1
	}

	// p must not be nil.
	if p == nil {
		return nil, fmt.Errorf("Expecting params")
	}

	if p.Tags != nil {
		// Compatibility with go-check.
		if p.Tags["os"] != "" {
			p.OS = p.Tags["os"]
		}
		if p.Tags["arch"] != "" {
			p.Arch = p.Tags["arch"]
		}
	}

	appVersion, err := semver.New(p.AppVersion)
	if err != nil {
		return nil, fmt.Errorf("Bad version string: %v", err)
	}

	if p.Checksum == "" {
		return nil, fmt.Errorf("Checksum must not be nil")
	}

	if p.OS == "" {
		return nil, fmt.Errorf("OS is required")
	}

	if p.Arch == "" {
		return nil, fmt.Errorf("Arch is required")
	}

	// Looking if there is a newer version for the os/arch.
	var update *Asset
	if update, err = g.getProductUpdate(p.OS, p.Arch); err != nil {
		return nil, fmt.Errorf("Could not lookup for updates: %s", err)
	}

	// Looking for the asset thay matches the current app checksum.
	var current *Asset
	if current, err = g.lookupAssetWithChecksum(p.OS, p.Arch, p.Checksum); err != nil {
		// No such asset with the given checksum, nothing to compare.

		r := &Result{
			Initiative: INITIATIVE_AUTO,
			URL:        update.URL,
			PatchType:  PATCHTYPE_NONE,
			Version:    update.v.String(),
			Checksum:   update.Checksum,
			Signature:  update.Signature,
		}

		return r, nil
	}

	// No update available.
	if update.v.LTE(appVersion) {
		return nil, ErrNoUpdateAvailable
	}

	// A newer version is available!

	// Generate a binary diff of the two assets.
	var patch *Patch
	log.Debugf("Generating patch")

	if patch, err = GeneratePatch(current.URL, update.URL); err != nil {
		log.Debugf("Unable to generate patch: %q", err)
		return nil, fmt.Errorf("Unable to generate patch: %q", err)
	}

	log.Debugf("Patch was generated.")

	// Generate result.
	r := &Result{
		Initiative: INITIATIVE_AUTO,
		URL:        update.URL,
		PatchURL:   patch.File,
		PatchType:  PATCHTYPE_BSDIFF,
		Version:    update.v.String(),
		Checksum:   update.Checksum,
		Signature:  update.Signature,
	}

	return r, nil
}
