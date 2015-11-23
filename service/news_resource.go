package service

import (
  "log"
  "strconv"

  "github.com/wearp/news/api"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
)

type NewsResource struct {
  db gorm.DB
}

func (nr *NewsResource) CreateNews(c *gin.Context) {
  var news api.News

  if err := c.Bind(&news); err != nil {
    c.JSON(400, api.NewError("problem decoding body"))
    return
  }
 
  news.CalculateRisk()
  nr.db.Save(&news)
  c.JSON(201, news)
}

func (nr *NewsResource) GetNews(c *gin.Context) {
  id, err := nr.getId(c)
  if err != nil {
    c.JSON(400, api.NewError("problem decoding id sent"))
    return
  }

  var news api.News
  
  if nr.db.First(&news, id).RecordNotFound() {
    c.JSON(404, gin.H{"error": "not found"})
  } else {
    c.JSON(200, news)
  }
}

func (nr *NewsResource) DeleteNews(c *gin.Context) {
  id, err := nr.getId(c)
  if err != nil {
    c.JSON(400, api.NewError("problem decoding id sent"))
    return
  }

  var news api.News

  if nr.db.First(&news, id).RecordNotFound() {
    c.JSON(404, gin.H{"error": "not found"})
  } else {
    nr.db.Delete(&news)
    c.Data(204, "application/json", make([]byte, 0))
  }
}

func (nr *NewsResource) PutNews(c *gin.Context) {
  id, err := nr.getId(c)
  if err != nil {
    c.JSON(400, api.NewError("problem decoding id sent"))
  }

  var news api.News
  
  if err := c.Bind(&news); err != nil {
    c.JSON(400, api.NewError("problem decoding body"))
    return
  }
  
  news.Id = int32(id)
  news.CalculateRisk()

  var existing api.News

  if nr.db.First(&existing, id).RecordNotFound() {
    c.JSON(404, api.NewError("not found"))
  } else {
    nr.db.Save(&news)
    c.JSON(201, news)
  }
}

func (nr *NewsResource) getId(c *gin.Context) (int32, error) {
  idStr := c.Params.ByName("id")
  id, err := strconv.Atoi(idStr)
  if err != nil {
    log.Print(err)
    return 0, err
  }
  return int32(id), nil
}
