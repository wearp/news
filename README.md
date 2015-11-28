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

#### `GET` | `POST` | `PUT` | `DELETE` /news
---
_This will fetch, create, replace or delete a NEWS observation. Upon creation, the NEWS score and clinical risk will be calculated alongwith the time of calculation._

* **URL**

  _/news_ or _/news/:id_
  
* **Methods**

  `GET` | `POST` | `DELETE` | `PUT`

* **URL Params**
  
  * **Required (`GET`, `DELETE`):**
  
    `id=[integer]`

* **Data Params**

  Example:
  
  ```
  {
      "patient_id": [integer],
      "location_id": [integer],
      "spell_id": [integer],
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
  
  * **Code:** 200 (`GET` | `PUT`) <br />
    **Content:**
    
    ```
    {"id":6,"patient_id":3,"spell_id":5,"user_id":4,"location_id":1,"avpu":"A","heart_rate":60,"respiratory_rate":12,"o2_saturation":97,"o2_supplement":false,"temperature":20.1,"systolic_bp":30,"score":6,"risk":"Medium","completed":1448439536}
   ```
 
  * **Code:** 201 (`POST`) <br />
    **Content:**
    
    ```
    {"id":6,"patient_id":3,"spell_id":5,"user_id":4,"location_id":1,"avpu":"A","heart_rate":60,"respiratory_rate":12,"o2_saturation":97,"o2_supplement":false,"temperature":20.1,"systolic_bp":30,"score":6,"risk":"Medium","completed":1448439536}
    ```
 
  * **Code:** 204 (`DELETE`)  

* **Error Response**
  
  * **Code:** 400 (`GET` | `DELETE` | `PUT`) <br />
    **Content**:
    
    ```
    {"error": "problem decoding query parameter sent"}
    ```
  
  * **Code:** 400 (`PUT` | `POST`) <br />
    **Content**:
    
    ```
    {"error": "problem decoding body"} 
    ```

  * **Code:** 404 (`GET` | `DELETE` | `PUT`) <br />
    **Content**:
    
    ```
    {"error": "not found"}
    ```

* **Sample Call**

  ```
  curl -H "Content-Type: application/json" -X POST -d '{"location_id": 1, "patient_id": 3, "user_id": 4, "spell_id": 5, "avpu": "A", "heart_rate": 60, "respiratory_rate": 12, "o2_saturation": 97, "temperature": 36.1, "o2_supplement": false}' http://localhost:8080/news
  ```

* **Notes**

#### ``GET`` /news?
---
_Fetches a list of NEWS observations based on given querystring parameters. Returns an empty list if no News observations match the parameters provided._

* **URL**

  _/news?_
  
* **Methods**

  `GET` 

* **URL Params**
  
  * **Optional:**
  
    `risk=[string]` <br />
    `patient_id=[integer]` <br />
    `spell_id=[integer]` <br />
    `location_id=[integer]` <br />
    `user_id=[integer]`

* **Success Response**

  * **Code:** 200 (`GET`) <br />
    **Content:**
    
    ```
    [{"id":6,"patient_id":3,"spell_id":5,"user_id":4,"location_id":1,"avpu":"A","heart_rate":60,"respiratory_rate":12,"o2_saturation":97,"o2_supplement":false,"temperature":20.1,"systolic_bp":30,"score":6,"risk":"Medium","completed":1448439536},{"id":7,"patient_id":3,"spell_id":5,"user_id":4,"location_id":1,"avpu":"A","heart_rate":60,"respiratory_rate":12,"o2_saturation":97,"o2_supplement":false,"temperature":20.1,"systolic_bp":30,"score":6,"risk":"Medium","completed":1448439539}]
   ```

* **Error Response**
  
  * **Code:** 400 (`GET`) <br />
    **Content**:
    
    ```
    {"error": "problem decoding query parameter sent"}
    ```

* **Sample Call**

  ```

  curl -X GET http://localhost:8080/news?risk=medium&spell_id=3&location_id=1
  ```
