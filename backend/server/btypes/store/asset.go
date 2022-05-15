package store

import "io/fs"

// AssetStore is read only interface for assets
type AssetStore interface {
	GetSchema(name string) ([]byte, error)
	GetTemplate(name string) ([]byte, error)
	GetAsset(name string) ([]byte, error)
	AssetAdapter() fs.FS
}
