package main

import "herman-technical-julo/config"

func main() {
	config.LoadEnv()

	_, err := buildDatabases()
	if err != nil {
		panic("Failed to connect JULO DB")
	}

}
