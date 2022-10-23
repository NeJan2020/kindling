package metadataclient

import (
	"fmt"
	"testing"
)

func TestClient_ListAndWatch(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func() {
			cli := NewMetaDataWrapperClient("http://localhost:9504/listAndWatch")
			err := cli.ListAndWatch()
			if err != nil {
				fmt.Println(err)
			}
		}()
		fmt.Printf("init: %d\n", i)
	}

	select {}
}

// func TestClient_ListAndWatch(t *testing.T) {
// 	cli := NewMetaDataWrapperClient("http://localhost:9504/listAndWatch")
// 	log.Fatal(cli.ListAndWatch())
// }
