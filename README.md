# ✅ Task Manager Backend – Go + Fiber + PostgreSQL

This is the backend API for a **general-purpose Task Manager web application**. It allows users to register, log in, and manage their tasks efficiently with features like priority levels, due dates, and smart sorting.

Built with **Golang**, **Fiber**, and **PostgreSQL**, this RESTful API supports full authentication, task isolation per user, and intelligent task organization.

---

## 🚀 Features

### 👥 Multi-User Support
- Register and log in using email and password
- JWT-based secure authentication
- Users can only access their own tasks

### 🗂️ Task Management
- Add, view, update, and delete tasks
- Mark tasks as completed
- Set task **priority** (`High`, `Medium`, `Low`)
- Set **due dates**

### 📊 Smart Sorting
- Tasks are sorted as:
  1. High → Medium → Low priority
  2. Earlier due dates first
  3. Completed tasks appear last

---

## 🧱 Tech Stack

| Layer        | Technology   |
|--------------|--------------|
| Language     | Golang       |
| Framework    | Fiber        |
| Database     | PostgreSQL   |
| ORM          | GORM         |
| Auth         | JWT          |
| Passwords    | bcrypt       |
| Testing Tool | Postman      |

---

## 🔗 API Endpoints

### 🔐 Auth

| Method | Endpoint        | Description         |
|--------|-----------------|---------------------|
| POST   | `/api/register` | Register a new user |
| POST   | `/api/login`    | Log in and receive JWT token |

### 🧾 Tasks (Requires JWT)

| Method | Endpoint         | Description                   |
|--------|------------------|-------------------------------|
| GET    | `/api/tasks`     | Get all tasks for the user    |
| POST   | `/api/tasks`     | Create a new task             |
| PUT    | `/api/tasks/:id` | Update an existing task       |
| DELETE | `/api/tasks/:id` | Delete a task                 |

> 🔐 Auth required:  
> Include `Authorization: Bearer <JWT>` in headers

---

## 🔧 Project Structure
Task_manager/
├── config/ # DB connection setup
├── controllers/ # Business logic
├── middleware/ # Auth middleware
├── models/ # GORM models (User, Task)
├── routes/ # API route definitions
├── main.go # Entry point
└── go.mod # Dependencies
