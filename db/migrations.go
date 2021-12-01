// Code generated by go-bindata. DO NOT EDIT.
// sources:
// db/migrations/000001_create_department_table.down.sql
// db/migrations/000001_create_department_table.up.sql
// db/migrations/000002_create_student_table.down.sql
// db/migrations/000002_create_student_table.up.sql

package migrations


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

var _bindata000001createdepartmenttableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x49\x2d\x48\x2c\x2a\xc9\x4d\xcd\x2b\x29\xb6\x06\x04\x00\x00\xff\xff\x01\x7b\x09\x27\x21\x00\x00\x00")

func bindata000001createdepartmenttableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000001createdepartmenttableDownSql,
		"000001_create_department_table.down.sql",
	)
}



func bindata000001createdepartmenttableDownSql() (*asset, error) {
	bytes, err := bindata000001createdepartmenttableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000001_create_department_table.down.sql",
		size: 33,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652421680, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000001createdepartmenttableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x49\x2d\x48\x2c\x2a\xc9\x4d\xcd\x2b\x29\x56\xd0\xe0\x52\x50\x50\x50\xc8\x4c\x51\xc8\xcc\x2b\x49\x4d\x4f\x2d\x52\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\xd4\x01\xcb\xe5\x25\xe6\xa6\x2a\x94\xa4\x56\x94\x70\x69\x5a\x03\x02\x00\x00\xff\xff\x54\x0c\x06\x19\x55\x00\x00\x00")

func bindata000001createdepartmenttableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000001createdepartmenttableUpSql,
		"000001_create_department_table.up.sql",
	)
}



func bindata000001createdepartmenttableUpSql() (*asset, error) {
	bytes, err := bindata000001createdepartmenttableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000001_create_department_table.up.sql",
		size: 85,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652423379, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000002createstudenttableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x29\x4d\x49\xcd\x2b\x29\xb6\x06\x04\x00\x00\xff\xff\x62\xa9\x6d\xa5\x1e\x00\x00\x00")

func bindata000002createstudenttableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000002createstudenttableDownSql,
		"000002_create_student_table.down.sql",
	)
}



func bindata000002createstudenttableDownSql() (*asset, error) {
	bytes, err := bindata000002createstudenttableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000002_create_student_table.down.sql",
		size: 30,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652421949, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000002createstudenttableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8d\x41\x0a\xc2\x30\x14\x44\xf7\x3d\xc5\x2c\x5b\xf0\x06\xae\xaa\x44\x28\x88\x82\xcd\xc2\x9d\x04\x33\x4a\xb0\xf9\x2d\x3f\x5f\xb0\xb7\x97\x82\xa8\xb3\x9c\x99\xc7\xdb\x9e\x5c\xeb\x1d\x7c\xbb\xd9\x3b\x74\x3b\x1c\x8e\x1e\xee\xdc\xf5\xbe\x47\xb1\x67\xa4\x58\x41\x5d\xe1\x13\x1d\x87\x41\x46\x24\x31\xde\xa9\x98\x34\xe5\xa0\x33\x1e\x9c\xb1\xfa\x9e\x24\x64\xc2\xf8\xb2\x5f\x15\x39\xd9\xe5\x0f\x54\xde\xa8\x94\x2b\xcb\x32\x05\xb5\xbc\x88\xea\x14\x9b\xaa\x59\xbf\x03\x00\x00\xff\xff\x5a\xed\x7c\x17\x93\x00\x00\x00")

func bindata000002createstudenttableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000002createstudenttableUpSql,
		"000002_create_student_table.up.sql",
	)
}



func bindata000002createstudenttableUpSql() (*asset, error) {
	bytes, err := bindata000002createstudenttableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000002_create_student_table.up.sql",
		size: 147,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652423120, 0),
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
	"000001_create_department_table.down.sql": bindata000001createdepartmenttableDownSql,
	"000001_create_department_table.up.sql":   bindata000001createdepartmenttableUpSql,
	"000002_create_student_table.down.sql":    bindata000002createstudenttableDownSql,
	"000002_create_student_table.up.sql":      bindata000002createstudenttableUpSql,
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
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
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
	"000001_create_department_table.down.sql": {Func: bindata000001createdepartmenttableDownSql, Children: map[string]*bintree{}},
	"000001_create_department_table.up.sql": {Func: bindata000001createdepartmenttableUpSql, Children: map[string]*bintree{}},
	"000002_create_student_table.down.sql": {Func: bindata000002createstudenttableDownSql, Children: map[string]*bintree{}},
	"000002_create_student_table.up.sql": {Func: bindata000002createstudenttableUpSql, Children: map[string]*bintree{}},
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
