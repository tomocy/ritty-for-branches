package route

import (
	"fmt"
	"path"

	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/route"
)

func MapRoutes() {
	mapRoutes(API, apiRaw, config.Current.Self.Host, config.Current.Self.Port)
	mapRoutes(Web, webRaw, config.Current.Self.Host, config.Current.Self.Port)
	mapRoutes(RittyBranchAuthAPI, rittyBranchAuthAPIRaw, config.Current.RittyBranchAuth.Host, config.Current.RittyBranchAuth.Port)
}

func mapRoutes(routeMap route.RouteMap, rawMap route.RawMap, host, port string) {
	makePathsAbsolute(rawMap, host, port)
	routeMap.Map(rawMap)
}

func makePathsAbsolute(rawMap route.RawMap, host, port string) {
	for key := range rawMap {
		rawMap[key] = makePathAbsolute(host, port, rawMap[key])
	}
}

func makePathAbsolute(host, port, p string) string {
	return fmt.Sprintf("%s:%s%s", host, port, path.Join("/", p))
}

var (
	API                = make(route.RouteMap)
	Web                = make(route.RouteMap)
	RittyBranchAuthAPI = make(route.RouteMap)
)

var (
	apiRaw = route.RawMap{
		"branch.index": "/api/branches",
	}
	webRaw = route.RawMap{
		"menu.index":             "/menus",
		"menu.new":               "/menus/new",
		"menu.create":            "/menus",
		"menu.show":              "/menus",
		"menu.update":            "/menus",
		"menu.delete":            "/menus",
		"authorization_code.new": "/services/ritty/branch/auth/authorization_codes/new",
		"authorization.new":      "/services/ritty/branch/auth/authorizations/new",
	}
	rittyBranchAuthAPIRaw = route.RawMap{
		"authorization.prepare": "/api/authorizations",
		"authorization.create":  "/api/authorizations",
	}
)
