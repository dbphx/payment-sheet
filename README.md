# payment-sheet

A simple payment sheet / payment-gateway boilerplate built with Go + frontend (TypeScript/HTML).  

## 🚀 What is this

This project provides a minimal-setup example for payment processing — backend in Go, frontend in TypeScript/HTML — with Docker / docker-compose configuration. Useful as a starting point or reference for building payment flows using modern tools.

## 📦 What’s inside

- Backend: Go  
- Frontend: TypeScript / HTML  
- Docker / docker-compose setup (docker-compose.yml / docker-compose-prod.yml) for easy local dev or production deployment  
- Nginx config (nginx.conf)  
- Example configs under `conf/`  

## 🔧 Getting Started (Development)

### Prerequisites

- Docker & docker-compose installed  
- (Optional) Golang / Node.js / npm if you want to run parts locally  

### Run locally with Docker

```bash
git clone https://github.com/dbphx/payment-sheet.git
cd payment-sheet
docker-compose up --build
