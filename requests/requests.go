// Code generated by go-bindata. DO NOT EDIT.
// sources:
// requests\insert_storage_key.sql
// requests\insert_user.sql
// requests\requests.go

package requests

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataRequestsinsertstoragekeysql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\x28\x2e\xc9\x2f\x4a\x4c" +
		"\x4f\x8d\xcf\x4e\xad\xd4\xc8\x4c\xd1\x51\xc8\x4e\xad\xd4\x51\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\xd1\x54\x08" +
		"\x73\xf4\x09\x75\x0d\x56\xd0\xb0\xd7\x51\x00\x21\x4d\x6b\x40\x00\x00\x00\xff\xff\x66\x91\x95\x92\x3e\x00\x00\x00" +
		"")

func bindataRequestsinsertstoragekeysqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataRequestsinsertstoragekeysql,
		"requests/insert_storage_key.sql",
	)
}

func bindataRequestsinsertstoragekeysql() (*asset, error) {
	bytes, err := bindataRequestsinsertstoragekeysqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "requests/insert_storage_key.sql",
		size:        62,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1580596190, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataRequestsinsertusersql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x31\xab\x84\x30\x10\x84\x7b\xc1\xff\xb0\x58\x19\xd8\x7f\x60\x21" +
		"\xaf\xb0\x10\xc4\x07\x4f\x7d\xed\xb2\x98\x3d\x0d\x87\x46\x92\x58\xdc\xbf\x3f\xa2\x57\x28\x07\x53\xcc\xc0\x0c\xdf" +
		"\xd4\x6d\x57\xfd\xf5\x50\xb7\xfd\x2f\x6c\x1c\x66\xc8\x8d\x46\xd8\xd8\xc9\x1a\x28\xda\xd1\x09\x07\xd1\xc4\x01\x81" +
		"\xc7\x51\xbc\xff\x84\xc5\x6a\xf3\x30\x47\x50\xf0\xff\xd3\x0c\x55\x07\x79\x89\xd0\x0e\x4d\x83\x50\x9e\x52\x45\x9a" +
		"\x5c\x11\xd9\xe4\xec\xbe\x65\x27\x65\xe5\x45\x6e\xd3\xaf\xfa\xee\xc5\x9d\xdd\xe8\x28\x0e\x10\x66\xf6\x33\x82\x0f" +
		"\xd6\xf1\x24\xf4\x94\x17\x82\xb3\x36\x50\xbc\x7f\x5c\x5e\xd8\xac\x74\x80\xc8\x68\x95\x26\x37\xc4\x55\xaa\x78\x07" +
		"\x00\x00\xff\xff\x79\x6b\x78\xd8\xff\x00\x00\x00")

func bindataRequestsinsertusersqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataRequestsinsertusersql,
		"requests/insert_user.sql",
	)
}

func bindataRequestsinsertusersql() (*asset, error) {
	bytes, err := bindataRequestsinsertusersqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "requests/insert_user.sql",
		size:        255,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1580613999, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataRequestsrequestsgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataRequestsrequestsgoBytes() ([]byte, error) {
	return bindataRead(
		_bindataRequestsrequestsgo,
		"requests/requests.go",
	)
}

func bindataRequestsrequestsgo() (*asset, error) {
	bytes, err := bindataRequestsrequestsgoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "requests/requests.go",
		size:        0,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1580658347, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"requests/insert_storage_key.sql": bindataRequestsinsertstoragekeysql,
	"requests/insert_user.sql":        bindataRequestsinsertusersql,
	"requests/requests.go":            bindataRequestsrequestsgo,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"requests": {Func: nil, Children: map[string]*bintree{
		"insert_storage_key.sql": {Func: bindataRequestsinsertstoragekeysql, Children: map[string]*bintree{}},
		"insert_user.sql":        {Func: bindataRequestsinsertusersql, Children: map[string]*bintree{}},
		"requests.go":            {Func: bindataRequestsrequestsgo, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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