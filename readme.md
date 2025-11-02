# GO URL shortner API

A simple URL shortener built using **Go (Golang)** that provides:
- URL shortening and redirection
- Domain-based usage metrics
- In-memory data storage (no database)
- Optional Docker containerization

---

## Features
1. **Shorten URLs**
   - Accepts a long URL and returns a unique short URL.
   - Same input URL always returns the same shortened version.

2. **Redirect to Original URL**
   - Visit `/redirect/{shortID}` to be redirected to the original link.

3. **Top 3 Domains Metrics**
   - `/metrics` endpoint returns the top 3 most shortened domain names.

4. **In-Memory Storage**
   - URLs and domain stats are stored in memory (resets on restart).

5. **Docker Support**
   - Easily containerize and run the app anywhere.

---

## API endpoints
**POST** `/shorten`  
**Body:**
```json
{
  "url": "https://www.example.com"
}
```
**Response:**
```json
{
    "short_url": "a1b2c3d4"
}
```
**GET** `/redirect/{short_url}`  
 - Example:
http://localhost:3000/redirect/a1b2c3d4
→ redirects to https://www.example.com

**GET** `/metrices`
```json
[
  {"domain": "www.udemy.com", "count": 6},
  {"domain": "www.youtube.com", "count": 4},
  {"domain": "www.wikipedia.org", "count": 2}
]
```
## Project structure
```
.
├── main.go          # Main Go source code
├── go.mod           # Go module file
├── go.sum           # Dependencies
└── Dockerfile       # Containerization setup
```
## Run the app
 - go run main.go
####  the server will start on
 - http://localhost:3000

