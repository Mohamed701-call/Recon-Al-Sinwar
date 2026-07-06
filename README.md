# Recon Al-Sinwar

A fast and modular reconnaissance framework written in Go for security assessments and authorized penetration testing.

> **Status:** Alpha (Under Active Development)

---

## Features

* Subdomain Discovery
* Live Host Detection
* Port Scanning
* Web Crawling
* Smart Directory Fuzzing
* Screenshot Collection
* Technology Detection
* Smart Nuclei Scanning
* JSON Reports

---

## Project Structure

```text
Recon-Al-Sinwar/
├── cmd/
├── internal/
├── tools/
├── workspace/
├── go.mod
└── go.sum
```

---

## Requirements

* Go 1.24+
* Windows (Linux support is planned)

---

## Installation

Clone the repository:

```bash
git clone https://github.com/Mohamed701-call/Recon-Al-Sinwar.git
```

Enter the project directory:

```bash
cd Recon-Al-Sinwar
```

Download dependencies:

```bash
go mod tidy
```

Build the application:

```bash
go build -o alsinwar.exe ./cmd
```

Run:

```bash
.\alsinwar.exe
```

---

## Development

Run without building:

```bash
go run ./cmd
```

---

## Current Pipeline

```
Target
   │
   ▼
Subfinder
   │
   ▼
HTTPX
   │
   ▼
Naabu
   │
   ▼
Katana
   │
   ▼
Smart FFUF
   │
   ▼
GoWitness
   │
   ▼
Technology Detection
   │
   ▼
Nuclei
   │
   ▼
Reports
```

---

## Output

The tool stores all generated data inside:

```text
workspace/
```

Including:

* Subdomains
* Live Hosts
* Open Ports
* Crawled URLs
* FFUF Results
* Screenshots
* Nuclei Findings
* JSON Reports

---

## Roadmap

* HTML Dashboard
* Telegram Bot Integration
* Automatic Tool Updates
* Automatic Wordlist Updates
* Automatic Nuclei Template Updates
* Resume Scans
* Distributed Scanning
* Plugin System

---

## Disclaimer

This tool is intended **only** for authorized security testing and penetration testing activities. Users are responsible for ensuring they have permission before scanning or testing any target.

---

## License

This project is licensed under the MIT License.
