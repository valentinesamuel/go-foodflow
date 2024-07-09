It seems there was an issue retrieving the content due to an unexpected error related to the branch specification. I'll proceed to retrieve the contents without specifying the branch name explicitly, as it defaults to "main".

Let's try again.

Here is a generated README for the `go-foodflow` repository, based on the retrieved contents:

---

# go-foodflow

`go-foodflow` is a microservices-based food ordering and delivery application written in Go. It provides various services to manage users, orders, deliveries, payments, notifications, menus, reviews, and restaurants.

## Table of Contents
- [Project Overview](#project-overview)
- [Services](#services)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Running the Application](#running-the-application)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Project Overview
`go-foodflow` is designed to be a scalable and efficient solution for managing food orders and deliveries. It is composed of several microservices, each responsible for a specific functionality within the system. The project utilizes Go's powerful concurrency model and the Gin web framework to handle HTTP requests.

## Services
The application consists of the following services:

1. **UserService**: Manages user information and authentication.
2. **OrderService**: Handles order processing and management.
3. **DeliveryService**: Manages the tracking and status of deliveries.
4. **PaymentService**: Handles payment processing and records.
5. **MenuService**: Manages restaurant menus and items.
6. **NotificationService**: Sends notifications to users.
7. **ReviewService**: Manages user reviews for restaurants.
8. **RestaurantService**: Handles restaurant information and management.
9. **GatewayService**: Acts as a gateway for routing requests to appropriate services.
10. **Nginx**: Acts as a reverse proxy for the services.

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Docker
- Docker Compose

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/valentinesamuel/go-foodflow.git
    cd go-foodflow
    ```

2. Build and start the services using Docker Compose:
    ```bash
    docker-compose up --build
    ```

## Running the Application
After the services have been started, you can access the application via the configured ports. Each service runs on a specific port as defined in the `docker-compose.yml` file.

## Environment Variables
Each service requires certain environment variables to be set. These variables are defined in the respective `.env` files within each service directory.

Here are some key environment variables used:
- **UserService**: `USER_DB_NAME`, `PORT`
- **OrderService**: `ORDER_DB_NAME`, `PORT`
- **DeliveryService**: `DELIVERY_DB_NAME`, `PORT`
- **PaymentService**: `PAYMENT_DB_NAME`, `PORT`
- **MenuService**: `MENU_DB_NAME`, `PORT`
- **NotificationService**: `NOTIFICATION_DB_NAME`, `PORT`
- **ReviewService**: `REVIEW_DB_NAME`, `PORT`
- **RestaurantService**: `RESTAURANT_DB_NAME`, `PORT`
- **GatewayService**: `GATEWAY_DB_NAME`, `PORT`

## Contributing
Contributions are welcome! Please open an issue or submit a pull request with your changes. Ensure your code adheres to the project's coding standards and includes appropriate tests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Please note that only the deliveryservice works with nginx, you can replicate the same setup for the rest services