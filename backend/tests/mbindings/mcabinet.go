package mbindings

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
)

func (m *MockedBindings) AddFile(bucket string, file string, contents []byte) error {

	if bucket == "err_bucket" {
		return easyerr.Error("cabinet addfile err")
	}

	return nil
}

func (m *MockedBindings) ListFolder(bucket string) ([]string, error) {
	if bucket == "err_bucket" {
		return nil, easyerr.Error("cabinet listfolder err")
	}

	return nil, nil
}

func (m *MockedBindings) GetFile(bucket string, file string) ([]byte, error) {
	if bucket == "err_bucket" {
		return nil, easyerr.Error("cabinet getfile err")
	}

	return nil, nil
}

func (m *MockedBindings) DeleteFile(bucket string, file string) error {
	if bucket == "err_bucket" {
		return easyerr.Error("cabinet deletefile err")
	}

	return nil
}

func (m *MockedBindings) GenerateTicket(bucket string, ticket *rtypes.CabTicket) (string, error) {

	if bucket == "err_bucket" {
		return "", easyerr.Error("cabinet generate_ticket err")
	}

	return "", nil
}
