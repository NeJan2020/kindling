package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Kindling-project/kindling/collector/pkg/metadata/kubernetes"
	"github.com/Kindling-project/kindling/collector/pkg/metadata/metaprovider/service"
)

func main() {
	authType := flag.String("authType", "serviceAccount", "AuthType describes the type of authentication to use for the K8s API, support 'kubeConfig' or 'serviceAccount'. ")
	kubeConfigPath := flag.String("kubeConfig", "/root/.kube/config", "kubeConfig describe the filePath to your kubeConfig,only used when authType is 'kubeConfig'")
	httpPort := flag.Int("http-port", 9504, "port describe which port will be used to expose data")

	flag.Parse()

	config := &service.Config{
		KubeAuthType:  kubernetes.AuthType(*authType),
		KubeConfigDir: *kubeConfigPath,
	}

	if mdw, err := service.NewMetaDataWrapper(config); err != nil {
		log.Fatalf("create MetaData Wrapper failed, err: %v", err)
	} else {
		http.HandleFunc("/listAndWatch", mdw.ListAndWatch)
		log.Printf("[http] service start at port: %d", *httpPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil))
	}
}
