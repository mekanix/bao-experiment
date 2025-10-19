package main

import "log"
import "fmt"
import "context"
import openbao "github.com/openbao/openbao/api/v2"

func main() {
	config := openbao.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := openbao.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize OpenBao client: %v", err)
	}

	client.SetToken("dev-only-token")

	secretData := map[string]interface{}{
		"password": "OpenBao123",
	}

	_, err = client.KVv2("secret").Put(context.Background(), "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	fmt.Println("Secret written successfully.")
	d, err := client.KVv2("secret").Get(context.Background(), "my-secret-password")
	if err != nil {
		log.Fatalf("unable to get secret: %v", err)
	}
	fmt.Println("password:", d.Data["password"])
}
