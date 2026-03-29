# Transaction Processor

The **Transaction Processor** is an asynchronous worker responsible for the final settlement and fulfillment of financial transactions.

## Responsibility
-   Subscribe to transaction-initiated events via RabbitMQ.
-   Perform balance verification and final updates through the Account Service.
-   Emit fulfillment status events (success/failure) for further downstream consumption.

## Tech Stack
-   **Language**: Go 1.24+
-   **Transport**: AMQP (Subscriber) and gRPC (Internal Client)
-   **Database**: PostgreSQL (for fulfillment state tracking)

## Configuration
-   `RABBITMQ_URL`: Connection parameters for the message broker.
-   `ACCOUNT_SERVICE_URL`: gRPC endpoint for the Account Service.
-   `DATABASE_URL`: Connection string for fulfillment logging.

## Workflow
1.  **Consume**: Listen for new transactions on the message queue.
2.  **Process**: Call Account Service to adjust balances.
3.  **Complete**: Update transaction status in the database and acknowledge the message.

## Local Development
```bash
go run main.go
```
This service requires a running RabbitMQ instance and connectivity to the Account Service.
