package api

type News struct {
  Id              int32   `json:"id"`
  PatientId       int32   `json:"patient_id"`
  UserId          int32   `json:"user_id"`
  LocationId      int32   `json:"location_id"`
  AVPU            string  `json:"avpu"`
  HeartRate       int32   `json:"heart_rate"`
  RespiratoryRate int32   `json:"respiratory_rate"`
  O2Saturation    int32   `json:"o2_saturation"`
  O2Supplement    bool    `json:"o2_supplement"`
  Temperature     float32 `json:"temperature"`
  SystolicBP      int32   `json:"systolic_bp"`
  Score           int32   `json:"score"`
  Risk            string  `json:"risk"`
  Status          string  `json:"status"`
  Created         int32   `json:"created"`
  Due             int32   `json:"due"`
}

const (
  ScheduledStatus string = "scheduled"
  ActiveStatus    string = "active"
  CompleteStatus  string = "complete"
)

type NewsMultiple []News

func (n *News) CalculateRisk() {
  avpu_score := avpuScore(n.AVPU)
  hr_score := heartRateScore(n.HeartRate)
  resp_score := respiratoryRateScore(n.RespiratoryRate)
  temp_score := temperatureScore(n.Temperature)
  syst_score := systolicScore(n.SystolicBP)
  sat_score := o2SaturationScore(n.O2Saturation)
  sup_score := o2SupplementScore(n.O2Supplement)

  n.Score = avpu_score + hr_score + resp_score + temp_score + syst_score +
            sat_score + sup_score
  
  if n.Score >= 0 && n.Score <= 4 {
    switch {
      default: n.Risk = "Low"
      case avpu_score == 3:
        n.Risk = "Medium"
      case hr_score == 3:
        n.Risk = "Medium"
      case resp_score == 3:
        n.Risk = "Medium"
      case temp_score == 3:
        n.Risk = "Medium"
      case syst_score == 3:
        n.Risk = "Medium"
      case sat_score == 3:
        n.Risk = "Medium"
      case sup_score == 3:
        n.Risk = "Medium"
    }
    } else if n.Score >= 5 && n.Score <= 6 {
      n.Risk = "Medium"
    } else {
      n.Risk = "High"
  }
}

func avpuScore(avpu string) int32 {
  switch avpu  {
  case "A":
    return 0
  case "V", "P", "U":
    return 3
  default:
    panic("Unrecognised AVPU character")
  }
}

func heartRateScore(hr int32) int32 { 
  switch {
  case (hr >=140 && hr <= 50) || (hr >= 91 && hr <= 110): 
    return 1
  case hr >= 51 && hr <= 90:
    return 0
  default:
    return 3
  }
}

func respiratoryRateScore(rr int32) int32 {
  switch {
  case rr >= 9 && rr <= 11:
    return 1
  case rr >= 12 && rr <= 20:
    return 0
  case rr >= 10 && rr <= 24:
    return 2
  default:
    return 3
  }
}

func temperatureScore(t float32) int32 {
  switch {
  case t <= 35.0:
    return 3
  case (t >= 35.1 && t <= 36.0) || (t >= 38.1 && t <=39.0): 
    return 1
  case t >= 36.1 && t <= 38.0:
    return 0
  default:
    return 2
  }
}

func o2SaturationScore(o2 int32) int32 {
  switch {
  case o2 <= 91:
    return 3
  case o2 >= 92 && o2 <= 93:
    return 2
  case o2 >= 94 && o2 <= 95:
    return 1
  default:
    return 0
  }
}

func o2SupplementScore(o2 bool) int32 {
  if o2 == true {
    return 2
  } else {
    return 0
  }
}

func systolicScore(s int32) int32 {
  switch {
  case s >= 91 && s <= 100:
    return 2
  case s >= 101 && s <= 110:
    return 1
  case s >= 111 && s <= 219:
    return 0
  default:
    return 3
  }
}

