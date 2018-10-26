// Code generated by go-bindata.
// sources:
// static/index.html
// static/scripts.js
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xcd\x72\xdb\x36\x10\x3e\x8b\x4f\x81\x30\x57\x11\xa0\xe4\x9f\xa4\x8c\xa8\xa6\x93\xe6\xd0\x43\xa7\x39\xb4\x33\xed\x71\x05\xac\x48\x48\x20\xc0\x00\x4b\xc9\x8a\xc7\xef\xde\x01\x49\xd9\xb2\x65\x77\x3a\x6e\x3a\x3e\x98\xf8\x76\xb1\xbb\xdf\xfe\x41\x8b\x37\x3f\xff\xf6\xe9\xf7\xbf\xbe\x7c\x66\x35\x35\x66\x99\x2c\xe2\x3f\x66\xc0\x56\x65\x8a\x36\x5d\x26\x93\x45\x8d\xa0\x96\xc9\x64\xb2\x68\x90\x80\xc9\x1a\x7c\x40\x2a\xd3\x8e\xd6\xd9\xfb\xf4\x41\x50\x13\xb5\x19\x7e\xed\xf4\xae\x4c\xff\xcc\xfe\xf8\x29\xfb\xe4\x9a\x16\x48\xaf\x0c\xa6\x4c\x3a\x4b\x68\xa9\x4c\x7f\xf9\x5c\xa2\xaa\xf0\xe4\x9e\x85\x06\xcb\x74\xa7\x71\xdf\x3a\x4f\x27\xaa\x7b\xad\xa8\x2e\x15\xee\xb4\xc4\xac\x3f\x4c\x99\xb6\x9a\x34\x98\x2c\x48\x30\x58\xce\xd2\x65\x12\xed\x90\x26\x83\xcb\xad\x74\xda\xb2\x4f\xce\x92\x77\x86\x7d\x01\x8b\x66\x21\x06\x51\xaf\x65\xb4\xdd\xb2\xda\xe3\xba\x4c\x63\xac\xa1\x10\x42\x2a\xbb\x09\x5c\x1a\xd7\xa9\xb5\x01\x8f\x5c\xba\x46\xc0\x06\x6e\x84\xd1\xab\x20\x68\xaf\x89\xd0\x67\x2b\xe7\x28\x90\x87\x56\x5c\xf0\x0b\xfe\x4e\xc8\x10\xc4\x3d\xc6\x1b\x6d\xb9\x0c\x21\x65\x1e\x4d\x99\x06\x3a\x18\x0c\x35\x22\xa5\x4c\x2c\x5f\xf6\xcb\xb7\x6e\x0f\x06\x38\xa1\xac\xc5\x70\x49\x8c\xd0\x8b\xc6\x5e\xc7\x62\xed\x2c\x65\xb0\xc7\xe0\x1a\x14\x97\xfc\x1d\xcf\x7b\x02\xa7\xf0\x6b\x38\xc4\xfb\x81\x57\xce\x55\x06\xa1\xd5\xa1\xf7\x2a\x43\xf8\x71\x0d\x8d\x36\x87\xf2\xd7\x28\x47\xef\x81\x8a\x79\x9e\x4f\x2f\xf2\x7c\x7a\x99\xe7\xd3\xab\x3c\x9f\x5e\xe7\xf9\xcb\x04\x83\xf4\xba\x25\x16\xbc\xfc\xd7\x0c\x37\x5f\x3b\xf4\x07\x71\xc1\x67\x7c\x36\x1e\x7a\x46\x9b\x90\x2e\x17\x62\x30\xb8\xfc\x4f\xb6\x33\xeb\xe8\x20\xe6\xfc\x92\xcf\x44\x0b\x72\x0b\x15\xaa\xa3\xa7\x28\xe2\x47\xf0\xbb\xf9\x7d\xa9\xf7\x36\x4f\x5b\xef\x7b\x38\x6b\x5c\x83\x96\xf8\x26\x88\x39\x9f\xbd\xe7\xf9\x11\x38\xb7\xdf\x3b\x88\x45\x5b\x26\x71\x2f\xa0\x67\xb7\xc9\x5d\x32\x99\xac\x9c\x3a\x4c\xfb\xdd\x71\x9b\x30\xc6\x58\x03\xbe\xd2\xb6\x60\xf9\x87\xfe\xd8\x82\x52\xda\x56\xf7\xe7\x15\xc8\x6d\xe5\x5d\x67\x55\xc1\xde\xce\x7f\x80\xab\xf5\x7a\x10\x48\x67\x9c\x2f\xd8\xdb\xf5\x11\xc8\xf6\xb8\xda\x6a\xca\xfa\x76\x0d\x8d\x73\x54\xf7\x86\xc0\xc6\x45\xa0\x21\xa0\x1a\x15\x1b\xf7\x2d\x73\xe1\xe6\x4c\xb3\xf2\x70\xe8\xf7\xc5\x87\xe4\x2e\x49\xea\xd9\xb4\x9e\x4f\xeb\x0b\x76\x9b\x10\xde\x50\x46\x1e\x6c\x58\x3b\xdf\x14\x5d\xdb\xa2\x97\x10\x46\xbd\xf9\x94\xf5\x5a\x93\x81\x4a\x46\xae\x2d\xd8\x1c\x9b\x5e\x6a\x60\x85\x3d\xd5\x47\xf1\xde\x25\xc9\xaa\x23\x72\x76\x0a\x7c\xf8\x18\xb3\xf1\x40\x37\x3b\x5e\xc8\xf3\xcb\xcb\x77\xd7\xec\x8d\x6e\xe2\xda\x03\x4b\x4f\xf8\x63\xfc\x3b\x13\x2b\x1d\x5a\x03\x87\x82\x69\x6b\xb4\xc5\x6c\x65\x9c\xdc\x0e\xa2\x1d\x7a\xd2\x12\x4c\x06\x46\x57\xb6\x60\x8d\x56\x2a\x72\x8e\xb2\x3e\x25\xc3\x5c\xc6\xab\x35\x7a\x4d\x4f\x0a\xc3\xdf\x5f\x61\xc3\x66\x91\xdf\x69\xda\xa1\x6d\x11\x3c\x58\x89\x05\xb3\xce\x8e\xf6\x56\xce\x2b\xf4\x05\x9b\xb5\x37\x2c\x38\xa3\x15\xeb\xd3\xd8\x82\xc7\x63\xa4\x83\x4a\xe6\x41\xe9\x2e\x14\xec\xaa\xbd\x19\xf0\x5e\x51\x93\x76\xb6\x38\x4b\x0b\xe3\xf3\xab\xc0\x10\x02\x66\xae\xa3\xe9\x33\xd8\x09\x9d\xa0\xbf\x61\xc1\x66\xfe\x18\x71\x9f\x8f\x1a\x75\x55\x53\xc1\x66\xa3\xb3\x58\xe1\x31\x1f\x12\x2d\xa1\x1f\xb3\xdc\xf9\x10\xd3\xdc\x3a\xfd\x00\x3e\x53\xa4\xbe\xaa\x4f\xda\x84\x3d\xee\x93\x93\x4a\x8f\x9d\x52\xe4\xa7\x55\xbb\x8b\xb6\xe3\xd4\x88\x71\x6c\x92\xc9\x22\x0e\xcb\xf0\x00\x82\xb6\xfd\x87\xd2\x3b\x26\x0d\x84\x50\xa6\xf7\x75\x1c\x02\x4e\x99\x56\x65\x2a\x87\x57\x6d\x78\x36\x85\xd2\xbb\x47\xbb\x92\x0e\x2d\x96\x29\xb4\xad\xd1\x12\x62\x6a\xc5\x06\x76\x30\x08\xd3\xc7\x8b\xa0\xb3\xed\xb6\xea\xa7\xdf\x23\x48\xfa\x38\xbb\xe6\x39\xcf\x45\xd7\xa8\x01\xe0\xad\x77\xaa\x93\xd1\xc8\x3f\x6e\x97\xd7\xbb\xcc\x94\x6b\xce\xdc\x46\xf0\x7f\x75\xbd\x8a\x03\x9b\x05\x02\xab\xc0\x38\x8b\x1f\xaf\xf9\xfc\x9a\xe7\x03\xfe\xdc\x8e\x3b\xf5\x15\x3b\x60\xd0\x1c\xed\x8f\xaa\xe1\x3c\x46\x31\x96\x74\x21\x86\x22\x2f\xc4\xf0\xbb\xea\xef\x00\x00\x00\xff\xff\x4d\xef\xf9\x19\x68\x09\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticScriptsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x51\x8f\xdb\x36\x12\x7e\x8e\x7f\xc5\x54\x07\x24\x32\xba\x2b\x79\x1b\x14\x28\xbc\x92\x0e\xe9\x6e\xef\xd0\xc3\xe5\x1a\x24\x2d\xf2\x10\x04\x01\x57\x1a\xdb\xc4\x4a\xa4\x40\x52\xeb\x18\xae\xff\xfb\x81\x14\x29\x91\xb6\xec\x6e\xd1\xbe\x99\xc3\x99\x8f\xc3\xe1\xf0\xe3\x27\x97\x35\x91\x12\xde\x52\xa6\xfe\x4b\xa5\xfa\x89\x29\xb1\x03\xfc\xaa\x90\x55\x12\xde\x23\x29\x55\x72\xc7\x9b\x96\x33\x64\x0a\xf6\x33\x80\x92\x33\xa9\x44\x57\x2a\x2e\xe2\x56\xf0\x56\xce\x8d\x19\x40\x76\x2d\x3a\xd3\xad\xb1\xa8\x0d\x95\x09\x67\x77\x9c\xad\xa8\x68\x20\x3f\x32\x24\x0f\x94\x55\xb1\xb6\x19\xff\xc3\x0c\x40\x20\xab\x50\xc4\x0e\x52\xa0\xea\x04\x83\xd8\x0c\x00\x32\x25\x0a\xfb\x53\x0f\xaa\x62\x6f\x00\xcd\x92\x09\xea\xcc\x13\x5a\x1d\xb2\x54\x55\x7f\xe4\xa6\xf8\x94\x5b\x7c\xe2\x47\x1a\xde\x31\x05\x29\xdc\x24\x0b\xbc\xf9\x61\x9e\x28\xfe\x2f\xfa\x15\xab\xf8\xbb\xf9\x01\xe2\x53\xdc\xde\xff\x30\x7f\x4e\x0e\x65\x5f\x05\xac\xe0\x9f\x10\xdd\xb9\x41\x04\x4b\x88\xde\x21\xab\x28\x5b\x47\x13\x49\x0e\x03\x80\xfd\x37\x17\x30\x5f\xbe\xf4\x3c\x01\x32\x02\x1b\x81\xab\x3c\xfa\x47\x04\xe6\xc4\xf3\xe8\xa1\x53\x8a\xb3\x08\x38\xbb\xab\x69\xf9\x98\xef\xc3\xd3\x39\x14\xf6\x47\x96\x12\x7f\xd5\xc3\x98\x4e\x90\xdc\x6c\x30\xda\x43\x1a\x0e\x75\x80\x8c\xd1\x1d\x2c\x26\xad\xc0\x27\x64\xea\x1e\x57\xa4\xab\x55\x3c\x1f\x1b\xa6\xdf\x8f\x44\x56\xdd\x13\x45\xe2\xbd\x05\x26\xa5\xa2\x9c\x2d\x21\xb2\x7b\xfc\xd2\x50\xa6\xa2\x2b\x3b\xbb\xe6\x4f\x28\x18\x17\x4b\x1f\xc3\x19\x9d\x13\xad\x82\x69\xd7\x31\xfd\xf4\x61\x6e\xd2\x3d\xcc\x66\xe3\x8d\xa0\x6c\xfd\x77\xdc\x85\x71\x6f\x52\x11\x85\x90\xc3\xfe\x24\xed\x68\xd8\x0a\xa9\x2a\x81\x52\x06\x26\xd3\x56\xbe\xa5\x63\x54\x8f\x6f\x7e\xb0\x96\xc3\xb8\xc6\x86\xb0\xaa\xc6\x7f\x5b\xe4\xbb\x0d\x61\x6b\x74\x77\x6f\x6a\xee\xe8\x1a\x06\x28\x6f\xfa\x5c\xa6\x40\x82\xa9\x8b\x18\x26\xf9\x49\x08\x6f\xe6\x12\xc2\x6f\x8c\x4e\xc6\x8f\xf6\x4b\xd1\xef\x04\x6f\xb9\x44\x7d\x9e\x61\xb8\x37\x71\x42\x45\xe6\x70\xed\x51\x7f\xa4\x75\xfd\x1e\x4b\xa4\x4f\x26\x44\xc6\x0c\xbf\xaa\x77\xfe\x71\xd3\x15\xc4\xe3\xf9\x0e\x7d\x07\x79\x0e\x51\x04\x2f\x5f\xc2\x10\x91\x90\xb2\xd4\x9b\x96\xd3\xd6\xa4\x46\xb6\x56\x1b\x28\x60\x31\x1f\x9a\xa4\x47\x46\xf5\x41\x83\xc7\x7b\xaf\x69\x4e\x01\x3e\x2d\x3e\xc3\xc1\xd6\xe0\xe0\x76\xf2\x07\xac\xba\xa2\x58\x57\x12\x95\x47\x34\x9b\xd7\x85\xa9\x97\xe2\x8f\xc8\x64\x96\x6e\x5e\x7b\x93\x35\x79\xc0\x1a\x56\x5c\xe4\x91\x4b\xe5\x8b\x5d\x3f\x2a\x5c\x6f\x81\xb5\x64\xa9\x71\x2f\x66\x63\xbc\xc4\x1a\x4b\x05\x6a\xd7\x62\x1e\x29\xfc\xaa\x0c\x07\x99\x73\xb4\x24\x34\xd5\xa6\x07\x78\x22\x75\xe7\x3c\xc2\x42\x1f\x02\x5a\xf4\x69\xdc\xd5\xe5\xf7\xdf\x3f\x7d\x9e\x27\x0d\x69\xe3\xd8\x9a\xe6\x90\x17\x43\x0d\x6c\x66\xbc\xd5\x1c\x03\x8f\xb8\xcb\xf7\xd6\x6d\x58\xd6\x8d\x8b\xe1\x57\x96\xf6\xfe\xfe\xe2\xf3\xb9\xcf\x8f\xfd\x4e\xfd\xbd\xaf\xb8\x68\x80\xb3\x0f\xdd\x43\x43\x55\xb0\x5b\xaf\x1b\x83\xed\xf8\xe5\xd6\x94\xf7\xc5\xf2\x43\x54\xbc\xc7\x92\xb6\x54\x73\x51\x6f\x19\x4a\xed\x05\x53\xd6\x76\x61\xa9\x69\x75\x84\x33\x5d\xfc\xe0\x7a\x4f\xd5\xde\x86\x1f\xd2\xcb\xc9\x36\x7d\x5b\xf4\x57\xfd\xcf\x66\xd8\x07\x9f\x49\xd0\x23\x8f\xc9\xfc\xfa\xa7\xf8\x72\x7a\x9a\x46\xa3\x42\xf3\xc8\x54\x6a\x13\x8d\x3a\xe4\x66\x22\xa7\x33\x1b\x69\x69\x2a\x2f\x1d\x18\x1c\xf0\xd8\x77\xbd\xb3\x26\xf5\xa2\xf9\x55\xdf\xbc\xa9\x0e\x3b\x71\x5f\x38\x6f\xd8\x22\x9d\x8a\x18\xfb\xd0\xb3\xf5\xaf\x7f\x61\xbb\x2e\x4b\xed\xd8\x6b\x5e\xdd\xaa\xc5\x6c\xf6\xe2\xc5\xc0\x06\xbc\x45\x41\x34\xba\x65\x04\x4f\x98\x90\x87\x1a\x03\x7c\xb5\x41\x52\x1d\xe5\xed\xcb\xb7\xc1\xad\xf8\xf9\x3e\x4b\xd5\x66\x6a\xe6\x8d\xeb\xeb\x33\xd3\xb6\xa7\xa6\x67\x35\x5b\x76\xe7\x63\x4b\xbb\x8f\xe3\xe9\x51\xbf\x0c\xe3\xe3\x9d\x64\xea\x81\x57\xbb\x30\x2c\x20\x9d\xc6\x8a\x69\x8f\x74\x8c\xd6\x98\xa0\x1c\x80\x2c\x94\xde\x86\x7b\x06\x2d\x0b\xe6\x97\x35\x1c\x06\xe6\x9f\x26\x41\x70\x9a\x29\xdf\x4f\x08\xa9\x03\xa4\x61\xca\x3e\x55\x99\x7d\x86\xbb\xca\xd2\xe0\x50\xb3\x34\x7c\x26\xe6\x56\xdc\x4d\x51\xf5\xa8\xf3\x8e\x9e\xae\xf1\xe5\xc2\x44\x11\xb1\x46\x95\x98\x26\x3e\x84\x68\x01\xf7\x9c\x05\x1b\x74\xd2\x65\x2c\x8f\x26\xce\x43\x59\x7d\x75\x11\x69\xbc\xd6\x67\x71\x7a\x55\x36\x89\x32\xc0\x78\x3c\x7f\x49\x0f\xdf\x3e\x5f\x10\x5f\x12\xc2\x61\x8b\x38\x27\x9f\xff\x03\x47\x6b\x0b\xfd\x6c\x71\x4e\xb8\x35\xf0\xea\xb7\x7e\xc4\x73\x93\xc2\xfa\x4d\xdb\xfe\xad\x1f\x98\xc7\x9a\xda\x3d\xf9\x4b\xf8\xf4\xd9\xa5\x58\x72\xca\x1e\x88\x44\x5f\x42\x3f\xd4\xbc\x7c\x5c\xc2\xc2\x0a\xe8\x53\xdd\x77\x4f\xab\xb7\x1a\x29\x0e\x4e\x5b\x60\xc9\x19\xc3\xb2\xff\x66\xd1\xed\x71\xac\x15\x7f\x63\xcd\x69\xd8\x56\x26\x65\xcd\x25\x06\x27\xbb\x95\x90\x43\xc7\x2a\x5c\x51\x86\xd5\xed\xa8\xd7\x86\x25\x42\x08\xc8\x81\xe1\x16\x3e\xe2\xc3\x07\x5e\x3e\xa2\x8a\xe3\x78\x4b\x59\xc5\xb7\x49\xcd\x4b\xc3\xcc\xba\x59\x14\x2f\x79\x0d\xb9\x16\x9f\x1b\xa5\x5a\xb9\x8c\xe6\xfa\x13\x73\x2b\xe5\x32\x4d\xcd\xf7\xe5\xd6\xfc\x9a\xc3\xb7\x70\x1c\xbe\xe1\x52\xc1\xb7\x10\xa5\xa4\xa5\x51\x98\x6a\xc2\x99\xd9\x01\xe4\x10\x1b\x26\x73\x05\x97\xa8\x7e\xa5\x0d\xf2\x4e\xc5\x61\x85\x3c\x55\x7d\x05\xaf\x17\x8b\x85\x13\xa6\xc7\xb8\x0d\x4a\x49\x46\x75\xcf\xd9\x47\xf9\xb6\x37\x79\x10\xae\x3a\xde\x6c\x6c\xee\x8c\x2b\xd2\x13\x11\xd0\xc8\x35\xe4\xf0\x9f\x0f\xbf\xfc\x2f\x69\x89\x90\xd6\x23\xa9\x88\x22\x76\x71\xad\xd5\x8d\x57\x9e\x03\xeb\xea\x7a\xd4\xd9\xbd\x2e\x76\xda\xd9\x77\x4e\x4c\xaf\xc0\x37\xb9\x77\x5a\x67\xf5\xb9\xed\xab\x21\xcc\x97\xe3\x23\xa2\xeb\xc8\x67\x82\x8e\x0d\xec\x07\x4f\x43\x0f\x9f\x18\xcf\x83\x1e\xef\x8b\x1f\x3c\x0d\xed\x5e\xb6\x67\x42\x3b\xf7\x1e\xda\x8d\xa6\xa1\x51\x08\x2e\xce\xe1\x32\xae\x76\xf1\xbe\x26\x3b\xde\xa9\x25\xbc\x52\xbc\xbd\x43\xa6\x50\xbc\xba\x02\xad\xcb\x7a\x7c\x83\x70\x65\xe4\xda\x12\x5e\x99\x91\x9e\xef\x5b\x73\x09\xdf\x2f\x16\x8b\x2b\x68\x05\x5f\x6b\x8a\xfb\x91\x68\x8a\x14\xdd\x99\x1a\xca\xae\x2c\x51\x9e\x2d\xe1\xb3\xf2\xb1\x18\x43\x46\x76\xfc\x27\x73\x32\x9d\x38\x90\xbf\x69\xe4\x23\x5e\xd1\x93\xb1\x69\x79\xa9\x04\x65\x6b\xba\xda\xf5\x7e\xf3\xdb\x67\x7e\x01\x56\xf4\xc9\xfd\x2b\x54\x72\xa6\x08\x65\x28\xa2\x62\xf6\x22\xd3\xd2\x07\xfd\x3f\xdd\x3c\x4f\xc1\xb7\x51\x20\x8b\x02\x94\xfa\xba\x5e\x5f\xdf\x7c\x17\x1d\xa9\xab\xcd\x4d\xf1\xa8\xbb\x17\xee\x38\x53\x82\xd7\xf0\x8e\x30\xac\xb3\x74\x73\x13\x8a\xae\x8a\x3e\xf9\x52\x24\x1c\x86\x49\x80\xec\x9a\x86\x88\xdd\xc5\x64\x64\x73\xfd\x7d\x54\x64\xd4\xd9\x56\x04\x56\xe4\x5a\x57\x49\xdf\xa3\x08\x88\xa0\xe4\x7a\x43\xab\x0a\x59\x1e\xe9\x23\x88\x8a\x2c\xa5\x05\xf8\x42\x6b\xb8\x75\xc3\x73\x72\x94\xd8\xe4\xb2\xaf\xff\xea\xb2\x3d\x89\xf4\xcf\x95\x9c\xae\xcc\x8b\x2c\x75\x27\x35\x4e\xb9\x3f\xb0\xdc\x8d\x0e\xbf\x8e\xdc\x35\x07\x77\x2b\x83\xe9\xe1\xaa\x1e\x6b\x4a\x37\xf4\x78\xd9\x13\x96\x5e\x76\xc3\xa3\x6f\x9e\xf8\xfb\x5f\xde\x26\xb6\x07\x67\x00\x99\x96\x00\x69\xa1\x1f\xde\x8a\x97\x5d\xa3\x19\x7a\x8d\xea\xa7\x1a\xf5\xcf\x1f\x77\x3f\x57\xf1\xab\xb2\xef\x8f\x57\xf3\xd9\xfc\x76\x36\xfb\x7f\x00\x00\x00\xff\xff\xab\x11\x04\x75\xa5\x16\x00\x00")

func staticScriptsJsBytes() ([]byte, error) {
	return bindataRead(
		_staticScriptsJs,
		"static/scripts.js",
	)
}

func staticScriptsJs() (*asset, error) {
	bytes, err := staticScriptsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/scripts.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/index.html": staticIndexHtml,
	"static/scripts.js": staticScriptsJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"static": {nil, map[string]*bintree{
		"index.html": {staticIndexHtml, map[string]*bintree{}},
		"scripts.js": {staticScriptsJs, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
