package route

import (
	"fmt"
	"path"

	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/route"
)

func MapRoutes() {
	mapRoutes(Web, webRaw, config.Current.Self.Host, config.Current.Self.Port)
	mapRoutes(BranchAuthAPI, branchAuthAPIRaw, config.Current.BranchAuth.Host, config.Current.BranchAuth.Port)
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
	Web           = make(route.RouteMap)
	BranchAuthAPI = make(route.RouteMap)
)

var (
	webRaw = route.RawMap{
		"menu.new": "/menus/new",
	}
	branchAuthAPIRaw = route.RawMap{}
)
