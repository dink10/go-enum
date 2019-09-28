package assets

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

var _enum_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xcf\x6e\xdb\x38\x13\x3f\x5b\x4f\x31\x15\xd2\x0f\x52\x3f\x55\xce\x62\x6f\x2d\x7c\x6a\xbb\x41\x0b\xb4\x29\x90\xec\x5e\x82\x20\x60\xa4\x91\xc3\x8d\x44\xa9\x24\xed\x3a\xd0\xf2\xdd\x17\x43\x52\x32\xad\xc8\x49\x0b\x74\x2f\x86\xc9\xf9\x3f\xf3\x9b\xe1\xa8\xef\x5f\x43\x89\x15\x17\x08\xf1\x1d\xb2\x12\x65\x6c\x4c\xb4\x5c\xc2\xbb\xb6\x44\x58\xa3\x40\xc9\x34\x96\x70\xfb\x00\xeb\xf6\x35\x8a\x4d\x43\xc4\xf7\xe7\xf0\xe5\xfc\x12\x3e\xbc\xff\x78\xf9\x22\x8a\x3a\x56\xdc\xb3\x35\x42\xdf\xe7\xfe\xaf\x31\x51\xc4\x9b\xae\x95\x1a\x92\x08\x00\x80\xcc\xf0\x0a\xf2\xbf\x55\x2b\xc0\x98\x18\x45\xd1\x96\x5c\xac\x97\x74\x11\xf7\x3d\xa0\x28\xe1\xb5\x31\x96\x39\xae\x1a\x1d\x47\x69\xd4\xf7\xc3\x6d\x14\xba\x49\x4e\x90\x93\x45\x2b\x14\xe9\x27\xda\x09\x5d\x7e\x61\x0d\xc2\x9b\x15\xe4\x74\xc8\xed\x89\x84\x2d\x7d\xcb\xa4\x22\x5a\xc9\x0b\x0d\x71\xcd\x94\x6e\xab\x4a\xa1\x8e\xe1\xd4\x33\x81\x64\x62\x8d\x70\x22\x3f\x8a\x12\x77\x19\x89\xd4\x9b\x40\xdf\x5f\x74\x54\x60\x4c\xb4\xb0\x1a\x49\xc7\xb9\xd5\x41\x3c\x5d\xbd\x29\xee\x0f\x15\x3b\x9b\xff\x40\xc5\xa5\xd2\x60\x4c\xdf\xc3\x49\x3b\x0a\xa8\xcd\xad\x37\xe1\x34\x0f\x86\xbd\x01\xca\x16\x7e\x1b\x38\x6c\x2c\xf1\x4d\x6c\xcc\x72\x09\x17\xf7\xbc\xeb\xb0\x04\x4b\xea\x7b\xac\x15\xda\xfb\xbe\xf7\xdc\x5f\x25\x56\x7c\x87\x25\x49\x19\x03\x5c\x01\x23\xe2\x90\x22\x63\xa0\xad\x40\x3f\x74\xb8\x17\x71\xf7\x36\xe1\x43\x80\xbc\x1a\xac\xbf\x6b\x9b\x06\x85\x26\x42\x68\x26\xb8\x26\x7e\x27\xea\x8a\x3d\xef\xc8\x3e\x2a\x1f\xea\xa9\xcd\x4a\xe8\xd8\x0a\x78\xab\x99\x63\x14\x08\xa7\x63\xc6\x8c\x81\xff\x43\x90\xc1\xd1\x59\x97\x00\xcf\x1f\x16\x25\xe4\x7c\x6c\xe2\xa8\xb6\x93\x1b\x5b\x1d\x52\x60\xeb\x77\x58\x52\xf7\xc7\x83\xca\x45\x9c\x12\x3a\x41\x63\xd3\xd5\x4c\x23\xc4\x4a\x4b\x2e\xd6\x28\x63\xc8\xa9\x96\xd4\x2e\x5f\x99\x54\xd8\xf7\x7b\x5c\x1a\x03\x4c\x93\x88\x56\xa0\x5b\x28\x5a\xb1\x45\xa9\x81\x81\x13\xa6\x3b\x2a\x59\x28\x10\x55\x1b\x51\xcc\x69\x4a\x04\x81\xc3\x09\xa6\x90\x1c\x12\x33\x40\x29\x5b\x99\x42\x1f\x2d\x78\x05\xbb\x0c\xda\x7b\x8a\xef\xe6\x90\xcd\x22\xf0\x8a\x14\x5d\xbf\x25\x8e\x3e\x5a\x2c\x24\xea\x8d\x14\x24\x22\x78\x1d\x2d\x6c\x95\xa9\x85\x89\x4b\xd9\x9e\x19\x58\x26\xfe\x9c\xa6\x19\x54\x8d\xce\x3f\x90\xe5\x2a\x89\x5f\x2a\x82\xa0\x68\x29\xbe\x2d\xab\x79\x09\x53\x1f\xb5\x7c\x80\xab\x97\xea\x3a\xce\x80\xb4\x67\x3e\x1a\x95\x7f\x6a\xb9\x48\x26\xbe\xd2\xaf\xca\x20\xce\x20\x4e\x53\x0f\x3d\x42\xc0\x2f\xf4\xc8\xfb\x91\x86\xc0\xb6\x43\xc8\xce\xb0\x86\x49\x75\xc7\x6a\x70\x73\xf2\xb3\x3b\x5d\xe2\x4e\x03\x6f\xba\x1a\xa9\x27\x14\xe8\x3b\x04\x4d\x77\x9e\xbb\x46\x09\x0d\xea\xbb\xb6\x74\x85\x4c\x76\xf0\xea\xd0\x68\x1a\xaa\x4a\x52\x48\xae\xae\x6f\x1f\x34\x86\x15\xf4\xd1\x39\x42\xb2\xcb\x2f\x6c\x9a\x92\x34\x75\x35\x72\x60\xfb\x53\x34\xcf\xb8\xb4\x11\x3f\xe1\xd4\x81\xba\xc4\xca\x3b\xfb\xa9\x73\x8c\xfc\x12\x7e\xf0\xba\xb2\x59\xa6\x34\x5a\xe8\xa6\xb3\xce\x13\xe5\x18\x6e\x53\x0b\x4b\x62\x7a\xb1\xa2\x18\x42\xe4\xa1\x94\x16\x76\xaf\x76\xb0\x02\xdd\x74\x63\xfc\x2e\xd6\x61\x58\x0d\x65\x51\xdf\x86\x92\x5c\x14\x4c\x4c\x03\xa7\x3b\x81\x12\xb8\xd0\x28\x2b\x56\x60\x7e\x3c\x64\xe2\x4d\xdc\xf4\x1f\xd9\x7b\x13\xc4\xbb\x65\x12\x82\xae\x8b\xa2\x85\xfa\xce\x75\x71\x07\x5b\x8a\xd5\xcd\xbe\x84\x06\xac\xad\x5a\xc1\xd4\xc0\xf9\x26\x5a\xb8\x64\xad\x60\xeb\x09\x2e\x99\x01\xc1\x27\x71\x9b\x7a\x06\xc1\x6b\xa2\xda\x2c\x3c\xc2\xf5\x3e\x5b\xbe\x49\xff\xf3\xb4\x53\x82\xdd\x83\x35\xc9\x70\x29\xf9\x16\xa5\xa3\xcd\xe6\x79\x9a\x66\xcb\x49\x40\x77\x92\xee\x19\x9c\x81\xfb\x1e\xe7\xd9\x91\xd2\xfb\xad\x22\x68\xc7\x4f\x17\xe7\x5f\xa6\x0e\x12\x57\xee\xe9\x3f\xde\x8d\xa4\xe9\xc9\x6e\x0c\xd5\x86\x3d\x39\x6d\xc7\xa3\x2e\x8d\x1c\x3f\xd3\x8d\xd6\xad\xf9\x6e\x24\x74\xaa\x01\x9a\x43\x9d\xdf\xac\x26\xc6\xac\x70\x06\xff\x53\xe9\xdb\x67\x80\xf0\x1c\xa0\xd4\x2f\x6c\xe2\xaa\x66\xeb\xa1\x8b\xf1\xd1\xf4\x3a\x6b\x6b\x26\xd6\x40\x4c\x7e\x69\x1a\x61\x06\x94\xb2\xa7\x7a\x1a\x35\xb5\xf4\xf8\x4e\xee\xb3\xf5\x64\x6c\x5b\x56\xa7\xde\xf9\x6d\x14\xc6\xe4\x8a\x7b\xf6\xb4\x8f\x67\xa8\x75\xd8\x0b\xcf\x39\x79\x86\x34\xf8\x83\x99\x13\xc0\xec\xd5\xce\xdb\xbc\xa4\xc5\x6d\x62\x74\xcd\xf5\xdd\xe6\x36\x2f\xda\x66\xa9\xba\xea\xb7\xdf\x97\xdd\x1f\x94\xc8\x49\x8e\x9e\xb0\x4c\x4a\x93\x74\x58\x3f\xf6\x56\xe3\xc9\xc3\x78\x58\x32\xff\xe7\x60\x3b\x1f\xd7\x9f\x71\x43\x9f\x79\xbf\x61\x45\x9a\xbd\x39\x5e\x3d\xb8\x15\x9b\x3e\x0c\x46\x24\xb8\x35\xc3\x18\x42\xf3\xdc\x06\x60\xa7\xa1\x60\xcd\x28\xed\x77\xad\x39\x56\x17\x0c\xad\xc1\x35\x57\x9a\xb6\xdf\xae\x55\x8a\xdf\xd6\xc3\x5c\x76\x33\x5b\x11\xe5\x50\xde\xa7\x6c\x46\x69\x92\xc2\xd5\xf5\x3e\x5f\xba\xe9\x08\x43\x0d\xbb\xc7\x64\xb8\xcf\xa0\xc6\xf9\xf5\x85\x16\x97\xa2\xed\x1e\x12\xdb\x5b\xb3\x1c\x63\x09\xa8\x69\xec\xf7\xc9\xf8\x31\x34\x93\x92\xcf\xac\xb3\x09\x81\x86\x75\x61\x3e\x6d\x4a\xdc\x48\x7a\xf4\x22\xfa\x42\xfd\xc8\xa8\x1e\x86\x5a\x00\x10\x5e\xd1\xe1\xc8\x3e\xf9\x99\x75\x57\xbb\x47\xab\xa4\xd2\x6e\x16\xf8\x23\x6d\x64\x17\x9d\xe4\x42\x57\xc9\x04\x67\xc9\xcb\x32\x8d\x33\xd8\xd9\x21\x3a\x13\xae\x03\xb6\x0d\x98\x96\x99\x20\xe4\xbc\x6e\xbf\xa3\xb4\x0f\x67\x88\xd1\x7f\x03\x00\x00\xff\xff\xda\x5a\x5a\x47\xe7\x0e\x00\x00")

func enum_tmpl() ([]byte, error) {
	return bindata_read(
		_enum_tmpl,
		"enum.tmpl",
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
	"enum.tmpl": enum_tmpl,
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
	"enum.tmpl": &_bintree_t{enum_tmpl, map[string]*_bintree_t{
	}},
}}
