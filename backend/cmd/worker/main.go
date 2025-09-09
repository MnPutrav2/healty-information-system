package main

import (
	"fmt"
	"time"

	"github.com/MnPutrav2/healty-information-system/backend/internal/config"
	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
)

type Exp struct {
	ID      int
	UserID  string
	Expired string
}

func main() {
	db := config.SqlDb()
	defer db.Close()

	authRepo := repository.NewAuthRepository(db)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println("time now : ", t.Format("2006-01-02 15:04:05"))
		_ = authRepo.CheckUserToken()
	}

}
