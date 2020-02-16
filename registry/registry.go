package registry

import (
	"encoding/json"
	"net/url"
	"path"
)

// Registry interacts with a NPM registry
type Registry struct {
	client *Client
}

// NewRegistry creates a new Registry taking the registries url as argument
func NewRegistry(url *url.URL) *Registry {
	client := Client{BaseURL: url}
	return &Registry{client: &client}
}

// Metadata represents NPM package registry metadata
// See https://github.com/npm/registry/blob/master/docs/responses/package-metadata.md
type Metadata struct {
	Name     string
	Versions map[string]Manifest
}

// Metadata fetches the package metadata of a package (https://github.com/npm/registry/blob/master/docs/responses/package-metadata.md)
func (r *Registry) Metadata(pkg string) (*Metadata, error) {
	req, err := r.client.NewRequest("GET", pkg, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var m *Metadata
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}

// Manifest fetches the manifest of a package
func (r *Registry) Manifest(pkg, version string) (*Manifest, error) {
	req, err := r.client.NewRequest("GET", path.Join(pkg, version), nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var m *Manifest
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}

// Manifest represents the package.json
type Manifest struct {
	Name         string
	Version      string
	Dependencies map[string]string
	Dist         struct {
		Tarball string
	}
}
