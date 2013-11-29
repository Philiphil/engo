package rog

import (
	"bytes"
	"compress/gzip"
	"github.com/ajhager/eng"
	"io"
	"reflect"
	"unsafe"
)

var (
	defaultFont *eng.Font
)

var _Terminal = "" +
	"\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x97\x79\x38\x14\xea" +
	"\xdb\xc7\x9f\x61\x64\x10\x92\x6c\x95\xa5\x23\x42\x9c\x24\x4e\xb2" +
	"\xcd\x24\xd9\xb2\x2b\xc3\x8c\x18\x14\x83\xc4\x41\xd9\xcd\xa0\xb2" +
	"\x1e\x4b\x19\x4a\x76\x09\x31\x96\x6c\xc9\x36\x64\x29\x0a\x65\x8b" +
	"\x62\x6c\x93\xdd\x64\x30\x8c\x31\xbc\xce\xf5\xfe\x7e\xef\xbf\xef" +
	"\xe7\xba\xee\xeb\xbe\x9e\xe7\x8f\xfb\xfa\xde\xdf\x3f\x9e\xe7\xbe" +
	"\x63\x2d\xcc\x0c\x78\xb9\x4f\x72\x03\x00\x78\x8d\x0c\xf5\xac\x00" +
	"\x80\x80\x7f\x03\x76\xe4\x30\xd9\xfb\x16\xbb\x1e\x26\x2e\x1f\x43" +
	"\x3b\x3f\x00\x78\x04\xff\x0d\x08\xc8\xc8\x14\x3d\xbc\xe4\xf4\x37" +
	"\x32\xbd\xce\xf9\x93\x9b\x97\x97\x8b\xdf\x16\x19\xc5\x0e\xc0\x31" +
	"\x8c\x91\xde\xd5\x9b\x81\xe3\xab\x59\x9a\x69\x4a\xf5\xbc\xd7\x99" +
	"\x93\xb7\x62\xa3\xa5\xcc\xa7\x1f\xf7\x89\x07\x29\x67\xa7\x68\x89" +
	"\x1a\xc4\x51\x7e\x9b\xb6\x59\xfb\x7c\x37\x0b\xa5\xa0\x6a\x97\xae" +
	"\xd1\xf0\x50\xc4\x5a\xd6\x42\xed\x7b\xbb\xce\x19\x4e\x28\x82\x6d" +
	"\x6f\x97\xed\xb0\xb0\xf2\x34\xf3\x90\x35\x37\xac\xa7\x01\xf0\x9c" +
	"\x54\x85\xe8\x42\xa7\xa3\x00\xec\x0a\x8c\x4e\xf5\x2e\x89\x57\xd9" +
	"\xb9\xa5\xf4\x2b\x24\xda\x8d\xd9\xec\x50\x81\xbc\x31\x36\xba\xbc" +
	"\xa5\x0b\x9b\x13\x9e\x93\x99\xf5\x7e\xd9\x5c\xa5\x73\xf0\x67\x16" +
	"\x8c\x69\xb4\xc1\x08\x58\xd5\xd3\xd8\x6f\x56\x11\x80\xdc\x9c\x26" +
	"\x6b\xed\xab\x78\x00\xe4\xdd\x54\x9a\x52\x69\x6b\x80\x26\xaa\xd9" +
	"\x3a\x28\x11\xc5\xa9\x9a\xfc\x70\xb5\x7a\x53\x0c\xbd\xc9\x5f\x77" +
	"\x12\x5a\x78\xdc\xa7\x62\xb7\xae\x6c\xff\xe8\xc8\x06\x3d\x5f\xbc" +
	"\xfb\xe5\xdf\x32\x8f\x4e\x80\xee\x58\xe4\x20\x4c\xb8\xae\xf4\x8d" +
	"\x86\xf8\x72\xe5\x37\x12\x5c\x30\xc3\x64\xff\xdd\xbd\x6a\xbe\xf6" +
	"\x02\xc1\x3b\x73\x40\xeb\xe5\xde\x79\xd3\xc8\x66\x70\x6f\x3d\x47" +
	"\xab\xb6\x30\xbd\x47\xfb\xfd\xc1\xcf\xee\x84\x74\x5e\x7a\x58\x9b" +
	"\xa1\xb5\x33\x1c\xef\xd9\xe3\xc0\x3b\xf5\xd5\x6b\x71\x91\x5d\x3d" +
	"\xfc\x12\x3a\xc5\x85\x59\xd9\xc5\x79\x6c\xff\x83\x10\xa7\xe1\xb2" +
	"\x5f\xdf\x91\xfb\xf1\x82\x5f\xab\xd7\x9b\xa1\xf8\xeb\x4c\x69\xb0" +
	"\x92\x55\x81\xec\x0f\xe9\x1f\x81\x48\x6e\x6d\xbc\x8b\x0c\x29\xb1" +
	"\x72\xf0\xc2\x4e\xf3\xb8\xbb\xd7\x97\xb6\x6e\x63\x0f\x60\xf3\xca" +
	"\x28\x8e\x9c\x6f\x0e\xdd\x2d\xe6\xf5\xbc\xd8\x5a\x1e\x4c\xd5\x52" +
	"\x8d\xff\xcc\x8b\xb1\xa7\x96\x3c\xcf\x9d\x38\xfa\x14\x80\xeb\x31" +
	"\xbc\xf2\x0b\xf1\xd8\x43\xdf\x64\xe1\xe0\x7f\x41\xa8\xdc\xdf\x3e" +
	"\x42\x60\x7d\x2c\x60\x7f\x4d\xbf\xe2\x11\x44\xd1\x2c\x4b\x8c\xac" +
	"\x93\x70\xa9\x0d\xc4\xbd\xa2\x47\xaf\x6f\x18\x88\xe1\x70\x27\xe8" +
	"\x1a\x71\xd8\xa0\x8c\xf2\x07\x92\xe4\xf7\xec\x06\x3a\x17\xac\x98" +
	"\x06\x99\x26\x5f\xc9\x6e\x5f\xf7\x0b\xcf\x2c\x7a\x2f\x13\xe0\xa8" +
	"\xee\x90\x2c\x84\x4b\xe7\x10\x7f\xbe\xed\xe9\x3b\xeb\x0b\x3c\xd4" +
	"\x7b\xc6\xdd\x86\xe3\x8d\xad\x13\x8a\x37\x27\xb0\x14\xb7\x97\x8f" +
	"\x60\xa9\x55\xab\xae\x4e\x91\xf2\x0d\x67\xcb\x76\x83\x73\xbe\xbe" +
	"\x31\x06\x9b\x16\xd1\xfa\xe3\xca\xc2\x82\xb2\xd5\xbc\x61\x23\x8b" +
	"\x4a\x51\xa3\xf9\x12\x0e\x94\xe1\xd0\x98\x34\x16\xe6\xa2\xb0\x92" +
	"\x0b\xbb\xbd\x97\x18\xff\xa9\xc0\x1a\x53\x6f\xaf\x6b\xdb\x26\x5c" +
	"\x42\xc3\x89\xd9\x25\xe5\x0e\xeb\x8a\x32\x55\x72\xd4\x6c\xc9\xd5" +
	"\x92\x3a\x09\x76\xdf\xb7\x1a\x55\x16\x07\x1b\x63\x53\x2f\x76\xbd" +
	"\xe3\x77\xbe\x0a\x2b\xf0\x3f\x93\xeb\xeb\x29\xba\x2f\x73\xde\x7d" +
	"\xac\x29\x14\xd8\xcc\x96\xd7\x2f\xb5\x5e\x34\xb1\x7b\x50\x50\x17" +
	"\xfd\xfe\x69\x3d\x77\x84\xa1\xd9\xc2\xc8\x2e\xae\x47\x9b\x10\x46" +
	"\x30\x26\xdd\xe3\xb4\x1b\xd2\x4f\x98\x81\xd6\xa0\x3a\xf5\xc5\x9d" +
	"\xe3\xe4\x42\xf5\x79\x94\x7b\xb7\x83\x24\xba\xd8\x61\x4e\xba\xe0" +
	"\xff\x40\x48\x27\x31\x40\x4b\xcb\xa6\x92\x72\xbc\x5b\x60\x57\xfb" +
	"\xa8\x4e\x77\xc8\xa9\xcd\x7b\xcb\xc5\xae\x9a\x25\x06\x07\xc3\x75" +
	"\x39\x28\xdc\xb2\x97\xcb\x16\xcb\xc3\x8c\xfb\xc1\xa9\x16\x79\x49" +
	"\xd2\x20\xbe\xbc\x04\xef\xa6\xc5\x5c\x08\x05\x93\xeb\x53\x36\x55" +
	"\xc0\x21\x48\xe0\x0f\xba\xa3\x17\xf3\x3c\x59\xf9\xed\x41\x63\x27" +
	"\xdb\x23\x2b\x3f\xb6\x2c\x5e\xce\x30\x35\x89\x2d\x87\x1b\x57\x45" +
	"\x09\xf9\x8f\xa4\x04\x64\x07\x96\x8a\x9e\xdb\xec\xed\x7f\x36\x78" +
	"\x5b\xad\xa8\x81\xee\xef\xeb\xd5\xc2\x57\xcf\x0f\xf5\xb6\x78\x04" +
	"\x99\xef\xb3\xa0\x7d\xc5\x57\xd3\x12\x5b\x89\xa3\xa0\x47\x22\x45" +
	"\xe4\x4c\xe9\xd7\x2a\x49\x8f\x67\xc1\x9f\xcf\x1e\x53\xf1\x84\x38" +
	"\x44\xdc\x7c\x7a\x84\x04\xe2\x64\x62\x6c\x3c\x71\xfd\x9f\xaa\x9f" +
	"\x8e\x11\xba\xb8\x62\x00\xaa\x22\xc7\xd5\x67\x52\xff\xe0\xc9\xb6" +
	"\x42\x94\x4a\xb5\xaa\xa3\xdd\x4e\x55\xd4\x60\x7a\xb2\x19\x42\xad" +
	"\xa4\xd0\x59\x93\x6c\xfa\xde\x52\x77\x6e\x87\x55\x70\xae\x6d\x4c" +
	"\xd8\x76\xb1\x3e\x67\x78\xe2\xc3\xe9\xdf\x2b\xf0\x5b\x31\x76\x09" +
	"\xb8\xff\xf4\x6d\xcb\x7e\xba\x08\x40\x64\xf3\xb6\xc5\x23\x36\xd9" +
	"\x5e\x65\x6b\x39\x44\xd9\x86\xe1\x52\x31\xf9\x45\x01\x22\x85\x1c" +
	"\xfd\xa9\xfc\x53\x6b\x51\x06\x1d\x7e\x8e\xac\x89\x8f\x81\x34\xd8" +
	"\x16\x72\xbc\x95\x1e\x96\x76\x57\x29\xfb\x6e\x0a\x3f\xe6\x04\xba" +
	"\xa2\x27\x84\x77\x8a\x45\xf6\x69\x8b\xe6\x47\xb2\xdd\xa1\x50\x93" +
	"\x34\xf8\xc8\x47\x47\xaf\x70\x4e\x59\x72\x78\x69\x7f\x7c\x66\x99" +
	"\xec\x1b\xec\xa3\x0b\x49\xbe\x0b\xfa\x94\x23\x82\xbf\xe7\x15\xb4" +
	"\x4a\x9e\x19\xc0\xb2\x6d\xfa\xea\xd4\x48\x70\x38\x84\xcd\xd4\xd4" +
	"\x45\x7c\x7b\xc2\x2c\xa8\x9f\x14\x69\xd6\xb5\xe2\xc5\xe0\x2e\x3a" +
	"\x2d\x83\xfd\xa3\x88\xb3\x0d\x39\xb2\xcf\x82\x68\x05\x1c\x50\x65" +
	"\x88\xe2\xfd\xc0\x2f\x23\x29\xcc\x5d\x85\xf4\x10\x05\x72\x79\x5a" +
	"\x7f\x31\x3d\x2a\x0f\xfc\x9e\x8c\x2a\xf4\xac\x64\x7c\x65\xcb\x84" +
	"\xe2\xb7\x96\xaa\x9c\xcf\x15\x2b\x6c\x5e\x94\x2f\x55\x81\x20\x73" +
	"\xa0\xb0\x08\xcf\xba\xd2\x2e\x5a\x93\xe4\xa4\xa6\x44\x46\xe0\x8a" +
	"\x0e\xea\xb4\x97\xa8\x85\xfa\x88\xbf\xac\x38\x4c\x82\xad\x83\xe9" +
	"\x80\x4f\xa1\xd6\xab\x7f\xe2\xfe\x4f\xf7\xca\x8f\x62\xd0\x7f\x4e" +
	"\xfe\x4d\x11\x6b\x55\x1b\xf7\x23\x49\xe4\x43\x34\xcd\x3b\x58\x29" +
	"\x58\xd1\x73\xca\xea\xba\x89\x00\x63\x15\x63\xcb\x55\xfe\x31\x60" +
	"\x96\x68\xb6\x4e\xc9\x5d\xa2\x9c\x2c\xf0\x1e\xad\x6d\x11\xed\xd8" +
	"\x78\xe8\x28\x73\x29\x71\x3c\x7c\x29\x00\xad\xe5\x36\xf7\x75\xa6" +
	"\x94\xc7\x0e\xf5\x56\x2c\xe8\x1f\x1c\x83\xe5\x56\x6a\xfa\x85\xdb" +
	"\xff\x75\xf8\x70\x1a\x92\x5c\x51\x1c\x65\x43\x26\x67\x20\xcd\xdc" +
	"\x35\xe3\xe9\xdb\x74\x88\x2d\x65\xd0\xba\xb5\xa1\x56\x42\xa0\x0b" +
	"\xb7\xc8\x24\x37\x0f\x9d\xf3\x8c\x25\x6a\xf5\xcf\xca\x6d\xe8\x9c" +
	"\x38\x29\xa4\xbe\xdf\x03\x47\x33\x7b\x47\xa3\x2a\xd6\x10\x6d\x11" +
	"\x1d\x73\x8e\x23\x9c\x3e\x84\x54\xbf\xc2\x27\xb1\xa9\x1f\x9d\x3b" +
	"\x63\x31\x29\xc9\xd0\x6c\x27\x11\xa9\x9d\xb3\xb6\x41\x6c\xc7\xba" +
	"\x4e\xe9\x81\x40\x8d\x64\xc5\x6c\xb7\xda\x8a\x95\xbc\x23\x15\x2d" +
	"\x39\xff\x1c\x5b\xd0\x74\xa8\xb7\x12\xaa\xa9\x13\x3c\xa1\x66\xf3" +
	"\x4f\x37\x72\x99\x9a\xc5\x4f\x62\x3a\xc7\x5e\xa2\x06\x40\x9f\x15" +
	"\x53\x9c\x9f\x26\xe4\xad\x08\xcf\xf2\x5f\xf2\x55\x59\xcc\x72\x17" +
	"\x4a\xf0\x55\x12\x9e\x2d\x42\x58\x0c\xd8\x5f\xd1\xb8\xf8\x12\x7a" +
	"\x2e\x36\xc1\xfb\xc7\xe0\x1a\x80\xb1\xd5\x5c\xc0\x13\xce\xe1\x06" +
	"\x57\x49\xa0\xb8\xdd\xef\x79\x82\xf9\x27\xbb\x2b\x9c\x20\x8f\xf3" +
	"\x41\x97\x16\x9f\x4d\x95\x68\xed\x4f\xae\x46\x99\xfe\x54\x0c\x3b" +
	"\x71\x22\x40\xa2\xbf\x12\x26\x6b\xb5\xdc\x20\x82\x40\x1e\x48\xf2" +
	"\x02\x0c\x29\x6e\xef\x85\xbf\xb4\x0d\xb9\x31\x2b\xdd\x03\xe2\xaf" +
	"\xf5\x78\xe4\x78\x5f\x71\xe6\x1d\xeb\x0f\xee\x4a\x1f\xfe\x46\xb7" +
	"\xb6\x16\xee\x0d\x93\x83\x9e\x58\xd7\xa7\x22\x99\x24\xe7\xe1\xcd" +
	"\x3f\xe1\xe3\xf6\x0d\x19\xd5\x01\x28\x84\xa5\x72\xb2\xef\x95\xbb" +
	"\xf1\xd7\xf4\xb5\xb9\xea\x03\xf0\x77\xea\x42\xde\x8e\x90\xef\x70" +
	"\xb0\xb9\x27\xd8\xdf\x36\x0b\x19\x32\x9e\x24\x8a\x8f\x17\x7c\xfe" +
	"\x38\x4f\xcc\x0d\xd6\x17\x38\xab\x44\x4d\x92\xda\x53\xe9\x38\x93" +
	"\xb9\xb5\xd5\xb2\x4a\x17\x0c\x6f\xdd\xc5\x85\xb7\x06\x95\x73\xcb" +
	"\x48\x12\x1e\x70\xfc\x0c\x88\xec\xc6\x24\x7e\x7b\x92\x69\x7c\xea" +
	"\x02\x94\x3e\x4b\x55\xe7\x4b\x9a\x9d\x16\x67\x14\xf4\xfd\xa9\xf0" +
	"\x58\xcd\xe4\x39\xac\xd4\x5b\x7f\xfa\x81\x1b\x5f\x26\x6b\x2d\x6b" +
	"\x3a\x0a\xc6\xe5\x25\xd9\x24\x9c\x92\x29\xac\x23\x25\x7a\x99\xd1" +
	"\x9f\x0f\x4a\x7d\xd9\x1c\x8c\x92\x53\xa2\x2e\x8f\xdf\xf1\xb8\xdd" +
	"\xa6\x25\xdb\x4e\x98\x75\xbb\xc9\x52\x6a\x11\x93\x44\xa0\x5f\x4e" +
	"\x0c\x6f\xaf\x04\x4c\xd0\x20\xb6\x97\x43\x4b\xa5\xdf\x76\xf2\xd1" +
	"\xf7\xf6\x5d\xb2\xb4\x57\x03\xd5\x08\xf5\x94\x8d\xdc\xf2\x8a\xf2" +
	"\x87\x90\xd7\x76\x84\x33\x2d\xc2\x43\x21\xd7\x97\xed\xe6\xc5\x42" +
	"\xcd\x65\xb6\xc8\xa2\xe3\x16\x2a\x81\x36\x37\xbc\x74\xa2\xea\x66" +
	"\x76\x79\x25\xda\x8e\x0d\x14\x1d\x2c\x34\x06\x8d\x95\x71\xfe\x53" +
	"\x3b\x11\x23\xce\xa5\x6c\x01\xb4\x64\x8f\x2d\x68\x41\xa8\x89\xfc" +
	"\xd5\xc7\x19\xc5\x49\xb0\xdb\x8e\xbc\xf7\xbe\x37\x67\xca\x08\x96" +
	"\xf9\xbf\xcc\xeb\x84\x06\x2a\x44\x01\xfc\xb0\x2c\xae\xf7\x5d\x85" +
	"\x7c\xeb\x17\xfa\xfa\x1b\x23\x5f\x11\xfb\x8b\x1a\x6a\xef\xf2\x83" +
	"\xee\x9a\x2b\x2c\x35\xc7\xa1\x26\xfa\x56\x04\x0c\xac\x6b\x29\xc4" +
	"\x0c\xd0\x03\x0d\x8b\xdb\xb1\xfe\x83\x9b\xc0\xb7\xb0\x3e\x99\xf2" +
	"\x35\x79\x72\x9b\x3b\xb3\x79\xf0\x1a\x1e\xd5\xdc\xef\x71\xfa\xf1" +
	"\x65\x08\xea\x4b\x24\x56\x16\x7f\x57\xf1\xfd\x53\xe1\x1e\xe3\xc7" +
	"\xd6\x93\xcb\x8b\x5f\x92\x68\x6b\x23\xb4\xa5\x2f\x26\x05\xbd\xb5" +
	"\x5a\xde\x1a\x82\x53\x1e\xc3\x25\x84\x2e\xe1\x74\xeb\xba\x87\xaf" +
	"\x3a\x8d\x0c\xa9\x58\x3d\xea\x36\xba\x7c\xc7\x58\x45\x28\xe9\x47" +
	"\x90\xc7\x78\xc8\x79\xc6\x47\xcb\x6e\x8d\xfe\xa1\xfb\xa5\x6d\x83" +
	"\xbe\xc4\xd4\xf1\x08\xf3\x08\xd1\x82\x9b\x3b\x84\x15\x5c\xbc\x67" +
	"\x5b\x6a\x01\x39\x6e\x11\x4b\xf9\x50\xd2\x20\x7c\x2f\x64\x57\xfa" +
	"\xc2\xfc\xe8\xc5\xbc\xb6\xe2\x88\x7b\xc3\x11\xe7\x67\x1b\x76\xd1" +
	"\xcd\x97\x99\xf2\x1d\x1e\xbc\xa4\x1f\x7c\x9a\x3c\xf8\x86\xa7\x1e" +
	"\xda\xe2\x2a\x69\x27\xf1\x45\x5c\x53\xf2\x1d\xdb\x19\xe8\x78\xe1" +
	"\x80\xe1\xfb\x12\xe9\x05\xe5\xc5\xd7\x3f\x3d\x7b\xd7\x39\x66\x35" +
	"\xcb\x29\xc7\x15\xb2\xe9\x3f\x63\x5f\x92\x2f\xbd\x7b\x26\x4e\x70" +
	"\x60\x99\xa9\xdd\xf2\x60\xf7\x53\x85\x6b\xcd\x2d\xcd\x1f\xbd\x2f" +
	"\x3d\x18\x93\x26\x76\x61\xd2\xdc\x10\xb5\xf3\xeb\x1f\xbd\xeb\x7a" +
	"\xe2\xb4\x22\x89\x5d\x7c\x92\x84\x1f\x5a\x8d\xef\x1f\xd8\xe0\x3a" +
	"\x63\x2f\xbd\xdf\x3b\x1a\x9f\x36\x7a\x25\xc6\x5f\xbe\xc3\xf8\x64" +
	"\xc5\x90\x44\xf9\xfb\xf4\xa6\xf6\xc5\x98\x4b\x7c\x18\x27\x1c\x3b" +
	"\xfd\x34\x00\x8a\x0d\x52\x74\x7b\x40\x2e\x24\x00\x3c\x23\x10\x6e" +
	"\xba\x3f\xec\xe5\xbd\x15\x87\x85\xca\xe9\xb9\x2e\xab\x31\xdc\xd4" +
	"\x4d\xf2\x4c\xc2\xcd\x4e\xe8\x1e\x4a\xf6\xa5\xb9\x9d\xde\xbb\x64" +
	"\xaf\xf8\x06\x76\xa9\x30\xf4\xd3\x74\xf9\xa4\xae\x68\x42\x28\x28" +
	"\x9d\x09\xcb\xc9\x56\x81\x34\x55\x74\xfe\xfa\xab\x88\xcb\x04\x0c" +
	"\xc8\xe2\x0d\x44\x17\x86\x99\xe8\x6b\x79\x42\x5e\x1f\x15\xa4\x3e" +
	"\xa2\x79\x42\xde\x34\x3e\xae\xdc\x54\x75\x12\x00\x10\x35\x98\xda" +
	"\xf3\xae\x6f\xa9\x96\xdf\x50\x76\x7a\x7f\xa3\x38\x8e\x13\x46\x66" +
	"\xd7\x44\x6b\x83\xdb\x2f\xcd\xa8\xff\x0c\x6e\x1e\x52\x01\x77\x8e" +
	"\xcc\x4f\xc7\x5d\x64\x38\xb8\xe2\x29\xf2\xc3\x61\x0e\x71\xf7\x4e" +
	"\x22\x60\xf0\x38\xdf\xe7\x4e\x5b\x5e\xaa\x37\x1c\x11\x5d\x35\xd8" +
	"\xb4\xfd\xa4\x11\x7c\x9d\x40\xf4\xdf\xc1\xf5\x07\xb7\x19\x03\x66" +
	"\xa4\x24\xee\xa6\xa5\xb9\x7a\xe3\xfb\x4e\x29\x76\x41\x73\xe2\x61" +
	"\x21\x1b\xdd\x2e\x42\x68\xf7\xaf\x4e\x22\x18\x5f\x48\xfe\xa1\x4f" +
	"\x61\xfc\x62\xf5\xd6\xda\x15\x0f\x25\x79\xcd\x48\x5c\x48\xdc\x96" +
	"\xbe\xcc\x72\xcd\x82\x18\x72\xe8\x96\xaa\x39\x94\x34\x57\x23\x2b" +
	"\xab\x5d\x2d\x4d\xb8\xf1\xa7\x47\x96\x4f\x8f\x93\x0a\xd5\xef\x98" +
	"\x0f\xd6\xc4\xdc\x4d\x22\x8b\xeb\xc5\x7c\x22\xe1\xf7\xc3\x41\xd0" +
	"\x74\xb3\xa4\x0c\x1d\x03\x0c\x2b\x8d\x7a\xdc\xb0\xad\x69\x8f\xca" +
	"\x02\xa8\x58\x06\xd5\xe1\x5a\x19\x4b\x4c\xf5\x3b\x07\xc9\xe0\xb7" +
	"\xc6\xe0\x73\x8f\x70\x9d\xe0\x3d\xcf\x7d\x4f\x4b\x4e\x04\xc0\xa7" +
	"\x94\x0c\xd1\xaa\xff\x80\x2a\x2e\x01\x18\x40\xb0\x03\x28\xc8\x85" +
	"\xec\x1d\x34\x0e\xfd\xc8\x43\xfe\xbe\x2c\x05\x06\xc0\xeb\x41\x40" +
	"\x2f\x12\x45\xfb\x0a\xca\x89\x94\xda\x77\xc2\x6b\xba\x74\x6d\x1d" +
	"\xfc\xc2\xe9\x4c\xe2\x26\xc7\x9b\x66\x8a\xd2\xf5\x90\x66\xdc\x50" +
	"\xc0\x96\xd6\xcf\xb7\x70\x5a\x88\x16\x9b\xfd\x6e\x35\x8d\x5c\xa0" +
	"\x5c\xc2\x60\x78\xd0\x72\xc2\x5c\x3e\x39\xbd\x2e\x12\x7f\x56\x44" +
	"\x73\xa3\x56\xa9\x1b\x95\xcc\x59\x27\x77\x17\x07\x68\x14\x95\x51" +
	"\xfb\xb1\xdc\x35\x7b\x65\xf7\x21\xd9\xc4\x9b\x85\x70\xc7\xb0\xee" +
	"\xdb\x2c\x63\x9b\xcf\x8d\xb8\x6d\x87\x36\xe2\xb7\x12\x42\x28\xb3" +
	"\xe0\x48\x72\x18\x05\xb9\x9c\x5b\x26\xea\x7a\xdc\xf6\xe6\xfa\x8d" +
	"\x25\xf5\xf8\x23\x3e\x23\xde\xc4\x36\x0b\xa5\xd9\xf9\xe7\x5c\x6f" +
	"\x6a\x5a\x9a\xbb\x98\xfb\xc1\x70\xbf\xc5\x94\xbb\xb7\x66\xa5\x2c" +
	"\xc7\x7a\xa2\x8f\x1c\xbb\xc6\x6e\xd1\x33\xea\xc1\xaf\x0e\x11\x71" +
	"\x1b\xb9\x0d\xd4\x5e\x6d\x3f\x7f\x4a\x2c\x2c\xcb\x87\x18\x9f\x0f" +
	"\xe0\x80\xe1\x3b\x5f\x78\x89\x4c\xed\xfc\x1c\xaf\x5c\x2a\xd8\x20" +
	"\xfe\xd5\xb0\x3c\x16\x82\x7a\x78\x06\xa5\x96\x64\xa1\xa0\x2a\x8b" +
	"\x08\x15\xc4\x8c\xd8\xe2\xc5\xfa\x20\xef\xc8\xc6\x8f\x01\x89\x33" +
	"\xf9\x0b\xa7\x72\x41\xd0\x13\x2e\xb0\x03\x5a\x46\xe2\x39\x44\xd7" +
	"\xc3\x60\x20\x14\x74\xb0\x95\xfd\x3c\x10\xa3\x65\x5c\xb3\x57\xab" +
	"\x46\x1e\x38\xba\x9f\x24\x52\x88\x4b\xd3\xb7\x6c\xae\x98\x06\x22" +
	"\xd3\x32\xf1\xc7\xac\xe4\x4a\x96\x5a\x50\x0e\x42\x52\xc6\xce\xc8" +
	"\x78\xaa\x01\x3b\x86\x23\xa7\xcb\x8c\x23\x14\x80\x01\x88\x09\xbe" +
	"\x13\xf9\x5b\x80\xfc\x6a\xd2\x33\x77\x07\xa4\x6a\x94\x59\x48\x40" +
	"\xc0\x84\x29\x7b\x7c\xaf\xa9\x3e\x77\x32\xa1\x47\x1a\xd8\xb6\x45" +
	"\xc3\xa2\x84\xca\xc7\xda\x39\x01\x46\x51\x15\x1b\xdb\xa7\x83\x1f" +
	"\x2e\x54\x89\x2a\x33\x0f\x79\xda\x20\xc2\x0b\x84\xeb\x58\x4a\xe5" +
	"\x0f\x0d\x78\x44\x73\x21\x95\x91\xec\xb9\x10\x30\xc5\x83\x38\xac" +
	"\x1f\xad\x13\x96\x70\x41\xaa\x18\x7f\x1d\xe0\x1f\x7f\xc1\x8d\x16" +
	"\x9d\x86\xc6\xa1\xf7\x3b\xa1\x47\x3c\x21\x87\xca\x81\x54\xe3\x84" +
	"\x6b\xd5\x7b\xf4\x2b\xfd\x3a\xb9\xd2\x57\x48\x7a\x31\x9c\x55\x91" +
	"\x7e\x6b\xb7\x9f\x86\x2a\xbf\x47\xbe\x6f\x25\x50\x0b\x51\x3a\x57" +
	"\x11\xac\x19\x6f\xc0\x58\x89\x24\x55\x1f\x28\xfd\x68\x55\xbb\x7c" +
	"\x70\x94\xcc\x77\xd5\xc6\xf2\xf0\x93\xf5\x8b\x5b\x6b\xb6\x6e\x57" +
	"\x91\xaf\x20\x9b\x6e\x1b\xb8\xf1\x8e\x1f\x69\xfc\x86\xf2\x68\xb7" +
	"\x74\x3b\x5f\x51\xe7\xf8\x9a\x76\x66\x25\x99\xd0\x8b\x9e\x30\xb0" +
	"\x50\x91\xf7\x90\x1d\x50\xa0\x20\x3d\x38\xeb\x20\xec\xbb\xbc\xdf" +
	"\xd1\x9c\xdb\x57\xb1\xcd\x91\x16\x2d\x65\x4d\xc3\x74\xe9\xf3\x92" +
	"\x27\x11\x46\x41\x2a\x6c\x1c\x22\xa4\xda\x74\x6b\x3a\x0a\x6a\xa3" +
	"\x69\xdc\xc8\x7d\xaa\x14\xd1\x44\x23\x32\xe6\x94\xaa\x62\x28\xfd" +
	"\xf6\xe1\x41\x10\x1f\xe1\x84\x3c\x7f\x40\xb9\xc2\x0f\xd6\x49\x5c" +
	"\xef\x72\xed\x20\xf8\x43\x01\x99\xf8\x17\x36\x6b\x56\x8d\x74\xc5" +
	"\xf8\xc4\x25\xbd\xd2\x30\xfd\x5f\x34\x37\xb3\xef\x23\x01\x6e\x8a" +
	"\xfc\xc5\xfb\x20\x39\x6c\xe6\x12\xe3\x0c\x0b\x7c\x3e\x6d\xfd\xe5" +
	"\x2c\x14\xda\x16\xb9\x3e\x67\x12\x5a\x23\x85\xa7\xa8\xc9\xf5\x27" +
	"\x98\x08\xa8\xb0\xdd\x1d\xc5\x5c\x90\x84\x61\x2d\xd4\xee\x36\xbf" +
	"\x1d\x0f\x7c\xbd\x65\x9c\x23\xb4\x10\x62\xb0\x7a\x0d\xf6\xeb\x2a" +
	"\x5e\x97\xaf\xae\xb4\xed\xcd\x24\x95\x0f\x18\x71\x20\x62\x2d\xf0" +
	"\x7e\x3a\xd1\x3e\x37\x21\x3d\x9d\xc2\xa9\x01\xe7\x71\xae\x70\x0c" +
	"\x5b\x4d\x3a\xf2\x1e\x76\x3c\x89\xa6\xdd\x2b\xc8\x33\x60\xc9\x83" +
	"\xbd\xcd\x42\xf2\x77\xca\x9c\x50\x8f\xe1\xc2\x8f\x7b\x7b\x72\x17" +
	"\xa3\x00\x45\xde\xfb\x76\x36\x26\x17\x6f\x7f\x3b\xd1\x88\x71\x91" +
	"\x82\xfc\x16\x39\x59\xb6\x0f\x96\x69\x5d\x4f\x7a\xf2\x04\x2c\x26" +
	"\x2b\xf6\x3e\x7b\xfd\x7e\x8f\x66\x11\x1e\xa5\x9e\x25\x0c\xd5\x14" +
	"\xbb\xfc\x18\xd5\x89\xa7\x3d\xa4\xa4\x9f\x6a\xcd\x44\xb1\xff\x77" +
	"\x86\x73\xce\xf3\x9a\x19\x93\x3e\x2b\x64\xa7\x26\x09\xfc\x48\xb9" +
	"\xac\xb4\x7f\x1e\x0b\xb7\x76\xfe\x21\xa1\xe3\x27\xa9\xf7\x89\x73" +
	"\x4a\x1e\xe6\x9a\xc1\x7e\xc3\x53\x57\x55\x8a\xd4\xfe\x4c\x52\x9e" +
	"\xfb\xd2\x46\xe0\x95\xe2\xe0\x64\x19\xcf\x20\x8a\xea\x59\xce\x92" +
	"\x99\xd3\x1d\x1e\x61\xdb\x76\x8a\x57\xbf\x74\xfe\xf1\xa2\xf7\x25" +
	"\x31\xf8\x45\x2f\x2e\x27\xf7\x3b\xae\xa3\x20\xe9\xbc\x37\x91\xf5" +
	"\xe5\x67\xb8\x1e\xbb\x8d\xe1\x80\xe5\x49\x28\xbe\x91\x0b\x24\xdb" +
	"\xe1\xa5\x53\x05\x40\xa5\x9c\xc9\xa6\xaf\x21\xd7\x91\x43\x83\x9f" +
	"\x0f\x9e\x90\x12\x21\x85\x31\x90\xee\x11\xb7\xd6\x9e\xb4\x06\xeb" +
	"\xaf\x32\x87\xbd\x44\xeb\x20\x5b\x5e\x88\x52\xa3\x6a\x67\x1d\xe8" +
	"\x7f\x67\xcd\x84\x95\xe9\x4c\xe0\x63\x5a\x8f\x91\x8a\x39\xd3\x04" +
	"\x3f\x78\x2d\xe9\xb4\xc8\x95\x8b\x85\xdb\xc4\x8a\xeb\xa7\x93\xe7" +
	"\x2f\x55\x8e\x0b\x46\x37\x04\x1a\x64\xbd\xeb\x6e\x8b\x97\xf2\x9c" +
	"\xdf\x82\xdb\x78\x46\xc1\x44\xe2\xf4\x70\x07\x8d\x9b\x9d\x05\xed" +
	"\x62\x15\xa1\x9c\x44\xf7\x92\xa6\x41\x23\x3b\xd3\xa9\x94\x20\x4a" +
	"\xaa\xc5\x42\xa9\x5b\x31\x45\xc9\x79\x57\x7b\xb7\xcf\x73\x8f\xd0" +
	"\x27\x9e\xc8\x06\x37\x69\xb2\x61\x5e\x31\xd6\xcf\x33\x9f\xf0\x84" +
	"\xfe\x76\x70\xd9\xd9\x6b\xa4\xf8\xf6\x69\x1b\x33\x46\x38\x0c\x3f" +
	"\xcd\x57\x1e\x84\x17\x52\x87\x68\xc4\x67\x38\xb8\x2b\x75\x49\xf9" +
	"\xae\xea\xd4\xd8\x9c\x75\xe2\xf4\x88\xc6\x60\xf9\x49\xb4\x1d\x21" +
	"\x75\xb8\xcc\xf5\xf8\x96\x8a\x50\xea\x67\xb3\xda\x67\x86\xda\xbb" +
	"\xc1\x8a\x45\x12\x2e\xe5\xab\xcc\x59\xf2\x48\x1a\xb1\x84\x7d\x9a" +
	"\xb4\x73\x34\x63\xb8\x7b\x61\xfc\xe5\x28\x23\x40\x8d\xef\xce\x85" +
	"\xdf\xa8\xde\xb4\x91\x69\xf1\xaa\xd9\x21\x21\x3d\xc2\x94\xb0\x2c" +
	"\xe2\xd7\x27\x9c\x2b\x5b\x60\x4d\x84\x46\x84\xf4\x26\x7a\xf9\xb2" +
	"\x9c\x0b\xb5\xf8\xe2\xc4\xf0\xf4\xed\x5d\x50\xcb\x26\x7d\x4b\x99" +
	"\xe7\x8c\xbf\x4d\xf4\xb7\x60\x5b\xd8\xc5\xf8\x89\xa4\xe6\xa5\x7e" +
	"\xe1\xdc\x24\xb0\x7e\x1c\x63\x20\x49\x1a\x26\x37\xe7\xc8\x56\xfd" +
	"\xd1\x36\x44\x7b\x81\xa4\x56\xdf\xf2\xef\x50\x28\x57\xef\xd7\x7f" +
	"\x21\x24\x6d\x83\x93\x73\x59\x96\x87\xb8\x8d\x85\xdc\x17\x4f\xf0" +
	"\x08\x5a\xfe\x4b\x31\x68\x84\x40\x1f\xf1\xf8\x50\x97\xe1\x63\x40" +
	"\xe4\x70\x52\x4e\x71\xe1\x12\xbd\x9c\x40\x90\x72\x35\x73\x3f\x77" +
	"\xa7\x77\x81\xed\xa2\xac\x0e\x2b\xfb\x26\xb8\x33\x4d\xc6\x31\x1a" +
	"\x6b\x50\x7b\x0c\xc7\x03\x0c\xbe\xfd\xad\xcf\xf1\x63\xff\x6e\xf0" +
	"\xe7\x55\x3a\x8d\x48\xad\xac\xf9\xfa\xc1\x83\x83\x7e\x38\xab\x76" +
	"\x61\x1b\x60\xc1\x39\x5b\x87\xf2\x31\xab\xc9\x7c\x09\x30\x90\x0b" +
	"\x91\x0b\xf9\xee\x3a\x2a\x88\xa4\x6d\xcc\xf7\x57\xfc\x8b\x24\x3c" +
	"\x5c\x95\x25\x01\x78\xa7\x6e\xf0\x44\xff\x34\xf0\x5e\x45\xe2\x92" +
	"\x2a\x0d\x01\x76\x00\x94\x45\xb9\xeb\x8b\x2f\x87\xa2\xc6\xec\x18" +
	"\x10\x85\xc3\x97\xa9\xb2\x98\x9f\x14\x69\x59\xd4\x96\xdd\xc4\x08" +
	"\x64\x42\x14\xc2\x40\x00\x7c\x34\x03\x80\x64\x43\x20\x27\x66\xb4" +
	"\x96\x3d\xac\xa0\x65\x4c\x3e\xe8\xa5\xb6\x86\xb3\x68\xcc\x8a\x4e" +
	"\x38\xbb\x14\x90\x90\xc5\xac\xce\xc1\x6f\x18\xbd\x72\xf0\xfe\xbc" +
	"\x06\xc1\x47\xf1\x23\xd6\xf7\x05\x79\x61\xc9\xf6\x40\xce\x56\x56" +
	"\xf0\x6a\x20\x38\xfa\x17\x04\xfc\x7f\x60\xb4\x0e\x38\x3e\xac\xf0" +
	"\xff\x72\x0c\xaf\x8a\xf8\xf7\x6c\x74\xdd\x4c\xaf\x5c\x17\x13\xf9" +
	"\x3f\x01\x00\x00\xff\xff\x98\x71\x93\xdf\xc1\x10\x00\x00"

// Terminal returns the raw, uncompressed file data data.
func dfontimg() *bytes.Buffer {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&_Terminal))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(_Terminal)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var buf bytes.Buffer
	io.Copy(&buf, gz)
	gz.Close()

	return &buf
}

func DefaultFont() *eng.Font {
	if defaultFont == nil {
		defaultFont = eng.NewGridFont(dfontimg(), 16, 16, "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ ✵웃世界¢¥¤§©¨«¬£ª±²³´¶·¸¹º»¼½¾¿☐☑═║╔╗╚╝╠╣╦╩╬░▒▓☺☻☼♀♂▀▁▂▃▄▅▆▇█ÐÑÒÓÔÕÖÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏØÙÚÛÜÝàáâãäåèéêëìíîïðñòóôõö÷ùúûüýÿ♥♦♣♠♪♬æçø←↑→↓↔↕®‼ꀥ")
	}
	return defaultFont
}
