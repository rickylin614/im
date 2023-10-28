package server

// Kind represents type of server
type Kind string

const (
	WEB Kind = "WebServer"
	JOB      = "JobServer"
	WS       = "WsServer"
	ALL      = "All"
)

func Run(st ...Kind) func(digIn) {
	return func(in digIn) {
	out:
		for _, v := range st {
			switch v {
			case WEB:
				in.ServerRunner.Register(&WebServer{In: in})
			case JOB:
				in.ServerRunner.Register(&JobServer{In: in})
			case WS:
				// TODO servers = append(servers, &WsServer{In: in})
			case ALL:
				in.ServerRunner.Register(&WebServer{In: in})
				in.ServerRunner.Register(&JobServer{In: in})
				break out
			}
		}

		in.ServerRunner.Run()
	}
}
