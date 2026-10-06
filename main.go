package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Konfigurasi WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Mengizinkan koneksi dari mana saja (Flutter)
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade koneksi HTTP biasa menjadi WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrade:", err)
		return
	}
	defer ws.Close()

	fmt.Println("Client terhubung!")

	// Looping untuk menerima dan membalas pesan terus-menerus
	for {
		var msg map[string]string
		// Baca pesan JSON dari client
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Client terputus:", err)
			break
		}

		fmt.Printf("Pesan diterima: %v\n", msg)

		// Kirim balasan langsung back ke client (Echo)
		err = ws.WriteJSON(msg)
		if err != nil {
			fmt.Println("Error kirim pesan:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	fmt.Println("Server Go berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}