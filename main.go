package main

import (
	"html/template"
	"net/http"
	"log"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessiond", 3600)
	go globalSessions.GC()
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("public/count.html")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func main() {
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}