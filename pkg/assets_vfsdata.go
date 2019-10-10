// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package pkg

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

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 10, 10, 12, 35, 56, 875477646, time.UTC),
		},
		"/handlers.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "handlers.tmpl",
			modTime:          time.Date(2019, 10, 10, 11, 49, 2, 894731047, time.UTC),
			uncompressedSize: 3826,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x5f\x6f\xdb\x36\x10\x7f\xb6\x3f\xc5\x55\x70\x3b\x69\x50\xd8\x0e\x28\xfa\xe0\x42\x03\xda\x24\x4d\xbd\xb5\x49\x16\xbb\xdb\x43\x51\x24\xac\x74\xb2\xb9\xc9\x94\x42\x52\x69\x0b\x55\xdf\x7d\x38\x52\x96\x25\x5b\x4b\xeb\x01\x7b\x98\x9e\xc4\xd3\x8f\xf7\xef\x77\x3c\x9e\xaa\x0a\x26\x85\xc2\x54\x7c\x86\x69\x04\xfe\x5b\xfe\x17\xce\x12\x94\x46\xa4\x02\x15\xb0\x99\x4c\x73\xb6\x10\x26\xc3\x00\xea\x7a\x3c\x1e\x57\xd5\x11\x24\x98\x0a\x89\xe0\xad\xb8\x4c\x32\x54\x27\x18\x67\x1e\x7d\x4d\x4b\x19\x83\xaf\x51\xdd\xa1\x82\xaa\x62\x4e\x6f\x5d\xcf\xad\x24\x80\xeb\x8e\xcc\xbe\x73\xb3\x62\x17\x05\x2a\x6e\x44\x2e\x67\x27\x75\x7d\xfd\xda\xa9\xf4\x63\xf8\x71\x29\x24\x3b\xce\xa5\xc1\xcf\x26\x80\x6a\x0c\x00\x20\x64\x08\xa8\x14\x79\x4a\x8e\x0e\xaa\xb8\xc2\x5b\x3f\x0e\x1c\x3c\x85\x0c\xa5\x8f\x4a\x05\xf0\x33\x3c\x69\x94\x6c\x9e\x3b\xae\x48\xd9\x5b\xbd\x04\x6d\x94\x90\x4b\xcd\x5e\x96\x22\x4b\x50\xf5\x60\x69\xae\x40\x90\x45\xc5\xe5\x12\xad\xf9\xbe\x1e\x0b\x5a\x1b\xf6\xaa\x50\x42\x9a\x4c\xfa\x8f\x9c\x56\xeb\xea\x7b\xf1\x81\x9d\x2a\x95\x2b\x3f\x08\x7a\xbb\x6a\xbb\x1a\x8d\x62\xf6\xcb\xfc\xe2\xdc\x5f\x19\x53\xb0\xb9\xe1\xa6\xd4\x2f\x79\x72\x85\xb7\x25\x6a\x13\x92\x63\x65\x6c\x1a\x83\xa3\xd1\xc8\xaa\x6a\xdc\x85\x9b\x3f\x75\x2e\xa7\x1e\x92\xcc\xbb\x69\x20\x75\xe5\xac\xb3\xb9\x05\xf9\x41\x1d\x34\x5f\x14\x9a\x52\x49\xb7\x70\xd6\x1d\x53\x6c\xae\xee\xd8\x60\x2a\x7d\xca\x77\x1c\x8c\x6b\xcb\x3a\xca\x64\xaf\x02\x14\xde\xce\xad\x8b\x6d\x0d\x90\x5e\x42\x4c\x9c\xeb\xe7\x7c\x8d\xb6\xb0\x6c\x72\xa0\xab\x1f\xbc\x2b\xbc\xf5\x02\x38\xaa\xeb\xb1\xf9\x52\x20\x50\x29\x76\x76\xd5\x75\x3f\x7c\xd2\xea\x38\x98\x14\x5c\xf1\x35\x1a\xb4\x95\xc0\x2e\x37\x2b\xdd\xb5\xdf\x62\x5a\x17\x76\x6a\x7b\x8b\x60\x04\x09\x60\x70\xf7\x82\x3c\x9b\x46\x5d\xf4\x59\x6e\x85\x2d\x7c\xd7\x56\x5d\xf7\x85\x7d\x74\x9b\xc8\xcd\xf2\x93\x30\x2b\x60\x0d\xe5\x2f\xf3\xe4\xcb\x06\xbb\x13\x74\x4c\x67\x41\x1a\xd2\x16\xc2\xe4\x23\x01\x29\xfa\x63\x27\xde\x6c\xb2\x0a\xc8\x3c\x01\xd8\x3c\x5e\xe1\x9a\xef\x78\xdc\xf7\xa3\xbf\x6a\xa9\xde\x3f\xed\x71\x2e\x1d\x1f\xb9\xea\xb1\x7d\x28\xdf\x9b\xd8\x3b\xd1\x2f\xf2\xe3\xad\x72\xeb\x2a\xb3\x65\xe1\x3a\x8a\x3b\xea\x3b\xb5\xb1\xd7\x21\x7c\x85\x1a\xaa\xaa\x83\xaa\x6b\x7b\x04\x73\xf5\x46\x68\x03\xef\x3f\xd8\xf7\x4d\x2b\x69\x4e\xbf\x03\x58\xc9\x35\x44\xb4\x72\xce\x3d\x7e\x0c\xef\xae\x66\xf0\x51\xc8\x44\xc8\xe5\x78\xd3\x08\xae\x43\xb0\xb4\x6e\xfb\x41\xec\xca\x4f\x77\x9a\x82\xfe\x24\x4c\xbc\x72\x40\xf6\x2b\x7e\xe9\x7c\xea\xd0\x29\x43\x98\xd8\x54\xb1\x99\xbc\xe4\x66\xd5\x65\x27\xe6\x1a\xc1\xa3\x98\x25\xd4\xb5\x37\xed\x75\x0e\x85\x9a\x55\x15\xec\x56\x33\x21\x21\x6a\xac\xfe\xce\xb3\x12\xc7\xe3\xd1\x68\xa7\xde\xe8\x49\x30\xe5\x65\x66\xfa\x4a\xb7\x79\x8a\x80\x17\x05\xca\xc4\x6f\x45\xa1\x6d\x6f\xb6\xf5\xa4\xbe\xf7\x4e\xe2\xe7\x02\x63\x83\x09\x94\x4a\x40\x5b\xe5\x53\xb8\x79\xa8\x6f\xbc\x70\x1b\x77\xa7\xe3\xb9\xd0\x3a\x45\x3f\x90\x84\xdf\x4a\x54\x7b\xb5\x2f\x52\xc0\x5b\x98\x20\xb3\x55\xe1\x09\x69\x70\x89\xca\xeb\xc2\xf6\xa0\xaf\x72\xb5\xe6\xc6\x82\x9f\x3d\x6d\xa1\xf7\xa4\xcd\x5d\x2a\x11\xb5\x9b\x38\x97\x77\xc4\xa8\xc6\x99\x34\x7e\xcc\xac\x53\xfe\x96\x8a\x20\x84\x9f\x9e\x84\xf0\xec\x69\x7b\xc3\xd0\xd6\x07\x11\x48\x91\x75\x78\xbe\x3f\x9d\xf6\x55\xb3\x3f\x14\x2f\x48\x1c\x82\x77\x6a\x53\x4a\x4d\xdd\x3a\x0d\x42\xc2\xd6\x64\x30\xde\xa6\xb0\x77\x7e\x33\x8d\x87\x07\xf7\xc2\xe4\x62\x28\xb0\xff\x30\x20\xa2\xec\x3b\x43\x6a\x5b\xd2\x61\x41\x42\x04\x03\x21\x0d\x74\xbb\xdd\x03\xf1\x0f\xb5\xf8\x1a\x79\x82\xea\x3b\x4a\xc7\x1a\x6e\x7a\x37\x73\xbb\xd8\x19\x9a\x9e\x17\x5d\x9b\xa3\xa6\x52\x19\xf5\x69\xda\x27\x14\x26\x6d\x4b\x14\xe9\x56\x1b\x44\x8e\x83\xaf\x5f\x3b\x16\x6c\x77\x8f\x0e\x23\xa7\x7b\x78\xed\x7e\xa1\x81\x7f\xd4\x28\x4d\x9f\x87\x81\xcc\x88\x14\x64\x6e\x9c\xb3\x07\x3a\xf9\x20\x02\x3b\xd3\x9c\xe7\xee\x46\xfa\x17\xce\x76\x3a\x0d\x5d\x66\xdf\xf4\x76\xf0\x9e\x74\x8c\xf6\x02\x68\xfa\x73\x6c\xe8\xd3\x30\x7b\xcd\xa5\x7a\x44\x4a\xbc\xe0\x39\xb4\x43\x88\xeb\xcc\xbc\x28\x32\x11\xdb\x7b\xed\x31\x0d\x61\x9d\x0e\x9d\x60\x4c\x6a\x49\xca\xce\xf1\xd3\x09\xc6\x79\x42\xf3\x6c\x2f\x3b\xcd\x48\xd6\x1c\xb6\x69\x44\xbb\x98\x83\xfa\x8f\xa8\xda\x2c\xe8\xf9\xfe\x51\x1c\x8d\x0e\x3d\x88\xc7\x5c\xfe\x60\xc0\x76\x34\xa0\x49\x73\x93\xc5\x66\x04\xdc\xbb\x0c\x0e\xa1\x47\x97\x45\x91\x2b\xe2\xa7\xc9\xf9\x11\xcd\x71\x53\x78\x78\xe7\x85\x10\x9b\xfb\xf8\x6a\x86\xd1\x81\xc1\xa3\x9d\x3b\x1a\x3a\x4b\x95\xb5\x34\xd2\x3d\xa9\x07\x28\x5f\xa3\x59\xe5\x49\x08\x93\xc2\xce\x1e\x89\x88\x0d\x78\x97\x17\xf3\x85\x47\xf7\xc1\x65\xae\x0d\x78\x67\xa7\x6e\x75\x86\x06\xbc\x93\xd3\x37\xa7\x8b\x53\xbb\x3e\xc1\x0c\x0d\x82\x77\xf9\xae\x41\x97\xb4\xf7\xc5\xe2\xf8\xb5\x5b\x72\xaa\x15\xef\xe2\x72\x31\xbb\x38\x9f\x5b\xd1\x45\x41\xc4\xeb\x60\xf7\xae\xb2\xa3\xcc\xa4\x20\x71\x55\x19\x5c\x17\x19\x37\xbb\x7f\x49\x1b\xe7\xdc\x6f\x90\xd7\xfe\x7b\x79\x34\x7d\xd3\x12\xf6\xd4\x0e\x8d\x69\x87\x24\xeb\xff\x94\xab\x9d\xff\x09\x07\xe8\x21\xf6\xa6\xd0\x06\xf3\xad\x8c\x6d\xdf\xfe\x0e\x00\x00\xff\xff\x34\x82\xc6\xce\xf2\x0e\x00\x00"),
		},
		"/interface.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "interface.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 932831877, time.UTC),
			uncompressedSize: 586,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\x6e\xab\x30\x10\x85\xf7\x3c\xc5\x91\xc5\x22\xb9\x4a\xc8\xfe\x4a\x77\x71\x95\xa0\x14\xa9\x29\xa8\xd0\x07\xb0\x60\x08\x56\x89\xa1\x66\x92\xb6\xb2\xfc\xee\x15\x3f\x49\xab\x26\x52\xbb\xf3\x8c\xe7\xe7\x3b\x67\x3c\x6b\x97\x28\xa8\x54\x9a\x20\x94\x66\x32\xa5\xcc\x69\x47\x5c\x35\x85\xc0\xd2\x39\x0f\x00\xac\x0d\xe2\x96\x8c\x64\xd5\xe8\x68\xe3\xdc\x4c\xe9\xab\xdc\x23\xbd\x2c\x90\xe3\xcf\x5e\xe9\x60\xdd\x68\xa6\x37\x9e\x0f\xd3\x49\x17\xce\x79\x1e\xbf\xb7\x04\x6b\xb1\x93\xcf\x14\x15\xa4\x59\x95\x8a\x0c\x82\x48\x97\x4d\x90\x29\xae\x09\xce\xe1\x82\x00\x3b\x6d\x5e\xc2\x48\xbd\x27\xf8\x47\x53\x2f\xe0\x13\xfe\xfe\x43\x90\x48\xae\xba\x09\xee\x5b\xd9\x61\x60\x5f\xc0\x6f\xfb\xca\x59\xa1\x72\x86\x48\xe2\x34\x13\xf0\x29\x48\x9a\x8e\x21\xb6\xe1\x18\x6d\x89\x21\x36\xe1\x7d\x98\x85\x43\xbc\xa1\x9a\x98\x20\x92\xa7\xa9\xfa\xd8\xf7\xfe\xcf\xd6\x77\x63\x28\x39\xaf\x20\xe2\x24\x8b\xe2\x87\x74\x48\xc5\x6d\xaf\xbf\x9b\xe3\x0b\xcc\x19\xe8\x55\x71\xd5\x63\x4c\x5f\xab\x55\x2f\xbf\x35\x4a\xf3\x99\x12\x02\x62\x10\x86\x8b\xd1\x4c\x87\xb6\x96\x7c\xeb\x1a\xe3\xa8\xab\x35\xa3\xbf\xb7\x32\x9f\xef\xdf\xfa\x9f\x92\x39\x91\x41\xc7\xe6\x98\xf3\x74\x82\xd4\x9c\xf0\x63\xa3\xe7\x3e\x02\x00\x00\xff\xff\x86\x24\xe5\x8a\x4a\x02\x00\x00"),
		},
		"/main.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "main.tmpl",
			modTime:          time.Date(2019, 10, 10, 12, 35, 56, 874917379, time.UTC),
			uncompressedSize: 335,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x31\x4e\x43\x31\x0c\xc6\xf1\xb9\x39\x85\x95\x09\x86\x26\x07\x60\xa4\x0c\x2c\x74\xe9\x05\xd2\x3c\xd7\x2f\xf4\xc5\x8e\x1c\x17\xa9\x7a\xea\xdd\x51\xa4\x0e\x08\x24\x3a\xf9\x2f\x7d\xbf\xc1\x31\xc2\xab\x4c\x08\x84\x8c\x9a\x0c\x27\x38\x5e\x81\x44\xb1\xdb\x0b\xec\xf6\xf0\xb1\x3f\xc0\xdb\xee\xfd\x10\x9c\x6b\x29\x9f\x13\x21\xac\x6b\xb8\xe7\xed\xe6\x5c\xa9\x4d\xd4\xe0\xc9\x6d\x3c\x72\x96\xa9\x30\xc5\xcf\x2e\xec\xdd\xc6\x9f\xaa\x8d\x43\xc5\xe6\xcb\x31\x64\xa9\x91\x0a\x6f\x49\xb8\xe4\x51\xbf\xb6\x76\xa6\x88\xaa\xa2\x7d\x0c\x8c\x16\x67\xb3\x36\xba\x9b\x66\xe1\xaf\x7b\x16\xa6\xee\xdd\xb3\x73\xeb\x6a\x58\xdb\x92\x0c\xc1\x17\x36\xd4\x53\xca\x18\xac\xb6\xc5\x43\xe8\x0d\x33\x6c\xc7\x87\x3f\xd9\x9c\x78\x5a\x50\xfb\xff\xca\xae\x0d\x1f\x10\x95\x8b\xa1\xfe\x31\xdf\x01\x00\x00\xff\xff\x11\x14\x09\xbf\x4f\x01\x00\x00"),
		},
		"/router.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "router.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 933059739, time.UTC),
			uncompressedSize: 479,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x50\xd1\x4a\xc3\x40\x10\x7c\xee\x7d\xc5\x70\x04\x69\x25\xde\x07\x08\x3e\x48\x1b\xda\x80\x9a\xd0\xa6\xcf\x72\x34\x9b\xe4\x30\x5e\xc2\xe5\x52\x85\x70\xff\x2e\x77\x29\xda\x82\xfb\x36\xb3\xbb\x33\xb3\x3b\x4d\x88\x7a\x43\x95\xfa\xc6\xe3\x13\x96\xaf\xf2\x83\xd2\x92\xb4\x55\x95\x22\x03\x91\xea\xaa\x13\x85\xb2\x2d\xad\xe0\x1c\x63\xd5\xa8\x4f\xd8\x53\xad\x06\x4b\x66\xdf\x8d\x96\x86\xa5\xc1\x7d\xad\xb4\x48\x74\xad\x34\xc5\x90\xbd\xc2\x95\xaa\x73\x2b\x4c\x0c\x20\xaf\x7f\x77\xd3\x38\x90\x39\x93\x99\x64\xaf\x1c\x63\x8b\x69\x7a\x80\x91\xba\x26\x44\xa3\x69\x63\x44\x61\x43\xe4\xd2\x36\x83\x73\x6c\x01\x00\x57\x33\x9f\x64\x9b\xae\x8c\x11\xf5\x21\x78\xa9\x4e\x16\x3c\xcf\x0e\x05\x47\x44\x22\xef\x06\x0b\xbe\x4d\x66\xb4\x25\x0b\xbe\x49\x5e\x92\x22\x09\x78\x43\x2d\x59\x02\xcf\x8f\x97\xe9\xd1\xef\x3e\x17\xeb\xdd\x0c\xa5\x3d\x35\xe0\x59\x5e\xa4\xd9\xdb\x21\x50\x59\x6f\x55\xa7\x87\xf0\x84\xdf\x24\x5f\xca\x36\xde\xdf\x39\x06\x18\xb1\x93\xba\x6c\x69\xc9\xfd\x8d\x73\x3a\x38\xc7\x63\x78\x62\xdd\xe9\x33\x19\x7b\x34\x6d\xb8\x6e\x6e\x90\x78\xbf\xf9\xc7\x8c\x44\xd6\x93\x91\xde\x2e\xdd\x78\x72\x96\x35\x2b\x86\x4b\x79\x6b\xd2\x65\x70\xfd\x8f\xf9\x43\x8e\xfd\x04\x00\x00\xff\xff\xbe\x70\x58\x84\xdf\x01\x00\x00"),
		},
		"/types.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "types.tmpl",
			modTime:          time.Date(2019, 10, 10, 10, 51, 17, 977911234, time.UTC),
			uncompressedSize: 1037,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\xcb\x4e\xc3\x30\x10\xbc\xe7\x2b\x56\x51\x0f\x09\x2a\xf9\x80\x4a\x39\xa0\x1e\x80\x43\x11\xe2\x71\xa6\x4b\xb2\x69\x5d\x62\x3b\xd8\x2e\x08\x59\xfe\x77\x94\x87\xf3\x28\xa9\xe0\x42\x4e\xf6\x7a\x76\x66\x77\xb2\x6b\xed\x25\xe4\x54\x30\x41\x10\x66\x92\x57\x52\x90\x30\x8f\xd9\x9e\x38\x86\xe0\x5c\x00\x00\x50\x63\x58\x01\xf4\x0e\x89\x6e\x5e\x92\xa7\xaf\x8a\x20\x94\xaf\x07\xca\x4c\x0f\x1b\x41\x3d\xee\x2a\xcf\x99\x61\x52\x60\x79\xaf\x64\x45\xca\x30\xd2\x63\x78\xfd\x99\x9a\xcb\xda\x2e\xe5\x0e\x39\x39\x07\x29\x1c\xb4\x14\xc9\x03\x7e\x6e\x48\x6b\xdc\xd1\x44\x81\x4a\x4d\x7f\xa2\xd1\x46\x1d\x33\x03\x76\x82\x6c\x39\x14\x8a\x1d\xc1\xa2\x52\xb2\x12\xc8\x69\xd9\x1e\x61\x95\xf6\xc5\x9f\x2f\xd9\x73\x2c\x2a\x54\xc8\xc9\x90\xaa\xf5\xea\xdc\x68\x83\x6f\x74\x9b\x93\x30\xac\x60\xa4\x06\xfe\xf8\x77\x8e\xc6\xd4\x55\xda\xe6\x24\xd7\xb2\xb9\xcf\x66\x9d\x0a\x3b\x37\x0d\x76\x99\xb0\xad\x4d\x5c\x85\xcd\x5b\x57\x07\x38\x17\x6e\x67\xec\x20\x91\x9f\x48\xb9\x20\xf8\x0f\xcf\x58\xd1\x35\x78\x83\x7a\x7d\xd4\x46\xf2\x73\x7d\xd6\xdf\x07\x2a\x78\x69\x67\xe1\x59\x70\x54\x7a\x8f\x25\x29\x48\x21\xba\xf0\x5d\x0d\x4e\xc5\x91\x60\x65\x3c\xab\xfa\xb3\xbd\x2e\x7a\x3a\xbc\xa3\xd0\xdc\xa4\xcd\x4e\x99\xb5\x7d\xff\xd3\xbf\x36\x22\x1c\x1d\x03\x6b\xbd\x93\x03\xcd\xd2\x5f\x1a\x37\xd7\x7e\x13\x75\xd2\xee\xa2\x1e\x31\x1a\xe2\x55\x89\x66\x6e\x61\xa3\x9c\x65\x06\xc2\x81\x36\x1c\x6b\xf8\x87\x3e\x08\x71\x5b\x9a\xaf\xec\x3b\x00\x00\xff\xff\xb2\x7a\xfc\x4f\x0d\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/handlers.tmpl"].(os.FileInfo),
		fs["/interface.tmpl"].(os.FileInfo),
		fs["/main.tmpl"].(os.FileInfo),
		fs["/router.tmpl"].(os.FileInfo),
		fs["/types.tmpl"].(os.FileInfo),
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
