package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	listenAddr      string
	SnoflakeService *Snowflake
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.Handle("GET /id", handleGetID(s.SnoflakeService))

	log.Printf("[service.idgen] listening for requests on port %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, mux)
}

func handleGetID(s *Snowflake) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := s.GenerateID()

		// Create a byte buffer to store the binary representation
		var buf bytes.Buffer
		// Write the int64 into the buffer in big-endian byte order
		err := binary.Write(&buf, binary.BigEndian, id)
		if err != nil {
			fmt.Println("Error writing int64 to buffer:", err)
		}

		fmt.Fprint(w, base64.StdEncoding.EncodeToString(buf.Bytes()))
	})
}
