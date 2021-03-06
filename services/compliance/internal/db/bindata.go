// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/01_init.sql
// migrations/02_auth_data.sql
// migrations/03_table_names.sql
// DO NOT EDIT!

package db

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5b\x73\xda\xc8\x12\x7e\x3e\xfc\x8a\x7e\xc3\xae\x23\x38\xe0\x18\x27\x81\xf2\x83\x02\xf2\x09\xb5\x20\x1c\x24\x36\x76\xd5\x56\x4d\x0d\xd2\x48\x4c\x59\x37\x8f\x86\x38\xe4\xd7\x6f\x8d\x04\xba\x21\x81\xbc\x68\xf3\x88\xa6\xa7\xfb\xeb\xab\xbe\x16\x9d\x4e\xab\xd3\x81\x47\x3f\xe4\x36\x23\xda\xb7\x19\x98\x98\xe3\x35\x0e\x09\x98\x5b\x37\x68\x75\x3a\x2d\x71\x3e\xd9\xba\x01\x31\xc1\x62\xbe\x9b\x0a\xfc\x20\x2c\xa4\xbe\x07\x9f\xbb\x77\xdd\x7e\x46\x6a\xbd\x83\xc0\x46\xe2\x7a\x41\xa4\xa5\x29\x3a\x84\x1c\x73\xe2\x12\x8f\x23\x4e\x5d\xe2\x6f\x39\xdc\x43\x6f\x14\x1d\x39\xbe\xf1\x72\xfc\x94\x9a\x0e\x41\xd4\x43\x9c\x61\x2f\xc4\x06\xa7\xbe\x87\x42\x12\x0a\xbd\xc7\xc2\x86\x43\x85\x6a\xe2\x19\xbe\x49\x3d\x1b\xee\xa1\xbd\xd2\x1f\x3e\xb5\x47\x07\xdb\x9e\x89\x99\x89\x0c\xdf\xb3\x7c\xe6\x52\xcf\x46\x21\x67\xd4\xb3\x43\xb8\x07\xdf\xdb\xeb\xd8\x10\xe3\x05\x59\x5b\x2f\xb6\xb5\xf6\x4d\x4a\xc4\xb9\x85\x9d\x90\xe4\xcc\xb8\xd4\x43\x2e\x09\x43\x6c\x47\x02\x6f\x98\x79\xd4\xb3\x63\x11\xe6\xbf\xa1\x90\x18\x5b\x46\xf9\x4e\x28\xb7\xac\x91\x08\xa5\x88\x93\x8a\x5d\x32\x84\xc0\x09\xec\xf0\xd5\x19\x81\xbe\x0b\xc8\x10\x94\x27\x5d\x51\xb5\xe9\x42\x1d\x81\x66\x6c\x88\x8b\x87\xd0\x19\xc1\xe2\xcd\x23\x6c\x08\x51\x1e\xc6\x4b\x45\xd6\x95\x54\x10\xa6\x0f\xa0\x2e\x74\x50\x9e\xa6\x9a\xae\x1d\xf4\xc1\xf7\xa9\xfe\x15\xb4\xf1\x57\x65\x2e\x8b\x3c\x18\x98\x63\xc7\xb7\x47\xad\xbc\xf5\x54\x4b\x01\xc7\x78\x31\x9f\x2b\xaa\x5e\x8d\x22\x3e\x87\x85\x7a\xac\x03\xa6\x1a\xb4\x1f\x67\xff\x0b\x6c\x51\x49\x01\xf3\x0d\x62\x6e\x19\x76\xc0\xc1\x9e\xbd\xc5\x36\x69\x0b\x18\x51\x26\x08\x66\xc6\x06\x05\x98\x6f\xe0\x1e\x82\xed\xda\xa1\x86\x94\x87\x2b\xc4\x4c\x62\xe1\xad\xc3\x11\xc7\x6b\x87\x84\x01\x36\x88\xc8\x68\xbb\x70\xfa\x46\xf9\x06\xf9\xd4\xcc\x24\x29\xe7\x2b\x76\x1c\xff\x8d\x98\xc8\xa2\x07\x27\x75\xf9\xcb\x4c\x49\x5d\x8c\xed\x27\x7e\x32\xdf\xe7\xd9\x88\x47\xd2\x19\x2d\x70\xd5\x02\x00\xa0\x26\xac\xa9\x4d\x3d\x1e\x65\x41\x5d\xcd\x66\x52\xf4\xdc\xc3\x2e\x01\x63\x83\x19\x36\x38\x61\xf0\x03\xb3\x1d\xf5\xec\xab\x9b\xc1\xe0\xba\x20\x69\xfa\x2e\xa6\x5e\x3d\xd9\x18\x23\x7a\x21\xbb\x54\xfe\x6a\x70\x57\x14\x3b\xa0\xc4\x1c\x44\x73\x84\x1c\xbb\x01\x88\x00\x89\x36\x11\x4f\xe0\x97\xef\x91\xe4\x52\xeb\x5a\x64\x44\x9e\xe9\xca\xf2\xd8\xcd\xc5\x77\x55\x3c\x5e\x44\x01\xa9\x88\xe9\x36\x24\xac\x89\xa8\x0a\x3d\x67\xe2\x6a\x51\x54\x3f\xb4\x16\x45\xef\x89\xae\x45\x51\xcd\x00\x0b\xa0\x88\x9a\xf5\xd4\x36\x96\x8d\x28\x3c\x35\xf2\x61\x51\x44\x4d\x14\x92\xd7\x43\x4e\x34\xe5\xdb\x4a\x51\xc7\xf5\xd3\x72\xb8\x70\xa4\x31\x72\x48\xd3\xe5\xa5\x1e\x8f\x98\x7e\xf4\x60\xaa\x8e\x97\x4a\x34\x10\xbe\x3c\xef\x1f\xa9\x0b\x98\x4f\xd5\x3f\xe5\xd9\x4a\x49\x7e\xcb\x4f\xe9\xef\xb1\x3c\xfe\xaa\x40\xbf\xc2\xd5\xc4\xde\x45\xee\x46\x97\x27\xf0\xe5\xf9\xbc\xdf\x31\x86\x4a\xb7\x13\x4d\x99\xce\xe8\x52\xb3\x38\x4f\xf7\x87\xfb\xea\x68\x34\x03\x19\x9d\xbf\x29\x07\x19\x8b\x75\xb2\x70\xc2\xe9\x8b\xf3\x50\x84\x92\xcb\x84\x38\x2c\xcb\xc5\x96\x6f\x90\xa0\x29\x17\x0d\xa6\x83\x92\x33\x53\x89\x91\xd7\x2d\x09\x79\xed\x91\xf0\x9e\xa9\x94\x62\xe0\xe4\x27\x3f\x35\x26\x12\xc1\x93\xe9\xda\xf2\x8d\x10\x6a\xae\x40\xf3\x0a\xff\xf5\xea\xcc\x9b\xbb\xc4\xd7\x0b\xea\xb2\x04\x44\x5c\x94\x87\x1c\x54\x54\xa4\xcf\xe8\x2f\x62\x66\x69\xec\xa5\xe5\x79\xac\xf1\x4c\xad\x66\x29\x74\x69\xbd\xde\xdd\x16\x4b\xd0\x25\xae\x5f\x4b\x30\xab\xfb\xa7\xc9\xf2\x15\x9b\x96\xf3\x1e\x73\xcd\x17\xe2\xbe\x65\x6a\x76\x40\x49\x3c\xce\x95\x48\x7c\x25\x1f\x97\xe6\x7a\xa3\x4a\xfb\xef\x68\x94\x2a\xdb\x8d\x85\xe4\xb2\x16\x3a\x09\x2f\xe9\xa7\xe3\x8c\x96\x34\x97\xed\xb3\x00\xb9\xd4\x66\x58\x08\x84\x97\x74\x55\x41\x55\xda\x4e\x65\xd5\x1c\x04\x0e\x2d\x2b\xe5\xb4\x8e\x8f\xcb\xb4\x68\xa0\x0e\xc5\xb6\x28\x50\xf3\xe0\xd5\x44\x79\x90\x57\x33\xbd\x6e\xcc\x63\xab\x0b\x75\x96\xa5\x2f\x10\x9f\x8d\x17\xb3\xd5\x5c\x15\xde\x89\x65\x6a\xaf\x18\x3c\xf2\x93\xff\xc0\xce\x55\xbb\x48\x84\xda\xc3\x21\x23\xb6\xe1\xe0\x30\xbc\xae\x60\x3f\x31\x4d\x6d\x12\x6c\xa4\xf0\x3d\x70\x33\x7c\xe1\x24\xe0\xe4\x85\xd9\x08\xda\x44\x5b\x3d\xa8\xf9\x57\xc8\x39\x9c\x25\x63\xad\x29\xd0\x25\xaa\x6b\x7b\x50\xd5\xc1\xe5\xee\x4c\x44\x74\x2c\x9f\x9d\x5e\xc9\x61\x22\xeb\x72\x8d\x9e\x5d\x3c\xe6\xea\xf9\x8a\x9a\x52\xb4\x76\x4b\x7b\x7a\x25\x65\xd6\x65\x29\xb3\x85\x5d\xc3\xc3\x72\x31\x87\x90\x9b\xd4\x1b\xb5\xfe\xea\x9e\x83\x77\xb4\xdd\xfe\x13\x80\xf1\x6a\x2b\x20\xee\x37\x58\x29\xdd\x4e\xa5\xfc\xe6\x29\x1d\xd6\xcb\x5a\xa0\x6b\xae\x42\x9a\x52\xa3\x3c\x34\x65\xa6\x8c\xf5\xcc\x17\x98\x6e\x48\x2a\xe6\x80\x04\x7d\x29\xfe\xce\x52\x35\x07\x4e\x2d\x04\xcd\xa0\xc9\xb6\x79\x29\x9e\x62\x42\x4b\x57\x82\x77\x65\x33\xdd\x07\x44\x7a\x52\xda\x9f\xd6\x5c\x22\x72\x2e\x63\x67\xb8\xe9\xa5\x21\x2a\x8c\x97\xba\xf1\x39\x43\x50\xdf\x1b\xac\x32\x76\x2a\xa2\x95\x9f\x18\x52\xc4\x30\xa5\x22\x7d\x94\xf2\x64\x51\x82\x9a\x61\xad\x4d\x5e\x1a\x88\x71\xe5\x00\xac\x13\xf0\x53\xa4\xe5\x3d\x91\x3e\x62\x2c\xd1\xec\x48\xa8\x49\x3e\x64\xbd\x3e\xa2\x1e\xe5\xdd\xf0\xd5\xf9\xcf\x4d\xaf\xff\xa9\xd3\xbb\xeb\xdc\x7c\x86\xfe\x60\xd8\xbf\x19\x0e\x06\xdd\xc1\xe7\xdb\xc1\xc7\x4f\xff\xed\xdd\xb4\x7a\x37\x28\xdd\x67\xaa\xc4\xef\x7a\xfd\x8f\x83\x0f\x91\xf8\x87\xf8\x13\x6d\x34\xde\xc2\x13\x17\x6e\x3f\x0e\x6e\xc5\x85\xaa\x29\x26\xa6\x79\x3a\x70\x82\x17\xb2\x4b\xbf\x4b\xab\x9a\xbe\x94\xa7\xea\x65\xdc\x27\xa2\x6e\xf2\x64\x92\xd1\x57\x30\x08\x8f\xcb\xe9\x5c\x5e\x3e\xc3\x1f\xca\xb3\x08\xe7\x69\xb2\x93\x9d\x47\x8d\xc3\x15\x5a\x4f\x00\x4e\x8c\x9e\x87\x9c\x0c\xaf\x64\x38\x34\x07\xf6\xa0\xbb\x14\x69\xd6\x5c\x2d\x98\x25\x63\xa3\xbc\xd9\x1a\x75\xe0\xd8\x6a\x95\x37\xe5\x40\xce\xba\x56\xec\xd3\xc2\xef\xe6\xbc\x29\x28\x2e\x73\xa3\xcc\xf6\xf9\xdc\x58\x14\xad\x77\x7b\xda\x72\x40\x3a\x55\x27\xca\x53\xed\x1d\x6b\xa5\x4e\xbf\xad\x94\xf8\x52\x5e\x1f\x2c\xd4\xec\x04\x58\x69\x53\xf5\xff\xb0\xe6\x8c\x10\xb8\x8a\x25\xaa\xf0\xa4\xdc\xa9\x39\x4c\x99\x7f\x02\x4e\xe0\x4a\xa5\x4a\xea\x58\xa8\xc9\x31\x3b\xb4\x67\x2c\x4d\xa0\xac\xd6\x9e\xc5\x1b\x8d\xa7\x1c\xe2\x72\xaa\x79\x84\x3e\xa5\x36\x0d\x80\xcd\x7c\x1e\x15\xd8\x92\x39\x94\x03\x96\x0a\x65\xc0\x54\xfd\x27\x0d\x86\xef\x06\x0e\xe1\x24\xb2\xf9\x77\x00\x00\x00\xff\xff\x3c\x3d\xcd\xce\xc0\x1e\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 7872, mode: os.FileMode(420), modTime: time.Unix(1530278032, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x41\x6f\x82\x40\x10\x85\xef\xfb\x2b\xe6\x28\xa9\x5e\x9a\xea\x85\x13\xad\x34\x21\xb5\x68\x09\x24\xf5\xb4\x19\xdd\x45\x27\x65\xc1\x2c\x4b\xd5\xfe\xfa\x86\x5a\x85\xad\xa2\xe9\x75\xdf\xdb\x99\xf7\x3e\xd8\xc1\x00\xee\x14\xad\x34\x1a\x09\xc9\x86\x3d\x45\xbe\x17\xfb\x10\x7b\x8f\x13\x1f\xbc\xca\xac\x0b\x4d\x5f\x52\xc4\x1a\xf3\x12\x97\x86\x8a\x1c\x7a\x0c\x80\x04\x2c\x68\x55\x4a\x4d\x98\xf5\x19\x80\x69\x74\x4e\x02\x3e\x51\x2f\xd7\xa8\x7b\xa3\x07\x07\xc2\x69\x0c\x61\x32\x99\xd4\x36\x25\x55\xd1\x29\xb6\x67\xec\x84\x06\x23\x77\xc6\x32\xe0\x29\x0e\x47\x03\x86\x94\x2c\x0d\xaa\x8d\xe5\x11\x68\xf0\xfc\x26\x03\x98\x45\xc1\xab\x17\xcd\xe1\xc5\x9f\x43\x8f\x84\xc3\x1c\x97\xfd\x69\x9b\x65\xc5\x56\x8a\xe7\xe0\x62\xc3\x1c\x95\x3c\x45\xbf\x1f\x0e\xed\xec\xa2\x50\x48\x79\xb7\xbe\xa9\x16\x19\x2d\xf9\x87\xdc\xc3\x8f\x61\x38\xb2\x75\x3c\xec\xee\xee\x75\x16\x9f\x39\xd0\x14\x48\xc2\xe0\x2d\xf1\x21\x08\xc7\xfe\x3b\x60\x4a\x7c\xb1\xe7\xbf\x91\xa6\x61\xbb\xd8\xe1\xd0\x71\xaf\x5d\x6c\x65\xb5\x2f\x37\x42\x17\xbb\xa4\x94\xfa\x22\xbd\x94\xf8\x75\x80\x29\xf1\x5b\x0c\x53\xe2\xb7\x30\x56\xa5\xd4\xed\xff\xef\x6c\xc6\xff\x39\x3b\x5d\x94\xab\x9a\x95\x95\x89\x1f\xd7\x37\xd8\x0e\x40\x2c\x57\xff\x98\xb2\x9e\xcc\xda\xcf\x6f\x5c\x6c\x73\x36\x8e\xa6\xb3\x6b\xcf\xcf\xb5\x1c\xc7\x8f\x73\xe9\xb4\xde\xed\xb2\xef\x00\x00\x00\xff\xff\x02\xc5\x23\x8a\xe0\x03\x00\x00")

func migrations01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initSql,
		"migrations/01_init.sql",
	)
}

