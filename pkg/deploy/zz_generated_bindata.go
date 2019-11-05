// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/ccm.yaml
// manifests/coredns.yaml
// manifests/local-storage.yaml
// manifests/metrics-server.yaml
// manifests/rolebindings.yaml
// manifests/traefik.yaml
package deploy

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ccmYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x3f\x4f\x33\x31\x0c\xc6\xf7\x7c\x0a\xab\x4b\xa5\x57\xba\x56\xef\x86\x6e\x84\x81\xbd\x12\xec\xbe\xc4\xb4\xa1\xb9\x38\xb2\x9d\xab\xe0\xd3\xa3\x6b\xcb\xc0\x15\xfa\x6f\xb3\xac\x3c\x3f\x3f\x8e\x6d\x2c\xf1\x95\x44\x23\xe7\x16\xa4\x43\xbf\xc0\x6a\x1b\x96\xf8\x89\x16\x39\x2f\xb6\x0f\xba\x88\xbc\x1c\xfe\xbb\x6d\xcc\xa1\x85\xa7\x54\xd5\x48\x56\x9c\xc8\xf5\x64\x18\xd0\xb0\x75\x00\x19\x7b\x6a\xc1\x27\xae\xa1\xf1\x9c\x4d\x38\x25\x92\xa6\xc7\x8c\x6b\x12\x27\x35\x91\xb6\xae\x01\x2c\xf1\x59\xb8\x16\x1d\x45\x0d\xcc\x66\x0e\x40\x48\xb9\x8a\xa7\x63\x8e\x06\xca\xa6\x0e\x60\x20\xe9\x8e\x39\x2f\x84\x46\xfb\xb0\xa0\xf9\xcd\x3e\xaa\x25\x8c\xc9\xab\xa0\x99\x03\x4d\x98\xf3\x7f\xf3\x1b\xb4\x4b\x35\xb4\x3a\x41\x1c\xbc\x5c\x05\x51\x92\x21\xfa\xa9\x87\x14\xd5\x7e\xef\x6a\x0c\x77\x37\xe3\xd1\x7b\xae\x7f\xfd\xde\x55\xa0\x32\x2e\x83\x1a\x65\x1b\x38\xd5\x7e\x6a\x78\x4d\xf6\xd3\xf8\x7d\x76\x29\x87\xc2\xf1\xdc\x98\x4f\x0a\xed\x4e\xe6\xde\x34\xee\xfe\xed\x7d\x8c\x39\xc4\xbc\xbe\x69\x89\x39\xd1\x8a\xde\xc6\x97\xdf\x2d\x9e\xa9\xea\x00\x4e\x4f\xe6\x62\x0d\xad\xdd\x3b\x79\xdb\xdf\xca\x41\xfe\xa2\x24\x97\x75\x87\x07\x5a\xd0\x53\x0b\xdb\xda\x51\xa3\x1f\x6a\xd4\xbb\xaf\x00\x00\x00\xff\xff\xed\x32\x7e\x6b\xe0\x03\x00\x00")

func ccmYamlBytes() ([]byte, error) {
	return bindataRead(
		_ccmYaml,
		"ccm.yaml",
	)
}

