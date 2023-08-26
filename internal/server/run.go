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
		// Define the servers
		var servers []IServer
	out:
		for _, v := range st {
			switch v {
			case WEB:
				servers = append(servers, &WebServer{In: in})
			case JOB:
				servers = append(servers, &JobServer{In: in})
			case WS:
				// TODO servers = append(servers, &WsServer{In: in})
			case ALL:
				servers = []IServer{&WebServer{In: in}, &JobServer{In: in}}
				break out
			}
		}

		for _, server := range servers {
			in.ServerRunner.Register(server)
		}
		in.ServerRunner.Run()
	}
}
