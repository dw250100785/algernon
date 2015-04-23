package main

import (
	"github.com/xyproto/permissions2"
	"github.com/yuin/gopher-lua"
	"net/http"
	"path"
)

// Expose functions for serving other files to Lua
func exportServeFile(w http.ResponseWriter, req *http.Request, L *lua.LState, filename string, perm *permissions.Permissions, luapool *lStatePool) {

	// Serve a file in the scriptdir
	L.SetGlobal("servefile", L.NewFunction(func(L *lua.LState) int {
		scriptdir := path.Dir(filename)
		serveFilename := path.Join(scriptdir, L.ToString(1))
		if exists(serveFilename) {
			filePage(w, req, serveFilename, perm, luapool)
		}
		return 0
	}))

}