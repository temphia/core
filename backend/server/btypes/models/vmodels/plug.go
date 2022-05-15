package vmodels

import "encoding/json"

type PlugRaw struct {
	Slug            string                    `json:"slug,omitempty"`
	Name            string                    `json:"name,omitempty"`
	Executor        string                    `json:"executor,omitempty"`
	HandlerHints    map[string]string         `json:"handler_hints,omitempty"`
	ResourceHints   map[string]ResourceHint   `json:"resource_hints,omitempty"`
	AgentHints      map[string]AgentHint      `json:"agent_hints,omitempty"`
	ConnectionHints map[string]ConnectionHint `json:"conn_hints,omitempty"`
}

type ConnectionHint struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Schema string `json:"schema,omitempty"`
}

type AgentHint struct {
	Type         string            `json:"type,omitempty"`
	InvokePolicy string            `json:"invoke_policy,omitempty"`
	EntryName    string            `json:"entry_name,omitempty"`
	EntryScript  string            `json:"entry_script,omitempty"`
	EntryStyle   string            `json:"entry_style,omitempty"`
	ExecLoader   string            `json:"exec_loader,omitempty"`
	Resources    map[string]string `json:"resources,omitempty"`
	ServeFiles   map[string]string `json:"serve_files,omitempty"`
}

type ResourceHint struct {
	Name    string            `json:"name,omitempty"`
	Type    string            `json:"type,omitempty"`
	SubType string            `json:"sub_type,omitempty"`
	Payload string            `json:"schema,omitempty"`
	Policy  string            `json:"policy,omitempty"`
	Meta    map[string]string `json:"meta,omitempty"`
}

// used in template rendering

type SubOriginData struct {
	LoaderJS       string
	LoaderOptsJSON string
	BaseURL        string
	Token          string
	Plug           string
	Agent          string
	EntryName      string
	ExecLoader     string
	JSPlugScript   string
	StyleFile      string
	ExtScripts     map[string]string
}

func (s *SubOriginData) BuildJSONOpts() error {
	opts := &LoaderOptions{
		BaseURL:      s.BaseURL,
		Token:        s.Token,
		EntryName:    s.EntryName,
		ExecLoader:   s.ExecLoader,
		JSPlugScript: s.JSPlugScript,
		StyleFile:    s.StyleFile,
		ExtScripts:   s.ExtScripts,
		Plug:         s.Plug,
		Agent:        s.Agent,
	}

	out, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	s.LoaderOptsJSON = string(out)
	return nil
}

type LoaderOptions struct {
	BaseURL      string            `json:"base_url,omitempty"`
	Token        string            `json:"token,omitempty"`
	EntryName    string            `json:"entry,omitempty"`
	ExecLoader   string            `json:"exec_loader,omitempty"`
	JSPlugScript string            `json:"js_plug_script,omitempty"`
	StyleFile    string            `json:"style,omitempty"`
	ExtScripts   map[string]string `json:"ext_scripts,omitempty"`
	Plug         string            `json:"plug,omitempty"`
	Agent        string            `json:"agent,omitempty"`
}
