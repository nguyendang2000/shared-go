package main

import (
	"os"

	"github.com/nguyendang2000/shared-go/minio"
)

func main() {
	con, _ := minio.NewService(minio.Config{
		Address:   "localhost:9000",
		AccessKey: "Iv4U0ddTKBa7xiF5i5m3",
		SecretKey: "aVXgpnjKT0YeXxORIyyQrkHPPayuoWHkQtzGJjY3",
	})

	// err := con.FPutObject("spreadsheet", "cat.jpeg", "./cat.jpeg", nil)

	// if err != nil {
	// 	panic(err)
	// }

	r, err := con.GetObject("spreadsheet", "cat.jpeg")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("file.jpeg", r, 0644)
	if err != nil {
		panic(err)
	}
}
