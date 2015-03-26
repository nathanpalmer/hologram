package agent

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _test_ssh_key = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x95\x37\x12\xa4\x8c\xae\x46\x73\x56\x31\x39\xf5\xaa\xf1\x26\xc4\x7b\xef\xc9\xa0\xf1\xbe\xf1\xb0\xfa\x37\xff\xc4\x57\xa9\x92\x4f\xd2\x29\x9d\xff\xfb\xaf\x58\x41\x52\xcc\x3f\xae\xc7\xfc\xb1\x5d\x25\x64\x7c\xe1\x8f\x26\x24\xff\x3a\x80\xa1\x28\xc2\x7c\x29\x2c\xc3\x68\x1c\xe3\x08\x0c\x74\x9b\xa7\x97\x7a\x5b\x01\xde\xfc\x00\x69\x3b\x26\x14\xf8\x87\xae\x55\x99\x40\xc5\x9b\x35\x57\xfd\x73\x5a\x3d\xe8\x24\x3d\x9f\x24\xdf\x58\x7f\x01\x3b\xbd\x4b\xae\xfb\xfc\x10\x9c\x84\x17\x50\xc6\x83\x47\x6f\x62\x3e\x30\x8d\xf4\xda\x72\x58\x27\x5a\xdf\xe9\x94\xb4\xb4\x6f\x2b\x3a\xd5\x51\x2e\xe7\x2a\xc0\x23\xd6\x8f\x77\x64\xb6\x3e\x9d\xa4\x00\x9e\x81\x1d\x55\xde\xe4\xec\x89\x4c\x08\x52\x59\xf2\x28\x85\x8e\x36\x14\xed\xa0\x7e\x09\xb1\x2b\x47\x3d\xb5\x02\x69\x58\xdb\xa4\x29\x7d\x51\x94\xae\x32\xe1\xc2\x6c\x6d\xe4\xc2\x4d\xb1\xe8\xd6\x80\x92\x15\x9b\xf9\x43\x3d\x51\x16\x12\x16\x43\x72\x21\xd9\xba\x43\x5b\x0f\x86\x2a\x27\x09\xdc\x60\xe1\xbe\xed\xbf\xa4\x97\xb8\x6c\x4e\x27\xbf\xc0\xd0\xfd\x9e\x16\x4a\x52\xf8\xf4\xe4\x2b\x2b\x5e\x81\xe8\x1a\xa4\xf4\x5e\x1a\xfa\x3c\x4f\x5a\xd4\x15\x73\x2b\xd1\xc9\xf8\xaa\x2d\x43\xe7\xa6\x5b\x7d\x27\x23\x0b\x6e\x6a\x5a\x74\xf7\x24\x67\xe4\xbc\x95\x2c\x21\x57\xb7\x1d\x4c\x67\xb7\x1c\x64\x10\x81\x9b\x5c\x4a\x27\xf6\xbe\xa2\x86\xbd\xd2\xf3\xec\x93\x54\x7c\x4e\x50\x9a\x9a\x5e\x9f\xef\xcf\x6f\xb6\xfc\x36\xde\x1c\x85\x67\x1c\x86\x65\xe6\xbf\xcb\xe6\xaa\x2f\xba\x69\xc5\x36\xb6\x9f\xc1\x65\x65\x80\x9d\x45\x1c\xed\x40\xc1\x14\xde\xd9\xc3\x13\xdc\xd2\x18\xcd\xf5\x6f\x89\xfa\x71\xcd\xb2\x29\xa5\x2d\xdb\x8d\xe7\xc2\x41\x7e\xc0\x83\x75\x20\xe3\xfb\x6b\xe0\xc7\x3a\x33\xd1\xe5\xad\x62\x97\x73\xe0\x6b\x17\x7f\x27\x76\x98\x5b\x0b\x1a\x14\xa3\x50\x65\x46\x9a\x03\x06\x31\xf5\xd3\x13\x6b\xec\x3c\xa6\x6a\xfb\xa9\xaa\xc6\x1b\x2d\xb2\xba\x54\x78\xe1\x38\x0b\xdf\xd2\x73\x43\x7a\xfe\xae\x95\x70\x02\x15\xaa\x89\x10\xb6\x10\xfa\x10\x83\x3c\xdb\xeb\xab\xbc\x28\xfc\x4d\x3a\x1b\x05\x4d\xdc\xb0\x2d\x66\x26\x3a\x61\x68\x39\x7e\xb7\x5c\x93\x3e\x41\xaa\x69\x09\xb6\x8b\x9d\x8a\x5e\x61\xd4\xea\xf0\x80\x26\x4e\xd2\xa1\xa5\x60\x4a\x86\x46\x69\xa8\x4d\x3a\x3b\xa3\x45\xb1\x5e\x3f\x52\xbb\x6c\x96\x9d\xe3\xd2\x5c\xc2\xc1\xc6\x96\xbd\xda\xe7\x0d\x0b\x0d\x39\x4c\x39\xe5\xbe\xb9\x4b\xd2\x01\xe4\xe9\x00\xba\xb0\x95\xe5\xf5\xf4\xf1\xa2\x47\xbb\x0d\xbc\x0a\x3b\x5d\x2f\x12\x83\x4c\x8c\x45\xa8\xda\x4f\x25\xea\xd5\x03\xfd\xbd\x0e\x67\x78\x02\xbe\xcd\x83\x16\x49\xfd\x6b\x79\xa9\x58\x20\x3f\x74\x35\x80\x7a\x39\xfb\xde\x0c\xb8\x3a\x11\x98\xcf\xea\xa7\xf4\x76\xf8\xec\xbb\xb4\x86\x31\xb6\xdf\x8e\x57\x14\x8c\x0d\x05\x02\xbb\x28\x17\xaa\x63\xfc\xeb\x6c\x53\x5c\xa9\xbe\x03\xb7\x29\xae\x3c\x36\x31\x03\x4c\x2d\xbe\x17\x18\x47\x33\x0f\x4e\x75\x14\x6b\xec\xeb\x62\x6e\x53\x8c\x39\x74\x54\xc1\xa0\x67\x4d\x08\x15\x4c\xa4\x15\x38\xa6\x57\xd6\x58\xe8\x0c\x3e\x90\x19\x42\x6b\x81\xfe\xa4\x3a\xac\xab\x80\x98\x86\x34\x96\x4a\x39\x6a\xd8\x68\x68\xdf\x4a\xb3\xf7\xad\xa4\xd6\x0a\x5c\x48\x8e\xd8\x6d\x2f\x06\xff\x7c\xb0\x47\x0b\xb3\x0a\xc7\x1b\xbb\xec\xf6\x26\xbd\x14\x35\xfe\x25\x86\xb9\x38\x07\xae\xaa\x2a\x8a\x5f\x6c\x47\xbf\x3c\x6d\xa3\xf1\x82\x73\xcc\x13\xb7\x99\x66\x9e\x46\xf6\x0c\x06\x84\x95\xe4\x58\xd3\xa7\xfa\x19\x61\xaf\x48\x98\x32\xec\x9e\x96\x9d\x88\x36\xdb\xda\xee\x88\x82\x0d\x10\x04\xd7\x44\x43\x70\xbe\x5e\x47\x3c\xe2\x0a\x3a\x96\x53\x29\xb7\x6c\xfd\x1e\xda\x5e\x6c\x02\xb9\x22\x63\xb0\x4d\x79\xd9\x27\x46\x40\xb5\x16\x36\x87\xb7\x1e\x6e\x22\xc8\xce\xe5\xfc\x64\xfa\x08\xd8\xd0\x90\x71\x62\xad\x21\x63\x4e\x94\xb0\xbf\xcc\x8d\xf0\xfd\x55\x7a\x1b\x2e\x1d\xba\x91\x54\x4c\x4d\x85\x15\x40\xdb\xbf\xc8\x3b\xa5\x96\xcb\x22\x0a\x7f\x21\x0d\x99\x4b\x0e\xa9\x97\xe3\x80\x81\x5b\x92\xbf\x2f\x64\x0d\x7f\x3d\x88\x7c\xea\x53\xc8\x25\x6a\x69\x5b\x70\x42\xfc\x00\xec\x50\x7b\x90\xf5\xfc\x96\xc0\x6e\xde\x22\x15\x2b\x69\x86\x97\xb2\xad\x81\x03\xa9\xe7\x85\x0d\xdf\xe2\x0e\xc0\xcb\x8a\x5d\x91\x7a\xfe\x31\x4d\x93\x19\xf4\x45\x06\x68\x2f\x3f\x9e\xbb\x3c\x8a\x26\x4d\xb7\x39\x5e\x79\x9a\x1a\x58\x56\x9c\x72\x27\x7a\x7b\x94\x30\xe6\x5b\x80\xe4\xac\x58\xdb\x2f\x3a\xc1\x09\x28\x8a\x08\xaf\x5d\x8d\x11\xd8\x2a\xca\x1c\x88\xad\xbf\x7b\x8d\x06\xff\x71\xc1\xd8\xa1\x06\x19\x30\xfc\x2c\xfb\x20\x8b\x07\xd4\xd1\xaa\x97\xb4\x3e\x3b\x73\x2d\xbe\x16\x32\x76\xc9\x66\x01\x07\x40\x35\xa0\x8d\x5e\xb1\xf9\x55\x56\x44\xfe\xe9\x0f\xd2\x83\xfc\xae\x23\x91\x64\x31\xda\x7e\xed\xfd\xb7\xae\xb2\x0f\xf4\xc2\xec\x2c\x24\x60\xbc\x8c\x66\x93\xa8\xd4\x4a\x53\x13\x42\x34\x9f\x6f\x82\x02\x06\xaa\xd3\xdf\x16\xbd\x6b\x2f\x4f\x1e\xb8\x31\xc5\xa8\x1f\xea\x85\x41\xb5\x43\xde\xa8\xa1\xbd\xad\xe6\x94\xab\x25\xd4\x8c\x27\x1f\xaf\xf5\xaa\xca\x06\x4a\xb4\x33\xb8\x7c\x13\x49\x27\x32\xca\x81\x80\xb1\x53\xc6\xd1\xd8\x5a\x0c\x43\xde\xf4\x64\xb3\xa9\x0d\x74\x38\x2d\xad\xf5\x9e\xac\xbf\x83\x03\xbc\xb7\xb2\xab\x7a\x8e\x4b\xea\xa6\xd6\xf1\xf8\x27\x55\x3c\x26\x76\xd1\x44\xed\x26\x3b\x94\x80\xc5\x69\x56\x9c\xbd\x2a\x04\xe5\x94\x57\xb0\x50\x72\xb1\xe8\xc7\x5e\x96\x82\xd9\x13\xcc\x17\xc6\xf8\x38\xb6\x77\x4c\xb1\x5f\x91\x9f\xe7\x11\x97\x01\x09\xd5\x4d\xe7\xbf\x7c\x5c\x14\x41\xd2\xa8\x00\x38\x52\x28\x41\xa4\x60\x81\x4b\x2b\xc5\xbb\xe6\xbd\x40\x51\x1e\xc8\xaf\x4f\x8f\xc8\xfb\xb1\xe9\x80\x60\xbc\x37\x3b\xc2\x2a\x15\xb9\x4a\x17\xad\xac\x16\x6f\xe0\x9f\x52\x04\x93\xff\xdf\xaa\xf9\xff\x00\x00\x00\xff\xff\x1c\x9f\xfc\x7c\x8b\x06\x00\x00")

func test_ssh_key_bytes() ([]byte, error) {
	return bindata_read(
		_test_ssh_key,
		"test_ssh_key",
	)
}

func test_ssh_key() (*asset, error) {
	bytes, err := test_ssh_key_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_ssh_key", size: 1675, mode: os.FileMode(420), modTime: time.Unix(1423483025, 0)}
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
	"test_ssh_key": test_ssh_key,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"test_ssh_key": &_bintree_t{test_ssh_key, map[string]*_bintree_t{}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}