package network

//functions for execution of utilitys.

//run execute commands.
func (a *Airport) run(command, param string) (stdout, stderr string, err error) {
	cmd := exec.Command(command, param)
	cmd.Stdout = &a.stdout
	cmd.Stderr = &a.stderr
	err = cmd.Start()
	if err != nil {
		return
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeDuration):
		err = cmd.Process.Kill()
	case err = <-done:
		stdout = a.stdout.String()
		stderr = a.stderr.String()
	}
	return
}
