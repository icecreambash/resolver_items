package helpers

import "AinedIndexChessCLI/db_models"

type Type int

func MapChunkSlice(slice []map[string]interface{}, chunkSize int) [][]map[string]interface{} {
	var chunks [][]map[string]interface{}
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func InterfChunkSlice(slice []interface{}, chunkSize int) [][]interface{} {
	var chunks [][]interface{}
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func ItemChunkSlice(slice []db_models.Item, chunkSize int) [][]db_models.Item {
	var chunks [][]db_models.Item
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func TenantChunkSlice(slice []db_models.Tenant, chunkSize int) [][]db_models.Tenant {
	var chunks [][]db_models.Tenant
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func ChessChunkSlice(slice []db_models.Chess, chunkSize int) [][]db_models.Chess {
	var chunks [][]db_models.Chess
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func FloorChunkSlice(slice []db_models.Floor, chunkSize int) [][]db_models.Floor {
	var chunks [][]db_models.Floor
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func FloorItemChunkSlice(slice []db_models.FloorItems, chunkSize int) [][]db_models.FloorItems {
	var chunks [][]db_models.FloorItems
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
