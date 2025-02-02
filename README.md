# Public API Service

A service for expose endpoint from User and Listing Service

## Prerequisite Requirements

Before running this project, ensure you have the following tools installed:

1. **Golang** (version 1.22.10 or higher) - for building the application.
2. **Git** - for version control.
3. **User Service** already running on your local machine
4. **Listing Service** already running on your local machine

---

## Project Setup and Usage

### 1. Clone the Repository

```bash
git clone https://github.com/okaaryanata/public-api.git
cd public-api
```

### 1. Install Dependencies

```bash
git mod tidy
```

### 2. Create File .env

- create file **.env** base on file **.env.example**

```
go run cmd/public/main.go
```

## API Endpoints

### Service Route

**`{{url}}/public-api`**

### Listing Endpoints

| Method | Endpoint                                      | Description                                                                                                                          |
| ------ | --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `POST` | `/listings`                                   | Create a listing                                                                                                                     |
| `GET`  | `/listings?page_num=1&page_size=10&user_id=5` | Get list of listing, optional: query param **page_num**, **page_size** for pagination and **user_id** for filter specific user by id |

### User Endpoints

| Method | Endpoint | Description                     |
| ------ | -------- | ------------------------------- |
| `POST` | `/users` | Create a user with param `name` |

## [Postman Document Link](https://documenter.getpostman.com/view/7748154/2sAYX3r3Yb#f46501c5-9d5f-45bf-adab-ed78e87e720a)

## Environment Variables

The application uses the following environment variables (defined in the `.env` file):

```plaintext
APP_HOST=localhost
APP_PORT=9191
URL_LISTING="http://0.0.0.0:6000"
URL_USER="http://0.0.0.0:9090"
```

---

## License

This project is licensed under the MIT License.
