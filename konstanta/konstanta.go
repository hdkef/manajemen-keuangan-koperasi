package konstanta

var homeroute = "/"
var HomeHTML = "home.html"

type route struct {
}

func GetRoute() *route {
	return &route{}
}

func (r *route) Home() string {
	return homeroute
}
