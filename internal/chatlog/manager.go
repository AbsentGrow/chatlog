package chatlog

import iwechat "github.com/sjzar/chatlog/internal/wechat"

type Manager interface {
	Run(configPath string) error
	Switch(info *iwechat.Account, history string) error
	StartService() error
	StopService() error
	SetHTTPAddr(text string) error
	GetDataKey() error
	DecryptDBFiles() error
	StartAutoDecrypt() error
	StopAutoDecrypt() error
	RefreshSession() error
	CommandKey(configPath string, pid int, force bool, showXorKey bool) (string, error)
	CommandDecrypt(configPath string, cmdConf map[string]any) error
	CommandHTTPServer(configPath string, cmdConf map[string]any) error
	GetWeChatInstances() []*iwechat.Account

	GetKey(configPath string, pid int, force bool, showXorKey bool) (*KeyData, error)
	Decrypt(configPath string, cmdConf map[string]any) error
}

type MangerType int

const (
	ManagerTypeCli MangerType = iota
	ManagerTypeGRPC
)

func New(mangerType MangerType) Manager {
	switch mangerType {
	case ManagerTypeCli:
		return &CliManager{}
	case ManagerTypeGRPC:
		return &GRPCManager{}
	default:
		return nil
	}
}
