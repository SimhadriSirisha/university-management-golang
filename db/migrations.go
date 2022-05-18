// Code generated by go-bindata. DO NOT EDIT.
// sources:
// db/migrations/000001_create_department_table.down.sql
// db/migrations/000001_create_department_table.up.sql
// db/migrations/000002_create_student_table.down.sql
// db/migrations/000002_create_student_table.up.sql
// db/migrations/000003_create_staff_table.down.sql
// db/migrations/000003_create_staff_table.up.sql
// db/migrations/000004_create_staff_department_relation_table.down.sql
// db/migrations/000004_create_staff_department_relation_table.up.sql
// db/migrations/000005_create_student_attendence_table.down.sql
// db/migrations/000005_create_student_attendence_table.up.sql

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
		modTime: time.Unix(1652689066, 0),
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

var _bindata000003createstafftableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x49\x4c\x4b\x2b\xb6\x06\x04\x00\x00\xff\xff\x0a\xc5\x80\x80\x1c\x00\x00\x00")

func bindata000003createstafftableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000003createstafftableDownSql,
		"000003_create_staff_table.down.sql",
	)
}



func bindata000003createstafftableDownSql() (*asset, error) {
	bytes, err := bindata000003createstafftableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000003_create_staff_table.down.sql",
		size: 28,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652689477, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000003createstafftableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x49\x4c\x4b\x2b\x56\xd0\xe0\x52\x80\x82\xcc\x14\x85\xcc\xbc\x92\xd4\xf4\xd4\x22\x85\x82\xa2\xcc\xdc\xc4\xa2\x4a\x85\xec\xd4\x4a\x05\x1d\xb8\x82\xbc\xc4\xdc\x54\x85\x92\xd4\x8a\x12\x2e\x4d\x6b\x40\x00\x00\x00\xff\xff\x3e\x7c\xa2\x29\x59\x00\x00\x00")

func bindata000003createstafftableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000003createstafftableUpSql,
		"000003_create_staff_table.up.sql",
	)
}



func bindata000003createstafftableUpSql() (*asset, error) {
	bytes, err := bindata000003createstafftableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000003_create_staff_table.up.sql",
		size: 89,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652695933, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000004createstaffdepartmentrelationtableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x49\x4c\x4b\x8b\x4f\x49\x2d\x48\x2c\x2a\xc9\x4d\xcd\x2b\xb1\x06\x04\x00\x00\xff\xff\x2b\x4b\x8b\xce\x26\x00\x00\x00")

func bindata000004createstaffdepartmentrelationtableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000004createstaffdepartmentrelationtableDownSql,
		"000004_create_staff_department_relation_table.down.sql",
	)
}



func bindata000004createstaffdepartmentrelationtableDownSql() (*asset, error) {
	bytes, err := bindata000004createstaffdepartmentrelationtableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000004_create_staff_department_relation_table.down.sql",
		size: 38,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652689477, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000004createstaffdepartmentrelationtableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xbf\x4e\x80\x30\x10\x87\x77\x9e\xe2\x37\x42\xe2\x1b\x38\x55\x53\x93\x46\xfc\x93\xd2\x41\x26\x42\xe8\x95\x34\xe0\x41\xda\x73\xf0\xed\x8d\x40\x24\x38\x18\x3b\xf6\xbe\xfb\xee\xbb\xb7\x5a\x39\x0d\xa7\xee\x6a\x0d\xf3\x80\xe7\x17\x07\xfd\x66\x1a\xd7\x20\x4b\x1f\x42\xe7\x69\xed\x93\xbc\x13\x0b\xca\x02\xc7\x8b\x1e\x8d\xb6\x46\xd5\x78\xb5\xe6\x49\xd9\x16\x8f\xba\xbd\xf9\x19\x6f\x9b\xc6\x23\xb2\xd0\x48\x09\xbc\x08\xf8\x63\x9e\x4f\xc2\xd3\x2a\x7f\x02\xc3\xc2\x59\x52\x1f\x59\x10\xa6\xee\x77\xca\xfe\x81\xb0\x24\x8a\x23\x63\xa2\x4f\x94\xc7\xd1\x0a\x89\x02\x25\xe2\x81\xf2\x1e\x92\xcb\xe8\xab\x7f\x9b\xbf\xcb\xae\xe2\xbd\xf5\xe2\x3d\xf9\x4d\x5e\x54\xb7\x5f\x01\x00\x00\xff\xff\x0f\xd0\xd2\xdb\x48\x01\x00\x00")

func bindata000004createstaffdepartmentrelationtableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000004createstaffdepartmentrelationtableUpSql,
		"000004_create_staff_department_relation_table.up.sql",
	)
}