func migrations01_initSql() (*asset, error) {
	bytes, err := migrations01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_init.sql", size: 992, mode: os.FileMode(420), modTime: time.Unix(1530096833, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations02_auth_dataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xb1\x4f\x85\x30\x10\x87\xf7\xfb\x2b\x7e\x23\x44\xdf\x62\xf2\x26\xa6\x6a\x3b\x10\xb1\x60\x43\x13\x99\xc8\x69\x1b\x68\x22\xa0\xa5\xa8\x7f\xbe\x61\x30\xa2\xc9\x5b\xef\xbb\xcb\x7d\xdf\xe9\x84\xab\x29\x0c\x91\x93\x87\x7d\xa3\x3b\xa3\x44\xab\xd0\x8a\xdb\x4a\x41\x6c\x69\x94\x9c\x18\x19\x01\xc1\xe1\x39\x0c\xab\x8f\x81\x5f\xaf\x09\x88\xfe\x7d\xf3\x6b\xea\x83\xc3\x07\xc7\x97\x91\x63\x76\x73\x3e\xe7\xd0\x75\x0b\x6d\xab\x6a\xdf\x71\xcb\xc4\x61\xbe\xcc\x79\x4b\x63\xef\xf6\x0f\xc9\x7f\xa5\x3f\x88\x80\xc6\x94\x0f\xc2\x74\xb8\x57\x1d\xb2\xe0\x72\xca\x0b\xfa\x11\xb4\xba\x7c\xb4\x0a\xa5\x96\xea\xe9\xa8\x52\xeb\x83\xf5\xef\x7c\xbf\x3c\x96\xca\xe5\x73\x26\x69\xea\xe6\x5f\x69\x41\xdf\x01\x00\x00\xff\xff\x0e\x07\x5f\x45\x10\x01\x00\x00")

