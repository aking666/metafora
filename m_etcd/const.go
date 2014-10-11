package m_etcd

const (
	TasksPath    = "tasks"
	NodesPath    = "nodes"
	CommandsPath = "commands"
	MetadataKey  = "_metafora" // _{KEYs} are hidden files, so this will not trigger our watches
	OwnerMarker  = "owner"

	ClaimTTL   = 120   //seconds
	NodeIDTTL  = 86400 // 24 hours
	ForeverTTL = 0     //Ref: https://github.com/coreos/go-etcd/blob/e10c58ee110f54c2f385ac99764e8a7ca4cb13df/etcd/requests.go#L356

	//Etcd Error codes are passed directly through go-etcd from the http response,
	//So to find the error codes use this ref:
	//       https://github.com/coreos/etcd/blob/master/error/error.go#L67
	EcodeKeyNotFound = 100
)