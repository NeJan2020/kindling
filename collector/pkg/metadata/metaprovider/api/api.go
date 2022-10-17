package api

type MetaDataService interface {
	List()
	Watch()
	ListAndWatch()
}

type MetaDataRsyncResponse struct {
	Code int
	Msg  string
}

type MetaDataRsyncRequest struct {
	// ...
}
