package service

import (
  "log"
  "strings"
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

func (nr *NewsResource) SearchNews(c *gin.Context) {
  var news api.News
  var multipleNews []api.News
 
  params := []string{"patient_id", "spell_id", "location_id", "user_id"}

  for index, param := range params {
    id, err := nr.getParamId(c, param)
    if err != nil {
      c.JSON(400, api.NewError("problem decoding query parameter sent"))
      return
    }
    if id > 0 {
      switch {
      case index == 0:
        news.PatientId = int32(id)
      case index == 1:
        news.SpellId = int32(id)
      case index == 2:
        news.LocationId = int32(id)
      case index == 3:
        news.UserId = int32(id)
      }
    }
  }
  
  if risk := c.Query("risk"); risk != "" {
    risk := strings.Title(risk)
    news.Risk = risk
  }
  
  nr.db.Where(&news).Find(&multipleNews)
  c.JSON(200, multipleNews)
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
    c.JSON(200, news)
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

func (nr *NewsResource) getParamId(c *gin.Context, param string) (int32, error) {
  paramStr := c.Query(param)
  if paramStr != "" {
    id, err := strconv.Atoi(paramStr)
    if err != nil {
      log.Print(err)
      return 0, err
    }
    
    if id != 0 {
      return int32(id), nil
    } 
  }
  return 0, nil
}
