package service

import (
  "github.com/wearp/news/api"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

type Config struct {
  SvcHost     string
  DbUser      string
  DbPassword  string
  DbHost      string
  DbName      string
}

type NewService struct {
}

func (s *NewService) getDb(cfg Config) (gorm.DB, error) {
  connectionString := "user=" + cfg.DbUser + " dbname=" + cfg.DbName + " sslmode=disable"

  return gorm.Open("postgres", connectionString)
}

func (s *NewService) Migrate(cfg Config) error {
  db, err := s.getDb(cfg)
  if err != nil {
    return err
  }
  db.SingularTable(true)

  db.AutoMigrate(&api.News{})
  return nil
}

func (s *NewService) Run(cfg Config) error {
  db, err := s.getDb(cfg)
  if err != nil {
    return err
  }
  db.SingularTable(true)
  db.CreateTable(&api.News{})

  newsResource := &NewsResource{db: db}

  r := gin.Default()

  r.POST("/news", newsResource.CalculateNews)
  r.GET("/news/:id", newsResource.GetNews)
  r.DELETE("/news/:id", newsResource.DeleteNews)

  r.Run(cfg.SvcHost)

  return nil
}
