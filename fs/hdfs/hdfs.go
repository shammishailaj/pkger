package hdfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/pkger/fs"
	"github.com/markbates/pkger/here"
	"github.com/markbates/pkger/internal/maps"
)

var _ fs.FileSystem = &FS{}

type FS struct {
	infos   *maps.Infos
	paths   *maps.Paths
	current here.Info
}

func New() (*FS, error) {
	info, err := here.Current()
	if err != nil {
		return nil, err
	}
	return &FS{
		infos: &maps.Infos{},
		paths: &maps.Paths{
			Current: info,
		},
		current: info,
	}, nil
}

func (fx *FS) Create(name string) (fs.File, error) {
	name, err := fx.locate(name)
	if err != nil {
		return nil, err
	}
	if err := fx.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return nil, err
	}
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return NewFile(fx, f)
}

func (f *FS) Current() (here.Info, error) {
	return f.current, nil
}

func (f *FS) Info(p string) (here.Info, error) {
	info, ok := f.infos.Load(p)
	if ok {
		return info, nil
	}

	info, err := here.Package(p)
	if err != nil {
		return info, err
	}
	f.infos.Store(p, info)
	return info, nil
}

func (f *FS) MkdirAll(p string, perm os.FileMode) error {
	p, err := f.locate(p)
	if err != nil {
		return err
	}
	return os.MkdirAll(p, perm)
}

func (fx *FS) Open(name string) (fs.File, error) {
	name, err := fx.locate(name)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return NewFile(fx, f)
}

func (f *FS) Parse(p string) (fs.Path, error) {
	return f.paths.Parse(p)
}

func (f *FS) ReadFile(s string) ([]byte, error) {
	s, err := f.locate(s)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(s)
}

func (f *FS) Stat(name string) (os.FileInfo, error) {
	name, err := f.locate(name)
	if err != nil {
		return nil, err
	}
	return os.Stat(name)
}

func (f *FS) Walk(p string, wf filepath.WalkFunc) error {
	fp, err := f.locate(p)
	if err != nil {
		return err
	}

	pt, err := f.Parse(p)
	if err != nil {
		return err
	}
	err = filepath.Walk(fp, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		path = strings.TrimPrefix(path, fp)
		pt, err := f.Parse(fmt.Sprintf("%s:%s", pt.Pkg, path))
		if err != nil {
			return err
		}
		return wf(pt.String(), fs.WithName(path, fs.NewFileInfo(fi)), nil)
	})

	return err
}

func (f *FS) locate(p string) (string, error) {
	return f.current.FilePath(p), nil
}

func (fx *FS) Remove(name string) error {
	name, err := fx.locate(name)
	if err != nil {
		return err
	}
	return os.Remove(name)
}

func (fx *FS) RemoveAll(name string) error {
	name, err := fx.locate(name)
	if err != nil {
		return err
	}
	return os.RemoveAll(name)
}
