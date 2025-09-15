// tasks/task4-maverics/dev_harness/main.go
package main

import (
	stdctx "context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	seapp "github.com/strata-io/service-extension/app"
	sebundle "github.com/strata-io/service-extension/bundle"
	secache "github.com/strata-io/service-extension/cache"
	sehttp "github.com/strata-io/service-extension/http"
	seid "github.com/strata-io/service-extension/idfabric"
	selog "github.com/strata-io/service-extension/log"
	"github.com/strata-io/service-extension/orchestrator"
	serouter "github.com/strata-io/service-extension/router"
	sesecret "github.com/strata-io/service-extension/secret"
	sesession "github.com/strata-io/service-extension/session"
	setai "github.com/strata-io/service-extension/tai"
	seweblogic "github.com/strata-io/service-extension/weblogic"
)

// ----- minimal logger that satisfies selog.Logger -----
type devLogger struct{}

func (devLogger) Debug(kv ...any) { log.Println(append([]any{"DEBUG"}, kv...)...) }
func (devLogger) Info(kv ...any)  { log.Println(append([]any{"INFO"}, kv...)...) }
func (devLogger) Error(kv ...any) { log.Println(append([]any{"ERROR"}, kv...)...) }

// ----- devAPI: stub that satisfies orchestrator.Orchestrator -----
type devAPI struct{}

func (devAPI) Logger(opts ...selog.Option) selog.Logger { return devLogger{} }
func (devAPI) Session(opts ...sesession.SessionOpt) (sesession.Provider, error) {
	// your SE only checks that Session() doesn't error
	return nil, nil
}

func (devAPI) SecretProvider() (sesecret.Provider, error) { return nil, errors.New("not implemented") }
func (devAPI) IdentityProvider(name string) (seid.IdentityProvider, error) {
	return nil, errors.New("not implemented")
}
func (devAPI) AttributeProvider(name string) (seid.AttributeProvider, error) {
	return nil, errors.New("not implemented")
}
func (devAPI) Metadata() map[string]any                                 { return map[string]any{} }
func (devAPI) Router() serouter.Router                                  { return nil }
func (devAPI) App() (seapp.App, error)                                  { return nil, errors.New("not implemented") }
func (devAPI) TAI() setai.Provider                                      { return nil }
func (devAPI) WebLogic() seweblogic.Provider                            { return nil }
func (devAPI) Context() stdctx.Context                                  { return nil }
func (devAPI) WithContext(ctx stdctx.Context) orchestrator.Orchestrator { return devAPI{} }
func (devAPI) Cache(ns string, opts ...secache.Constraint) (secache.Cache, error) {
	return nil, errors.New("not implemented")
}
func (devAPI) ServiceExtensionAssets() sebundle.SEAssets { return nil }
func (devAPI) HTTP() sehttp.HTTP                         { return nil }

// ----- minimal server using default mux -----
func main() {
	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		hdr, err := CreateEmailHeader(devAPI{}, w, r)
		if err != nil {
			http.Error(w, fmt.Sprintf("CreateEmailHeader error: %v", err), http.StatusBadRequest)
			return
		}

		// Echo the headers so you can see them in the response
		for k, vs := range hdr {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"headers": hdr})
	})

	srv := &http.Server{
		Addr:              "127.0.0.1:8080",
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("dev harness on http://127.0.0.1:8080/headers")
	log.Fatal(srv.ListenAndServe())
}
