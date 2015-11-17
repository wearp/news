package service

import (
  "time"
  "strconv"

  "github.com/wearp/news/api"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
)

type NewsResource struct {
  db gorm.DB
}

// POST News and calculate
// PUT News and calculate

func (nr *NewsResource) CalculateNews(c *gin.Context) {
  var news api.News

  c.Bind(&news)
 
  news.Status = api.CompleteStatus
  news.Created = int32(time.Now().Unix())
  
  news.CalculateRisk()

  nr.db.Save(&news)
  c.JSON(201, news)
}

func (nr *NewsResource) GetNews(c *gin.Context) {
  idStr := c.Params.ByName("id")
  idInt, _ := strconv.Atoi(idStr)
  id := int32(idInt)

  var news api.News
  
  if nr.db.First(&news, id).RecordNotFound() {
    c.JSON(404, gin.H{"error": "not found"})
  } else {
    c.JSON(200, news)
  }
}

func (nr *NewsResource) DeleteNews(c *gin.Context) {
  idStr := c.Params.ByName("id")
  idInt, _ := strconv.Atoi(idStr)
  id := int32(idInt)

  var news api.News

  if nr.db.First(&news, id).RecordNotFound() {
    c.JSON(404, gin.H{"error": "not found"})
  } else {
    nr.db.Delete(&news)
    c.Data(204, "application/json", make([]byte, 0))
  }
}
