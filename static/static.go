package static

import (
	"embed"
	"io/fs"
	"os"
	"strings"

	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	//go:embed swagger
	EmbedFs embed.FS
	aLog    = log.WithField("module", "swagger")
)

// FileData is the file type
type FileData struct {
	ContentType string
	Content     []byte
}

// GetFileSystem returns the embedded filesystem
func GetFileSystem() http.FileSystem {
	logf := aLog.WithField("fn", "GetFileSystem")

	logf.Print("using embed mode")
	fsys, err := fs.Sub(EmbedFs, "static")
	if err != nil {
		logf.Error(err)
		panic(err)
	}

	return http.FS(fsys)
}

func IsDir(path string) bool {
	for _, s := range GetPathTree("swagger") {
		if s == "[DIR]"+path {
			return true
		}
	}
	return false
}

// GetPathTree reads the directory tree
func GetPathTree(path string) []string {
	logf := aLog.WithField("fn", "GetFileSystem")
	logf.Infof("Into %s", path)

	var entries []os.DirEntry
	var err error
	if strings.HasPrefix(path, "./") {
		entries, err = EmbedFs.ReadDir(path[2:])
	} else {
		entries, err = EmbedFs.ReadDir(path)
	}
	ret := make([]string, 0)
	if err != nil {
		logf.Error("err: ", err)
		return ret
	}
	logf.Infof("Path %s %d entries", path, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			ret = append(ret, "[DIR]"+path+"/"+e.Name())
			ret = append(ret, GetPathTree(path+"/"+e.Name())...)
		} else {
			ret = append(ret, path+"/"+e.Name())
		}
	}
	return ret
}

// GetFile
func GetFile(path string) (*FileData, error) {
	bytes, err := EmbedFs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	mimeType, err := MimeForFileName(path)
	if err != nil {
		return &FileData{
			ContentType: http.DetectContentType(bytes),
			Content:     bytes,
		}, nil
	}
	return &FileData{
		ContentType: mimeType,
		Content:     bytes,
	}, nil
}
