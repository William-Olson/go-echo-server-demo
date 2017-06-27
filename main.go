package main

func main() {

	db := DB{}
	server := Server{db: &db}

	// setup
	db.start()
	server.start()

	defer db.client.Close()

}
