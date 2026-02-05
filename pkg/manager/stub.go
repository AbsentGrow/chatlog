// Package manager provides stub implementations for the Manager interface.
// TODO: Replace this with your actual implementation.
package manager

import "errors"

// StubManager is a stub implementation of the Manager interface.
// Replace this with your actual implementation.
type StubManager struct {
	// Add your fields here
}

// NewStubManager creates a new StubManager.
func NewStubManager() *StubManager {
	return &StubManager{}
}

// Ensure StubManager implements Manager interface
var _ Manager = (*StubManager)(nil)

func (m *StubManager) Run(configPath string) error {
	// TODO: Implement
	return errors.New("not implemented: Run")
}

func (m *StubManager) Switch(info *Account, history string) error {
	// TODO: Implement
	return errors.New("not implemented: Switch")
}

func (m *StubManager) StartService() error {
	// TODO: Implement
	return errors.New("not implemented: StartService")
}

func (m *StubManager) StopService() error {
	// TODO: Implement
	return errors.New("not implemented: StopService")
}

func (m *StubManager) SetHTTPAddr(text string) error {
	// TODO: Implement
	return errors.New("not implemented: SetHTTPAddr")
}

func (m *StubManager) GetDataKey() error {
	// TODO: Implement
	return errors.New("not implemented: GetDataKey")
}

func (m *StubManager) DecryptDBFiles() error {
	// TODO: Implement
	return errors.New("not implemented: DecryptDBFiles")
}

func (m *StubManager) StartAutoDecrypt() error {
	// TODO: Implement
	return errors.New("not implemented: StartAutoDecrypt")
}

func (m *StubManager) StopAutoDecrypt() error {
	// TODO: Implement
	return errors.New("not implemented: StopAutoDecrypt")
}

func (m *StubManager) RefreshSession() error {
	// TODO: Implement
	return errors.New("not implemented: RefreshSession")
}

func (m *StubManager) CommandKey(configPath string, pid int, force bool, showXorKey bool) (string, error) {
	// TODO: Implement
	return "", errors.New("not implemented: CommandKey")
}

func (m *StubManager) CommandDecrypt(configPath string, cmdConf map[string]any) error {
	// TODO: Implement
	return errors.New("not implemented: CommandDecrypt")
}

func (m *StubManager) CommandHTTPServer(configPath string, cmdConf map[string]any) error {
	// TODO: Implement
	return errors.New("not implemented: CommandHTTPServer")
}

func (m *StubManager) GetWeChatInstances() []*Account {
	// TODO: Implement
	return nil
}

func (m *StubManager) GetKey(configPath string, pid int, force bool, showXorKey bool) (*KeyData, error) {
	// TODO: Implement
	return nil, errors.New("not implemented: GetKey")
}

func (m *StubManager) Decrypt(configPath string, cmdConf map[string]any) error {
	// TODO: Implement
	return errors.New("not implemented: Decrypt")
}
