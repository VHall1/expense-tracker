package main

func main() {
	// TODO: pull port from config
	srv := NewServer(":8002")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
