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

#### NEWS observation
---
_This will create, fetch or delete a NEWS observation. Upon creation, the NEWS score and clinical risk will be calculated._

* **URL**

  _/news_ or _/news/:id_
  
* **Method**

  `GET` | `POST` | `DELETE`

* **URL Params**
  
  **Required (`GET`, `DELETE`):**
  
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

  * **Code:** 200 (`GET`) <br />
    **Content:**

  * **Code:** 204 (`DELETE`) <br />
    **Content:**

* **Error Response**

  * **Code:** 404 (`GET` | `DELETE`)
    **Content**:

* **Sample Call**

* **Notes**

