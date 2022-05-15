package pacman

import (
	"context"
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
	"github.com/temphia/temphia/backend/server/btypes/slugger"
	"github.com/ztrue/tracerr"
	"gopkg.in/yaml.v2"
)

func (b *PacMan) extractSchema(tenantId string, model *entities.BPrint, inline string) ([]byte, error) {
	if inline != "" {
		return []byte(inline), nil
	}
	if model.InlineSchema != "" {
		return []byte(model.InlineSchema), nil
	}

	return b.blobStore(tenantId).GetBlob(context.Background(), btypes.BprintBlobFolder, ffile(model.ID, "schema.json"))
}

func (b *PacMan) extractDynGroupSchema(tenantId string, model *entities.BPrint, inline string) (*entities.NewTableGroup, error) {

	var data []byte
	var isYaml bool

	data, err1 := b.extractSchema(tenantId, model, inline)
	if err1 != nil {
		pp.Println("cannot read schema.json")

		// we only support yaml as fallback
		data2, err2 := b.blobStore(tenantId).GetBlob(context.Background(), btypes.BprintBlobFolder, ffile(model.ID, "schema.yaml"))
		if err2 != nil {
			pp.Println("cannot read schema.json")
			return nil, tracerr.Wrap(err1)
		}
		isYaml = true
		data = data2
	}

	if isYaml {
		tg := &entities.NewTableGroup{}
		err := yaml.Unmarshal(data, tg)
		if err != nil {
			return nil, err
		}
		return tg, nil
	} else {
		tg := &entities.NewTableGroup{}
		err := json.Unmarshal(data, tg)
		if err != nil {
			return nil, err
		}
		return tg, nil

	}

}

func (b *PacMan) installPlug(tenantId string, model *entities.BPrint, opts *vmodels.PlugInstallOptions) (*vmodels.PlugInstallResponse, error) {
	schema, err := b.extractSchema(tenantId, model, opts.Schema)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	raw := &vmodels.PlugRaw{}
	err = json.Unmarshal(schema, raw)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	if opts.NewPlugId == "" {
		opts.NewPlugId = xid.New().String()
	}

	plug := &entities.Plug{
		Id:        opts.NewPlugId,
		TenantId:  tenantId,
		Name:      raw.Name,
		Executor:  raw.Executor,
		Live:      true,
		Dev:       true,
		ExtraMeta: nil,
		Owner:     "",
		BprintId:  model.ID,
		Handlers:  raw.HandlerHints,
	}

	err = b.syncer.PlugNew(tenantId, plug)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	resp := &vmodels.PlugInstallResponse{
		Agents:       make(map[string]string),
		Resources:    make(map[string]string),
		ErrAgents:    make(map[string]string),
		ErrResources: make(map[string]string),
	}

	for _, resName := range opts.Resources {

		pp.Println("creating resource...", resName)

		hint := raw.ResourceHints[resName]

		resid := xid.New().String()

		resource := &entities.Resource{
			Name:      hint.Name,
			TenantId:  tenantId,
			Type:      hint.Type,
			SubType:   hint.SubType,
			Payload:   hint.Payload,
			Policy:    hint.Policy,
			PlugId:    opts.NewPlugId,
			Id:        resid,
			ExtraMeta: hint.Meta,
		}

		err := b.syncer.ResourceNew(tenantId, resource)
		pp.Println("@@res", err)

		if err != nil {
			resp.ErrResources[resName] = err.Error()
			continue
		}
		resp.Resources[resName] = resid
	}

	pp.Println("All agents", opts.Agents)

	for _, aname := range opts.Agents {

		hintAgent := raw.AgentHints[aname]
		aid := slugger.Slugify(aname)

		for k := range hintAgent.Resources {
			actualResId, ok := resp.Resources[k]
			if !ok {
				continue
			}
			hintAgent.Resources[k] = actualResId
		}

		agent := &entities.Agent{
			Id:           aid,
			Name:         aname,
			Type:         hintAgent.Type,
			InvokePolicy: hintAgent.InvokePolicy,
			PlugID:       opts.NewPlugId,
			Resources:    hintAgent.Resources,
			ServeFiles:   hintAgent.ServeFiles,
			EntryName:    hintAgent.EntryName,
			EntryScript:  hintAgent.EntryScript,
			EntryStyle:   hintAgent.EntryStyle,
			ExtScripts:   nil,
			ExtraMeta:    entities.JsonStrMap{},
			TenantId:     tenantId,
			ExecLoader:   hintAgent.ExecLoader,
		}

		pp.Println(agent)

		err := b.syncer.AgentNew(tenantId, agent)
		if err != nil {
			pp.Println(err)
			resp.ErrAgents[aname] = err.Error()
			continue
		}

		resp.Agents[aname] = aid
	}

	return resp, nil
}

func (b *PacMan) installDGroup(tenantId string, model *entities.BPrint, opts *vmodels.DGroupInstallOptions) error {
	tg, err := b.extractDynGroupSchema(tenantId, model, opts.Schema)
	if err != nil {
		return err
	}

	pp.Println("@ => 1")

	if opts.GroupSlug != "" {
		tg.Slug = opts.GroupSlug
	}

	if opts.GroupName != "" {
		tg.Name = opts.GroupName
	}

	// fixme =>
	if opts.DyndbSource == "" {
		opts.DyndbSource = "default_db"
	}

	dynSrc := b.dynHub.GetSource(opts.DyndbSource, tenantId)

	err = dynSrc.NewGroup(tg)
	if err != nil {
		return err
	}

	for _, tbl := range tg.Tables {
		for _, view := range tbl.Views {

			dynSrc.NewView(&entities.DataView{
				Id:          0,
				Name:        view.Name,
				Count:       view.Count,
				FilterConds: view.FilterConds,
				Selects:     view.Selects,
				MainColumn:  view.MainColumn,
				SearchTerm:  view.SearchTerm,
				TableID:     tbl.Slug,
				GroupID:     opts.GroupName,
				TenantID:    tenantId,
				ExtraMeta:   nil,
			})

		}

	}

	seeder := Seeder{
		tg:     tg,
		model:  model,
		pacman: b,
		source: dynSrc,
		tenant: tenantId,
		group:  opts.GroupSlug,
		userId: opts.UserId,
	}

	switch opts.SeedFrom {
	case "data":
		return seeder.dataSeed()
	case "autogen":
		return seeder.generatedSeed(200)
	}

	return nil
}

func (b *PacMan) installDTable(tenantId string, model *entities.BPrint, opts *vmodels.DTableInstallOptions) error {
	schema, err := b.extractSchema(tenantId, model, opts.Schema)
	if err != nil {
		return err
	}

	tbl := &entities.NewTable{}
	err = json.Unmarshal(schema, tbl)
	if err != nil {
		return err
	}
	err = b.dynHub.GetSource(opts.TargetSource, tenantId).AddTable(opts.TargetGroupId, tbl)
	if err != nil {
		return err
	}

	if opts.SeedRandom {
		// fixme => seed here
	}

	return nil
}
