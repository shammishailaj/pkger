package memfs

// func Test_Open(t *testing.T) {
// 	r := require.New(t)
//
// 	fs, err := New(here.Info{})
// 	r.NoError(err)
//
// 	_, err = fs.Open("/i.dont.exist")
// 	r.Error(err)
//
// 	f, err := fs.Create("/i.exist")
// 	r.NoError(err)
// 	_, err = io.Copy(f, strings.NewReader(radio))
// 	r.NoError(err)
// 	r.NoError(f.Close())
//
// 	f, err = fs.Open("/i.exist")
// 	r.NoError(err)
// 	b, err := ioutil.ReadAll(f)
// 	r.NoError(err)
// 	r.NoError(f.Close())
// 	r.Equal([]byte(radio), b)
// }
