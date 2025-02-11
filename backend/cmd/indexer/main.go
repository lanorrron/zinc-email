package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/email/handler"
	"awesomeProject/internal/email/repository"
	"awesomeProject/internal/email/service"
	"awesomeProject/internal/zincsearch"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {

	cpu, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpu)
	defer func ()  {
		pprof.StopCPUProfile()
		cpu.Close()
	
	}()

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var indexName = ""
	fmt.Println("write name index")
	fmt.Scan(&indexName)

	cfg := config.LoadConfig()
	client := zincsearch.NewZincClient(cfg.ZincSearchHost, cfg.ZincSearchUser, cfg.ZincSearchPassword, indexName)
	emailRepo := repository.NewEmailRepository(client)
	emailService := service.NewEmailService(emailRepo)
	emailHandler := handler.NewEmailHandler(emailService)

	emailHandler.IndexEmailToZinc()

	runtime.GC()
	mem, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer mem.Close()

	if err := pprof.WriteHeapProfile(mem); err != nil {
		log.Fatal("failed to write heap profile: ", err)
	}
}
