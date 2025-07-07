package meta

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/usememos/memos/internal/profile"
)

type MetaService struct {
	Profile *profile.Profile
}

func NewMetaService(profile *profile.Profile) *MetaService {
	return &MetaService{
		Profile: profile,
	}
}

func (s *MetaService) RegisterRoutes(e *echo.Echo) {
	apiSkipper := func(c echo.Context) bool {
		req := c.Request()
		userAgent := req.Header.Get("User-Agent")
		lowerUserAgent := strings.ToLower(userAgent)
		// 检查是否包含 "bot"
		if strings.Contains(lowerUserAgent, "bot") {
			uri := req.RequestURI
			if uri == "/" || strings.HasPrefix(uri, "/memos") {
				return false
			}
			return true
		}
		return true
	}

	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if apiSkipper(c) {
				return next(c)
			}

			req := c.Request()
			return c.Redirect(302, "https://fix-memos.xtaolabs.com/"+s.Profile.InstanceURL+req.RequestURI)
		}
	})
}
