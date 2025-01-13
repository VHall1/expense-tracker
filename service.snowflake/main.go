package main

func main() {
	// TODO: should pull this from .env instead
	srv := NewServer(":8002")

	snow, err := NewSnowflake(int64(0))
	if err != nil {
		panic(err)
	}

	srv.SnoflakeService = snow

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