func ccmYaml() (*asset, error) {
	bytes, err := ccmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ccm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\x1b\x37\x13\xbe\xeb\x57\x10\xfb\x22\xb7\x77\x65\x0b\x86\x53\x97\xb7\x44\x72\x13\x03\xb1\x22\x58\x76\x2e\x45\x11\x8c\xb8\x23\x89\x35\x97\xc3\x92\x5c\xc5\x6a\xea\xff\x5e\x70\xbf\x44\xca\xeb\x20\x31\xd2\x3d\x91\x1c\xce\x33\xe4\x7c\x3c\x9c\x05\x23\x3f\xa1\x75\x92\x34\x67\xbb\xc9\xe8\x5e\xea\x82\xb3\x25\xda\x9d\x14\xf8\x46\x08\xaa\xb4\x1f\x95\xe8\xa1\x00\x0f\x7c\xc4\x98\x86\x12\x39\x13\x64\xb1\xd0\xae\x9d\x3b\x03\x02\x39\xbb\xaf\x56\x98\xbb\xbd\xf3\x58\x8e\xf2\x3c\x1f\xc5\xd0\x76\x05\x62\x0c\x95\xdf\x92\x95\x7f\x83\x97\xa4\xc7\xf7\x17\x6e\x2c\xe9\x64\x37\x59\xa1\x87\xce\xf2\x54\x55\xce\xa3\xbd\x21\x85\x89\x59\x05\x2b\x54\x2e\x8c\x58\x6d\xc7\x6a\xf4\x58\xeb\xaf\x88\xbc\xf3\x16\x8c\x91\x7a\xd3\x18\xca\x0b\x5c\x43\xa5\xbc\xeb\xcf\xdb\x9c\x8a\x77\xc7\xb6\x95\x42\xc7\x47\x39\x03\x23\xdf\x59\xaa\x4c\x8d\x9c\xb3\x2c\x1b\x31\x66\xd1\x51\x65\x05\xb6\x6b\xa8\x0b\x43\x52\xd7\x60\x39\x73\x8d\x67\x9a\x89\xa1\xa2\x19\xf4\x4e\x08\xd3\x1d\xda\x55\xab\xab\xa4\xf3\xf5\xe0\x0b\x78\xb1\xfd\x3e\x7b\x9a\x8a\x63\x98\x0d\xfa\x9f\xe1\xd0\xb7\x52\x17\x52\x6f\x12\xbf\x82\xd6\xe4\x6b\xf5\xd6\xb9\x43\xb8\x89\xbf\xa1\xf2\x54\x99\x02\x3c\x72\x96\x79\x5b\x61\xf6\xf3\xc3\x43\x0a\x6f\x70\x5d\x9f\xaf\x75\xd8\x37\x2e\x3c\x62\xec\x69\xee\x3c\x83\xec\xaa\xd5\x9f\x28\x7c\x1d\xfb\xc1\x54\x7f\x71\x82\xf7\xb5\x33\x25\xbd\x96\x9b\x6b\x30\x2f\x29\x9b\x6e\xfb\x94\x2c\xae\xa5\x42\xce\xfe\xa9\x7d\x3a\xe6\xe7\x67\xec\x6b\x3d\x0c\x1f\x5a\x4b\xd6\xf5\xd3\x2d\x82\xf2\xdb\x7e\x6a\x11\x8a\x7d\x3f\x3b\x84\x83\xbd\xfa\x3a\xfd\x70\xb7\xbc\xbd\xbc\xf9\x3c\xfb\x78\xfd\xe6\x6a\xfe\xf8\x8a\x49\x9d\x43\x51\xd8\x31\x58\x03\x4c\x9a\xd7\xcd\xe0\x60\x89\xd5\x49\xce\xa4\x76\x28\x2a\x8b\xd1\x7a\x65\x9c\xb7\x08\x65\xb4\xb4\x06\xa5\xfc\xd6\x52\xb5\xd9\x0e\x03\xf7\x7b\x1f\x0f\x67\x27\xe7\x1d\x3b\x41\x2f\x4e\x5a\xef\x9c\xcc\xa9\xc0\xf7\xf5\x72\x7c\x0e\x8b\x8a\xa0\x60\x13\x37\x6c\x70\x00\xda\x58\x2a\xd1\x6f\xb1\x72\x8c\xff\x3a\x39\x3f\xeb\x05\x6b\xb2\x5f\xc0\x16\x6c\xdc\xd8\x0d\x05\xa8\x76\x63\x41\x7a\xdd\x6f\x11\x20\xb6\xc8\xce\x4e\xfb\x05\x45\x64\x46\xe9\x59\x22\x19\x14\x2b\x50\xa0\x45\xe3\xa0\xc7\x27\xd9\x01\xc6\xb8\x93\x3e\x45\x66\x68\x14\xed\x4b\x7c\x19\xb5\x1e\x55\xdb\x85\xcb\xc1\x98\x76\x4b\xa3\x78\x5c\x83\x0d\x70\x16\x92\x6a\x36\x5f\x66\x23\x67\x50\x04\xed\xff\x59\x34\x4a\x0a\x70\x9c\x4d\x46\x8c\x85\x32\xf5\xb8\xd9\x37\xc0\x7e\x6f\x90\xb3\x1b\x52\x4a\xea\xcd\x5d\x5d\xf0\x0d\x41\xc4\x2b\xbc\xf5\x41\x09\x0f\x77\x1a\x76\x20\x15\xac\x42\xd6\xd6\x70\xa8\x50\x78\xb2\xcd\x9e\x32\x30\xe0\x87\xe8\xe0\xc3\x47\xf7\x58\x1a\xd5\x03\xc7\xde\xa9\x1d\x9d\xe8\x3f\x77\xf9\xee\x7a\xf5\x38\x29\xef\xf9\x91\x87\xeb\x7b\x92\x42\x1b\x33\x60\xf8\x72\x76\x8f\xfb\xe0\x32\x2b\xbd\x14\xa0\xde\x14\x05\x69\xf7\x51\xab\x7d\x16\xa5\x1f\x99\xa0\x49\x96\xb3\xec\xf2\x41\x3a\xef\x3a\x61\xe0\xf0\x65\x72\xfd\xf0\x05\x5a\x3e\x22\x53\x72\x9c\x29\xa9\xab\x87\x76\x93\x20\xed\x41\x6a\xb4\xfd\x59\xf2\x27\x69\xd1\x7c\xb2\x84\xcd\x61\xb9\xab\x1d\x3e\x19\xbf\x1e\x9f\xa5\x9b\x16\x95\x52\x0b\x52\x52\xec\x39\xbb\x5a\xcf\xc9\x2f\x2c\x3a\xac\xb9\xae\xcb\xe6\xe8\x01\xea\x73\x5a\x96\xd2\x27\x2b\x21\x1c\x25\xd9\x3d\x67\x93\x5f\x4e\xaf\x65\x52\x9a\x7f\x55\xe8\x8e\x77\x0b\x53\x71\x36\x39\x3d\x2d\x07\x31\x12\x08\xb0\x1b\xc7\xd9\xef\x2c\xcb\x43\x11\x66\xff\x67\x59\x42\x08\x1d\x1b\x66\xec\x8f\x5e\x65\x47\xaa\x2a\xf1\x3a\x44\x35\x89\x5b\xe7\xad\x40\xc2\x79\xb3\x29\xb2\x5f\x86\xfd\x0b\xf0\x5b\x9e\x50\x4e\x72\x17\x28\x42\x9c\x39\x0b\x6f\xdb\x81\x4b\xc8\xa6\x76\xfa\x48\x2d\xc8\x7a\xce\x22\x76\xe9\x0a\x39\xc5\x35\x96\x3c\x09\x52\x9c\xdd\xcd\x16\x3f\x8a\x93\x7b\x61\x06\xb1\x6e\xa7\xdf\xc0\x4a\x38\xaf\x43\x2b\xd1\x5b\x29\x86\x4f\x16\xa3\xd5\x7c\x2f\xfd\x7e\x4a\xda\xe3\x83\x8f\x43\x0b\x4a\xd1\x97\x85\x95\x3b\xa9\x70\x83\x97\x4e\x80\xaa\xeb\x87\x07\x3e\x76\xb1\xbb\x05\x18\x58\x49\x25\xbd\xc4\xa3\xe4\x80\xa2\x48\x17\x72\x36\xbf\xbc\xfd\xfc\xf6\x6a\x3e\xfb\xbc\xbc\xbc\xf9\x74\x35\xbd\x4c\xc4\x85\x25\x73\xac\x00\x4a\x0d\x04\xee\x86\xc8\xff\x26\x15\xb6\x2f\x7f\x1a\x46\x25\x77\xa8\xd1\xb9\x85\xa5\x15\xc6\x78\x5b\xef\xcd\x3b\xf4\xa9\x09\xd3\x24\xca\xd1\xf3\xca\xda\x74\xe0\xec\xe2\xf4\xe2\x34\x59\x76\x62\x8b\xc1\xc9\xef\x6f\x6f\x17\x91\x40\x6a\xe9\x25\xa8\x19\x2a\xd8\x2f\x51\x90\x2e\x1c\x67\xaf\x63\x55\x2f\x4b\xa4\xca\xf7\xc2\xf3\x48\xe6\x2a\x21\xd0\xb9\xdb\xad\x45\xb7\x25\x55\x34\xec\xda\x7d\x6b\x90\xaa\xb2\x18\x49\xcf\x93\x46\x40\xfe\xf0\x75\xd3\xf6\x21\xba\xed\xe4\x62\xf2\xe2\xdb\x4e\xfe\xeb\xdb\x16\xda\x75\x24\x37\x6b\xda\xcb\x56\xd0\x70\xc0\x0f\x70\x84\xe8\x1a\xb8\xd4\x3b\xc3\x34\x5c\x5f\xd8\x63\xe9\x8e\x93\xb3\x7e\x3f\x3a\xe2\x4a\x64\x9d\xa3\x07\x85\xad\x62\xdf\x02\x0d\x6a\x1e\xa4\xcf\x76\xa1\x6d\x5b\x3b\xd0\x5f\x44\x4f\xe5\xb3\x0d\xc6\x93\xbf\x82\x43\x2b\x15\x5e\xad\x26\x1f\xb2\xc0\x2f\xd9\x80\xd8\x09\x0b\xe6\xd9\xbf\x83\xef\xe8\x57\x44\xd3\xc8\xe7\xed\xe3\x1d\x21\x7d\x6f\x67\x93\xf6\x1e\x43\x36\x5b\x1b\x57\x0b\x1e\xb7\xc5\xf3\xe5\xe3\xab\x51\xc4\xf6\xf9\x11\x97\x9b\x98\xa4\x8f\x29\x3d\x1f\x20\xec\x67\x14\x1a\xa6\xcd\x07\x38\xd9\xa4\xd4\x9d\xaa\xfc\x1b\x00\x00\xff\xff\x4f\x8a\x56\xe7\xad\x0f\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\xe3\x36\x10\xbe\xeb\x57\x10\x02\xf6\x52\x2c\xed\xf5\xad\xe0\xcd\xf5\x63\xbb\x80\x5f\x88\xd3\xf6\x10\x2c\x8c\x31\x35\xb6\xb9\xa1\x48\x81\x1c\xa9\x48\xd3\xfc\xf7\x82\x7a\x45\x8a\x93\xd8\x41\xb1\xbc\x28\x19\xce\x37\xdf\x70\xf8\xcd\xd0\x90\xa9\x3f\xd1\x79\x65\x8d\x60\xc5\x28\xba\x57\x26\x11\x6c\x8b\xae\x50\x12\xc7\x52\xda\xdc\x50\x94\x22\x41\x02\x04\x22\x62\xcc\x40\x8a\x82\x69\x2b\x41\xf3\x0c\xe8\xc4\x33\x67\x0b\x15\xf0\xe8\xb8\xaf\x70\x1c\x6a\x60\xe5\xee\x33\x90\x28\xd8\x7d\xbe\x47\xee\x1f\x3c\x61\x1a\x71\xce\xa3\x2e\xb3\xdb\x83\x1c\x40\x4e\x27\xeb\xd4\x3f\x40\xca\x9a\xc1\xfd\xaf\x7e\xa0\xec\xb0\x18\xed\x91\xa0\x49\x6c\xa2\x73\x4f\xe8\x6e\xac\xc6\xeb\xb3\x72\xc1\xdb\xe5\x1a\xbd\x88\x38\x83\x4c\x7d\x75\x36\xcf\xbc\x60\x77\x71\xfc\x3d\x62\xcc\xa1\xb7\xb9\x93\x58\x5a\x8c\x4d\xd0\xc7\x9f\x59\x9c\x85\xdc\x3c\xa1\xa1\xc2\xea\x3c\x45\xa9\x41\xa5\xbe\x04\x14\xe8\xf6\xa5\xf3\x11\x29\xb8\x6a\xe5\xcb\xef\xdf\x40\xf2\x14\x7f\xbf\x4c\x82\x26\xc9\xac\x32\xf4\x2a\x51\x65\xb4\xc9\x0b\xae\x5f\xae\x0a\x5c\x60\x88\xda\x03\x4a\x87\x40\x58\x06\x7d\x3d\x3f\x4f\xd6\xc1\x11\xeb\x8a\x9f\x07\xad\xf7\xa5\x06\xef\xf1\xca\x0a\xfc\xff\xfb\xfd\x4d\x99\x44\x99\xe3\xf5\xd7\xbc\x57\x26\x89\xc2\x5d\xdf\xe0\x21\x38\x37\x67\x7c\x87\x3d\x62\xec\x5c\x57\xd7\xa8\xc9\xe7\xfb\x1f\x28\xa9\x14\xd4\xab\x2d\xf3\xb3\x1a\x05\xb2\xcc\x0f\xdb\x3e\x9d\x62\xa6\xed\x43\x8a\x1f\xe8\xd1\xb7\xa9\x7c\x86\x52\x94\x77\x9f\x69\x25\xc1\x0b\x36\x8a\x18\xf3\xa8\x51\x92\x75\x61\x87\xb1\x34\xdc\xef\x02\xf6\xa8\x7d\x65\x08\x65\xce\xde\xe1\x22\x4c\x33\x0d\x84\x35\xbc\x93\x64\x58\xba\x17\xe9\x52\x2c\xc6\x9a\x14\xcb\xbf\x7b\x05\x5f\x7d\xac\xda\x61\x49\x6b\x08\x94\x41\xd7\x26\xc0\x2f\x55\xae\x5a\x2a\x85\x23\x0a\xe6\xc0\xc8\x13\xba\xe1\xeb\xde\xa2\xf8\x32\xf8\x32\x18\x8d\xfa\xa8\x4d\xae\xf5\xc6\x6a\x25\x1f\x04\xfb\x76\x58\x59\xda\x38\xf4\xd8\xa6\x14\x92\x4a\x53\x30\xc9\x73\x49\xf8\xa5\x6c\x38\xf3\x04\x8e\x3a\xff\x73\x2e\xad\x39\xa8\x63\xc7\x34\x44\x92\xc3\xca\x5a\x7f\x06\x3f\xbc\x35\xad\x47\x35\x79\x96\xa1\x3a\xbe\xcb\x5d\xd5\xa3\x42\xf0\xca\xa9\xdd\x65\x2c\x0d\xfe\x1b\xa0\x93\xe8\x11\xb4\x1e\x68\x8a\xf3\x60\x9b\xf5\x74\xb7\x1a\x2f\x67\xdb\xcd\x78\x32\xeb\x04\x2b\x40\xe7\x38\x77\x36\x15\x1d\x23\x63\x07\x85\x3a\xa9\x3b\xfa\xcc\x5e\x71\x37\xaa\x1a\xb4\xc2\x8e\xba\xa7\xfa\xc0\x81\x2a\xfb\x12\xb2\x3e\xdb\x99\x2a\xea\xfa\xbe\x6c\xce\xfe\x20\x7d\x6e\xd3\x6d\x65\x9f\x84\x01\xfa\x6e\xa3\x86\xa9\x65\x8c\xa5\x72\x46\xd5\x79\x77\xa7\xef\x20\xf4\xab\x33\x48\x58\x12\x28\xcf\x13\x3c\x40\xae\x89\x97\xdb\x82\xc5\xe4\x72\x8c\xa3\xae\x0e\x1b\x9d\x06\x40\x87\xa9\x3a\x7b\x3d\x64\x97\x36\x41\xc1\xfe\x02\x45\x73\xeb\xe6\xca\x79\x9a\x58\xe3\xf3\x14\x5d\xe4\xaa\x87\xaf\x11\xed\x14\x35\x12\x96\x27\xaf\x27\x67\x53\xb2\xe8\xc5\x2f\x89\x77\x07\x52\x2b\xd0\x37\x66\x51\x03\xec\x68\x55\xb0\x7f\x79\x7b\x2b\x8f\x2f\xd4\xc0\x58\xf9\x70\x07\x3d\x2c\x21\x8b\xc5\xdd\xd9\xfe\x39\xa2\x87\x8c\x45\x3c\x9d\xcd\xc7\x7f\x2c\x6e\x77\x9b\xf1\xed\xef\xbb\xf9\xfa\x66\xb7\x5a\xaf\x76\x8b\x6f\xdb\xdb\xd9\x74\xb7\x5a\x4f\x67\xdb\xf8\xf3\xdb\x31\xc2\xa9\x7c\x2c\xee\xe2\x4f\x8f\x4d\x9c\xc5\x7a\x32\x5e\xec\xb6\xb7\xeb\x9b\xf1\xd7\x59\x19\xf5\xe9\x53\xf9\x78\xf6\xd7\xd3\x99\xe5\xd9\xe7\x29\xfa\x2f\x00\x00\xff\xff\x90\x65\x00\x29\x9f\x09\x00\x00")

func localStorageYamlBytes() ([]byte, error) {
	return bindataRead(
		_localStorageYaml,
		"local-storage.yaml",
	)
}

func localStorageYaml() (*asset, error) {
	bytes, err := localStorageYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local-storage.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x6f\x1b\x37\x10\xbd\xef\xaf\x18\x6c\xaf\xa5\xe4\x20\x41\x11\xf0\x96\x26\x40\x51\x20\x2e\x84\x38\xf0\xc5\xf0\x81\xe2\x8e\x56\xac\x96\x1f\x18\xce\xca\x55\x8b\xfe\xf7\x82\xfb\x25\xad\x3e\xb6\xb2\x0d\x03\x39\x69\x45\xf2\xcd\xbc\x79\x7c\x33\x14\x42\x64\x2a\x98\x7b\xa4\x68\xbc\x93\x40\x4b\xa5\x67\xaa\xe6\xb5\x27\xf3\xb7\x62\xe3\xdd\x6c\xf3\x31\xce\x8c\x9f\x6f\xdf\x65\x1b\xe3\x0a\x09\x9f\xab\x3a\x32\xd2\x37\x5f\x61\x66\x91\x55\xa1\x58\xc9\x0c\xc0\x29\x8b\x12\xe2\x2e\x32\x5a\x69\x91\xc9\xe8\x28\x22\xd2\x16\x29\xa3\xba\xc2\x28\x33\x01\x2a\x98\xdf\xc8\xd7\x21\x26\x84\x80\x3c\xcf\x00\x08\xa3\xaf\x49\x63\xb7\x16\x7c\x11\x9b\x0f\xe7\x0b\x3c\xf8\x9a\x47\x56\xdc\xfd\x57\x16\x63\x50\xba\xd9\xde\x22\x2d\x3b\x68\x89\xdc\xfc\x56\x26\xb6\x1f\x4f\x8a\xf5\x3a\x7b\x5d\x91\xbf\x1a\x57\x18\x57\x5e\x5f\xab\xaf\xf0\x1b\xae\xd2\xb1\xbe\xda\x89\x94\x19\xc0\xa9\xac\xd3\x09\x62\xbd\xfc\x13\x35\x37\x7a\xb6\xd8\x3b\xa4\xad\xd1\xf8\x49\x6b\x5f\x3b\x1e\xe0\x47\x38\xd8\xeb\x26\x61\x53\x2f\x51\xb4\xf1\x9f\xa5\xcf\x12\x59\xf5\x22\x4d\xab\x33\xce\x2e\x52\x3c\x41\xa8\x8a\x29\x26\x2f\xd4\x6e\x24\x1a\xfe\xc5\xe8\x52\x21\x42\x05\x73\x90\x1c\x1d\x1b\xdd\xc0\x7b\x1a\x3f\x8c\x90\x57\xb9\x6d\x4c\x42\x76\xde\x68\x64\x2d\xb0\xc2\x52\xb1\x7f\x0b\xf3\x1d\x25\x78\x13\xcd\x4e\x49\x5c\xdd\xae\x17\xbb\x52\x95\x25\x25\xce\x58\x88\x9e\xca\xe0\xbe\x4a\x2d\xb1\x6a\x66\x06\x4c\x44\x1f\x22\x08\xf6\x62\x6b\xf0\x49\x42\xce\x54\x63\xfe\x1c\x1c\x16\x86\x5f\x82\x53\x85\x35\x6e\x00\x9e\x99\x9f\xf0\x90\x77\x75\x75\xf8\xfc\x71\x34\x4b\xe1\x21\x4f\x93\x34\xff\x19\xf2\x66\x7c\x36\xdb\xed\xac\x84\x87\xbc\x44\x4e\x3b\x69\x50\xa6\xdf\x66\x4e\xe6\x8f\x27\x06\x56\xc1\x10\x96\x26\x32\x4d\xb8\xf7\xd3\xe2\xf7\xce\x03\x67\x6e\xa3\x3b\x3a\x1b\x73\xcd\x62\x40\x9d\x0e\xc5\x16\xd8\x5e\xc5\x05\xe3\x5c\xb4\x0e\x40\xd9\x1a\xfc\x28\x78\x53\x67\x5b\x40\xcf\x14\xc0\xb8\x88\xba\x26\xbc\xdb\x98\xf0\xfd\xeb\xdd\x3d\x92\x59\xed\x24\x24\x7d\xfb\x40\x0b\x32\x9e\x0c\xef\x6e\x8d\x33\xb6\xb6\x12\xde\xdd\xdc\xec\x83\xf5\xbb\xed\xf2\xb1\x52\xc3\xcb\x71\x59\x8a\x6b\x3b\x62\xec\xcf\xb4\x43\x0e\x19\x1b\xe5\xdb\x48\xf9\xed\x28\x54\x7e\xe6\xa0\x6e\x9b\x49\xf4\x02\xf7\x4e\xda\x0b\x5f\xa1\x66\x4f\x5d\x92\x8f\x51\xa8\x10\xce\x70\x0c\x9e\x78\x78\x98\x89\x25\x7c\xf8\xf0\xbe\x81\x04\xf2\xec\xb5\xaf\x24\x7c\xff\xbc\x68\x56\x58\x51\x89\xbc\x18\x4e\xfd\x9f\x44\xfd\xc4\x78\xb9\x52\xa7\x76\x0d\x71\xff\x86\x7f\xc1\x50\xf9\x9d\xc5\x57\xa5\x38\xba\x8c\x0b\x3a\x9d\x57\xd5\xa6\xae\xfa\x7a\x00\x9f\x12\x9a\xd1\x86\x4a\x71\xd7\x0a\x87\x84\x27\x5b\x63\xcc\x6f\x3a\x05\x40\xcf\xb3\xf9\x1e\xdd\xc2\x1f\x97\x13\x6c\x7d\x55\x5b\x1c\x32\xfc\x04\x36\x01\xc0\x38\x60\x1b\x20\x7a\x78\x42\xd0\xca\x41\x54\x2b\xac\x76\x50\x47\x84\x15\x79\x2b\xa2\xa6\x24\x00\x18\xab\x4a\x8c\xa0\x5c\x31\xf7\x04\x69\x08\x0b\xef\xaa\x1d\x68\xef\x58\x19\x87\x14\xbb\xc8\xa2\x2b\x93\x6d\x10\x85\xa1\xa1\x22\xb4\x81\x77\x5f\x0c\x49\xf8\xe7\xdf\x6e\x71\x8f\x95\x47\xe0\xb3\x25\x40\x4b\x42\x02\x29\xa7\xd7\x48\xf3\xa3\xc7\x74\x7b\x33\x7b\x3f\xfb\x65\x38\xdc\x56\x7c\x9b\xca\x3c\x10\xf6\x12\x3d\x68\x05\x59\x28\x5e\x4b\x98\xb3\x0d\xd9\x7f\x01\x00\x00\xff\xff\x67\x49\xf7\x31\x56\x0b\x00\x00")

func metricsServerYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerYaml,
		"metrics-server.yaml",
	)
}

