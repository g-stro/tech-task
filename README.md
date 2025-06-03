# Technical task

This repository implements a solution for the Software Engineer recruitment scenario, enabling retail customers to invest in ISAs.

---

## Requirements
1. Customers can select a single fund from available options (with future expansion to multiple funds)
2. Customers can specify the amount they wish to invest
3. Investment details are recorded for future retrieval
4. Handle the specific use case of a £25,000 into a Cushon ISA all into the "Cushon Equities Fund".

---

## Assumptions
1. The customer has an existing account, is already logged in and authenticated, and is eligible to make a deposit into a Cushon ISA
2. The customer has not made any contributions to any existing ISAs in the current tax year
3. The service must adhere to the UK ISA allowance of £20,000.00 per year
4. The service would integrate within a much larger system but is out of the scope of the assignment
5. Funds already exist within the system, presumable via another service

---

## Solution
This solution focuses primarily on handling customer investment requests and reporting investment history. While the brief mentioned the customer should be able to select a single fund from a list of available options, I chose not to implement this or any kind of frontend. Instead, I decided to direct focus toward the functionality of the backend.
I have chosen a layered architecture approach following Domain-Driven Design principles. The system separates investment processing from reporting to allow for future decoupling into microservices. 
I have created domain models with appropriate relationships reflecting business domains, and used interfaces to abstract data access, allowing for easier testability and flexibility.

I have chosen PostgreSQL for database management. This is because there is a clear relationship between domains.  
### Architecture
- **API Layer**: REST endpoints for processing and reporting on investments
- **Service Layer**: Containing core business logic:
   - **Investment Service**: Handles processing investment requests, validating against ISA rules
   - **Reporting Service**: Manages data persistence and retrieval
- **Repository Layer**: Data access abstraction for domain models
- **Domain Layer**: Core business entities (Customer, Account, Fund, Investment)

### **Features**
- REST API
- PostgreSQL for database management
- Fully containerised with Docker and Docker Compose
- Unit tests using Go's standard testing package

---

## Domains
1. **Customer**
   - Represents a retail customers
   - Contains personal information (name, email)
   - Can have multiple accounts

2. **Account**
   - Represents a customer account (ISA)
   - Contains a unique account number
   - Has a many-to-one relationship with a Customer
   - Can have multiple investments

3. **Investment**
   - Represents an investment transaction
   - Associates an account with a fund
   - Contains the amount invested, status, and timestamps
   - Requires validation

4. **Fund**
   - A product representing investment options
   - Contains category and risk level
   - Multiple investments can reference a single fund, but may extend to multiple in the future

---

## Example Use Cases

**A customer who wishes to deposit **£25,000** into a Cushon ISA all into the Cushon Equities Fund.**
1. Browse the available funds
2. Select Cushon Equities Fund
3. Enter £25,000 as the investment amount
4. Submit the investment request
5. Request fails as exceeds maximum deposit allowance (£20,000)

**A customer who wishes to deposit **£20,000** into a Cushon ISA all into the Cushon Equities Fund.**
1. Browse the available funds
2. Select Cushon Equities Fund
3. Enter £20,000 as the investment amount
4. Submit the investment request
5. View their investment details

---

## Future Enhancements
In a real-world system, an event-driven, microservices architecture would be highly suitable for processing ISA investments.
When a user makes an investment request, an event would be created and processed asynchronously in the background. This design is both reliable and scalable.
In its current form, this service couples investment processing and reporting too closely. Ideally, this service would be split into two separate microservices; an investment service for processing and a reporting service to persist and retrieve data.

### Investment service
Listen for and process investment request events, responsible for:
- **Validating investment request**:
   - With multiple products (funds), accounts, and regulations, additional services would be a good fit:
     - Rules service
     - Compliance service
- **Processing investments**:
   - Orchestrate processing with other services:
     - Payment service
     - Fulfillment service 
- Send events downstream for reporting once fulfilled

In this scenario, I would consider using a NoSQL database for storing events, offering better performance, scalability, and flexibility compared to RDBMS's.  

### Reporting service
A reporting service would be responsible for persisting data from multiple services and be optimised for retrieval.
- Provide investment reports
- Data and performance analyses
- CQRS for optimisation

### Additional Enhancements
The following are enhancements I would have implemented to my current solution given more time.
- Integration tests
- Structured logging
- Tracing using Go's context package
- Improved error handling
- Standardised HTTP responses

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
DB_SSL_MODE=disable
DB_TIMEZONE=UTC

SERVICE_PORT_TEST=8081
DB_PORT_TEST=5433
```

---

### **2. Run the Application Locally**

#### Using Docker Compose
1. **Build and Start Services**:
   ```bash
   make build
   ```

2. **Access the API**:  
   The service will be available at `http://localhost:8080`.

3. **Stop Docker Containers**:
   ```bash
   make down
   ```

---

## Endpoints:
1. **`POST /investments/`**  
   Creates a new investment  
   Request body:
   ```json
   {
      "account_id": "d7ee4877-7645-461a-b2cc-2f2c8f6a7284",
      "fund_id": "cb91e975-d8bc-423b-bc99-fa6f396c2eaf",
      "amount": 20000.00
   }
   ```
   Example response:
   ```json
   {
      "id": "b1e80bb6-dd4e-4981-839f-b578582c6285"
   }
   ```

2. **`GET /accounts/{id}/investments/`**  
   Retrieves all investments for a specific account  
   Example response:
   ```json
   [
       {
           "id": "22be3b86-8c97-4429-855f-c3478d14230a",
           "amount": 20000,
           "status": "PENDING",
           "created_at": "2025-06-03 15:24:05",
           "updated_at": "2025-06-03 15:24:05",
           "fund": {
               "id": "cb91e975-d8bc-423b-bc99-fa6f396c2eaf",
               "name": "Cushon Equities Fund",
               "category": "EQUITY",
               "currency": "GBP",
               "risk_return": "LOW",
               "created_at": "2025-06-03T15:21:09.802108Z",
               "updated_at": "2025-06-03T15:21:09.802108Z"
           }
       }
   ]
   ```
---

## **Running Tests**
### **Unit Tests**
```bash
make unit-tests
```

### **Integration Tests -- Not implemented**
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
- `make integration-tests`: Run integration tests. -- **Not yet implemented**
- `make tests`: Run all tests.