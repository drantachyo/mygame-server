package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWS(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Разрыв:", err)
            break
        }
        fmt.Println("Сообщение:", string(msg))
        conn.WriteMessage(websocket.TextMessage, msg)
    }
}

func main() {
    http.HandleFunc("/ws", handleWS)
    fmt.Println("Сервер на http://localhost:8080/ws")
    http.ListenAndServe(":8080", nil)
}