func metricsServerYaml() (*asset, error) {
	bytes, err := metricsServerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rolebindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x21\xbb\x72\x38\xdc\x72\xf0\xd8\x0e\xdd\x03\xb4\x3b\x6d\xb3\x09\x6b\x59\x14\x48\x2a\x41\xfb\xeb\x0b\xa7\x6e\x82\xa4\x76\x90\xb4\xdd\x24\x41\x7c\x1f\x1f\xf9\x20\xd3\x13\x8a\x12\xa7\xca\x4b\x0d\xcd\x12\x8a\x6d\x58\xe8\x0d\x8c\x38\x2d\xbb\xff\xba\x24\xfe\xb3\xfd\xeb\x3a\x4a\x6d\xe5\xef\x63\x51\x43\x59\x71\xc4\x3b\x4a\x2d\xa5\xb5\xeb\xd1\xa0\x05\x83\xca\x79\x9f\xa0\xc7\xca\x77\xa5\xc6\x00\x99\x14\x65\x8b\x12\x86\x6b\x44\x0b\xd0\xf6\x94\x9c\x70\xc4\x15\x3e\x0f\xbf\x21\xd3\x83\x70\xc9\x17\xc8\xce\xfb\x2f\xe0\x03\x47\x5f\xd5\xb0\xaf\x0e\xfa\x99\x46\x86\x96\xfa\x05\x1b\xd3\xca\x85\x9b\x20\x8f\x8a\x32\xe3\xc2\xb9\x10\x82\xfb\xfe\xb4\x26\xc6\xf4\xd9\xfe\x3f\x0d\x0d\x27\x13\x8e\x11\xc5\x49\x89\x78\xd2\xb8\x0e\x15\xc1\x2f\x16\xce\x7b\x41\xe5\x22\x0d\x8e\x6f\x89\x5b\x54\xe7\xfd\x16\xa5\x1e\x9f\xd6\x68\x57\xd6\x42\x8f\x9a\xa1\x39\x17\x88\xa4\xb6\x3f\xec\xc0\x9a\xcd\x84\x56\x42\xdb\xb1\x74\x94\xd6\xa3\xdf\x29\xf1\x8f\x3f\x99\x23\x35\x74\x33\x61\x42\x10\x53\x9b\x99\x92\xe9\xfe\x96\xb9\x9d\xd3\x1c\xfc\x1f\xb5\x7f\xb8\xb4\xf9\x88\xcf\xec\xee\xf7\xb3\x7d\x0a\x38\x06\x7b\xf0\x78\x1d\xe3\x2c\xdc\x97\x01\xef\x01\x00\x00\xff\xff\x46\xd3\x6d\x9d\x0f\x04\x00\x00")

func rolebindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_rolebindingsYaml,
		"rolebindings.yaml",
	)
}

