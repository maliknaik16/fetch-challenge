# Receipt Processor API

A web service that processes receipt data and calculates the number of points awarded based on the provided rules on the request data.

## Prerequisites

- Docker and Docker Compose
- Go 1.23+ (for local development)

## Running server with Docker

1. Clone the repository
2. Run the following command to start the container get the server running:

```bash
docker compose up -d
```

## Running server with Go

1. Clone the repository
2. Run the following command to start the server:

```bash
go build -o main . && ./main
```

Once the server is running it should be accessible at on port `8080` in localhost.

## API Endpoints

The following endpoints are supported:

1. Process Receipts
   * **Path**: `/receipts/process`
   * **Method**: `POST`
   * **Payload**: Receipt JSON
   * **Response**: JSON object with an ID for the processed receipt

2. Get Points

   * **Path**: `/receipts/{id}/points`
   * **Method**: `GET`
   * **Response**: JSON object containing the number of points awarded for the receipt

    
## Points Calculation Rules

These rules collectively define how many points should be awarded to a receipt.

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of 0.25.
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm.

## Testing Endpoints - Postman

This repository includes a `postman.json` file for testing the API endpoints. To get started:

1. Open Postman
2. Click "Import" and select the `postman.json` file
3. You'll see three requests:
   - Process Example 1 (POST)
   - Process Example 2 (POST)
   - Get Points (GET)

Note: The POST requests include post-response scripts that automatically store the returned ID as a Postman variable. 
After making a POST request, you can use the "Get Points" request to view the points associated with that ID.

## Testing using curl

Here is the POST request command to the /receipts/process

```bash
curl --location 'http://localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}'
```

Sample response:
```jquery
{
    "id": "b5a2ee27-4985-494f-901d-308ffc742fee"
}
```

To get the points make a GET request on the /receipts/{id}/points. Replace the ID with the ID returned from the previous request.

```bash
curl --location 'http://localhost:8080/receipts/b5a2ee27-4985-494f-901d-308ffc742fee/points'
```

Sample response:
```json
{
   "points": 28
}
```

Note: These curl commands were exported from Postman directly.
