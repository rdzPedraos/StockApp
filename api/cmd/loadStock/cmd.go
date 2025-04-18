package loadStock

import (
	"app/services/stockRating"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "load-stock",
	Short: "Carga datos de ratings de acciones",
	Long:  `Obtiene y guarda los datos de ratings de acciones desde la API externa.`,
	Run:   loadStock,
}

func loadStock(cmd *cobra.Command, args []string) {
	startTime := time.Now()

	// Paso 1: Obtener los datos de la API
	items, err := stockRating.Service().GetAll()
	if err != nil {
		log.Fatalf("Error al cargar los datos: %v", err)
	}

	size := len(items)
	log.Default().Printf("Se cargaron %d registros de ratings de acciones desde la API\n", size)

	if size == 0 {
		log.Default().Println("No se encontraron datos de ratings de acciones")
		return
	}

	// Paso 2: Cargar los datos a la base de datos
	result := processAndSaveData(items)

	// Paso 3: Mostrar resultados
	elapsedTime := time.Since(startTime)
	log.Default().Printf("Resultados: %d tickers, %d brokers, %d recomendaciones agregadas (tiempo: %v)",
		result.TickersCount, result.BrokersCount, result.RecommendationsCount, elapsedTime)
}
