// Code generated by go-bindata.
// sources:
// faucet.html
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\xff\x73\xdb\xb6\x92\xff\xd9\xf9\x2b\xb6\xbc\xe4\x49\x3a\x9b\xa4\x6c\x27\x79\x3e\x89\x54\x27\x2f\xed\xeb\xe5\xbe\xb4\x9d\x26\x9d\xbb\x37\x6d\xe7\x06\x24\x56\x22\x62\x10\x60\x01\x50\xb2\xea\xd1\xff\x7e\x03\x80\xa4\x48\x49\x76\x93\x26\x77\xd3\xfc\xe0\x90\xc0\x62\x77\xb1\xfb\x59\xec\x62\xc5\xe4\x8b\xaf\xbe\x7b\xfd\xee\x1f\xdf\x7f\x0d\x85\x29\xf9\xe2\x49\x62\xff\x03\x4e\xc4\x2a\x0d\x50\x04\x8b\x27\x67\x49\x81\x84\x2e\x9e\x9c\x9d\x25\x25\x1a\x02\x79\x41\x94\x46\x93\x06\xb5\x59\x86\x37\xc1\x7e\xa2\x30\xa6\x0a\xf1\xd7\x9a\xad\xd3\xe0\xbf\xc3\x1f\x5f\x85\xaf\x65\x59\x11\xc3\x32\x8e\x01\xe4\x52\x18\x14\x26\x0d\xde\x7c\x9d\x22\x5d\x61\x6f\x9d\x20\x25\xa6\xc1\x9a\xe1\xa6\x92\xca\xf4\x48\x37\x8c\x9a\x22\xa5\xb8\x66\x39\x86\xee\xe5\x02\x98\x60\x86\x11\x1e\xea\x9c\x70\x4c\x2f\x83\xc5\x13\xcb\xc7\x30\xc3\x71\x71\x7f\x1f\x7d\x8b\x66\x23\xd5\xed\x6e\x37\x83\x57\xb5\x29\x50\x18\x96\x13\x83\x14\xfe\x4e\xea\x1c\xcd\x55\x12\x7b\x52\xb7\x8a\x33\x71\x0b\x85\xc2\x65\x1a\x58\xdd\xf5\x2c\x8e\x73\x2a\xde\xeb\x28\xe7\xb2\xa6\x4b\x4e\x14\x46\xb9\x2c\x63\xf2\x9e\xdc\xc5\x9c\x65\x3a\x36\x1b\x66\x0c\xaa\x30\x93\xd2\x68\xa3\x48\x15\x5f\x47\xd7\xd1\x5f\xe3\x5c\xeb\xb8\x1b\x8b\x4a\x26\xa2\x5c\xeb\x00\x14\xf2\x34\xd0\x66\xcb\x51\x17\x88\x26\x80\x78\xf1\xc7\xe4\x2e\xa5\x30\x21\xd9\xa0\x96\x25\xc6\xcf\xa3\xbf\x46\x53\x27\xb2\x3f\xfc\xb8\x54\x2b\x56\xe7\x8a\x55\x06\xb4\xca\x3f\x58\xee\xfb\x5f\x6b\x54\xdb\xf8\x3a\xba\x8c\x2e\x9b\x17\x27\xe7\xbd\x0e\x16\x49\xec\x19\x2e\x3e\x89\x77\x28\xa4\xd9\xc6\x57\xd1\xf3\xe8\x32\xae\x48\x7e\x4b\x56\x48\x5b\x49\x76\x2a\x6a\x07\x3f\x9b\xdc\x87\x7c\xf8\xfe\xd0\x85\x9f\x43\x58\x29\x4b\x14\x26\x7a\xaf\xe3\xab\xe8\xf2\x26\x9a\xb6\x03\xc7\xfc\x9d\x00\xeb\x34\x2b\xea\x2c\x5a\xa3\xb2\xd0\xe5\x61\x8e\xc2\xa0\x82\x7b\x3b\x7a\x56\x32\x11\x16\xc8\x56\x85\x99\xc1\xe5\x74\xfa\x6c\x7e\x6a\x74\x5d\xf8\x61\xca\x74\xc5\xc9\x76\x06\x4b\x8e\x77\x7e\x88\x70\xb6\x12\x21\x33\x58\xea\x19\x78\xce\x6e\x62\xe7\x64\x56\x4a\xae\x14\x6a\xdd\x08\xab\xa4\x66\x86\x49\x31\xb3\x88\x22\x86\xad\xf1\x14\xad\xae\x88\x38\x5a\x40\x32\x2d\x79\x6d\xf0\x40\x91\x8c\xcb\xfc\xd6\x8f\xb9\x70\xee\x6f\x22\x97\x5c\xaa\x19\x6c\x0a\xd6\x2c\x03\x27\x08\x2a\x85\x0d\x7b\xa8\x08\xa5\x4c\xac\x66\xf0\xb2\x6a\xf6\x03\x25\x51\x2b\x26\x66\x30\xdd\x2f\x49\xe2\xd6\x8c\x49\xec\x4f\xae\x27\x67\x49\x26\xe9\xd6\xf9\x90\xb2\x35\xe4\x9c\x68\x9d\x06\x07\x26\x76\x27\xd2\x80\xc0\x1e\x44\x84\x89\x76\x6a\x30\xa7\xe4\x26\x00\x27\x28\x0d\xbc\x12\x61\x26\x8d\x91\xe5\x0c\x2e\xad\x7a\xcd\x92\x03\x7e\x3c\xe4\xab\xf0\xf2\xaa\x9d\x3c\x4b\x8a\xcb\x96\x89\xc1\x3b\x13\x3a\xff\x74\x9e\x09\x16\x09\x6b\xd7\x2e\x09\x2c\x49\x98\x11\x53\x04\x40\x14\x23\x61\xc1\x28\x45\x91\x06\x46\xd5\x68\x71\xc4\x16\xd0\x3f\xff\x4e\x1e\x7f\x49\x5c\x5c\xb6\x7a\xc5\x94\xad\x9b\x6d\xf5\x1e\x0f\x76\xf8\xf0\x26\x6e\xa0\x79\x90\xcb\xa5\x46\x13\xf6\xf6\xd4\x23\x66\xa2\xaa\x4d\xb8\x52\xb2\xae\xba\xf9\xb3\xc4\x8d\x02\xa3\x69\x50\x2b\x1e\x34\xe7\xbf\x7b\x34\xdb\xaa\x31\x45\xd0\x6d\x5c\xaa\x32\xb4\x9e\x50\x92\x07\x50\x71\x92\x63\x21\x39\x45\x95\x06\x6f\x65\xce\x08\x07\xe1\xf7\x0c\x3f\xfe\xf0\x1f\xd0\xb8\x8c\x89\x15\x6c\x65\xad\xe0\xdf\xe5\x86\x70\x02\x84\x52\x0b\xd6\x28\x8a\x7a\x6a\x38\xe4\x1e\x2b\x1a\x66\x46\xec\xa9\xce\x92\xac\x36\x46\x76\x84\x99\x11\x90\x19\x11\x52\x5c\x92\x9a\x1b\xa0\x4a\x56\x54\x6e\x44\x68\xe4\x6a\x65\x13\x9d\xdf\x82\x5f\x14\x00\x25\x86\x34\x53\x69\xd0\xd2\xb6\x1e\x24\xba\x92\x55\x5d\x35\x3e\xf4\x83\x78\x57\x11\x41\x91\x5a\x8f\x73\x8d\xc1\xe2\x1b\xb6\x46\x28\x11\x6e\x7f\x7c\xfb\x15\x9c\x1d\x02\x22\x27\x0a\x4d\xd8\x67\x7a\x04\x8b\x24\xf6\xca\xf8\x2d\x41\xf3\x2f\xa9\x79\xcb\xa9\xdb\x42\x89\xa2\x86\xc1\x5b\xa8\xec\xa9\x12\x2c\xee\xef\x15\x11\x2b\x84\xa7\x8c\xde\x5d\xc0\x53\x52\xca\x5a\x18\x98\xa5\x10\xbd\x72\x8f\x7a\xb7\x1b\x70\x07\x48\x38\x5b\x24\xe4\x31\x70\x83\x14\x39\x67\xf9\x6d\x1a\x18\x86\x2a\xbd\xbf\xb7\xcc\x77\xbb\x39\xdc\xdf\xb3\x25\x3c\x8d\x7e\xc0\x9c\x54\x26\x2f\xc8\x6e\xb7\x52\xed\x73\x84\x77\x98\xd7\x06\xc7\x93\xfb\x7b\xe4\x1a\x77\x3b\x5d\x67\x25\x33\xe3\x76\xb9\x1d\x17\x74\xb7\xb3\x3a\x37\x7a\xee\x76\x10\x5b\xa6\x82\xe2\x1d\x3c\x8d\xbe\x47\xc5\x24\xd5\xe0\xe9\x93\x98\x2c\x92\x98\xb3\x45\xb3\x6e\x68\xa4\xb8\xe6\x7b\xbc\xc4\x16\x30\x1d\xca\x5d\xd0\x38\x55\xfb\x9a\x9e\x88\x81\x55\xd8\x69\xdf\xe0\x41\x33\x83\xb7\xb8\x4d\x83\xfb\xfb\xfe\xda\x66\x36\x27\x9c\x67\xc4\xda\xc5\x6f\xad\x5b\xf4\x1b\x5a\x9c\xae\x99\x76\x15\xd5\xa2\xd5\x60\xaf\xf6\x07\x06\xf5\xc1\xb1\x65\x64\x35\x83\xeb\xab\xde\x99\x75\x2a\xde\x5f\x1e\xc4\xfb\xf5\x49\xe2\x8a\x08\xe4\xe0\xfe\x86\xba\x24\xbc\x7d\x6e\xa2\xa5\x17\x7c\x87\x8b\x42\x7b\x42\x77\xaa\x75\x27\xfd\x74\x0e\x72\x8d\x6a\xc9\xe5\x66\x06\xa4\x36\x72\x0e\x25\xb9\xeb\xb2\xdd\xf5\x74\xda\xd7\xdb\x56\x82\x24\xe3\xe8\xce\x16\x85\xbf\xd6\xa8\x8d\xee\x4e\x12\x3f\xe5\xfe\xda\x03\x85\xa2\xd0\x48\x0f\xac\x61\x25\x5a\xd3\x3a\xaa\x9e\xeb\x3b\x63\x9e\xd4\x7d\x29\x65\x97\x40\xfa\x6a\x34\xac\x7b\xb9\x2e\x58\x24\x46\xed\xe9\xce\x12\x43\x3f\x2a\x01\x28\x5b\xe0\x3d\x74\xfe\xfb\x13\xcd\xee\xbd\x42\x54\xbe\xba\xb0\x90\x05\xf7\x9a\xc4\x86\x7e\x82\x64\x0b\xc2\x8c\x68\xfc\x10\xf1\x2e\xcf\xef\xc5\xbb\xd7\x4f\x95\x5f\x20\x51\x26\x43\x62\x3e\x44\x81\x65\x2d\x68\x6f\xff\xf6\xec\xfc\x54\xf9\xb5\x60\x6b\x54\x9a\x99\xed\x87\x2a\x80\x74\xaf\x81\x7f\x1f\xaa\x90\xc4\x46\x3d\x0e\xb5\xfe\xcb\x67\x8a\xed\xdf\xab\x47\xae\x17\xff\x2a\x37\x40\x25\x6a\x30\x05\xd3\x60\x33\xeb\x97\x49\x5c\x5c\x77\x24\xd5\xe2\x9d\x9d\x70\xf9\x68\xe9\xca\x0a\x60\x1a\x54\x2d\x5c\xd6\x95\x02\x4c\x81\xc3\x52\xa4\x49\xd0\x11\xbc\x93\xb6\x9c\x5b\xa3\x30\x50\x12\xce\x72\x26\x6b\x0d\x24\x37\x52\x69\x58\x2a\x59\x02\xde\x15\xa4\xd6\xc6\x32\xb2\x87\x07\x59\x13\xc6\x5d\x24\x39\x87\x82\x54\x40\xf2\xbc\x2e\x6b\x5b\x8e\x8a\x15\xa0\x90\xf5\xaa\xf0\xaa\x18\x09\x3e\x2b\x71\x29\x56\x9d\x3a\xba\x22\x25\x10\x63\x48\x7e\xab\x2f\xa0\x3d\x12\x80\x28\x04\xc3\x90\xda\x55\xb9\x2c\x4b\x29\xe0\x5a\x51\xa8\x88\x32\x5b\xd0\xc3\xb2\x82\xe4\xb9\x4b\x71\x11\xbc\x12\x5b\x29\x10\x0a\xb2\x76\x0a\xc2\x3b\x7f\x93\xb8\x80\x6f\xa4\x5c\x71\x3c\xb7\xfa\xfd\x9d\xe4\x98\x49\xd9\x2d\x83\x92\x6c\x5b\xb9\xcd\x2e\x36\xcc\x14\xcc\x9b\xa9\x42\x55\x5a\x1e\x14\x38\x2b\x99\xd1\x51\x12\x57\xfb\x73\x75\x9f\xa1\x79\x58\x48\xc5\x7e\xb3\xc5\x0d\xef\x1f\xa2\xe6\xe0\x88\x69\x4f\x48\xe7\x7c\x8e\x4b\x33\x83\xe7\xfe\x84\x3c\x84\x73\x73\x0b\x3a\x85\xe5\x96\xa7\xbb\x5d\xda\xb4\x33\x83\x6b\x5f\xd2\xfa\x72\x82\x9a\x9e\x06\xf4\x00\x71\x5e\xe8\xcd\x4d\x75\xd7\xe9\xd1\xd5\xc5\xd3\x8e\x89\x45\xc2\xd0\x28\x6b\xd6\xb3\x67\x49\x6e\x11\x08\x24\xe4\xe0\x96\xdc\x28\xed\xee\x58\xcc\x35\x09\x62\xb3\x41\x34\x5f\xda\x08\x4e\x7f\xf0\x0c\x99\x58\x3d\xbb\x9a\x7a\x64\xda\x07\xcb\xfe\xd9\xd5\x94\x09\x23\x9f\x5d\x4d\xa7\x77\xd3\x0f\xfc\xf7\xec\x6a\x2a\xc5\xb3\xab\xa9\x29\xf0\xd9\xd5\xf4\xd9\xd5\x75\x1f\xd3\x7e\xc4\x57\x97\x96\x06\xb5\x95\xd5\x02\x3d\x00\x43\xd4\x0a\x4d\x1a\xfc\x0f\xc9\x64\x6d\x66\x19\x27\xe2\x36\x58\x38\x65\x6d\xc5\xe1\x30\x70\xaa\x42\x85\x8a\x68\x0b\x07\xab\xad\x43\x48\xd3\x0c\xd1\x30\xd6\xb5\x52\xb2\x16\x36\x2f\x82\xdd\xaf\x0b\x52\x31\xb2\x08\xb3\x46\x99\x44\x49\xa6\xe2\xc5\x6b\x59\x6d\x43\xc7\xc4\x2d\x3f\x32\xa1\xae\xab\x4a\x2a\x13\xf5\x4d\x49\xec\x3d\x88\xa3\x8e\x6f\xa6\x2f\x6e\x5e\x3e\xaa\xbc\xb6\x55\xb6\xdb\x41\xa7\x21\xc9\xe4\x1a\xc1\xd7\xf4\x99\xbc\x03\x22\x28\x2c\x99\x42\x20\x1b\xb2\xfd\x22\x89\xa9\xbb\x81\x7d\x3a\x62\x57\x2e\xc8\xc2\x8a\xd7\xda\x16\x21\xcc\x06\xe9\x9f\x0a\xbe\xfe\x14\x80\xef\x79\xad\x2f\xa0\xaa\x33\xce\x74\x01\x04\x04\x6e\x20\xd1\x46\x49\xb1\x5a\xb8\xd1\xdc\xde\x50\xdd\x2b\x54\x52\x9b\x87\xb1\x80\x65\x86\x94\x9e\x40\xc3\x1f\x04\x83\x95\xe6\x1c\xf8\xff\xef\xbc\x65\x73\x2c\xfe\xa9\x1c\xd6\x9e\xd5\x7f\x4e\x6f\x1d\x85\xee\x66\xb3\x89\x5a\x3b\xba\xb8\x2d\x90\x57\xb1\x4d\x5f\xb5\x60\x66\x1b\xfb\xd3\x4f\x8a\xf8\x4b\x46\xd3\xab\x9b\xab\x97\x2f\xaf\x9e\xff\xcb\xcd\x8b\x17\x57\x37\xcf\x5f\x3c\x14\xd4\x1d\x24\xfe\x78\x4c\xfb\x3b\xd0\xb7\xf2\x55\x6d\x8a\xee\x02\xe4\xd1\xd2\x16\xde\xb6\xbe\xa2\xf6\x02\xa9\x82\x3f\x8c\xa0\x5a\xd8\x2a\x32\x24\xfc\x64\x01\xf8\x11\x18\x72\x20\x7a\x44\xb3\x4f\x04\x56\x0b\x1e\x8b\x13\x59\x1b\xbb\xc3\xb6\x0f\xc3\xa4\xe8\xc0\x74\x01\x9a\x95\x15\xdf\x42\xbe\xf7\xfa\x29\x54\x3d\xe8\x92\xdf\x05\xd5\xd0\x69\x1e\x62\xae\x72\x2b\x25\x45\x5b\xb2\xe9\x5a\xe7\x58\xb9\xfe\xbc\xad\x83\xfe\xb6\xfd\x8d\x08\xc3\x04\xb6\xf5\x52\x04\xdf\x09\xbe\x85\x5a\x23\x2c\xa5\x02\x8a\x59\xbd\x5a\xb9\x1a\x4f\x41\xa5\xd8\x9a\x18\x6c\x8b\x24\xdd\x60\xa2\x83\x44\xef\x52\x6a\xcb\x55\xde\xab\x1e\xff\x21\x6b\xc8\x89\x00\xa3\x48\x7e\xeb\xe3\xa4\x56\xca\xc6\x49\x85\x7e\x37\x5d\x99\x96\x21\x97\x1b\x47\xe2\xf7\xbd\x64\xc8\x5d\xcd\xa6\x11\xa1\x90\x1b\x28\xeb\xdc\x05\xa3\xad\xc9\xdc\x26\x36\x84\x19\xa8\x85\x61\xdc\x5b\xd3\xd4\x4a\xd8\x0a\x0f\x07\xa5\xd5\xd1\xb5\x3d\xc1\x72\xf1\xae\xc0\x13\xf5\x6c\x77\xe1\x06\x85\xaf\x3d\x39\x54\x4a\x1a\xcc\xad\x3b\x81\xac\x08\x13\xda\x7a\xc4\x15\x6f\x58\x7e\xc0\x85\xbc\x7b\x6a\x1e\xf6\xad\x65\x37\x1d\xc7\xf0\x0d\x97\x19\xe1\xb0\xb6\x38\xcf\xb8\x2d\xc5\x25\x14\xd2\x6e\xbd\x67\x2d\x6d\x88\xa9\x35\xc8\xa5\x1b\xf5\x9a\xdb\xf5\x6b\xa2\xac\x07\xb1\xac\x0c\xa4\x4d\x63\xd4\x8e\x69\x54\xeb\xa6\xdd\x6b\x5f\x0d\x43\x35\x98\xef\xac\x9e\xc2\x4f\xbf\xcc\x9f\x34\xaa\x7c\x85\x4b\x07\x09\x8b\x6e\xbf\x65\x53\x10\x03\xb9\x42\x62\x50\x43\xce\xa5\xae\x95\xd7\x90\x2a\x59\x81\xd5\xb2\xe5\xd4\x72\xb6\x13\x95\x93\xd6\x32\x19\x17\x44\x17\x93\xa6\xaf\xab\xd0\x79\xa9\x9b\x6b\xc7\xcf\x2c\xea\xc6\x96\x01\x4b\xa7\x73\x60\x49\xcb\x37\xe2\x28\x56\xa6\x98\x03\x3b\x3f\xef\x88\xcf\xd8\x12\xc6\x2d\xc5\x4f\xec\x97\xc8\xdc\x45\x56\x0a\xa4\x29\xf4\xa5\x39\x81\x0d\x1f\x5d\x71\x96\xe3\x98\x5d\xc0\xe5\x64\xde\xce\x66\x0a\xc9\x6d\xfb\xd6\xf8\xd1\xff\xe7\xfe\xee\xe6\x43\xcb\x38\xe3\x0f\x6c\xe3\xdb\x36\x1a\x08\xac\x98\x36\x50\x2b\x0e\x4d\x0c\x7b\x17\x74\x0e\x71\x74\x7d\xab\x1c\xe1\xb2\x79\x68\x30\xd5\x6e\xc1\xb3\x89\x34\x0a\x3a\xfe\xb7\xb7\xdf\x7d\x1b\x69\xa3\x98\x58\xb1\xe5\x76\x7c\x5f\x2b\x3e\x83\xa7\xe3\xe0\x9f\x6a\xc5\x83\xc9\x4f\xd3\x5f\xa2\x35\xe1\x35\x5e\x38\x7f\xcf\xdc\xdf\x23\x29\x17\xd0\x3c\xce\x60\x28\x70\x37\x99\xcc\x4f\xb7\xb8\x7a\x1d\x39\x85\x1a\xcd\xd8\x12\x76\xc0\x3f\xb4\x11\x81\x12\x4d\x21\x5d\xe8\x2a\xcc\xa5\x10\x98\x1b\xa8\x2b\x29\x1a\x93\x00\x97\x5a\xef\x81\xd8\x52\xa4\xc7\xa0\x68\xe8\x53\x97\xa8\xff\x0b\xb3\xb7\x32\xbf\x45\x33\x1e\x8f\x37\x4c\x50\xb9\x89\xb8\xf4\x07\x6d\x64\x83\x54\xe6\x92\x43\x9a\xa6\xd0\xe4\xd0\x60\x02\x5f\x42\xb0\xd1\x36\x9b\x06\x30\xb3\x8f\xf6\x69\x02\xe7\x70\xb8\xbc\xb0\xb9\xfe\x1c\x82\x98\x54\x2c\x98\xf8\x70\x68\x0d\x2f\x45\x89\x5a\x93\x15\xf6\x15\x74\xd7\xda\x0e\x64\x76\x1f\xa5\x5e\x41\x0a\xce\x41\x15\x51\x1a\x3d\x49\x44\x89\x21\x2d\xda\x2c\x66\x1d\x59\x9a\x82\xa8\x39\xdf\x83\xd4\x07\xc5\xbc\x85\xdf\x80\x3c\xf2\x99\xe6\x8b\x34\x85\x5a\x50\x67\x62\xba\x5f\x69\x9d\xef\xfb\x1f\x93\xc8\xe6\x85\xfd\x8a\xc9\xbc\x8f\xe6\x01\x37\xa4\xbf\xc7\x0e\xe9\x21\x3f\xa4\x0f\x30\x74\xed\xa6\xc7\xf8\xf9\xf6\x54\x8f\x9d\x1b\x78\x80\x9b\xa8\xcb\x0c\xd5\x63\xec\x7c\xbb\xa9\x61\xe7\x4c\xfd\x46\x98\xde\xda\x0b\xb8\x7c\x39\x79\x80\x3b\x2a\x25\x1f\x64\x2e\xa4\xd9\x8e\xef\x39\xd9\xda\x8a\x09\x46\x46\x56\xaf\x5d\x7b\x68\x74\xe1\x32\xee\x0c\x3a\x0e\x17\xae\xef\x3f\x83\x91\x7b\xb3\xf3\xac\x44\xb7\xea\xc5\x74\x3a\xbd\x80\xf6\xe7\xb2\xbf\x11\x1b\x84\xaa\xc6\xdd\x03\xfa\xe8\x3a\xcf\x6d\xde\xff\x14\x8d\x1a\x1e\x9d\x4e\xcd\xfb\x27\x68\xd5\xe5\x86\x81\x5a\xf0\x97\xbf\xc0\xd1\xec\x10\xc6\x71\x0c\xff\x49\xd4\xad\xeb\xe6\x54\x0a\xd7\xae\xe3\xd3\xd1\x97\x4c\x6b\xd7\x4a\xd1\x40\xa5\xc0\x66\xcd\xc7\x1d\xfb\x47\x3a\x36\x64\xb0\x80\xe9\xa1\x82\xf6\x38\xec\xa5\x85\x13\xd9\xa2\xc7\x77\x98\x08\xce\x76\x7d\x79\x83\x95\xac\x44\xf8\x22\x85\x20\xe8\x2f\x3e\xa2\xb0\x04\x1d\xb3\x33\x8d\xe6\x9d\xf7\xc5\xb8\xc9\x8e\xa7\x72\xd7\xe4\x02\xae\xa7\xd3\xe9\xe4\x48\x89\xdd\xde\xbc\xaf\x2a\x5b\x36\x01\x11\x5b\x77\x24\x76\xb6\x75\x85\xa3\x2d\x81\xec\x91\xc6\x21\x97\x9c\xfb\x9a\xa5\x59\x6a\x0d\xdc\xb4\xbe\x52\x08\x2f\xe7\x27\xb2\x68\xcf\x92\xbd\xad\x1d\xba\xe7\x84\xed\x0f\x5d\x34\xb4\xd9\x01\x71\x78\x39\x70\xca\xc0\x5f\xa7\x1d\x73\xd6\xe9\xcd\xf6\x16\x3d\x70\xd7\xde\x5f\x87\x36\xeb\xe9\xef\xf9\x9c\x5f\x7e\xe0\x36\xba\xe9\xaa\xd6\xc5\xf8\x40\xd1\xc9\xfc\xd8\x37\x6f\x0c\x2a\x5b\x25\x4b\x9b\xb2\xac\x2f\xec\x45\x40\xe1\x91\x4b\x5c\xa9\xae\x30\x54\x28\x28\xaa\xb6\xa4\xf0\x95\xbd\x2d\x00\x07\x2e\xf3\x77\xca\x3e\x9c\x3e\x32\x60\x5c\x49\x26\x05\x02\x00\x1c\x04\x81\x03\xea\x00\xa9\x96\x18\x39\xa9\x34\x52\x48\xc1\x7f\xbd\x30\x9e\x44\xb5\x60\x77\xe3\x49\xd8\xbc\x1f\xf2\x68\xe7\xe7\xdd\x25\xb1\x55\xfb\x3c\x85\x20\x31\x0a\x18\x4d\x47\x01\x9c\x9f\x0a\x41\x9b\x75\x47\x8b\xbd\x06\xfd\xa5\x00\x89\xa1\x0b\xd7\xc3\xf6\xb7\xb5\x9f\x83\x8c\xe4\xb7\x2b\x77\x11\x9a\xd9\x52\x6b\x7c\xc4\x96\xac\x89\x21\xca\x71\x9d\xcc\x61\x4f\xde\x5c\x13\x73\xeb\x9c\x39\xf8\xfb\xa8\x6b\x95\x43\xf7\xeb\x92\x7b\xcb\xa4\xa2\xa8\x42\x45\x28\xab\xf5\x0c\x9e\x57\x77\xf3\x9f\xdb\x5f\xdf\x5c\x43\xff\x51\x55\x2b\x85\x8b\x23\x8d\x9a\xd6\xf0\x39\x04\x49\x6c\x09\x7e\x8f\x4d\xb7\xd9\xfe\x57\x13\x70\xe2\x67\x0b\xe8\xbe\x69\x68\xc6\x4b\x46\x29\x47\xab\xf0\x9e\xbd\x0d\x46\xeb\xff\x7e\x48\x0d\x45\x42\xf3\x7b\xc5\x7e\xcd\x0e\x90\x6b\x7c\x64\x41\xf7\xd3\xc7\xc8\x02\x20\xb4\x5b\x66\xce\xe6\xcd\x55\xdb\x0d\xab\x91\xb3\x45\xf3\x0d\x0c\xad\x95\xab\xb5\xc6\x61\x03\xb0\x0b\x18\x69\x5b\xfb\x51\x3d\x9a\x44\x45\x5d\x12\xc1\x7e\xc3\xb1\xcd\x4b\x13\x6f\x2b\xf7\x5b\x4a\x70\x7c\x24\x1f\x29\xb3\xff\x91\x63\xd4\xe6\xb8\x51\x63\xc4\x51\xeb\xdd\xe7\xfb\x9b\xfd\x0c\xa6\xf3\xd1\x47\x5a\xe8\xb4\x94\x30\x23\x0a\xfa\x2f\x61\x9b\x7c\x41\x49\x2b\xbd\x9d\xcb\x88\x1a\xf9\x3e\x86\xab\xcf\x85\xdc\xa4\xa3\xeb\x69\xa7\xa4\x77\xb4\xf3\xf3\xa8\xc1\xda\x91\x33\xac\x96\x6d\x68\x2e\xe0\x7a\xfa\x39\xb4\xf5\xbd\x90\x83\x1d\x18\xc5\x2a\xa4\x40\x72\xc3\xd6\xf8\x7f\xb0\x91\xcf\x60\xe4\x8f\x56\xd1\xe2\xb0\x35\x9e\x83\xe9\x40\x5f\x3b\xdb\xd9\xf6\x9f\x6d\xbc\x41\xec\x2c\x7c\x0e\xc1\xc9\x8d\x3c\x88\xc4\x03\xc2\x83\xd0\x7e\x38\xee\xdd\x8f\x83\xc1\x61\x4e\xb1\xd5\x6e\xf7\xbb\xf6\x24\x2a\x4c\xc9\xc7\x41\x62\xdc\xd7\x4d\x56\xe7\x8e\x83\x63\xe0\x87\x87\x25\xdd\x6e\x78\x91\xb1\xf7\x77\x3c\xb8\x67\x41\xaf\x38\xe9\xee\x62\x6d\x25\x02\xbb\xfd\x47\x60\x71\x0c\x6f\x0d\x51\x06\x08\xfc\xf8\x06\xea\x8a\x12\x63\xb3\x97\x04\x9b\x1f\x7d\xc7\xb9\xfd\x4a\x2c\x23\x4a\xc3\x52\xaa\x0d\x51\xb4\xe9\xcf\x98\x02\xb7\xee\x87\xb8\xb6\xf4\xd3\x68\xde\xd8\x53\x6c\x4d\xf8\xf8\xe8\xde\xf7\x74\x3c\x8a\xfa\x2e\x1f\x4d\x22\x24\x79\x71\x4c\xe8\x32\x56\x27\x37\x85\x6f\xdd\x15\x60\xfc\x74\x6c\x0a\xa6\x27\x11\x31\x46\x8d\x47\x03\x30\x8c\x26\xd6\xaf\x97\xbd\x2b\x59\xb7\x3c\x19\x84\xd5\x63\x3c\xf6\xc5\x74\x57\x08\xb4\xe4\xb9\xd6\x63\x8f\xab\xd1\x45\x8f\xf7\x10\x56\xa3\x67\xa3\xce\x51\xfb\xf0\xde\xef\x23\x3d\xa9\xc9\x80\xf5\xc8\x46\xd9\xe8\x48\x3c\xa1\xf4\xb5\x8d\x9f\x71\x70\x22\xd2\x0f\xd1\x31\xe9\x8c\xed\xcf\xeb\x47\xad\xec\xbf\xa8\x79\xc0\xc4\x8c\x8e\x26\x91\xae\x33\xdf\x9b\x18\xbf\xe8\x2e\x60\x2d\x99\x03\xef\x61\x2a\x38\x2a\x28\xac\x88\x61\x51\x11\x1e\x14\x21\x8f\x64\x8d\x46\xa4\xdf\xd5\xee\xc2\x1a\x7c\x3a\xe9\x5a\x5b\x5f\x6b\x5b\x5c\xf9\xb6\xff\x06\x33\xed\x3a\x09\xd0\xe0\xdd\x75\x73\x7c\xd7\xe6\xd5\xf7\x6f\x7a\x9d\x9b\x2e\x22\xc6\x8e\x7b\xf7\x01\xe7\xa9\x3e\xc9\xc9\x2f\x46\x37\x9b\x4d\xe4\x7f\xcb\x72\x4d\xfc\xae\x91\x12\x93\x8a\x45\xef\x75\x00\x44\x6f\x45\x0e\x14\x97\xa8\x16\x3d\xf6\x4d\x77\x25\x89\xfd\xb7\x8c\x49\xec\xbf\xd7\xfe\xdf\x00\x00\x00\xff\xff\x0a\x69\xbc\xea\xc0\x2d\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"faucet.html": faucetHtml,
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
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
