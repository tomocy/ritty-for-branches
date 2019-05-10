package route

import (
	"fmt"
	"path"

	"github.com/tomocy/route"
)

func MapRoutes() {}

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
