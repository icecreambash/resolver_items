package services

import (
	"AinedIndexChessCLI/db_models"
	"AinedIndexChessCLI/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"sync"
)

func GetChesses(db *gorm.DB, indexer *helpers.Indexer, tenantID uuid.UUID, chessCounter *int) {
	var chesses []db_models.Chess
	db.Table("chesses").Find(&chesses)

	chessChunks := helpers.ChessChunkSlice(chesses, 20)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, chessChunk := range chessChunks {
		wg.Add(1)
		go func(chesses []db_models.Chess) {
			defer wg.Done()

			for _, chess := range chesses {
				if err := indexer.IndexData(
					"chesses",
					map[string]interface{}{
						"id":        chess.ID,
						"tenant_id": tenantID.String(),
						"model_id":  chess.ModelID,
						"title":     "Шахматка",
						"entrances": GetEntrancesByChess(chess.ID, db, indexer, tenantID),
						"sections":  GetSections(chess.ID, db, indexer, tenantID),
					}); err != nil {
					mu.Lock()
					log.Fatalf("Error indexing data: %v", err)
					mu.Unlock()
				}
				*chessCounter++
				mu.Lock()
				log.Println("Шахматка ", chess.ID, " занесена в индекс")
				mu.Unlock()
			}

		}(chessChunk)
	}
	wg.Wait()
}

func GetSections(chessID uuid.UUID, db *gorm.DB, indexer *helpers.Indexer, tenantID uuid.UUID) []string {
	var chessSections []db_models.Section
	var chessSectionIds []string
	db.Table("chess_sections").Where("chess_id = ?", chessID).Find(&chessSections)
	for _, section := range chessSections {

		if err := indexer.IndexData(
			"sections",
			map[string]interface{}{
				"id":        section.ID,
				"tenant_id": tenantID.String(),
				"chess_id":  section.ChessID,
				"idx":       section.IDx,
				"title":     section.Title,
				"entrances": GetEntrancesBySection(section.ID, db, indexer, tenantID),
			}); err != nil {
			log.Fatalf("Error indexing data: %v", err)
		}
		chessSectionIds = append(chessSectionIds, section.ID.String())
	}
	return chessSectionIds
}

func GetEntrancesByChess(chessID uuid.UUID, db *gorm.DB, indexer *helpers.Indexer, tenantID uuid.UUID) []string {
	var chessEntrances []db_models.ChessSyncEntrance
	var chessEntranceIds []string
	db.Table("chess_sync_entrances").Where("chess_id = ?", chessID).Find(&chessEntrances)
	for _, chessEntrance := range chessEntrances {
		var entrance db_models.ChessEntrance
		db.Table("chess_entrances").Where("id = ?", chessEntrance.EntranceID).First(&entrance)
		GetFloors(entrance.ID, db, indexer, tenantID)

		viewport := "default"
		if entrance.IsLazy {
			viewport = "lazy"
		}

		if err := indexer.IndexData(
			"entrances",
			map[string]interface{}{
				"id":        entrance.ID,
				"tenant_id": tenantID.String(),
				"idx":       entrance.IDx,
				"is_lazy":   entrance.IsLazy,
				"title":     entrance.Title,
				"viewport":  viewport,
			}); err != nil {
			log.Fatalf("Error indexing data: %v", err)
		}

		chessEntranceIds = append(chessEntranceIds, chessEntrance.EntranceID.String())
	}
	return chessEntranceIds
}

func GetEntrancesBySection(sectionID uuid.UUID, db *gorm.DB, indexer *helpers.Indexer, tenantID uuid.UUID) []string {
	var sectionEntrances []db_models.ChessSyncEntrance
	var sectionEntranceIds []string
	db.Table("chess_section_entrances").Where("section_id = ?", sectionID).Find(&sectionEntrances)
	for _, sectionEntrance := range sectionEntrances {
		var entrance db_models.ChessEntrance
		db.Table("chess_entrances").Where("id = ?", sectionEntrance.EntranceID).First(&entrance)
		GetFloors(entrance.ID, db, indexer, tenantID)

		viewport := "default"
		if entrance.IsLazy {
			viewport = "lazy"
		}

		if err := indexer.IndexData(
			"entrances",
			map[string]interface{}{
				"id":        entrance.ID,
				"tenant_id": tenantID.String(),
				"idx":       entrance.IDx,
				"is_lazy":   entrance.IsLazy,
				"title":     entrance.Title,
				"viewport":  viewport,
			}); err != nil {
			log.Fatalf("Error indexing data: %v", err)
		}

		sectionEntranceIds = append(sectionEntranceIds, sectionEntrance.EntranceID.String())
	}
	return sectionEntranceIds
}

func GetFloors(entranceID uuid.UUID, db *gorm.DB, indexer *helpers.Indexer, tenantID uuid.UUID) []string {
	var floors []db_models.Floor
	db.Table("chess_entrance_floors").Where("entrance_id = ?", entranceID).Find(&floors)
	var entranceFloorIds []string
	floorChunks := helpers.FloorChunkSlice(floors, 5)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, floorChunk := range floorChunks {
		wg.Add(1)
		go func(chunk []db_models.Floor) {
			defer wg.Done()
			for _, floor := range chunk {
				if err := indexer.IndexData(
					"floors",
					map[string]interface{}{
						"id":              floor.ID,
						"tenant_id":       tenantID.String(),
						"entrance_id":     floor.EntranceID,
						"idx":             floor.IDx,
						"is_active_limit": floor.IsActiveLimit,
						"active_limit":    floor.ActiveLimit,
						"items":           GetFloorItems(floor.ID, db, indexer, tenantID),
					}); err != nil {
					mu.Lock()
					log.Fatalf("Error indexing data: %v", err)
					mu.Unlock()
				}

				entranceFloorIds = append(entranceFloorIds, floor.ID.String())
			}
		}(floorChunk)
	}
	wg.Wait()
	return entranceFloorIds
}

func GetFloorItems(floorID uuid.UUID, db *gorm.DB, indexer *helpers.Indexer, tenantID uuid.UUID) []string {
	var items []db_models.FloorItems
	db.Table("chess_entrance_floor_items").Where("floor_id = ?", floorID).Find(&items)
	var floorItemIds []string

	chunks := helpers.FloorItemChunkSlice(items, 3)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []db_models.FloorItems) {
			defer wg.Done()
			for _, item := range chunk {
				if err := indexer.IndexData(
					"floor_items",
					map[string]interface{}{
						"id":        item.ID,
						"tenant_id": tenantID.String(),
						"floor_id":  item.FloorID,
						"idx":       item.IDx,
						"model_id":  item.ModelID,
						"is_ghost":  item.IsGhost,
						"is_vanish": item.IsVanish,
						"slave":     item.Slave,
					}); err != nil {
					log.Fatalf("Error indexing data: %v", err)
				}
				mu.Lock()
				floorItemIds = append(floorItemIds, item.ID.String())
				mu.Unlock()
			}
		}(chunk)
	}
	wg.Wait()
	return floorItemIds
}
