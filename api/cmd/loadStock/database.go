package loadStock

import (
	"app/database"
	"app/enums/actions"
	"app/enums/ratings"
	"app/lib/helper"
	"app/models"
	"app/services/stockRating"

	"gorm.io/gorm/clause"
)

type items = []stockRating.GetAllItem

type ResultStats struct {
	TickersCount         int
	BrokersCount         int
	RecommendationsCount int
}

func processAndSaveData(items items) ResultStats {
	stats := ResultStats{}

	// 1. Precargar tickers y brokers existentes
	tickerMap, brokerMap := preloadEntities()

	// 2. Preparar lotes para inserción
	newTickers, newBrokers := prepareEntitiesBatch(items, tickerMap, brokerMap)

	// 3. Insertar tickers y brokers en lotes
	stats.TickersCount = saveTickers(newTickers)
	stats.BrokersCount = saveBrokers(newBrokers)

	// 4. Actualizar mapas si se agregaron nuevos brokers
	if stats.BrokersCount > 0 {
		_, brokerMap = preloadEntities()
	}

	// 5. Procesar e insertar recomendaciones en lotes
	stats.RecommendationsCount = saveRecommendations(items, brokerMap)

	return stats
}

// preloadEntities carga los tickers y brokers existentes de la base de datos
func preloadEntities() (map[string]bool, map[string]uint) {
	db := database.DB()

	var existingTickers []models.Ticker
	var existingBrokers []models.Broker

	db.Find(&existingTickers)
	db.Find(&existingBrokers)

	// Crear mapas para búsqueda eficiente
	tickerMap := make(map[string]bool)
	brokerMap := make(map[string]uint)

	for _, ticker := range existingTickers {
		tickerMap[ticker.ID] = true
	}

	for _, broker := range existingBrokers {
		brokerMap[broker.Name] = broker.ID
	}

	return tickerMap, brokerMap
}

// prepareEntitiesBatch prepara los lotes de tickers y brokers para inserción
func prepareEntitiesBatch(items items, tickerMap map[string]bool, brokerMap map[string]uint) ([]models.Ticker, []models.Broker) {
	var newTickers []models.Ticker
	var newBrokers []models.Broker

	// Usamos mapas para evitar duplicados en el lote
	uniqueTickerIDs := make(map[string]bool)
	uniqueBrokerNames := make(map[string]bool)

	for _, item := range items {
		// Procesar tickers nuevos
		if !tickerMap[item.Ticker] && !uniqueTickerIDs[item.Ticker] {
			newTickers = append(newTickers, models.Ticker{
				ID:      item.Ticker,
				Company: item.Company,
			})
			uniqueTickerIDs[item.Ticker] = true
		}

		// Procesar brokers nuevos
		if _, exists := brokerMap[item.Brokerage]; !exists && !uniqueBrokerNames[item.Brokerage] {
			newBrokers = append(newBrokers, models.Broker{
				Name: item.Brokerage,
			})
			uniqueBrokerNames[item.Brokerage] = true
		}
	}

	return newTickers, newBrokers
}

// saveTickers guarda los tickers en la base de datos
func saveTickers(tickers []models.Ticker) int {
	if len(tickers) == 0 {
		return 0
	}

	db := database.DB()
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"company"}),
	}).Create(&tickers)

	return int(result.RowsAffected)
}

// saveBrokers guarda los brokers en la base de datos
func saveBrokers(brokers []models.Broker) int {
	if len(brokers) == 0 {
		return 0
	}

	db := database.DB()
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&brokers)

	return int(result.RowsAffected)
}

// saveRecommendations guarda las recomendaciones en la base de datos
func saveRecommendations(items items, brokerMap map[string]uint) int {
	const batchSize = 100
	db := database.DB()
	total := 0

	var batch []models.Recommendation
	for i, item := range items {
		recommendation := models.Recommendation{
			TickerID:   item.Ticker,
			BrokerID:   brokerMap[item.Brokerage],
			Action:     actions.Parse(item.Action),
			TargetFrom: helper.ParseFloat(item.TargetFrom),
			TargetTo:   helper.ParseFloat(item.TargetTo),
			RatingFrom: ratings.Parse(item.RatingFrom),
			RatingTo:   ratings.Parse(item.RatingTo),
			Time:       item.Time,
		}

		batch = append(batch, recommendation)

		// Procesar en lotes o el último lote
		if len(batch) >= batchSize || i == len(items)-1 {
			result := db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "ticker_id"}, {Name: "broker_id"}, {Name: "time"}},
				DoNothing: true,
			}).Create(&batch)

			total += int(result.RowsAffected)
			batch = nil
		}
	}

	return total
}
