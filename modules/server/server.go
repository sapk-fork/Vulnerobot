package server

import (
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/linuxisnotunix/Vulnerobot/modules/assets"
	"github.com/linuxisnotunix/Vulnerobot/modules/collectors"
	"github.com/linuxisnotunix/Vulnerobot/modules/settings"
	"github.com/linuxisnotunix/Vulnerobot/modules/tools"
)

//HandlePublic handlerfunc to publish assets
func HandlePublic(res http.ResponseWriter, req *http.Request) {
	log.Debug("public: GET " + req.URL.Path)
	path := strings.TrimPrefix(req.URL.Path, "/public")
	data, hash, contentType, err := assets.Asset("public", path)
	if err != nil {
		/*
		   log.Debug("NOT FOUND :  " + path)
		   data = []byte(err.Error())
		*/
		///*
		data, hash, contentType, err = assets.Asset("public", "/index.html") //TODO better redirect to /public/index.html
		if err != nil {
			data = []byte(err.Error())
		}
		//*/
	}
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", contentType)
	res.Header().Add("Cache-Control", "public, max-age=31536000")
	res.Header().Add("ETag", hash)
	if req.Header.Get("If-None-Match") == hash {
		res.WriteHeader(http.StatusNotModified)
	} else {
		res.WriteHeader(http.StatusOK)
		_, err := res.Write(data)
		if err != nil {
			log.Warn(err)
		}
	}
}

//HandleAPI handlerfunc to publish assets
func HandleAPI(res http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/api")
	q := req.URL.Query()
	log.WithField("query", q).Info("api: GET " + req.URL.Path)

	res.Header().Set("Content-Type", "application/json")
	switch path {
	case "/collect":
		_, err := res.Write([]byte("{error:'not implemented'}"))
		if err != nil {
			log.Warn(err)
		}
	case "/list":
		data, err := ioutil.ReadFile(settings.ConfigPath)
		if err != nil {
			log.Fatalf("Fail to get config file : %v", err)
		}
		cl := collectors.Init(map[string]interface{}{
			"appList":       tools.ParseConfiguration(string(data)),
			"pluginList":    tools.ParseFlagList(q.Get("plugins")),
			"functionList":  tools.ParseFlagList(q.Get("functions")),
			"componentList": tools.ParseFlagList(q.Get("components")),
			/* TODO from args
			"outputFormat":  tools.ParseFlagList(c.String("format")),
			*/
		})
		cl.List(res)
		/*
			_, err := res.Write([]byte("{error:'not implemented'}"))
			if err != nil {
				log.Warn(err)
			}
		*/
	case "/status":
		_, err := res.Write([]byte("{error:'not implemented'}"))
		if err != nil {
			log.Warn(err)
		}
	default:
		_, err := res.Write([]byte("{error:'not found'}"))
		if err != nil {
			log.Warn(err)
		}
	}
}
