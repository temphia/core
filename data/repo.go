package data

import (
	"encoding/json"

	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/store"
)

type EmbedRepo struct {
	assetStore *AssetStore
	TenantId   string
	BaseURL    string
}

func NewEmbed(opts *store.RepoBuilderOptions) (store.Repository, error) {
	return &EmbedRepo{
		assetStore: defaultAssetStore,
		TenantId:   opts.TenantId,
		BaseURL:    opts.BaseURL,
	}, nil
}

func (lr *EmbedRepo) Name() string {
	return "embed"
}

func (lr *EmbedRepo) Query(tenantId string, opts *store.RepoQuery) ([]entities.BPrint, error) {

	out, err := lr.assetStore.tryRead("", "repo", "index.json")
	if err != nil {
		return nil, err
	}

	index := map[string]string{}
	err = json.Unmarshal(out, &index)
	if err != nil {
		return nil, err
	}

	bprints := make([]entities.BPrint, 0, len(index))
	for _, v := range index {
		bout, err := lr.assetStore.tryRead("", "repo", v)
		if err != nil {
			continue
		}
		bprint := entities.BPrint{}
		err = json.Unmarshal(bout, &bprint)
		if err != nil {
			continue
		}

		if opts.Group != "" {
			if bprint.Type != opts.Group {
				continue
			}
		}

		bprints = append(bprints, bprint)
	}
	return bprints, nil
}

func (lr *EmbedRepo) GetItem(tenantid, group, slug string) (*entities.BPrint, error) {
	bout, err := lr.assetStore.tryRead("", "repo", (slug + "_index.json"))
	if err != nil {
		return nil, err
	}
	bprint := &entities.BPrint{}
	err = json.Unmarshal(bout, bprint)
	if err != nil {
		return nil, err
	}

	return bprint, nil
}

func (lr *EmbedRepo) GetFile(tenantid, group, slug, file string) ([]byte, error) {
	return lr.assetStore.tryRead("", "repo", (slug + "_" + file))
}

func (lr *EmbedRepo) GetFileURL(tenantid, group, slug, file string) (string, error) {
	// fmt.Sprintf("%s", lr.BaseURL)
	return "", nil
}
