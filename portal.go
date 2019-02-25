package main
import (
	"bytes"	
	"io"
	"time"
	"net/http"
	"sort"
	"log"

	"github.com/nci/gsky/utils"
	"github.com/CloudyKit/jet"
)

const DefaultTerriaUrl = "http://nationalmap.gov.au"

type GSKYGetCaps struct {
	Name string
	URL string
}

func servePortal(configMap map[string]*utils.Config, w http.ResponseWriter, r *http.Request) {
	assetDir := utils.DataDir + "/static";
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.ServeFile(w, r, assetDir + r.URL.Path)
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

	datasets := make([]*GSKYGetCaps, len(configMap))
	var terriaUrl string
	var terriaInitLayer string

	ids := 0
	for ns, conf := range configMap {
		if len(terriaUrl) == 0 && len(conf.ServiceConfig.TerriaUrl) > 0 {
			terriaUrl = conf.ServiceConfig.TerriaUrl
		}

		if len(terriaInitLayer) == 0 && len(conf.ServiceConfig.TerriaInitLayer) > 0 {
			for _, layer := range conf.Layers {
				if layer.Name == conf.ServiceConfig.TerriaInitLayer {
					terriaInitLayer = ns + "/" + layer.Title
					break
				}
			}

			if len(terriaInitLayer) == 0 {
				log.Printf("portal error: terriaInitLayer '%v' not found", conf.ServiceConfig.TerriaInitLayer)
				http.Error(w, "portal internal error", 500)
				return
			}
		}

		nsUrl := "http://" + conf.ServiceConfig.OWSHostname + "/ows/" + ns
		datasets[ids] = &GSKYGetCaps {Name: ns, URL: nsUrl}	
		ids++
	}	

	if len(terriaUrl) == 0 {
		terriaUrl = DefaultTerriaUrl
	}
	vars.Set("terriaUrl", terriaUrl)

	if len(terriaInitLayer) > 0 {
		vars.Set("terriaInitLayer", terriaInitLayer)
	}

	sort.Slice(datasets, func(i, j int) bool { return datasets[i].Name < datasets[j].Name })
	vars.Set("datasets", datasets)

	if err = template.Execute(&resBuf, vars, nil); err != nil {
		log.Printf("portal error: %v", err)
		http.Error(w, "portal internal error", 500)
		return
	}

	w.Write(resBuf.Bytes())	
}
