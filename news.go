package main

import "time"

type Observation struct {
    Id              int       `json:"id"`
    PatientId       int       `json:"patient_id"`
    UserId          int       `json:"user_id"`
    AVPU            string    `json:"avpu"`
    HeartRate       int       `json:"heart_rate"`
    RespiratoryRate int       `json:"respiratory_rate"`
    O2Saturation    int       `json:"o2_saturation"`
    O2Supplement    bool      `json:"o2_supplement"`
    Temperature     float64   `json:"temperature"`
    SystolicBP      int       `json:"systolic_bp"`
    Score           int       `json:"score"`
    Risk            string    `json:"risk"`
    Completed       bool      `json:"completed"`
    Due             time.Time `json:"due"`
}

type Observations []Observation

func (o *Observation) calculateRisk() {
  avpu_score := avpuScore(o.AVPU)
  hr_score := heartRateScore(o.HeartRate)
  resp_score := respiratoryRateScore(o.RespiratoryRate)
  temp_score := temperatureScore(o.Temperature)
  syst_score := systolicScore(o.SystolicBP)
  sat_score := o2SaturationScore(o.O2Saturation)
  sup_score := o2SupplementScore(o.O2Supplement)

  o.Score = avpu_score + hr_score + resp_score + temp_score + syst_score +
            sat_score + sup_score
  
  if o.Score >= 0 && o.Score <= 4 {
    switch {
      default: o.Risk = "Low"
      case avpu_score == 3:
        o.Risk = "Medium"
      case hr_score == 3:
        o.Risk = "Medium"
      case resp_score == 3:
        o.Risk = "Medium"
      case temp_score == 3:
        o.Risk = "Medium"
      case syst_score == 3:
        o.Risk = "Medium"
      case sat_score == 3:
        o.Risk = "Medium"
      case sup_score == 3:
        o.Risk = "Medium"
    }
    } else if o.Score >= 5 && o.Score <= 6 {
      o.Risk = "Medium"
    } else {
      o.Risk = "High"
  }
}

func avpuScore(avpu string) int {
  switch avpu  {
  case "A":
    return 0
  case "V", "P", "U":
    return 3
  default:
    panic("Unrecognised AVPU character")
  }
}

func heartRateScore(hr int) int { 
  switch {
  case (hr >=140 && hr <= 50) || (hr >= 91 && hr <= 110): 
    return 1
  case hr >= 51 && hr <= 90:
    return 0
  default:
    return 3
  }
}

func respiratoryRateScore(rr int) int {
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

func temperatureScore(t float64) int {
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

func o2SaturationScore(o2 int) int {
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

func o2SupplementScore(o2 bool) int {
  if o2 == true {
    return 2
  } else {
    return 0
  }
}

func systolicScore(s int) int {
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
