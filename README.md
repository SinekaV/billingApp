Supermarket Billing System API
This is a simple API developed in Go for calculating bills at a supermarket. It includes functionality to log in a biller, add products, search for products, and generate bills with GST calculations.

Requirements
Go version 1.18 or higher.
Postman or any API testing tool (optional for testing).
APIs
1. Login API
Endpoint:
POST /login
Request:
{
  "username": "admin",
  "password": "password123"
}
Response:
200 OK: Login successful.
401 Unauthorized: Invalid credentials.
Description:
This API allows the biller to log in with pre-stored credentials. The session for the logged-in user is managed using cookies or an access token.

2. Add Product API
Endpoint:
POST /products
Request:
{
  "code": "P001",
  "name": "Apple",
  "price": 100.00,
  "gst": 18
}
Response:
200 OK: Product added successfully.
400 Bad Request: Invalid input data.
Description:
This API allows the biller to add a new product to the system. It stores the product's code, name, price (unit price), and GST percentage.

3. Search Product API
Endpoint:
GET /products/search
Request:
Query Parameters:
code (optional): Product code to search.
name (optional): Product name to search.
Response:
200 OK: Returns the product details if found.
{
  "code": "P001",
  "name": "Apple",
  "price": 100.00,
  "gst": 18
}
404 Not Found: If the product is not found.
Description:
This API allows searching for a product by its code or name. If found, it returns the product details.

4. Generate Bill API
Endpoint:
POST /bill
Request:
Body (JSON):
{
  "products": [
    {"code": "P001", "quantity": 2},
    {"code": "P002", "quantity": 3}
  ]
}
Response:
200 OK: Returns the total bill details, including tax.
{
  "total": 500.00,
  "tax": 90.00,
  "products": [
    {
      "code": "P001",
      "name": "Apple",
      "quantity": 2,
      "subtotal": 200.00,
      "gst": 36.00
    },
    {
      "code": "P002",
      "name": "Banana",
      "quantity": 3,
      "subtotal": 300.00,
      "gst": 54.00
    }
  ]
}
400 Bad Request: If the input data is invalid or products are not found.

Description:
This API generates the final bill by calculating the total amount based on product prices, quantities, and GST. The response includes the subtotal, tax (GST), and the final total amount.

Project Setup
1. Clone the repository
bash
Copy code
git clone https://github.com/your-username/supermarket-billing-api.git
cd supermarket-billing-api
2. Install Dependencies
This project only requires Go to be installed. Go modules are used for dependency management, so you don't need to install anything else.

bash
Copy code
go mod tidy
3. Run the application
bash
Copy code
go run main.go
The server will start on http://localhost:8080.

4. Test the APIs
You can test the APIs using Postman or any HTTP client.

Login: POST request to /login
Add Product: POST request to /products
Search Product: GET request to /products/search?code=P001
Generate Bill: POST request to /bill
Notes
Make sure to use valid product codes when adding or searching for products.
The Generate Bill API requires an array of products with their quantities.
The bill is calculated by considering the quantity, unit price, and GST of each product.

End of README
This is a straightforward and professional README.md that provides a detailed explanation of the supermarket billing system's API, setup instructions, and how to test the various functionalities. It’s clear and concise, ensuring that developers or users can easily understand and implement the system.