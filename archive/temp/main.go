package plugins

import (
	"encoding/json"
	"net/http"

	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
)

func init() {
	err := plugin.RegisterPlugin(&Suraj{})
	if err != nil {
		log.Fatalf("failed to register plugin say: %s", err)
	}
}

// Say is a demo to show how to return data directly instead of proxying
// it to the upstream.
type Suraj struct {
	// Embed the default plugin here,
	// so that we don't need to reimplement all the methods.
	plugin.DefaultPlugin
}

type SurajConf struct {
	Body string `json:"body"`
}

func (p *Suraj) Name() string {
	return "suraj"
}

func (p *Suraj) ParseConf(in []byte) (interface{}, error) {
	conf := SurajConf{}
	err := json.Unmarshal(in, &conf)
	return conf, err
}

func (p *Suraj) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	body := conf.(SurajConf).Body
	if len(body) == 0 {
		return
	}

	w.Header().Add("X-Resp-A6-Runner", "Suraj")
	_, err := w.Write([]byte(r.URL.Path[1:]))
	if err != nil {
		log.Errorf("failed to write: %s", err)
	}
}
