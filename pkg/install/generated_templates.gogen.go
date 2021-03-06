// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 7186,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x6d\x6f\x1b\x37\xf2\x7f\xef\x4f\x31\x50\xfe\x40\x62\x40\x5a\xd9\x75\xdb\xff\x61\x7b\x2a\x2e\xcd\x83\x9b\x4b\x93\x1a\x71\x72\x87\xbe\xaa\x29\xee\x48\x4b\x88\x4b\xee\x71\xb8\x52\x05\xa1\xdf\xfd\x30\xe4\x3e\x70\x65\xd9\x2e\xf2\xee\xf2\x22\xb6\xc9\xe1\x70\x9e\xe7\x37\xdc\xd9\x6c\x76\x26\x6a\xf5\x2f\x74\xa4\xac\xc9\x41\xd4\x35\xcd\xb7\x97\x67\x1b\x65\x8a\x1c\x5e\x63\xad\xed\xbe\x42\xe3\xcf\x2a\xf4\xa2\x10\x5e\xe4\x67\x00\x46\x54\x98\xc3\x4a\x37\x7f\x1c\x0e\xa0\x56\x90\x7d\x14\x15\x52\x2d\x24\xc2\x9f\x7f\xb6\xfb\xe1\xcf\x1c\x0e\x87\xf1\xee\xe1\x00\x68\x0a\x26\xa3\x1a\x25\x33\x73\x58\x6b\x25\x05\xe5\x70\x79\x06\x40\xa8\x51\x7a\xeb\x78\x07\xa0\x12\x5e\x96\xbf\x88\x25\x6a\x8a\x0b\xe9\xdd\x4c\xed\x9d\xf0\xb8\xde\xc7\x4d\xbf\xaf\x31\x87\x4f\x28\x1d\x0a\x8f\x67\x00\x1e\xab\x5a\x0b\x8f\x2d\xb3\x44\x03\xfe\x27\x8c\xb1\x5e\x78\x65\x4d\xcf\x1c\xa0\x76\xb6\x42\x5f\x62\x43\x99\xb2\xf3\xda\x3a\x9f\xc3\xe4\xea\xe2\xea\x72\x02\xcf\xc0\xa3\xd6\x09\x05\x78\x0b\x24\x9d\xa8\x11\xe6\x15\x7a\xa7\x24\xb1\x72\xb5\x55\xc6\x3f\x27\xe0\xc3\x59\xcb\x58\x8f\x74\x38\xd2\x02\xa0\xb3\x45\xd8\xb2\x05\xde\x8e\xac\xc0\xff\x96\xe8\x45\xb6\x69\x96\xe8\x0c\x7a\x0c\xc2\x59\xca\x41\x2b\xd3\xb2\x60\xd3\xb9\xad\x92\xf8\x52\x4a\xdb\x18\xff\x71\x7c\x03\xc0\xd6\xea\xa6\xc2\x5e\x86\x59\x2b\xc3\x5a\xf9\xd9\x06\xf7\xfd\x45\xc4\xe6\xf3\xc3\xc5\xdd\xca\xc0\x6f\xc6\x47\x8a\x10\x19\x09\x55\x81\x2b\xd1\x68\xff\xc1\x16\x98\xc3\xc5\xb7\x17\x17\xf0\x0c\x76\x25\x1a\xa8\x58\x1a\x2c\xc0\xa1\x28\x66\xd6\xe8\xfd\x14\x76\x08\x3b\x6b\x9e\x7b\x58\x22\x88\xa5\x46\x36\xa4\x2c\x2b\x5b\x9c\xb5\x0c\x9f\xc1\xe7\x52\x11\x28\x02\x01\xbe\xaa\x57\x04\x0d\x61\x01\x2b\xeb\x60\x8d\x06\x9d\xf0\xca\xac\xe1\xf6\xf6\x67\xd8\xe0\x9e\x32\x78\x67\xe0\xfd\xdf\x08\x7e\x5c\xc0\x65\x76\x79\x31\xed\xb9\x74\x77\x47\x15\x08\x84\xc3\x54\x0e\xb2\x2c\x8a\x41\x2c\x40\x00\x61\x2d\x38\x9a\x5a\x43\xc1\x0e\x7b\x36\x52\x18\xd8\x39\xe5\x59\xd0\xec\xb4\xfd\xd6\x68\x7a\x63\x60\x55\xfb\xfd\x6b\xe5\x52\x23\x56\x58\xa8\xa6\xca\xe1\x03\x56\xd6\xed\x53\x3d\x11\x56\x56\x6b\xbb\x63\x8d\xda\xab\x15\x05\x55\x1b\xe2\x35\x01\xb2\x21\x6f\x2b\xc5\x16\xd8\x18\xbb\x33\xbf\x97\x96\x3c\xf5\x2c\x56\x4a\xe3\x14\x76\xa5\x92\x25\xec\x6d\x03\x3b\xa5\x75\x54\xca\x5b\x28\x2c\x27\x28\x2f\xf3\x21\xfe\xc5\x81\xdd\x19\x16\xbb\x67\xe0\xb0\xb6\xe0\x84\x2f\xd1\x81\x2f\x85\x69\x2f\x5e\x2b\x5f\x36\x4b\xb0\xbc\x88\xa0\xd5\x06\x33\xf8\xcd\x36\xcf\xb5\x06\xa1\xc9\x76\x57\x8c\x8d\x0d\xca\x83\x32\xde\x86\x33\xd2\x1a\x2f\x94\x41\x37\x85\x25\x6a\xbb\xcb\xe0\x16\x07\xab\x96\xde\xd7\x94\xcf\xe7\x85\x95\x94\x71\x60\xc9\x82\xc3\x1a\xcd\x9c\x73\x96\xfc\x7c\xdd\xa8\x02\x69\xde\x10\xce\x6a\xa7\xb6\xc2\x63\x08\x3d\x56\x24\x2b\x7d\xa5\x7b\x4e\x9d\x2f\x88\xca\x99\xb4\x66\xa5\xd6\xfd\x16\x40\x5c\xf8\x20\xea\x3c\x59\x4c\x33\x70\x96\x1c\xfb\x5a\xbf\x84\xd4\x9c\x47\x26\x43\xf8\x3d\xe9\x93\x9d\xa2\x92\x57\x4a\xb1\x45\x10\x50\xa8\xd5\x0a\x1d\x57\xdb\x8e\x43\x9b\x55\x43\x45\x0d\x2e\x88\xec\x52\x27\x70\x55\xda\xaa\x02\x3b\xb3\xaf\xd4\xba\x12\xf5\x20\x88\xf2\x25\x08\x03\x68\xbc\xdb\x07\x1d\xee\x22\xd1\xdd\x14\x84\x29\xa0\x31\xd2\x56\x5c\xe6\xc3\xf9\xa8\xed\x87\xe0\x4e\x61\x8a\x9e\x0b\x9a\x6d\xe0\xa0\x90\x5a\x7f\xde\xf3\x00\x9b\xe1\x2b\x3c\x90\x1c\x7b\xd2\x03\xa1\x12\x78\x0b\xaa\xe2\x02\x0b\xd7\x37\xd7\xa1\x08\xc0\x0b\x56\x8b\xd4\xda\x28\x33\x5c\xce\xca\x6d\xd1\xa9\x95\x92\xa1\xd2\x43\xdd\xb8\xda\x12\xd2\xf9\x5f\x30\x64\xcf\x25\x96\x8f\x68\x45\x36\x10\xdf\xf7\x17\x0c\x07\xc2\xad\x87\x34\x7d\xc0\x62\xeb\x7a\xcd\xf5\x83\x12\xd3\x8c\x4b\xf0\xb3\x07\x8a\xf0\xfd\x73\x27\x8a\x70\x67\xce\x3e\x13\xef\xd5\xff\xa4\x43\xb4\x56\x77\x18\xea\xa4\xb1\x30\xc9\x63\x26\x4e\x40\x55\x62\x8d\x31\xfa\xf9\x40\x06\x6f\x95\x29\x82\xce\x15\x97\x15\x87\x72\x88\xda\x58\x52\x34\x0a\x42\x2e\x1e\xe1\x28\x3b\x81\x01\x06\x08\xdf\xe7\x7d\xd9\x2c\xb3\xc2\xca\x0d\xba\x4c\xda\x6a\xee\xe6\xb1\x06\x84\x1f\x73\x2f\x7a\xd3\x75\x7e\x64\xa0\xc0\x20\x82\x6f\xf5\x62\x0d\x2c\x69\xd6\xd3\x84\x6b\x72\x68\x19\x2a\x9b\x72\xcb\x2f\xb3\xcb\xff\xcf\x2e\xc6\xb4\x37\x8d\xd6\x37\x56\x2b\xb9\xcf\xe1\xdd\xea\xa3\xf5\x37\x0e\x29\xd5\xc2\x21\xd9\xc6\x49\xa4\xb4\x8e\x3b\xfc\x4f\x83\xe4\x47\x6b\x00\xb2\x6e\x72\xf8\xee\xa2\x1a\x2d\x56\xa1\xd4\xe7\xf0\xfd\xb7\x1f\xd4\x80\x2f\xac\x4b\x0f\xcf\x06\xcf\xdc\x04\xac\x71\x75\x71\xc5\x9d\x53\x99\x95\x75\x55\x08\x59\xa1\x7b\x6a\xad\xb6\x68\x90\xe8\xc6\xd9\x25\xa6\x12\xb0\x49\xaf\xc7\x5d\x3b\x5e\x15\x19\x8e\x97\x85\x2f\x73\x98\x8b\x5a\x45\x4b\x6f\xbf\x9f\xab\x02\x8d\x57\x7e\x9f\xd5\xcd\x32\xa1\x55\x46\x79\x25\xf4\x6b\xd4\x62\x7f\xcb\xf9\x59\x50\x0e\xdf\x25\x04\x5e\x55\x68\x1b\x7f\x62\x8f\x9b\xac\xfa\xdf\x10\x35\x49\xda\x91\x63\x4e\xc3\x23\x88\x6d\xee\x26\x4a\x86\x5e\x06\xc9\x8a\x39\x51\xc9\x00\xd1\x46\xc8\x0a\xda\xb6\xf5\x66\xcd\x2e\x03\x65\x62\xcc\x3d\xa7\x78\x86\xa8\x9c\x8f\xca\x64\x67\xb3\x5f\x8d\xde\xe7\xe0\x5d\x83\xcc\x8d\x31\x50\xa8\x50\xcb\xb6\xb0\x73\x4a\xd5\xe8\x56\xd6\x49\x64\xa6\x11\xf4\x30\xe6\x79\x48\xf0\x14\x97\x8c\x65\xdf\x0a\xd7\xca\x1e\xc9\xbe\x4e\xfc\x24\x47\xdf\x19\xa9\x9b\x50\x39\x19\xba\xc5\x06\xd7\x55\xd5\x88\x0d\x9e\x80\x32\x1d\x98\xf9\x81\x8f\x1e\xc1\x8c\xbe\xba\x42\x81\x52\x0b\xc7\x90\x6d\x69\xb7\x49\x01\x78\x04\x06\xc4\xf2\x98\x2a\xef\xac\xf5\xf3\x8c\xa8\x7c\x50\x01\x61\x46\xb7\x4e\x86\x16\x35\x89\x37\x4f\x3b\x92\x84\x03\x9a\xad\x72\xd6\x84\x86\x10\x7b\xed\xe4\xfd\x97\x9f\xde\xbc\xfa\xf5\xe3\xdb\x77\xd7\x93\xd8\x02\xa6\x6c\x0f\xbb\x45\xe7\xc6\xfd\x3a\x61\x13\x5a\xdc\x72\x1f\xbb\xa9\xd7\xa7\x74\xbc\xd7\x68\xef\xeb\x38\x04\x27\x13\x3f\xa8\x28\xf7\x3c\x9e\x58\xba\xdb\xb8\x44\x27\x50\xa4\x95\x2e\xf8\x24\x61\x71\x0c\x68\x52\xa7\x07\x34\xd3\x41\x6f\x61\x40\x68\x8f\xce\x30\xb4\xbe\x27\xf1\xca\xd9\x8a\xc3\xa2\x43\x2c\x53\x10\xc4\xe1\xd6\x76\x55\x36\x83\xb6\x72\x43\xf7\x9d\x8d\x66\x9b\x9f\xb0\xcb\x60\xee\x91\x5d\xb6\x42\x37\x78\xcf\x26\x4f\x05\xf1\x71\x0c\x74\x3d\xf7\x91\x08\xe0\x96\x3f\x6e\xf5\x8f\x34\xfb\x07\xe2\x92\xa9\x22\xba\x19\xd1\x8d\xeb\xc3\x53\x99\xb7\x13\x0c\x4a\x2c\x50\x53\xd7\x7a\x0f\x3f\x7f\xfe\x7c\x03\x4b\x41\x4a\x82\x68\x7c\x09\xd2\x61\xa8\xa4\x42\xc7\xae\x3e\xcc\x03\xcc\x70\xab\x44\x50\xfc\xee\xfa\xdd\xe7\xdf\x5f\x7e\xf9\xfc\xf3\x97\xdb\x37\x9f\xee\x82\xba\xfd\xd2\xfb\x37\xbf\xdd\x8d\x02\x7e\x2b\x9c\xe2\x69\x8e\x3a\x80\x9c\x30\x8c\xf0\xe5\xc8\x7f\x6f\x9d\xad\xc6\x3e\x8c\x64\x9f\x70\x95\x8f\x34\x1f\x61\x45\x2e\x6c\xac\xc2\x60\x00\xb6\x79\x3e\xb2\x47\x34\x41\x9c\x51\xb1\xe0\x4e\x2c\x85\x2c\xb1\xe0\xd0\x4a\x63\xbb\x87\xd5\x6c\x29\xe6\x3e\x4d\xb8\x58\xd7\xe2\xe6\xe4\x40\x3b\x63\x87\x83\xd3\x70\x09\xcf\x86\xad\x8d\x7d\x89\x94\xc6\xc2\x80\x5e\xfd\xce\xb2\x94\x0d\xdb\x29\x64\x5c\x78\x49\x08\x81\x08\xa5\xdd\x85\xf9\xd7\x1a\x83\x32\xb8\x4c\xf9\x71\xec\xcc\x66\xbd\x02\x61\xf8\xe1\xcb\x17\xfd\x52\xd6\x82\xbe\x8c\xb6\x32\x93\xba\x21\x8f\x2e\xe3\x02\xae\x53\x93\x7c\xa1\x58\x6b\x06\x53\xbc\x8a\xa4\xef\x6e\x46\x4a\x71\xd9\x21\xf4\x61\xbe\x1e\x47\xf6\x20\x43\x47\xcf\xd1\xe5\x1d\x53\x86\x89\x37\x69\x41\xa9\xc4\x2d\xf5\xe2\x6c\x84\x32\x15\x41\xd5\x50\x78\x01\x08\xd6\x53\x58\xc4\x74\x5a\x86\xc6\x16\x30\x5e\x18\xfc\x5f\x74\xd3\xf4\x79\x2a\x4b\x57\x5c\x62\x1a\x72\x00\x27\xf3\xff\x48\x10\x6e\x06\xb1\xc1\xcd\x0a\xe5\x16\xf7\xda\x5e\x2a\xd6\xa7\x04\x61\x0e\xce\xfb\xf2\xe9\x97\xf8\x40\x21\xcc\x3a\xee\x5d\x2b\x1f\x86\x66\x52\xde\xba\x7d\x5f\xae\xdf\x32\x32\x4e\xd8\x3d\x96\x73\x1c\x36\x89\xee\x6d\xca\x9c\x4c\xa7\x34\x17\x3a\xec\xfc\x7f\x2f\xd2\xcc\x3c\xcf\x87\xbf\xdf\xbf\xf9\xed\xfc\x1f\x71\x74\x0f\xb0\xba\x21\x74\xf3\x41\xd8\x2c\x4d\x74\xb6\x0f\xa7\x53\xe3\xf4\xe2\x70\x80\xec\x5a\x79\x56\x36\xbc\xe1\x8d\x29\x96\x4e\x18\x59\x76\x44\x3f\x85\xbf\xe2\x6b\x9e\x5a\x85\x25\xae\x5f\x74\xea\x24\x63\x38\x3e\x77\x1b\x22\x85\xfe\x69\x95\x49\x0e\x4c\xa6\x93\xf6\x51\x50\x13\xa6\xc7\x1f\x2f\x6a\x0e\x39\xf0\x64\x9c\xba\x2a\x61\xd4\x8a\x31\x39\xe7\x10\xa9\x02\x5d\x74\xc7\xd1\x64\x13\xde\x24\x2c\x21\x34\xa6\x40\x77\xe4\x63\x87\x5a\x78\xb5\xc5\x00\x39\xa9\x8b\xc0\xf5\xc8\xcf\x47\x39\xd9\x2b\x47\xcd\xb2\x50\xee\x72\x1a\x7f\x7e\xd3\xbf\x70\x0e\xc6\x09\x2f\x98\xa7\x8c\x13\x9e\x05\x3b\xab\x76\x54\x27\x18\x7c\x21\x74\xa7\xce\xb3\x73\x7b\xcf\x31\x0d\x9c\x3e\xff\xa6\x12\xea\xa4\x00\xc8\x1b\x1d\x87\x8e\x6a\x78\xa3\x3d\xe9\x0e\xe4\x52\xb2\xb3\x6c\x50\x34\xe1\xf9\x8e\xed\xc4\x1d\x5b\xf9\xa3\x01\x3c\xb5\x55\xdb\xfb\xda\xce\xb6\x78\xa4\xd5\x75\x27\x5a\x5e\x7c\x6a\xf1\xf7\x0d\xee\x41\x15\x3f\xf6\x64\x8f\xc0\x99\x44\x2a\x66\x21\x7c\xe3\x70\xf4\x0a\x70\xe2\xae\xb0\xbd\x9f\xf5\xf4\x34\x2a\x57\x5d\xb5\x06\xe5\xa1\x14\x14\x5a\xb1\x35\x7a\x0f\x42\x4a\xa4\x58\xd1\x4b\x8c\x0f\x69\x2f\xba\x37\x9b\xbb\x95\xd0\x84\x77\xe7\x67\x87\xc3\xac\x73\xc4\xa7\xb6\x87\x9f\xf2\x45\xc7\x34\xd0\xdf\xcf\x87\xd3\x64\x27\xfc\x44\xde\x35\xd2\x47\x79\x77\x61\x9c\x67\x88\xd7\x78\xa0\xbd\x91\xb0\xb4\x76\xb3\x41\xac\x39\xea\x7b\x51\x27\x6b\xe5\x27\x53\xa8\x50\xb0\xc1\xb9\xa0\x81\x08\x33\x76\x9b\x08\x4d\x4d\xde\xa1\xa8\xfa\x8c\x38\x3f\x12\x8c\x59\xcf\xc8\x0b\x8f\x0b\x2e\x30\x0f\xc6\x8d\xc1\x3f\x7c\x17\x3c\x49\xc7\x13\x06\x26\xdd\x1d\x93\xae\x1f\x25\x4c\x5e\x60\xb6\xce\xa6\xf0\x6f\x64\x64\xf9\x4a\xdb\xa6\x38\xcf\xc2\x03\x91\xb7\x1b\x9e\x4f\x08\x6a\xe1\xbc\x92\x8d\x16\xae\x73\x46\xcb\xe5\xb8\x95\xb6\xb7\x2e\x76\xc4\x75\x54\x32\xaf\x6c\xc7\x7c\xb3\x9d\x75\x1b\xea\x87\xcd\xa3\x63\xe1\xa2\x85\x58\xca\xcb\x6f\xae\xee\xff\x9f\x2a\xfc\x26\x46\x5f\x57\x95\xfa\x07\x6b\x6b\x1e\x09\x8d\x0f\x2d\xf5\xf5\x40\x7c\x14\x21\x1d\xbf\xd9\xc0\x6f\x11\x70\xe0\xc3\xd1\x72\xea\x48\xb8\xf8\x81\xd0\xb9\x45\xb7\x3d\xf1\x29\x83\x07\x82\x01\x01\x71\xae\xfe\x90\xb6\x62\xb1\xe1\x36\x16\xa3\x8c\xd0\x27\xdf\x47\x9e\x27\x9f\x58\x92\x6f\x25\xec\x9c\xf0\x74\x17\x40\x79\x36\xd2\x52\x2b\xf2\x68\x66\xad\x08\x8b\xfc\xea\xe2\xea\xf2\xac\xad\x63\x2f\x8b\x42\xc5\x07\x11\x6e\xb4\x2f\x19\x68\x8f\x54\x1e\xf6\x07\xac\x75\x38\x80\x0b\x6d\xfb\x89\xd3\xb3\xf0\xa1\x6a\x54\xfb\x86\xdf\xba\x0b\x7e\xad\x5b\xf6\xaf\x3f\xde\x76\x20\x89\xa6\xed\xf0\xd2\xb8\x16\x32\x81\x29\xac\x27\xb0\x81\x18\x2a\xb1\x0f\x0f\x49\x7a\x3b\x3c\x27\x1a\xd2\xd6\x6e\x9a\x1a\x14\x51\x83\x04\xd6\x00\xd9\x0a\xe1\x7d\xff\x79\x87\xb9\x37\x35\x0d\xaf\x85\x85\xa1\xee\xad\x6a\xf2\xd1\x1a\x9c\xa4\x3b\xaf\x82\x00\xe9\x7b\x61\xbc\x9c\xc6\x4f\x88\xdd\x10\x12\xe4\x1b\xed\xf4\xf3\xd1\xe4\x72\x72\xf6\xdf\x00\x00\x00\xff\xff\x19\xdf\x17\x3a\x12\x1c\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 931,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\xcb\x6e\x9c\x40\x10\xbc\xf3\x15\x25\xed\x35\xe0\x60\xc9\x17\x6e\x51\x9c\x44\x96\x12\x6b\xa5\xc8\xb9\xf7\x0e\x0d\x1e\x79\x5e\x99\xe9\xd9\x2c\x41\xfe\xf7\x08\xf6\x05\xb1\xfb\x04\x54\x55\x57\x75\xcf\x50\x96\x65\xb1\x81\x65\xab\x48\x3d\x73\x8b\x96\x83\xf1\x83\x65\x27\xc8\x89\x5b\xec\x06\x7c\x35\xf9\x00\xf1\x98\x19\xc5\x06\xca\x3b\x21\xed\x38\x42\x5b\xea\x19\x96\x85\x5a\x12\xaa\x0a\x0a\xfa\x17\xc7\xa4\xbd\x6b\x40\x21\xa4\x9b\x7d\x5d\xbc\x68\xd7\x36\xb8\xbf\xb4\x2d\xce\xf4\xa6\x00\x1c\x59\x6e\xae\xee\xe3\x08\xdd\xa1\x7a\x24\xcb\x29\x90\x62\xbc\xbe\x9e\x48\xf3\x6b\x83\x71\x5c\xa3\xe3\x08\x76\xed\x44\x4b\x81\xd5\xd4\x31\x72\x30\x5a\x51\x6a\x50\x17\x40\x62\xc3\x4a\x7c\x9c\x10\xc0\x92\xa8\xe7\xef\xb4\x63\x93\x8e\x1f\xde\x04\x28\x00\x61\x1b\x0c\x09\x9f\x24\x8b\xb0\x53\x99\x95\xfa\x3d\x3d\x70\x8e\x32\xe3\xbe\xe5\x9f\xab\x10\x53\xed\x58\xa8\x7a\xc9\x3b\x8e\x8e\x85\x53\xa5\xfd\x8d\x4f\x0d\x8c\x76\xf9\x70\x22\x5d\x96\x7c\x31\x2b\xdf\x35\x9b\x6a\x3e\x86\x05\xd0\xd4\xd5\x5d\x75\xfb\x71\x8d\x6f\xb3\x31\x5b\x6f\xb4\x1a\x1a\x3c\x74\x8f\x5e\xb6\x91\xd3\x74\x1e\x67\x16\xc5\x7e\x31\x58\x89\xd2\xe2\xae\xbe\x05\xb0\xc1\x0f\x3a\x68\x9b\xed\xe4\xe0\xe3\x30\xdd\x85\x9c\xf8\x03\xb4\x83\xe5\x9e\x76\x83\x70\x5a\x0a\x1f\x70\x67\xb1\x12\x26\xfd\x97\xd1\xf9\x08\xef\x18\x5a\xd8\x2e\xe9\x01\x75\x7d\x5b\xd7\xd8\xe0\x9e\x3b\xca\x46\x10\x7c\xbc\xe6\xda\x4c\x9c\xfd\xfe\xf8\xf8\xe4\x94\xb7\xf3\xed\x14\x8f\x9e\x05\xc6\xf7\x09\xbe\x03\x93\x7a\x46\xe4\xdf\x99\x93\x80\x5c\x8b\xc8\x29\x78\x97\xb8\xba\x34\x9a\xba\xae\x26\x3c\xee\x53\x19\xcd\x4e\xae\x03\x2c\x76\xbf\xf5\x51\x9a\x63\xba\x0b\x9c\x58\xe5\xa8\x65\xf8\xec\x9d\xf0\x41\x9a\x85\x2e\x66\xf7\x29\x3d\x25\x8e\xff\x6b\x4e\xd0\xb7\xe8\x73\x78\x8b\x91\x31\xfe\xcf\x36\xea\xbd\x36\xdc\xf3\x97\xa4\xc8\x90\xcc\xff\x50\x47\x26\x71\xf1\x2f\x00\x00\xff\xff\x68\x3b\xcf\xad\xa3\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
