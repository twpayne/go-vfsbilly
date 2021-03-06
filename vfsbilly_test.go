package vfsbilly_test

import (
	"testing"

	billy "github.com/go-git/go-billy/v5"
	billyutil "github.com/go-git/go-billy/v5/util"
	"github.com/twpayne/go-vfs/vfst"

	vfsbilly "github.com/twpayne/go-vfsbilly"
)

var _ billy.Filesystem = &vfsbilly.BillyFS{}

func TestBillyFS(t *testing.T) {
	fs, cleanup, err := vfst.NewTestFS(map[string]interface{}{
		"/home/user/.bashrc": "# contents of .bashrc\n",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	billyFS := vfsbilly.NewBillyFS(fs)
	if err := billyutil.WriteFile(billyFS, "/home/user/foo", []byte("bar"), 0o666); err != nil {
		t.Errorf("billyutil.WriteFile(...) == %v, want <nil>", err)
	}

	vfst.RunTests(t, fs, "",
		vfst.TestPath("/home/user/foo",
			vfst.TestContentsString("bar"),
		),
	)
}
