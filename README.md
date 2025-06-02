# Technical task

This repository implements a solution for the Software Engineer recruitment scenario, enabling retail customers to invest in ISAs.

---

## Requirements
1. Allow customers to select a single fund from available options (with future expansion to multiple funds)
2. Allow customers to specify the amount they wish to invest
3. Record these values and allow for later retrieval
4. Handle the specific use case of a £25,000 into a Cushon ISA all into the "Cushon Equities Fund".

---

## Assumptions
1. The customer has an existing account, is already logged in and authenticated, and is eligible to make a deposit into a Cushon ISA
2. The customer has not made any contributions to any existing ISAs
3. The service must adhere to the UK ISA allowance of £20,000.00 per year
4. The service would integrate within a much larger system but is out of the scope of the assignment
5. Funds already exist within the system, presumable via another service

---

## Example Use Cases

### 1. A customer who wishes to deposit £25,000 into a Cushon ISA all into the Cushon Equities Fund.
1. Browse the available funds
2. Select Cushon Equities Fund
3. Enter £25,000 as the investment amount
4. Submit the investment request
5. Request fails as exceeds maximum deposit allowance (£20,000)

### 2. A customer who wishes to deposit £20,000 into a Cushon ISA all into the Cushon Equities Fund.
1. Browse the available funds
2. Select Cushon Equities Fund
3. Enter £20,000 as the investment amount
4. Submit the investment request
5. View their investment details

---

## Domains
// TODO
1. Customer
    - Customer details
2. Account (ISA)
    - Account details
    -
3. Investment
    - Amount
    - Transaction history
4. Fund
    - Fund details
    - Available options

---

## Future Enhancements
// TODO
In a real-world system, an event-driven, microservices architecture would be highly suitable for processing ISA investments.
When a user makes an investment request, an event would be created and processed asynchronously in the background. This design is both reliable and scalable.
In its current form, this service couples investment processing and reporting too closely. Ideally, this service would be split into two separate microservices; an investment service for processing and a reporting service to persist and retrieve data.

### Investment service
Listen for and handle investment request events, primarily focusing on:
1. validating the request
2. Processing the investment
   - The investment service would likely orchestrate processing with other services, such as Payment and Fulfillment services. 
3. Send downstream for reporting

### Reporting service
A reporting service would be responsible for persisting data from multiple services and be optimised for retrieval. 

Given more time and resources, I would implement these services.

---

## **Features**
// TODO
- REST API:
    - `GET /customer`: ...
- PostgreSQL for database management.
- Fully containerised with Docker and Docker Compose.
- Unit and integration tests using Go's standard testing package.

---

## **Setup**

### **Prerequisites**
- [Go 1.24+](https://go.dev/dl/) installed
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL database (bundled with Docker Compose)

---

### **1. Setup Environment**
Modify the .env.example if required.

Example environment variables:
```dotenv
SERVICE_PORT=8080

DB_USERNAME=test_user
DB_PASSWORD=test_password
DB_NAME=test_db
DB_HOST=postgres
DB_PORT=5432
```

---

### **2. Run the Application Locally**

#### Using Docker Compose
1. **Build and Start Services**:
   ```bash
   make build
   ```
2. **Run the Services**:
   ```bash
   make up
   ```

3. **Access the API**:  
   The service will be available at `http://localhost:8080`.

4. **Stop Docker Containers**:
   ```bash
   make down
   ```

---

## **Using the API**
### Endpoints:
// TODO
1. **`GET /customer`**  
   ...  
   Example response:
   ```json
   {
     "status": "success",
     "data": {
       "customer": [
         {
           "id": "uuid"
         }
       ]
     }
   }
   ```

---

## **Running Tests**

### **Unit Tests**
```bash
make unit-tests
```

### **Integration Tests**
Integration tests require both the service and the database to be running. If using Docker Compose:
```bash
make integration-tests
```

### **All Tests**
```bash
make tests
```

---

## **Makefile Commands**
### **Basic Commands**:
- `make help`: Display available `make` targets.
- `make build`: Build and start Docker containers.
- `make up`: Start the service containers.
- `make clean`: Remove containers, images, volumes, and orphans for a clean environment.
- `make restart`: Restart the entire service stack.
- `make logs`: Tail logs from the running containers.

### **Testing Commands**:
- `make unit-tests`: Run unit tests.
- `make integration-tests`: Run integration tests.
- `make tests`: Run all tests.