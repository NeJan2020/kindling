package service

import (
	"net/http"

	"github.com/Kindling-project/kindling/collector/pkg/metadata/metaprovider/ioutil"
)

type MetaDataProvider struct {
}

func (mp *MetaDataProvider) List() {
	panic("not implemented") // TODO: Implement
}

func (mp *MetaDataProvider) Watch(w http.ResponseWriter, r *http.Request) {
	f := ioutil.NewWriteFlusher(w)

}

func (mp *MetaDataProvider) ListAndWatch(w http.ResponseWriter, r *http.Request) {
	f := ioutil.NewWriteFlusher(w)

}
