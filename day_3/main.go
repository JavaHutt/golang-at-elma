package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type requestIDCtxKey struct{}

func CounterMiddleware(next http.Handler) http.Handler {
	mutex := new(sync.Mutex)
	var counter int64
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		counter += 1
		requestID := counter
		mutex.Unlock()
		//context
		ctx := r.Context()
		ctx = context.WithValue(ctx, requestIDCtxKey{}, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestIDFromContext(ctx context.Context) int64 {
	id, _ := ctx.Value(requestIDCtxKey{}).(int64)
	return id
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := RequestIDFromContext(ctx)
	sleep := rand.Intn(300)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Fprintf(w, "Request #%d (%d ms)", requestID, sleep)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

var wss = new(sync.Map)

func onWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wss.Store(conn, nil)
	conn.SetCloseHandler(func(int, string) error {
		wss.Delete(conn)
		return nil
	})
}
func broadcast(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	msg := struct {
		Message string `json:"message"`
	}{string(text)}
	wss.Range(func(conn, _ interface{}) bool {
		wsConn := conn.(*websocket.Conn)
		wsConn.WriteJSON(msg)
		return true
	})
}
func main() {
	http.Handle("/hi", CounterMiddleware(http.HandlerFunc(hello)))
	http.HandleFunc("/ws", onWebsocket)
	http.HandleFunc("/send", broadcast)
	fmt.Println("start listen")
	http.ListenAndServe(":1234", nil)
}
