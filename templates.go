// Code generated by go-bindata.
// sources:
// templates/chat.html
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

var _templatesChatHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x54\x4f\x6f\xd3\x4a\x10\xbf\xf7\x53\xcc\xdb\xd7\x83\x7d\x78\x5e\xe5\xda\x38\x39\xbc\x13\x77\x0e\x1c\x10\x87\x8d\xb3\x49\xdd\xae\xbd\xc1\xbb\x6e\x12\x45\x91\x6a\x1b\x21\x0e\x15\x48\x20\x51\x51\x90\xaa\x52\xd1\x42\x7b\x40\x08\x51\x24\x54\xda\x0f\x33\x2d\x85\x13\x5f\x01\xd9\xeb\x38\x4e\x49\x0e\xce\x64\x67\xf6\xf7\x67\x66\x62\x77\x5d\x07\xa2\xbd\x02\xe0\xae\x73\xd6\xcd\x03\x00\x57\xfb\x5a\xf0\x36\x66\x09\x66\x87\x98\x65\x98\x3d\x71\xa9\x39\x33\x79\xa5\xc7\xb3\x18\xc0\x0f\x07\xb1\x9e\x74\x7d\x35\x10\x6c\xbc\x06\x1d\x21\xbd\xcd\xe6\xb4\x4c\xc6\x62\x22\x7c\xa5\xff\x2b\x6e\xac\x41\x28\x43\x5e\xe6\x5c\x5a\xa1\xb8\x74\xc6\xed\x76\x64\x77\x5c\x92\xc4\x02\xfc\x6e\x8b\x04\x5c\x29\xd6\xe7\x8a\xb4\x5d\x1a\x0b\x93\x1b\xf2\x8e\x92\xde\x26\xd7\x98\x3e\xbf\xfa\x7e\x89\xc9\x21\x26\xfb\x75\xb9\x98\xbe\xc5\x6c\x17\xb3\x13\x4c\x3f\x61\x76\x8e\xe9\x57\xcc\x8e\x31\xfb\x6c\x90\x7b\x32\x0a\x0a\x6c\x6f\x9d\xe9\x8e\x1c\x91\x99\x15\x57\xf3\x91\x66\x11\x67\x6d\x97\x56\xe1\x2c\x57\xf8\x04\x3d\x1e\xf0\x16\x51\x71\x27\xf0\x35\x81\x2d\x26\x62\xde\x22\xbf\xb6\x93\xab\xcb\x83\x12\xc6\xa5\x39\xbe\xf1\x35\xb3\xe3\x2a\x2f\xf2\x07\x1a\x54\xe4\xb5\x08\xa5\x6c\x83\x8d\x9c\xbe\x94\x7d\xc1\xd9\xc0\x57\x8e\x27\x83\xe2\x8c\x0a\xbf\xa3\xe8\xc6\xc3\x98\x47\x63\xda\x70\x1a\x0d\xa7\x51\xfe\x72\x02\x3f\x74\x36\x14\x31\xb0\x06\xad\x06\x6c\x98\x57\xad\x5e\x1c\x7a\xda\x97\xa1\x65\x4f\x4a\xd9\x5b\x2c\x02\xd3\x2b\x68\x41\x18\x0b\xd1\xac\x25\x02\xd5\xff\x5f\x8e\xa0\x05\xab\x16\xf9\xb7\x6c\x06\xcc\x8c\x13\x7b\xa1\xb4\x1c\x43\x59\x5c\x4d\xa5\x2a\xaa\x41\x10\xdb\x31\x0d\x5a\xa2\x07\xc0\xef\x81\xf5\x8f\x61\x76\xb6\x98\xb0\x6c\x1b\x22\xae\xe3\x28\x84\x1e\x13\x8a\x37\x17\x0b\x8d\xf6\xda\x75\x00\x26\x78\xa4\x2d\x82\xe9\x7b\xcc\x3e\x60\x76\xfe\xfb\x7c\x6f\xbe\x10\xc9\xe9\x8f\xa7\xef\x6e\xce\xf6\x30\xd9\xf9\x79\xb0\x83\xe9\x33\x4c\x77\x30\x39\xc2\xe4\x02\x93\xd7\x98\xbe\xc0\xed\x74\xae\x39\xff\x2c\xe7\x9e\x56\x91\xc1\x75\x14\x0f\xbb\xd6\x82\xea\x79\x71\xed\x98\xd4\xc1\x97\x41\x4f\xed\xe6\xca\x4a\xcd\xdf\xd0\x0f\xbb\x72\x78\x9f\xdc\xe3\x9d\xbb\x05\x13\x79\x50\x33\xbb\xc4\x2a\x54\x95\x98\x9c\x5e\x7f\xbc\xb8\xbe\x7c\x83\xc9\x6e\x61\xf1\x11\x26\x27\xf9\x33\x7b\x99\x57\xa7\x47\x98\x9e\x61\x72\x8c\xc9\xab\x05\xd3\x53\x2e\x14\x87\xc9\x2d\x83\xf9\x76\xf0\xe1\x1c\xdc\x22\x43\xb5\x46\xe9\x64\xe2\xdc\x91\x4a\x4f\xa7\x34\x92\x32\xa8\x9b\x2b\xfb\x22\x43\x4f\x48\xc5\xa1\x05\xcb\x66\x5d\x39\xa8\x86\x72\xf3\x25\xbd\xfa\xf6\xb8\x50\x7c\x51\x3c\xf7\x6f\x4d\xe4\xaf\xce\xcb\xb0\x5c\xb6\x3a\x07\x5f\x20\x99\x6d\xa3\xc3\x06\x83\x7c\x4e\xab\x16\x71\x85\xdf\x26\xb6\x93\x2f\xb3\xc5\x9d\x2e\xd3\xcc\x5e\x42\x62\xbe\xa7\x45\x66\xfe\xb7\x72\x69\xf1\x52\xfc\x13\x00\x00\xff\xff\x24\x54\xa7\x26\x1b\x05\x00\x00")

func templatesChatHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesChatHtml,
		"templates/chat.html",
	)
}

func templatesChatHtml() (*asset, error) {
	bytes, err := templatesChatHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/chat.html", size: 1307, mode: os.FileMode(420), modTime: time.Unix(1503110123, 0)}
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
	"templates/chat.html": templatesChatHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"chat.html": &bintree{templatesChatHtml, map[string]*bintree{}},
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
