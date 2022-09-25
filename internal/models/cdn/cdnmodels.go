package cdnmodels

import (
	"net"
	"net/http"
)

type Connection struct {
	Header http.Header `json:"header"`
	IP     net.IP      `json:"ip"`
}

type Traffic struct {
	Connections []*Connection `json:"connections"`
}
