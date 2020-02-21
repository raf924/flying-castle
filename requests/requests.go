// Code generated by go-bindata. DO NOT EDIT.
// sources:
// requests\find_by_name_in_parent.sql
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

var _bindataRequestsfindbynameinparentsql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x76\xf5\x71\x75\x0e\x51\xd0\x52\x48\x2b\xca\xcf\x55\x28\x48\x2c\xc9" +
		"\x50\x28\x50\x28\xcf\x48\x2d\x4a\x55\x28\xd0\x2b\x48\x2c\x4a\xcd\x2b\x89\xcf\x4c\x51\xb0\x55\xb0\x57\x48\xcc\x4b" +
		"\x51\x28\xd0\xcb\x4b\xcc\x4d\x05\x71\xad\x01\x01\x00\x00\xff\xff\x82\xbf\x8d\xc0\x3a\x00\x00\x00")

func bindataRequestsfindbynameinparentsqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataRequestsfindbynameinparentsql,
		"requests/find_by_name_in_parent.sql",
	)
}

func bindataRequestsfindbynameinparentsql() (*asset, error) {
	bytes, err := bindataRequestsfindbynameinparentsqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "requests/find_by_name_in_parent.sql",
		size:        58,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1581016883, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
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
		modTime:     time.Unix(1580722827, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataRequestsinsertusersql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\x50\x4a\x2f\xca\x2f\x2d" +
		"\x50\x52\xd0\xc8\x4c\xd1\x51\xc8\x4b\xcc\x4d\xd5\x54\x08\x73\xf4\x09\x75\x0d\x56\xd0\xb0\xd7\x51\xb0\xd7\xb4\xe6" +
		"\xe5\x42\x56\x5e\x5a\x9c\x5a\x04\x51\x0b\x62\xc5\x83\x34\xe8\x28\x64\x24\x16\x67\xe8\x28\x14\x27\xe6\x94\xe8\x28" +
		"\x14\xe5\xe7\x97\xc4\xa7\xe5\xe7\xa4\xa4\x16\xc5\x83\x94\xe5\x26\x66\xe6\xc5\x83\xed\x88\xcf\x4c\xd1\xe4\xe5\x42" +
		"\x31\x1d\x19\x69\x5a\x03\x02\x00\x00\xff\xff\x8a\xc2\x3f\xb7\x97\x00\x00\x00")

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
		size:        151,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1580989631, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataRequestsrequestsgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x5b\x6f\xdb\x38\x16\x7e\x16\x7f\x05\xd7\x40\x0b\x69\xd7\x6b\x4b" +
		"\xb2\x24\xcb\x06\xf2\xd2\x24\x0b\xf4\xa1\x2d\xb0\xd3\x79\x9a\x33\x30\x74\x21\x5d\x22\xb6\xe4\x48\x72\xe7\x38\x45" +
		"\xfe\xfb\xe0\x90\x94\xad\xb4\xce\x65\x3a\xed\x04\x50\x44\x91\x3c\xdf\xb9\x90\xdf\x47\x7a\x3a\xe5\x97\x75\x29\xf8" +
		"\x5a\x54\xa2\xc9\x3a\x51\xf2\xfc\xc0\xd7\xf5\x7f\x73\x55\x95\x59\x97\x4d\xf8\xd5\x07\xfe\xfe\xc3\x47\x7e\x7d\xf5" +
		"\xf6\xe3\x84\x4d\xa7\xbc\xad\xf7\x4d\x21\xda\x25\xb5\x1b\x71\xbb\x17\x6d\xd7\x82\x54\x55\xb9\xca\x0f\xab\x2a\xdb" +
		"\x8a\x95\xaa\x56\xbb\xac\x11\x55\x37\x69\x6f\x37\x0f\xa6\xa9\xaa\x15\x4d\xb7\x6a\xbb\xba\xc9\xd6\x62\x75\x23\x0e" +
		"\x8f\x4d\xd9\xb7\xa2\xf9\x66\xac\x6f\x4c\xd6\x35\x63\xbb\xac\xb8\xc9\xd6\xe2\x38\xca\x18\x53\xdb\x5d\xdd\x74\xdc" +
		"\x65\xce\x28\x3f\x74\xa2\x1d\x31\x67\x54\xd4\xdb\x5d\x23\xda\x76\xba\xbe\x53\x3b\xea\x90\xdb\x8e\x5e\xaa\x36\xff" +
		"\xa7\xaa\xde\x77\x6a\x43\x1f\xb5\x36\xd8\x65\xdd\xa7\xa9\x54\x1b\x41\x0d\xea\x68\xbb\x46\x55\x6b\x3d\xd6\xa9\xad" +
		"\x18\x31\x8f\x31\xb9\xaf\x0a\x6e\x4b\xf4\x7f\x91\x95\x2e\x35\xf8\x6f\xbf\x93\xdb\x31\xa7\x2a\x70\x63\xe6\x71\xb7" +
		"\xef\x15\x4d\x53\x37\x1e\xff\xc2\x9c\xf5\x9d\xfe\xe2\xcb\x0b\x4e\x51\x4d\xde\x8b\x3f\x08\x44\x34\xae\x0e\x9b\xbe" +
		"\xdf\xec\xa5\x14\x8d\x86\xf5\x3c\xe6\x28\xa9\x0d\xfe\x75\xc1\x2b\xb5\x21\x08\xa7\x11\xdd\xbe\xa9\xe8\x73\xcc\xe5" +
		"\xb6\x9b\x5c\x13\xba\x74\x47\x04\xc4\x5f\xdd\x2e\xf9\xab\xcf\x23\x13\x89\xf6\xe5\x31\xe7\x9e\x31\xe7\x73\xd6\xf0" +
		"\x7c\x2f\xb9\xf1\x63\x9c\x30\x67\x65\xc2\xb9\xe0\xaa\x9e\x5c\xd6\xbb\x83\xfb\x3a\xdf\xcb\x31\x5f\xdf\x79\xcc\x29" +
		"\x36\xd7\x7d\xa4\x93\xcb\x4d\xdd\x0a\xd7\x63\x3f\x2a\x1e\x82\x31\xf8\x8f\x00\x89\xa6\x31\x71\xdb\xce\x7c\x2f\x27" +
		"\x6f\x28\x74\xd7\x1b\xd3\x0c\x76\xcf\x18\xeb\x0e\x3b\xc1\xb3\xb6\x15\x1d\xd5\x7c\x5f\x74\x04\xa3\x13\xb4\x0b\xc2" +
		"\x1c\x55\xc9\x9a\x73\x5a\xd4\xb7\x95\xac\xaf\x91\xec\xb4\xd9\xa9\x8b\xab\xaa\x13\x8d\xcc\x0a\x41\xe6\x75\x3b\xf9" +
		"\x9f\x1d\x62\xce\xbb\xab\xf8\xf2\x93\x28\x6e\xda\xfd\xd6\xf5\xec\xba\x1e\x11\xec\x26\xe8\x67\x0f\x42\xd0\xbb\xc0" +
		"\xfe\x59\x23\xa7\x55\x77\xc7\x3e\x55\x75\x49\xc4\x9c\x2d\xb1\xcf\xfe\x59\xb7\xef\xea\x52\xe8\x81\x8f\xca\x42\xd0" +
		"\xc6\x9b\xd0\x17\x73\xb6\x65\x5c\xd8\x68\x06\xb1\xe8\x0d\xe9\x4a\xf5\x75\x3c\x1e\x7f\x9f\x6d\xc5\x31\x6c\x8a\xcb" +
		"\xd6\x52\xaa\x09\x45\xc8\xee\x9f\xb0\xfd\x45\xdd\x91\xad\x8e\xf4\xa1\x29\x25\xf2\xa4\x29\xe5\xe0\x7a\xc3\x8c\x1e" +
		"\x02\x50\xda\xcf\x01\x50\xc2\xae\x77\x4a\xfe\x1b\x04\x5d\x91\x27\x41\xce\x2c\xdd\x57\x28\xa7\x72\x3e\x89\xf4\xb6" +
		"\xbd\x52\x8d\xeb\xf1\xbc\xae\x37\x43\x84\x6c\xd3\x3e\x53\xc3\x43\x6b\x4a\x68\x76\xd7\x97\xfb\x81\xb5\xdd\xc2\xc4" +
		"\xca\xd5\x51\x4d\x8c\x9c\x91\xa2\xe6\x07\x5a\x21\x55\x19\x35\x6d\x6f\x37\xfc\xc2\x6e\x69\x92\x38\xc0\x40\x02\xa6" +
		"\x39\xa0\x9f\x02\xfa\xfe\xf9\x47\x4a\x40\x3f\x03\x9c\x27\x80\x32\x06\x9c\x07\x80\xf3\x18\xd0\x17\x80\x71\x00\x58" +
		"\xfa\x80\x71\x08\x18\xa5\x80\x61\x0e\x58\x64\x80\x85\x04\x8c\x63\xc0\x30\xb5\xfd\x05\x60\xb1\x18\xf1\xff\x68\xaf" +
		"\xb1\x6f\x46\xfa\x37\xcd\xd6\xb3\x4a\xc0\x28\x3b\x59\x12\x32\x21\xf6\x08\x34\x56\x94\xa6\x2f\x5d\x58\xbb\xc2\x44" +
		"\x91\xfb\xc6\x4e\xbf\xe7\xc6\xa6\x20\x9b\xfc\xe8\x35\x38\xa1\x16\x84\x9a\xdb\x19\x25\xa0\x6f\xf3\xca\xa8\x1d\xd8" +
		"\x67\x90\x3f\x3d\x69\x08\x98\xd3\xbb\x04\x2c\x7c\xc0\x59\x36\xac\xd3\xe8\x1b\x4d\x7f\x62\x15\xac\x02\x9d\x53\xf6" +
		"\x5e\xa7\x06\x27\x03\x73\x9c\x17\xad\xed\x98\x39\xce\xa8\x3f\xcc\xa6\x8f\x1f\xa8\xa3\x31\x73\x3c\x2d\x7c\x2f\x8f" +
		"\x98\x82\xfd\xb7\x96\xc8\x61\xb0\x5a\x23\x8f\x27\xd1\x5f\xc8\xfc\x39\xfd\x3f\xca\xb6\xd6\xdd\x13\x76\x4f\x0a\x9a" +
		"\x4f\xc8\x4b\xfe\xd2\x84\xb5\x70\x2e\x79\x9c\x52\x7b\x40\xd9\x25\x1f\xe9\x61\x52\x93\xe5\x50\x6c\xdc\x68\x96\x7a" +
		"\x76\x84\x54\x62\x69\x54\xe4\xd7\x4a\xa1\x1b\xc4\x69\xe0\x07\x49\x9a\xce\xc6\xdc\xa7\x49\x14\x6a\x46\x71\xbe\xd6" +
		"\x35\xfa\xa2\x0b\xb3\xe4\xb6\x3e\x94\xc4\x52\xff\x1f\x9c\x44\xd9\xf8\x29\xf2\x9a\x4b\x8c\xbd\xe6\xdc\x88\xc3\xf7" +
		"\x53\x57\x86\x80\x32\x02\xf4\x73\x43\x61\xbf\x34\x74\x26\x3a\x48\xff\x34\x26\x03\x43\x1c\xa2\x48\x28\x88\xae\x80" +
		"\xa1\x34\xa4\x8b\x8a\x9e\x44\x51\x4f\x00\x1a\x11\x86\x30\x65\x04\x58\xa4\x86\x8a\x65\x60\x90\xf5\xf7\x60\x9c\xfa" +
		"\x22\x8b\xac\x11\x4b\x4b\xe9\x85\x69\x13\x9d\x09\x3b\xec\x31\x28\xaa\xb4\xf7\x3a\x9f\xd9\x38\x17\x56\x7a\x4a\xc0" +
		"\x38\x31\x54\x26\xc2\x97\x73\xe3\x81\x72\x0e\x03\x83\x98\x10\xc5\xcf\xd5\x43\x02\x26\x09\xe0\x22\x00\x5c\xc4\x80" +
		"\x8b\x10\x70\x26\x1e\x50\x59\x7b\x7d\x8c\xd0\x67\x56\xe6\x6f\xd3\xf9\x0c\xe6\x43\x32\x9f\xbf\xf6\x3e\x43\xe4\x33" +
		"\xa8\xdf\x41\xe3\xc7\xf3\xfd\x19\x24\x7e\x3c\x51\x4b\xe0\x24\xfc\x41\x04\xf6\xe7\x61\x98\x86\xf3\x9f\x4b\x60\xfa" +
		"\x11\xf2\xcf\x51\x97\xce\x53\x22\x17\xd1\xb6\xb0\xef\xb0\x1c\x9e\xba\x74\x4e\xeb\xf3\xef\x1c\x5d\x07\xe7\x61\x19" +
		"\xf7\x14\x7c\x39\xf9\xfa\x76\x1e\x01\x8a\xa4\xf7\x2a\x62\xc0\x28\x34\x16\x31\xdd\x18\x32\xc0\x45\x61\xde\xbe\x15" +
		"\x06\xca\x21\x09\x01\x8b\x18\x30\x9d\x01\xce\x08\x21\x35\x52\x94\x44\x80\x61\x04\x18\x24\x80\xc9\xfc\xd4\x1f\x50" +
		"\x3f\x7d\x13\x95\x8f\xf3\x7b\xaf\x34\x4a\x9e\xc5\x1c\x70\x31\x07\x2c\x22\xc0\x6c\x7e\xea\xcb\x22\xf3\x10\x6a\xef" +
		"\x75\x61\x6d\xc2\xc4\x08\x84\x18\x8c\x09\x92\xa8\xf4\x74\xe3\xa0\xba\x89\xa8\xcf\xae\xf7\x3a\x0b\x00\x83\x12\x30" +
		"\x58\x00\x26\x0b\x9b\xe3\x0c\xd0\x0f\xcf\xdc\x26\xe8\x2e\x43\xd2\x23\x01\x73\x1b\xe5\xf3\xb7\x89\x07\xbb\xea\x07" +
		"\xc9\x8e\x45\x3b\x2b\x38\xfd\x8f\xe8\xa7\x94\xe6\xcf\x00\x00\x00\xff\xff\xec\xc4\x25\xc6\x00\x10\x00\x00")

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
		size:        8192,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1581076940, 0),
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
	"requests/find_by_name_in_parent.sql": bindataRequestsfindbynameinparentsql,
	"requests/insert_storage_key.sql":     bindataRequestsinsertstoragekeysql,
	"requests/insert_user.sql":            bindataRequestsinsertusersql,
	"requests/requests.go":                bindataRequestsrequestsgo,
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
		"find_by_name_in_parent.sql": {Func: bindataRequestsfindbynameinparentsql, Children: map[string]*bintree{}},
		"insert_storage_key.sql":     {Func: bindataRequestsinsertstoragekeysql, Children: map[string]*bintree{}},
		"insert_user.sql":            {Func: bindataRequestsinsertusersql, Children: map[string]*bintree{}},
		"requests.go":                {Func: bindataRequestsrequestsgo, Children: map[string]*bintree{}},
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
