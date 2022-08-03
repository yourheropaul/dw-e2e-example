package frontend

import (
	"bytes"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

type fileFinder struct {
	baseDir string
}

func newFileFinder(relativePath string) *fileFinder {
	_, filename, _, _ := runtime.Caller(0)
	return &fileFinder{
		baseDir: filepath.Join(
			filepath.Dir(filename),
			relativePath,
		)}
}

func (f *fileFinder) HTML(public string) string {
	return filepath.Join(
		f.baseDir,
		public+".html",
	)
}

func Test_ATemplateSiteGeneratoCanRenderTheHomePage(t *testing.T) {
	finder := newFileFinder("templates")
	generator, err := NewTemplateSiteGenerator(finder.HTML("index"))
	require.Nil(t, err)

	buf := bytes.Buffer{}
	generator.Render(&buf)
	require.Greater(t, buf.Len(), 742000)
}
