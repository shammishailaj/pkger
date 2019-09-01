package memfs

import (
	"testing"

	"github.com/markbates/pkger/fs/fstest"
	"github.com/markbates/pkger/here"
	"github.com/stretchr/testify/require"
)

func Test_FS(t *testing.T) {
	r := require.New(t)

	info, err := here.Current()
	r.NoError(err)

	myfs, err := New(info)
	r.NoError(err)

	suite, err := fstest.NewFileSystem(myfs)
	r.NoError(err)

	suite.Test(t)
}
