package vfsbilly

import (
	"testing"

	"github.com/twpayne/go-vfs/vfst"
	billy "gopkg.in/src-d/go-billy.v4"
	billyutil "gopkg.in/src-d/go-billy.v4/util"
)

var (
	_ billy.Filesystem = &BillyFS{}
)

func TestBillyFS(t *testing.T) {

	fs, cleanup, err := vfst.NewTestFS(map[string]interface{}{
		"/home/user/.bashrc": "# contents of .bashrc\n",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	billyFS := NewBillyFS(fs)
	if err := billyutil.WriteFile(billyFS, "/home/user/foo", []byte("bar"), 0666); err != nil {
		t.Errorf("billyutil.WriteFile(...) == %v, want <nil>", err)
	}

	vfst.RunTests(t, fs, "",
		vfst.TestPath("/home/user/foo",
			vfst.TestContentsString("bar"),
		),
	)

}
