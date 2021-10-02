package main

import (
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessiond", 3600)
	go globalSessions.GC()
}

func main() {
	
}