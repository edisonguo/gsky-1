package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/CloudyKit/jet"
	"github.com/nci/gsky/utils"
)

const DefaultTerriaUrl = "http://nationalmap.gov.au"

type GSKYGetCaps struct {
	Name    string
	Service string
	URL     string
}

func servePortal(configMap map[string]*utils.Config, w http.ResponseWriter, r *http.Request) {
	assetDir := utils.DataDir + "/static"
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.ServeFile(w, r, assetDir+r.URL.Path)
		return
	}

	view := jet.NewSet(jet.SafeWriter(func(w io.Writer, b []byte) {
		w.Write(b)
	}), "/")

	template, err := view.GetTemplate(assetDir + "/index.html")
	if err != nil {
		log.Printf("portal error: %v", err)
		http.Error(w, "portal internal error", 500)
		return
	}

	var resBuf bytes.Buffer
	vars := make(jet.VarMap)

	vars.Set("currentYear", time.Now().Year())

	dsGetCaps := make([]*GSKYGetCaps, 0)

	var terriaUrl string

	var terriaInitGroup string
	var terriaInitLayer string

	for ns, conf := range configMap {
		if len(terriaUrl) == 0 && len(conf.ServiceConfig.TerriaUrl) > 0 {
			terriaUrl = conf.ServiceConfig.TerriaUrl
		}

		if len(terriaInitLayer) == 0 && len(conf.ServiceConfig.TerriaInitLayer) > 0 {
			for _, layer := range conf.Layers {
				if layer.Name == conf.ServiceConfig.TerriaInitLayer {
					terriaInitGroup = ns
					terriaInitLayer = layer.Title
					break
				}
			}

			if len(terriaInitLayer) == 0 {
				log.Printf("portal error: terriaInitLayer '%v' not found", conf.ServiceConfig.TerriaInitLayer)
				http.Error(w, "portal internal error", 500)
				return
			}
		}

		nsUrl := conf.ServiceConfig.OWSProtocol + "://" + conf.ServiceConfig.OWSHostname + "/ows/" + ns

		if len(conf.Layers) > 0 {
			dsGetCaps = append(dsGetCaps, &GSKYGetCaps{Name: ns, Service: "wms-getCapabilities", URL: nsUrl})
		}

		if len(conf.Processes) > 0 {
			dsGetCaps = append(dsGetCaps, &GSKYGetCaps{Name: ns + " geometry drill", Service: "wps-getCapabilities", URL: nsUrl})
		}
	}

	if len(terriaUrl) == 0 {
		terriaUrl = DefaultTerriaUrl
	}
	vars.Set("terriaUrl", terriaUrl)

	if len(terriaInitLayer) > 0 {
		vars.Set("terriaInitLayer", terriaInitLayer)
		vars.Set("terriaInitGroup", terriaInitGroup)
	}

	sort.Slice(dsGetCaps, func(i, j int) bool { return dsGetCaps[i].Name < dsGetCaps[j].Name })
	vars.Set("datasets", dsGetCaps)

	if err = template.Execute(&resBuf, vars, nil); err != nil {
		log.Printf("portal error: %v", err)
		http.Error(w, "portal internal error", 500)
		return
	}

	w.Write(resBuf.Bytes())
}
