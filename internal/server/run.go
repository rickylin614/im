package server

func RunWeb() func(*SrvCtrl, WebDigIn) {
	return func(ServerRunner *SrvCtrl, in WebDigIn) {
		ServerRunner.Register(&WebServer{In: in})
		ServerRunner.Run()
	}
}

func RunJob() func(*SrvCtrl, JobDigIn) {
	return func(ServerRunner *SrvCtrl, in JobDigIn) {
		ServerRunner.Register(&JobServer{In: in})
		ServerRunner.Run()
	}
}

func RunWs() func(*SrvCtrl, WsDigIn) {
	return func(ServerRunner *SrvCtrl, in WsDigIn) {
		ServerRunner.Register(&WsServer{In: in})
		ServerRunner.Run()
	}
}

func RunAll() func(*SrvCtrl, WebDigIn, JobDigIn, WsDigIn) {
	return func(ServerRunner *SrvCtrl, webIn WebDigIn, jobIn JobDigIn, wsIn WsDigIn) {
		ServerRunner.Register(&WebServer{In: webIn})
		ServerRunner.Register(&JobServer{In: jobIn})
		ServerRunner.Register(&WsServer{In: wsIn})
		ServerRunner.Run()
	}
}
