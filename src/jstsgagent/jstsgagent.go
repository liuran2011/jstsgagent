package main

import (
	cfg "jstsgagent/tsgcfg"
    ctrl "jstsgagent/ctrl_platform"
)

func main() {
	cfg.ReadConfig()
   
    ctrl.ContInit()
}
