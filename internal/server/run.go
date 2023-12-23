package server

func RunWeb() func(WebDigIn) {
	return func(in WebDigIn) {
		in.ServerRunner.Register(&WebServer{In: in})
		in.ServerRunner.Run(in.Ctx)
	}
}

func RunJob() func(JobDigIn) {
	return func(in JobDigIn) {
		in.ServerRunner.Register(&JobServer{In: in})
		in.ServerRunner.Run(in.Ctx)
	}
}

func RunWs() func(WsDigIn) {
	return func(in WsDigIn) {
		in.ServerRunner.Register(&WsServer{In: in})
		in.ServerRunner.Run(in.Ctx)
	}
}