func bindata000004createstaffdepartmentrelationtableUpSql() (*asset, error) {
	bytes, err := bindata000004createstaffdepartmentrelationtableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000004_create_staff_department_relation_table.up.sql",
		size: 328,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652697002, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000005createstudentattendencetableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x29\x4d\x49\xcd\x2b\x89\x4f\x2c\x29\x49\xcd\x4b\x49\xcc\x4b\x4e\xb5\x06\x04\x00\x00\xff\xff\x94\x6c\x79\xb9\x28\x00\x00\x00")

func bindata000005createstudentattendencetableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000005createstudentattendencetableDownSql,
		"000005_create_student_attendence_table.down.sql",
	)
}



func bindata000005createstudentattendencetableDownSql() (*asset, error) {
	bytes, err := bindata000005createstudentattendencetableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000005_create_student_attendence_table.down.sql",
		size: 40,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652782404, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindata000005createstudentattendencetableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\xb1\x0a\xc2\x30\x14\x85\xe1\xbd\x4f\x71\xc6\x16\xfa\x06\x4e\x51\x22\x04\x8a\x42\x93\xc1\xad\x04\x73\x29\xc1\xf4\x46\xe2\xed\xd0\xb7\x17\x12\xd0\xe5\xc0\xf9\x86\xff\x32\x6b\xe5\x34\x9c\x3a\x4f\x1a\xe6\x8a\xdb\xdd\x41\x3f\x8c\x75\x16\x1f\xd9\x03\xb1\x2c\x5e\x84\x38\x78\x7e\x12\xfa\x0e\x00\x62\x80\xd5\xb3\x51\xd3\x58\x6f\xc9\x29\x71\x46\x64\xa1\x95\x4a\xb3\xe0\x85\xea\xb4\x9b\xf2\x1a\x79\x91\xb8\x11\xea\xfc\x34\xef\xf2\xe7\xa6\xef\x12\x37\x5f\x0e\xbc\xe8\x40\xdf\xda\x63\x4d\x0d\xdd\x70\xfa\x06\x00\x00\xff\xff\x2d\x87\xb1\x8a\xaf\x00\x00\x00")

func bindata000005createstudentattendencetableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_bindata000005createstudentattendencetableUpSql,
		"000005_create_student_attendence_table.up.sql",
	)
}



func bindata000005createstudentattendencetableUpSql() (*asset, error) {
	bytes, err := bindata000005createstudentattendencetableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "000005_create_student_attendence_table.up.sql",
		size: 175,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1652870760, 0),
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
	"000001_create_department_table.down.sql":                bindata000001createdepartmenttableDownSql,
	"000001_create_department_table.up.sql":                  bindata000001createdepartmenttableUpSql,
	"000002_create_student_table.down.sql":                   bindata000002createstudenttableDownSql,
	"000002_create_student_table.up.sql":                     bindata000002createstudenttableUpSql,
	"000003_create_staff_table.down.sql":                     bindata000003createstafftableDownSql,
	"000003_create_staff_table.up.sql":                       bindata000003createstafftableUpSql,
	"000004_create_staff_department_relation_table.down.sql": bindata000004createstaffdepartmentrelationtableDownSql,
	"000004_create_staff_department_relation_table.up.sql":   bindata000004createstaffdepartmentrelationtableUpSql,
	"000005_create_student_attendence_table.down.sql":        bindata000005createstudentattendencetableDownSql,
	"000005_create_student_attendence_table.up.sql":          bindata000005createstudentattendencetableUpSql,
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
	"000003_create_staff_table.down.sql": {Func: bindata000003createstafftableDownSql, Children: map[string]*bintree{}},
	"000003_create_staff_table.up.sql": {Func: bindata000003createstafftableUpSql, Children: map[string]*bintree{}},
	"000004_create_staff_department_relation_table.down.sql": {Func: bindata000004createstaffdepartmentrelationtableDownSql, Children: map[string]*bintree{}},
	"000004_create_staff_department_relation_table.up.sql": {Func: bindata000004createstaffdepartmentrelationtableUpSql, Children: map[string]*bintree{}},
	"000005_create_student_attendence_table.down.sql": {Func: bindata000005createstudentattendencetableDownSql, Children: map[string]*bintree{}},
	"000005_create_student_attendence_table.up.sql": {Func: bindata000005createstudentattendencetableUpSql, Children: map[string]*bintree{}},
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
