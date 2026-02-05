// Package manager provides the IPC wrapper for the Manager interface.
// This file defines the Manager interface and types that you need to implement.
package manager

// Account represents a WeChat account.
type Account struct {
	Name        string
	Platform    string
	Version     int
	FullVersion string
	DataDir     string
	Key         string
	ImgKey      string
	PID         uint32
	ExePath     string
	Status      string
}

// KeyData contains decryption key information.
type KeyData struct {
	Key    string
	ImgKey string
}

// Manager defines the interface for WeChat operations.
// TODO: Implement this interface with your business logic.
type Manager interface {
	Run(configPath string) error
	Switch(info *Account, history string) error
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
	GetWeChatInstances() []*Account
	GetKey(configPath string, pid int, force bool, showXorKey bool) (*KeyData, error)
	Decrypt(configPath string, cmdConf map[string]any) error
}
