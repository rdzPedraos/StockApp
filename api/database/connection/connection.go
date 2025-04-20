package connection

import (
	"app/config"
	"app/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second * 2, // Umbral de 2 segundos para detectar consultas realmente lentas
			LogLevel:                  logger.Info,     // Mostrar consultas
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Configuración optimizada para CockroachDB
	gormConfig := &gorm.Config{
		Logger:                                   newLogger,
		PrepareStmt:                              true, // Reutilizar consultas preparadas
		DisableForeignKeyConstraintWhenMigrating: true, // Reduce logs durante migraciones
	}

	// Usar la cadena de conexión optimizada
	db, err := gorm.Open(
		postgres.Open(config.DataBase.GetDSN()),
		gormConfig,
	)

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	autoMigrate(db)
	DB = db
}

func autoMigrate(db *gorm.DB) {
	// Silenciar temporalmente los logs durante la migración
	originalLogger := db.Logger
	db.Logger = db.Logger.LogMode(logger.Silent)

	db.AutoMigrate(models.Schemas...)

	// Restaurar el logger original después de la migración
	db.Logger = originalLogger
}
