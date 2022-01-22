package configs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

var _configs_config_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func configs_config_go() ([]byte, error) {
	return bindata_read(
		_configs_config_go,
		"configs/config.go",
	)
}

var _configs_config_yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\x4f\x6f\xdb\x30\x0c\xc5\xef\xfa\x14\x04\x76\xae\x63\x3b\xad\x9b\xf1\xd6\x3f\x19\x96\x61\xd9\x8c\xc5\x45\x8f\x03\x13\x33\x8a\x07\xd9\xd2\x24\x2a\x75\xf7\xe9\x07\x39\x45\x93\x61\x37\xe1\xc7\x87\x27\x3e\xbe\x0d\xfb\x23\x7b\x54\x00\x3f\xe2\xb0\xb6\x2d\x23\xb4\xbc\x8d\x5a\x01\x7c\x16\x71\xb5\xf5\x82\xb0\xc8\xf3\x3c\x29\x98\xda\xa6\xeb\xd9\x46\x41\xa8\x12\x79\xf6\x9d\xf0\x25\xba\x73\x2e\x79\x3d\xf2\x9e\xa2\x91\x9a\x34\x6f\xba\x3f\x8c\x50\x24\xf5\x9a\xc6\x4b\x92\xd0\x57\xab\x37\x74\xe4\x9a\xe4\x80\x10\xc4\x7a\xd2\x3c\x33\x56\x87\xd3\xec\x53\x67\xf8\x1b\xf5\x8c\x40\xce\x9d\xd1\x72\x14\x84\xcc\xd8\xb4\xe5\x93\x33\x96\xda\xff\x4d\xe2\xc4\xc3\x59\x31\x05\x7d\xf2\x06\xe1\x20\xe2\x70\x36\x2b\xca\xdb\x2c\xcf\xf2\xac\xc0\x94\x6f\x16\x84\xa4\xdb\xbd\xeb\x57\x3d\x69\x5e\xd3\x78\xda\xf6\x06\x3e\xc0\xfa\xfe\xdf\xe1\x9d\x31\xf6\x65\x39\x4a\x48\x89\x01\xae\x20\xfb\xe5\xf4\xf9\xc9\xef\x6f\x37\x68\xf5\x48\x42\x5b\x0a\x3c\x5d\xe7\xbe\x79\x75\x8c\xd0\xbf\x86\xdf\x26\x79\x06\xf6\xc3\x94\xd2\x5b\x2b\x0a\xa0\xa6\x10\x5e\xac\x6f\x11\x8a\x72\x7e\x7d\x53\xa5\x32\x6c\x10\x84\xe2\x63\x99\x15\xd5\x22\x2b\xca\xeb\xac\xa8\x70\x3e\xcf\xab\xc9\xef\x74\xa3\xad\xb1\xfa\x67\x60\x7f\xec\x76\xac\x00\x1a\xda\x1a\xae\x3d\xef\xbb\xf1\x6d\xa6\x00\x1e\x0e\xe4\x03\x0b\x42\x94\xfd\x62\xfa\xca\x87\xa9\x41\x84\xc6\x47\x3e\xb5\xb4\x6a\x0d\x3f\xd8\x61\x08\xe7\xe2\xbe\x3b\x1e\xde\xd0\x3c\x57\xea\xcb\x73\x93\x92\x6c\x78\xe7\x93\x59\x38\x50\x1c\x14\xc0\x2a\x84\xc8\x1e\x41\x77\xc3\x55\xcb\xbd\x55\x00\xcb\xd1\x75\x9e\x11\x6e\xcb\x3c\xff\x1b\x00\x00\xff\xff\x6d\x1d\x7e\xa5\x6d\x02\x00\x00")

func configs_config_yaml() ([]byte, error) {
	return bindata_read(
		_configs_config_yaml,
		"configs/config.yaml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"configs/config.go": configs_config_go,
	"configs/config.yaml": configs_config_yaml,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"configs": &_bintree_t{nil, map[string]*_bintree_t{
		"config.go": &_bintree_t{configs_config_go, map[string]*_bintree_t{
		}},
		"config.yaml": &_bintree_t{configs_config_yaml, map[string]*_bintree_t{
		}},
	}},
}}