func rolebindingsYaml() (*asset, error) {
	bytes, err := rolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xcd\x6a\xc3\x30\x10\x84\xef\x7e\x8a\x25\x90\x63\x64\x72\x0a\xe8\xd6\x1f\x43\x4b\xa1\x84\xa6\xed\xb5\xac\xe5\x49\x2c\x22\xc9\x42\xbb\x0a\xb4\xa5\xef\x5e\x1c\x7c\xcc\x71\x77\x3e\x3e\x66\x38\xfb\x4f\x14\xf1\x53\xb2\x34\x22\x44\xe3\x58\x35\xc0\xf8\xa9\xbd\x6c\x9b\xb3\x4f\x83\xa5\x27\x84\xf8\x30\x72\xd1\x26\x42\x79\x60\x65\xdb\x10\x25\x8e\xb0\xa4\x85\x71\xf4\xe7\xe5\x96\xcc\x0e\x96\xce\xb5\xc7\x46\xbe\x45\x11\x1b\xc9\x70\x33\xee\x66\x81\xa5\x51\x35\x8b\x6d\xdb\xf5\xef\xcb\xc7\x7d\xf7\xf6\xda\xbd\x77\x87\xaf\xbb\xfd\xf3\xdf\xba\x15\x65\xf5\xae\xbd\x82\xd2\x2e\xe2\xcd\xd6\xec\x76\x66\x6b\xf4\xf4\xd3\x10\x09\x74\x76\x11\x95\x9e\x9d\x41\xe2\x3e\x60\xb0\xb4\xd2\x52\xb1\xba\x06\x22\xe1\xe6\x3f\x42\x8b\x77\x62\x72\x99\x22\x74\x44\x95\x9b\xd8\xdc\xbc\x24\x28\xc4\xf8\x74\x2a\x10\xe9\xd2\x90\x27\x9f\xd4\x54\xc1\x23\x8e\x5c\x83\xee\x6b\x1f\xbc\x8c\x18\x0e\x28\x17\x3f\x0f\x5e\x0c\xff\x01\x00\x00\xff\xff\xd8\x40\x47\x8d\x4d\x01\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"ccm.yaml":            ccmYaml,
	"coredns.yaml":        corednsYaml,
	"local-storage.yaml":  localStorageYaml,
	"metrics-server.yaml": metricsServerYaml,
	"rolebindings.yaml":   rolebindingsYaml,
	"traefik.yaml":        traefikYaml,
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
	"ccm.yaml":            &bintree{ccmYaml, map[string]*bintree{}},
	"coredns.yaml":        &bintree{corednsYaml, map[string]*bintree{}},
	"local-storage.yaml":  &bintree{localStorageYaml, map[string]*bintree{}},
	"metrics-server.yaml": &bintree{metricsServerYaml, map[string]*bintree{}},
	"rolebindings.yaml":   &bintree{rolebindingsYaml, map[string]*bintree{}},
	"traefik.yaml":        &bintree{traefikYaml, map[string]*bintree{}},
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