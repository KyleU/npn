// Content managed by Project Forge, see [projectforge.md] for details.
package scripting

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/lib/search/result"
	"github.com/kyleu/npn/app/util"
)

type Service struct {
	FS   filesystem.FileLoader `json:"-"`
	Path string                `json:"path,omitempty"`
}

func NewService(fs filesystem.FileLoader, pth string) *Service {
	return &Service{FS: fs, Path: pth}
}

func (s *Service) ListScripts(logger util.Logger) []string {
	files, _ := s.FS.ListFilesRecursive(s.Path, nil, logger)
	return lo.FilterMap(files, func(x string, _ int) (string, bool) {
		return x, strings.HasSuffix(x, ".js")
	})
}

func (s *Service) ListScriptSizes(logger util.Logger) ([]string, map[string]int) {
	files := s.ListScripts(logger)
	return files, lo.SliceToMap(files, func(scr string) (string, int) {
		return scr, s.Size(scr)
	})
}

func (s *Service) LoadScript(pth string, logger util.Logger) (string, error) {
	logger.Infof("loading script [%s]", pth)
	filePath := filepath.Join(s.Path, pth)
	b, err := s.FS.ReadFile(filePath)
	if err != nil {
		b, err = s.FS.ReadFile(filePath + ".js")
		if err != nil {
			return "", err
		}
	}
	sc := string(b)
	return sc, nil
}

func (s *Service) SaveScript(pth string, content string, logger util.Logger) error {
	logger.Infof("saving script [%s]", pth)
	filePath := filepath.Join(s.Path, pth)
	return s.FS.WriteFile(filePath, []byte(content), filesystem.DefaultMode, true)
}

func (s *Service) DeleteScript(pth string, logger util.Logger) error {
	filePath := filepath.Join(s.Path, pth)
	return s.FS.Remove(filePath, logger)
}

func (s *Service) Size(scr string) int {
	filePath := filepath.Join(s.Path, scr)
	st, err := s.FS.Stat(filePath)
	if err != nil {
		return 0
	}
	return int(st.Size)
}

func (s *Service) SearchScripts(ctx context.Context, ps filter.ParamSet, q string, logger util.Logger) (result.Results, error) {
	return lo.FilterMap(s.ListScripts(logger), func(fn string, _ int) (*result.Result, bool) {
		scr, _ := s.LoadScript(fn, logger)
		res := result.NewResult("script", fn, "/admin/scripting/"+fn, fn, "file-code", scr, scr, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}
