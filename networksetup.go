package network

//NetWorkSetup is struct base for utility networksetup
type NetWorkSetup struct{}

//NewNetWorkSetup return new instance of networksetup utility.
func NewNetWorkSetup() *NetWorkSetup {
	return &NetWorkSetup{}
}

//ListNetWorkServiceOrder
func (n *NetWorkSetup) ListNetWorkServiceOrder() (string, error) {
	stdout, _, err := run(cmdNetWorkSetup, "-listnetworkserviceorder")
	if err != nil {
		return "", err
	}
	return stdout, nil
}

//ListAllNetWorkServices
func (n *NetWorkSetup) ListAllNetWorkServices() (string, error) {
	stdout, _, err := run(cmdNetWorkSetup, "-listallnetworkservices")
	if err != nil {
		return "", err
	}
	return stdout, nil
}

//ListAllHardwarePorts
func (n *NetWorkSetup) ListAllHardwarePorts() (string, error) {
	stdout, _, err := run(cmdNetWorkSetup, "-listallhardwareports")
	if err != nil {
		return "", err
	}
	return stdout, nil
}
