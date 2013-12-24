package appcache

import (
    "fmt"
    "net/http"
    "os"
    "regexp"
    "time"
)

func init() {
    setAppVersion()
    http.HandleFunc("/manifest.appcache", handler)
}

var appVersion = ""
func setAppVersion() {
    matcher := regexp.MustCompile("^.*/([\\d\\.]+)/$")
    match := matcher.FindStringSubmatch(os.Getenv("PWD"))
    if match != nil {
        appVersion = "v"+match[1]
    }
}

func getLastModifiedTimestamp() string {
    stat, _ := os.Lstat("static")
    if stat != nil {
        return fmt.Sprintf("%s", stat.ModTime())
    } else {
        return fmt.Sprintf("%s", time.Now())
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/cache-manifest")
    w.Header().Set("Cache-Control", "no-cache")
    version := appVersion
    if len(version) == 0 {
        version = getLastModifiedTimestamp()
    }
    fmt.Fprint(w, "CACHE MANIFEST\nNETWORK:\n*\n# ", version)
}
