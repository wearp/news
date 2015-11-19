## NEWS
A microservice for the NHS National Early Warning Score (NEWS)

### Installation

#### Build the Service
      
      make build

#### Build the Database (PostgreSQL)
   
      ./cli/server/server --config config.yaml migratedb

Before building the database, you will need to create one that corresponds to `dbuser`, `dbname`, etc. as defined in `config.yaml`.
    
#### Run the Service

      ./cli/server/server --config config.yaml server

### HTTP API
Documents the HTTP API for the service.

#### `GET` | `POST` | `DELETE` /news
---
_This will create, fetch or delete a NEWS observation. Upon creation, the NEWS score and clinical risk will be calculated._

* **URL**

  _/news_ or _/news/:id_
  
* **Methods**

  `GET` | `POST` | `DELETE`

* **URL Params**
  
  * **Required (`GET`, `DELETE`):**
  
    `id=[integer]`

* **Data Params**

  Example:
  
  ```
  {
      "patient_id": [integer],
      "location_id": [integer],
      "user_id": [integer],
      "avpu": [string],
      "heart_rate": [integer],
      "respiratory_rate": [integer],
      "o2_saturation": [integer],
      "o2_supplement": [bool],
      "temperature": [float],
      "systolic_bp": [integer]
  }
  ```

* **Success Response**

  * **Code:** 201 (`POST`) <br />
    **Content:**
    
    ```
    {"id":1,"patient_id":3,"user_id":4,"location_id":1,"avpu":"A","heart_rate":60,"respiratory_rate":12,"o2_saturation":97,"o2_supplement":false,"temperature":36.1,"systolic_bp":0,"score":3,"risk":"Medium","status":"complete","created":1447957870,"due":0}
    ```

  * **Code:** 200 (`GET`) <br />
    **Content:**
    
    ```
    {"id":4,"patient_id":3,"user_id":4,"location_id":1,"avpu":"A","heart_rate":60,"respiratory_rate":12,"o2_saturation":97,"o2_supplement":false,"temperature":36.1,"systolic_bp":0,"score":3,"risk":"Medium","status":"complete","created":1447959927,"due":0}
    ```
  
  * **Code:** 204 (`DELETE`) 

* **Error Response**

  * **Code:** 404 (`GET` | `DELETE`) <br />
    **Content**:
    
    ```
    {"error": "not found"}
    ```

* **Sample Call**

  ```
  curl -H "Content-Type: application/json" -X POST -d '{"location_id": 1, "patient_id": 3, "user_id": 4, "avpu": "A", "heart_rate": 60, "respiratory_rate": 12, "o2_saturation": 97, "temperature": 36.1, "o2_supplement": false}' http://localhost:8080/news
  ```

* **Notes**

