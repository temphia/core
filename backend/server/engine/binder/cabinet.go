package binder

import (
	"context"

	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/store"
)

func (b *Binder) resourceFolder(name string) string {
	return b.resourced(name)
}

func (b *Binder) cabSrc(name string) store.CabinetSourced {
	return b.factory.App.Cabinet().Default(b.namespace)
}

func (b *Binder) AddFile(bucket string, file string, contents []byte) error {
	bucket = b.resourceFolder(bucket)
	source := b.cabSrc(bucket)

	return source.AddBlob(context.TODO(), bucket, file, contents)
}

func (b *Binder) ListFolder(bucket string) ([]string, error) {
	bucket = b.resourceFolder(bucket)
	source := b.cabSrc(bucket)

	blobs, err := source.ListFolder(context.TODO(), bucket)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0, len(blobs))
	for _, bi := range blobs {
		files = append(files, bi.Name)
	}
	return files, nil
}

func (b *Binder) GetFile(bucket string, file string) ([]byte, error) {
	bucket = b.resourceFolder(bucket)
	source := b.cabSrc(bucket)
	return source.GetBlob(context.TODO(), bucket, file)
}

func (b *Binder) DeleteFile(bucket string, file string) error {
	bucket = b.resourceFolder(bucket)
	source := b.cabSrc(bucket)
	return source.DeleteBlob(context.TODO(), bucket, file)
}

func (b *Binder) GenerateTicket(bucket string, ticket *rtypes.CabTicket) (string, error) {
	return "", nil
}
