# 🚀 Go Backend Series – Build a Production-Grade Go Application (Step by Step)

Welcome to the **Go Backend Series**!  
This is a complete, step-by-step guide to building a **scalable, maintainable, and production-ready backend** using **Go (Golang)** and modern best practices.

Whether you're a student, a backend developer starting with Go, or someone looking to build real-world systems, this series is for you.

---

## 📚 Series Overview

This repository is structured by parts. Each part focuses on one key aspect of backend development. Follow them in order to build a fully functional backend.

| Part | Title                                                                                     | Description                                                                                 | Status         |
|------|-------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|----------------|
| 1    | [Project Setup with Gin, Live Reloading, and .env Support](./part1/)                      | Basic API setup, folder structure, environment config, and live reloading with Air         | ✅ Completed   |
| 2    | [Routing and Modular Architecture](./part2/)                                               | Organizing your project with clean and scalable routing and folder structure               | 🔄 In Progress |
| 3    | [Connecting to a Database (PostgreSQL)](./part3/)                                          | Database setup using GORM or pgx, configuration, and migration handling                    | 🔄 In Progress |
| 4    | [Dependency Injection and Services Layer](./part4/)                                       | Build a clean and testable services layer using interfaces and DI                          | 🔄 In Progress |
| 5    | [Authentication with JWT](./part5/)                                                       | Implement user login, signup, and protected routes using JWT tokens                        | 🔄 In Progress |
| 6    | [Middleware, Logging, and Error Handling](./part6/)                                       | Set up global middlewares, structured logs, and consistent error responses                 | 🔄 In Progress |
| 7    | [Unit Testing and Integration Testing](./part7/)                                          | Add tests for handlers, services, and database interactions                                | 🔄 In Progress |
| 8    | [Dockerizing the Application](./part8/)                                                   | Run your Go backend in containers with Docker, Docker Compose, and `.env` support          | 🔄 In Progress |
| 9    | [CI/CD with GitHub Actions](./part9/)                                                     | Automate builds, tests, and deployments using GitHub Actions                               | 🔄 In Progress |
| 10   | [Deployment to Production (Render / Railway / Fly.io)](./part10/)                         | Deploy your app to a live environment with domain and HTTPS                                | 🔄 In Progress |

> 🔄 More parts will be added soon, including:
> - Role-based access control (RBAC)
> - Pagination & filtering
> - File uploads & cloud storage
> - WebSocket integration
> - gRPC support
> - Metrics & monitoring

---

## 🛠️ Tech Stack

- **Go (Golang)**
- **Gin** – HTTP web framework
- **PostgreSQL** – Relational database
- **GORM / pgx** – ORM or DB driver
- **JWT** – Authentication
- **Air** – Live reload during development
- **Docker** – Containerization
- **GitHub Actions** – CI/CD
- **Fly.io / Railway / Render** – Deployment platforms

---

## 🚀 How to Use This Repo

Each part is self-contained and builds upon the previous one. You can:

- Start from `part1/` and follow the code and blog post together
- Run each part locally (`cd partX && go run main.go`)
- Read the corresponding blog post for detailed explanation (links below)

---

## 📖 Blog Series

You can follow the full blog series here:  
👉 [Go Backend Series on Medium](https://medium.com/@bharath0292)

---

## 🤝 Contributions

Feel free to fork, star ⭐️, or contribute!  
If you find a bug or want to improve something, please open an issue or pull request.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

Happy Coding! 🙌  
— [Bharath T](https://github.com/bharath0292/go-backend-series)
