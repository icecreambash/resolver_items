package main

import (
	"AinedIndexChessCLI/databases"
	"AinedIndexChessCLI/db_models"
	"AinedIndexChessCLI/helpers"
	"AinedIndexChessCLI/services"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

func init() {
	loadEnv()
}

func main() {
	// Ограничиваем количество доступных потоков
	runtime.GOMAXPROCS(2)
	start := time.Now()

	// Подключаемся к центральной базе
	db := databases.GetCon(os.Getenv("DATABASE_MASTER_TABLE"))

	//s3client, err := s3.NewS3Client()
	//if err != nil {
	//	log.Fatalf("Failed to create S3 client: %v", err)
	//}

	// Создание клиента Elasticsearch
	client, err := databases.NewClient(databases.Config{
		Addresses: []string{os.Getenv("ELASTICSEARCH_HOST")},
		Username:  os.Getenv("ELASTICSEARCH_USER"),
		Password:  os.Getenv("ELASTICSEARCH_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}
	// Создание экземпляра Indexer
	indexer := helpers.NewIndexer(client)

	fmt.Println("Starting delete old documents in index")
	if err := indexer.DeleteOldDocuments("chesses"); err != nil {
		log.Fatalf("Error deleting old chesses: %v", err)
	}
	if err := indexer.DeleteOldDocuments("sections"); err != nil {
		log.Fatalf("Error deleting old sections: %v", err)
	}
	if err := indexer.DeleteOldDocuments("entrances"); err != nil {
		log.Fatalf("Error deleting old entrances: %v", err)
	}
	if err := indexer.DeleteOldDocuments("floors"); err != nil {
		log.Fatalf("Error deleting old floors: %v", err)
	}
	if err := indexer.DeleteOldDocuments("floor_items"); err != nil {
		log.Fatalf("Error deleting old floor items: %v", err)
	}
	fmt.Println("Finished delete old documents in index")

	// Получаем все id тенатов
	var tenants []db_models.Tenant
	if err := db.Table("tenants").Find(&tenants).Error; err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	count := 0
	//Делим тенанты по чанкам
	tenantChunks := helpers.TenantChunkSlice(tenants, 10)
	for _, tenantChunk := range tenantChunks {
		wg.Add(1)
		go func(tenants []db_models.Tenant) {
			defer wg.Done()
			for _, tenant := range tenants {
				tenantDB := databases.GetCon(os.Getenv("DATABASE_PREFIX") + tenant.ID.String())
				mu.Lock()
				fmt.Println("Starting collect data in tenant: " + tenant.ID.String())
				mu.Unlock()
				services.GetChesses(tenantDB, indexer, tenant.ID, &count)

				mu.Lock()
				fmt.Println("Finish collect data in tenant: " + tenant.ID.String())
				mu.Unlock()
			}
		}(tenantChunk)
	}
	wg.Wait()

	fmt.Println("Indexed ", count, "chesses in ", time.Since(start))
}
