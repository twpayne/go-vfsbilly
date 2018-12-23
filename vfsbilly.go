// Package vfsbilly provides a compatibility layer between
// github.com/twpayne/go-vfs and github.com/src-d/go-billy.
package vfsbilly

import (
	"os"
	"path/filepath"

	vfs "github.com/twpayne/go-vfs"
	billy "gopkg.in/src-d/go-billy.v4"
)

// A BillyFS implements github.com/src-d/go-billy.Filesystem.
type BillyFS struct {
	vfs.FS
}

// A billyFile implements github.com/src-d/go-billy.File.
type billyFile struct {
	*os.File
}

// NewBillyFS returns a new BillyFS.
func NewBillyFS(fs vfs.FS) *BillyFS {
	return &BillyFS{
		FS: fs,
	}
}

// Capabilities implements billy.Filesystem.Capabilities.
func (b *BillyFS) Capabilities() billy.Capability {
	return billy.WriteCapability | billy.ReadCapability | billy.ReadAndWriteCapability | billy.SeekCapability | billy.TruncateCapability
}

// Chroot implements billy.Chroot.Chroot.
func (b *BillyFS) Chroot(path string) (billy.Filesystem, error) {
	return nil, billy.ErrNotSupported
}

// Create implements billy.Basic.Create.
func (b *BillyFS) Create(filename string) (billy.File, error) {
	f, err := b.FS.Create(filename)
	return newBillyFile(f), err
}

// Join implements billy.Basic.Join.
func (b *BillyFS) Join(elem ...string) string {
	return filepath.Join(elem...)
}

// MkdirAll implements billy.Dir.MkdirAll.
func (b *BillyFS) MkdirAll(filename string, perm os.FileMode) error {
	return vfs.MkdirAll(b.FS, filename, perm)
}

// Open implements billy.Basic.Open.
func (b *BillyFS) Open(filename string) (billy.File, error) {
	f, err := b.FS.Open(filename)
	return newBillyFile(f), err
}

// OpenFile implements billy.Basic.OpenFile.
func (b *BillyFS) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	f, err := b.FS.OpenFile(filename, flag, perm)
	return newBillyFile(f), err
}

// Root implements billy.Chroot.Root.
func (b *BillyFS) Root() string {
	return "/"
}

// TempFile implements billy.TempFile.TempFile.
func (b *BillyFS) TempFile(dir, prefix string) (billy.File, error) {
	return nil, billy.ErrNotSupported
}

func newBillyFile(f *os.File) *billyFile {
	if f == nil {
		return nil
	}
	return &billyFile{
		File: f,
	}
}

// Close implements billy.File.Close.
func (f *billyFile) Close() error {
	if f == nil {
		return os.ErrInvalid
	}
	return f.File.Close()
}

// Lock implements billy.File.Lock.
func (f *billyFile) Lock() error {
	return billy.ErrNotSupported
}

// Unlock implements billy.File.Unlock.
func (f *billyFile) Unlock() error {
	return billy.ErrNotSupported
}