func migrations02_auth_dataSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations02_auth_dataSql,
		"migrations/02_auth_data.sql",
	)
}

func migrations02_auth_dataSql() (*asset, error) {
	bytes, err := migrations02_auth_dataSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/02_auth_data.sql", size: 272, mode: os.FileMode(420), modTime: time.Unix(1530277947, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations03_table_namesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xd1\x0a\xc2\x20\x14\x86\xef\xf7\x14\xe7\x3e\xf6\x04\x5e\x19\x33\x08\x56\xc1\x70\xd7\xe3\x90\x56\xc2\xd2\xd0\x23\x83\x9e\x3e\x2a\x6a\xba\x79\xab\xff\xff\x9f\xef\xab\x6b\xd8\xdc\xcd\xd5\x23\x69\xe8\x1f\x15\x6f\xa5\xe8\x40\xf2\x6d\x2b\x80\x47\xba\x39\x6f\x9e\x5a\x49\x8f\x36\xe0\x99\x8c\xb3\xd0\x89\x23\x3f\x08\x90\x27\xc0\xff\xff\x40\x73\x80\xe5\x1b\xe3\xe8\x26\xad\x76\xfb\xb4\xf7\x7d\x1b\x2e\xa6\x98\xed\x83\xf6\x85\x74\x0c\xda\xb3\x15\x5f\x83\x84\x0b\xa4\x41\x21\x21\xab\xaa\xd4\xac\x71\x93\xcd\xba\x65\xf6\x64\xa9\x28\x9f\xdf\x9f\x3d\xd2\xde\x4f\xb8\x9c\x8d\xb9\x5c\xa2\xcc\x56\x7c\x1f\x91\x05\xd2\xdb\x97\xbd\x02\x00\x00\xff\xff\x6b\xd7\x46\x9f\xb4\x01\x00\x00")

func migrations03_table_namesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations03_table_namesSql,
		"migrations/03_table_names.sql",
	)
}

func migrations03_table_namesSql() (*asset, error) {
	bytes, err := migrations03_table_namesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/03_table_names.sql", size: 436, mode: os.FileMode(420), modTime: time.Unix(1530274484, 0)}
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
	"latest.sql":                    latestSql,
	"migrations/01_init.sql":        migrations01_initSql,
	"migrations/02_auth_data.sql":   migrations02_auth_dataSql,
	"migrations/03_table_names.sql": migrations03_table_namesSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_init.sql":        &bintree{migrations01_initSql, map[string]*bintree{}},
		"02_auth_data.sql":   &bintree{migrations02_auth_dataSql, map[string]*bintree{}},
		"03_table_names.sql": &bintree{migrations03_table_namesSql, map[string]*bintree{}},
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
