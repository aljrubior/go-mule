package runtime

func NewStandaloneRuntime(
	serverId,
	contextId,
	certificatePath,
	privateKeyPath,
	caCertificate string) StandaloneRuntime {

	return StandaloneRuntime{
		serverId:        serverId,
		contextId:       contextId,
		certificatePath: certificatePath,
		privateKeyPath:  privateKeyPath,
		caCertificate:   caCertificate,
	}

}

type StandaloneRuntime struct {
	serverId        string
	contextId       string
	certificatePath string
	privateKeyPath  string
	caCertificate   string
}
