# ecom
 go lang ecom project
## Description
This is a Go lang e-commerce project aimed at building an online shopping platform.

## Features
- User authentication and authorization
- Product catalog with search and filtering options
- Shopping cart functionality
- Order management
- Payment integration
- User reviews and ratings
- Admin dashboard for managing products, orders, and users

## Technologies Used
- Go programming language
- HTML/CSS for frontend
- PostgreSQL for database
- RESTful API architecture

## Installation
1. Clone the repository: `git clone https://github.com/your-username/ecom.git`
2. Navigate to the project directory: `cd ecom`
3. Install the dependencies: `go mod download`
4. Set up the database: 
    - Create a PostgreSQL database
    - Update the database configuration in the `.env` file
5. Run the application: `go run main.go`

## Contributing
Contributions are welcome! Please follow the guidelines in the [CONTRIBUTING.md](./CONTRIBUTING.md) file.

## License
This project is licensed under the [MIT License](./LICENSE).

## Contact
For any inquiries or support, please contact us at [email@example.com](mailto:email@example.com).

## Docker
docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=0845635040 -e POSTGRES_DB=postgres -p 5433:5432 -d postgres:14
