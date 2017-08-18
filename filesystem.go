package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type fileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
	Create(name string) (file, error)
	ReadFile(string) ([]byte, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer

	Stat() (os.FileInfo, error)
	WriteString(string) (int, error)
}

type osFS struct{}

func (osFS) Open(name string) (file, error)        { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func (osFS) Create(name string) (file, error)      { return os.Create(name) }
func (osFS) ReadFile(name string) ([]byte, error)  { return ioutil.ReadFile(name) }

type cmdRunner interface {
	Run(string, ...string) ([]byte, error)
}

type realRunner struct{}

func (realRunner) Run(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).Output()
}
