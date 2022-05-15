package binder

import (
	"io"
	"io/fs"
	"time"
)

type SelfFs struct {
	b *Binder
}

func (s *SelfFs) Open(name string) (fs.File, error) {
	return &SelfFsFile{
		name: name,
		b:    s.b,
	}, nil
}

type SelfFsFile struct {
	name      string
	b         *Binder
	dataCache []byte
	offset    int64
}

func (s *SelfFsFile) Stat() (fs.FileInfo, error) { return s, nil }
func (s *SelfFsFile) Read(b []byte) (int, error) {
	if s.dataCache == nil {
		err := s.fillData()
		if err != nil {
			return 0, err
		}
	}

	if s.offset >= int64(len(s.dataCache)) {
		return 0, io.EOF
	}

	if s.offset < 0 {
		return 0, &fs.PathError{Op: "read", Path: s.name, Err: fs.ErrInvalid}
	}

	n := copy(b, s.dataCache[s.offset:])
	s.offset += int64(n)
	return n, nil
}

func (s *SelfFsFile) Close() error {
	s.b = nil
	s.dataCache = nil
	return nil
}

func (s *SelfFsFile) fillData() error {
	data, err := s.b.GetSelfFile(s.name)
	if err != nil {
		return err
	}
	s.dataCache = data
	return nil
}

// fs.fsinfo
func (s *SelfFsFile) Name() string { return s.name }
func (s *SelfFsFile) Size() int64 {
	if s.dataCache == nil {
		err := s.fillData()
		if err != nil {
			return 0
		}
	}

	return int64(len(s.dataCache))
}
func (s *SelfFsFile) Mode() fs.FileMode  { return fs.FileMode(0666) }
func (s *SelfFsFile) ModTime() time.Time { return time.Time{} }
func (s *SelfFsFile) IsDir() bool        { return false }
func (s *SelfFsFile) Sys() interface{}   { return nil }
