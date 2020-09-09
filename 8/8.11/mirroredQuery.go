package mirroredQuery

var done = make(chan struct{})

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	select {
	case <-done:
		return "abort"
	case <-responses:
		return <-responses
	}
}
func request(hostname string) (response string) { /*......*/ }
