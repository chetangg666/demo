# Demo
#To run the program navigate to demo directory and run below command
# 
docker-compose up --build
# REST API

The REST API to the example app is described below.

## Create a new Record

### Request

`POST /operation`

   {
"LastName":"darshan",
"FirstName":"kumar"
}

### Response

    HTTP/1.1 200 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 Created
    Connection: close
    Content-Type: application/json
   {"ID":1,"LastName":"darshan","FirstName":"kumar"}

## Delete a Record

### Request

`DELETE /operation`

 {
"ID":1
}

### Response

    HTTP/1.1 200 No Content
    Date: Thu, 24 Feb 2011 12:36:32 GMT
    Status: 200 No Content
    Connection: close
    {
    "ID": 1
    }


## Try to delete same Thing again

### Request

`DELETE /operation`

    {
    "ID":1
   }

### Response

    HTTP/1.1 200 Not Found
    Date: Thu, 24 Feb 2011 12:36:32 GMT
    Status: 200 Not Found
    Connection: close
    Content-Type: application/json
    Content-Length: 35

    {
    "ID": 1,
    "err": "Not exist in db to delete"
    }
## Metrics 

### Request

` /metrics`
