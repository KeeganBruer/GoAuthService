package middleware

import (
	"fmt"
	"kbrouter"
	"os"
	"path/filepath"
)

// Global Middleware :: router.AddMiddleware(ServeStaticFiles("../public"))
// Serve static files, maps incoming request urls to a relative path :: /file.txt => ../public/file.txt
func ServeStaticFiles(folderDir string) kbrouter.KBRouteHandler {
	return func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		//build path from base folder and sub-router relative path
		relPath := fmt.Sprintf("%s%s", folderDir, req.CurrPath)

		//relative to absolute path
		absPath, err := filepath.Abs(relPath)
		if err != nil {
			return
		}

		//Check if file exists
		if stats, err := os.Stat(absPath); err != nil || stats.IsDir() {
			return
		}

		//If file exists
		res.SendFile(absPath)
		res.Close() //end the request handling here
	}
}

// Route Handler :: router.AddRoute("GET","/url/endpoint", ServeStaticFile("/path/to/file.html"))
// This is the route handler so it handles error responses
func ServeStaticFile(filePath string) kbrouter.KBRouteHandler {
	return func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		absPath, err := filepath.Abs(filePath)
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting absolute path: %v", err))
			return
		}
		res.SendFile(absPath)
		res.Close() //end the request handling here
	}
}
