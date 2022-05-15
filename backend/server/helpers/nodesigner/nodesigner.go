package nodesigner

type NodeSinger struct {
	clusterSecret string
	nodeSecret    string
}

func New() *NodeSinger {
	return &NodeSinger{
		clusterSecret: "",
		nodeSecret:    "",
	}

}

func (n *NodeSinger) Decrypt(payload string, salt string) (string, error) { return "", nil }
func (n *NodeSinger) Encrypt(payload string, salt string) (string, error) { return "", nil }
func (n *NodeSinger) SignMessage(message string) (string, error)          { return "", nil }
func (n *NodeSinger) VerifyMessage(message string) (string, error)        { return "", nil }
func (n *NodeSinger) GetNodeKey() (string, error)                         { return "", nil }
func (n *NodeSinger) GetClusterKey() (string, error)                      { return "", nil }
func (n *NodeSinger) GetPubP2PKey() (string, error)                       { return "", nil }
