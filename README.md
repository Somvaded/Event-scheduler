
# 📅 Event Scheduler API (Go + GORM + JWT + Email)

A simple backend service built in Go for scheduling user-specific events with automatic email reminders. Features include:

- 🧑 User registration & login (with JWT authentication)
- 🗓️ Create, list, and manage personal events
- 📧 Email reminders 24 hours before event time
- 🛡️ Secure routes with token-based access

---

## 🚀 Tech Stack

- **Go (Golang)**
- **Gin** - HTTP web framework
- **GORM** - ORM for SQLite
- **JWT** - Secure authentication
- **Gomail** - SMTP email sender

---

## 🔧 Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/event-scheduler.git
cd event-scheduler


### 2. Install dependencies

```bash
go mod tidy
```

### 3. Configure your SMTP credentials

Edit or create a `.env` file (or modify directly in `utils/email.go`):

```env
PORT="keep empty for local"
HOST_EMAIL = "hostemail@gmail.com"
APP_PASSWORD = "16 digit app password"
```

>  For Gmail, you must enable [App Passwords](https://myaccount.google.com/apppasswords).

### 4. Run the server

```bash
go run main.go
```

By default, it runs on: `http://localhost:8080`

---

## 📬 API Endpoints

### 🔐 Auth

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| POST   | `/register` | Register new user |
| POST   | `/login`    | Login and get JWT |

### 📅 Events (Protected)

| Method | Endpoint  | Description        |
| ------ | --------- | ------------------ |
| GET    | `/events` | List user’s events |
| POST   | `/events` | Create a new event |

#### 🔒 Authorization:

Add the JWT token in request headers:

```
Authorization: <token>
```

---

## 🔁 Reminder Service

The app checks every 1 minute for upcoming events (within 24 hours) and sends reminder emails automatically to the event owner.

---

## 📦 Example JSON Payloads

### Register / Login

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

### Create Event

```json
{
  "title": "Team Sync",
  "description": "Weekly team meeting",
  "time": "2025-05-23T15:30:00Z"
}
```

---

## 🧠 Database Models

### User

```go
type User struct {
    gorm.Model
    Email    string
    Password string
    Events   []Event
}
```

### Event

```go
type Event struct {
    gorm.Model
    Title       string
    Description string
    Time        time.Time
    Reminded    bool
    UserID      uint
}
```


