# 🚀 DockLab

> **Spin up a complete, isolated cybersecurity lab with a single command.**

DockLab is a Go-based CLI that automatically deploys a vulnerable penetration testing environment using Docker. It creates a dedicated Docker network, launches multiple intentionally vulnerable applications, and provides a pre-configured penetration testing container so you can start hacking immediately.

Perfect for:
- 🎓 Cybersecurity students
- 🏴 CTF practice
- 🔐 Pentesting labs
- 👨‍🏫 Security training
- 🧪 Safe exploit testing

---

# Features

- ✅ One-command lab deployment
- ✅ Automatic Docker network creation
- ✅ Preconfigured penetration testing toolkit
- ✅ Multiple vulnerable targets
- ✅ Interactive shell into the toolkit container
- ✅ Automatic cleanup
- ✅ Cross-platform (Linux, macOS, Windows with Docker)

---

# Included Containers

| Container | Purpose |
|-----------|---------|
| Pentest Toolkit | Attack machine containing common pentesting tools |
| DVWA | Damn Vulnerable Web Application |
| OWASP Juice Shop | Modern intentionally vulnerable web application |
| WebGoat | OWASP training platform |

All containers are connected to an isolated Docker network.

---

# Architecture

```
                     Docker Network
                  ┌──────────────────┐

        ┌──────────────────────────────────────────────┐
        │                                              │
        │   Pentest Toolkit                            │
        │   nmap                                       │
        │   ffuf                                       │
        │   sqlmap                                     │
        │   nikto                                      │
        │   hydra                                      │
        │   curl                                       │
        │   gobuster                                   │
        │   feroxbuster                                │
        │   searchsploit                               │
        │                                              │
        └──────────────┬───────────────────────────────┘
                       │
     ┌─────────────────┼─────────────────────┐
     │                 │                     │
┌───────────┐   ┌──────────────┐     ┌────────────┐
│   DVWA    │   │ Juice Shop   │     │ WebGoat   │
└───────────┘   └──────────────┘     └────────────┘
```

---

# Installation

## Requirements

- Docker
- Go (only if building from source)

Verify Docker:

```bash
docker --version
```

---

## Option 1 — Install Binary

Download the latest release.

```bash
chmod +x install.sh
./install.sh
```

The installer will automatically:

- install DockLab
- place it in your PATH
- make it executable

After installation:

```bash
docklab
```

---

## Option 2 — Build From Source

Clone the repository.

```bash
git clone https://github.com/YOUR_USERNAME/docklab.git

cd docklab
```

Build:

```bash
go build -o docklab
```

Run:

```bash
./docklab
```

---

# Usage

## Start the Lab

```bash
docklab up
```

Output:

```
Checking Docker...

Creating Docker network...

Starting DVWA...

Starting Juice Shop...

Starting WebGoat...

Starting Pentest Toolkit...

Lab Ready!
```

---

## Start and Enter the Toolkit

```bash
docklab up --shell
```

or

```bash
docklab shell
```

You'll be dropped directly into the pentest container.

```
root@toolkit:/#
```

---

## Stop Everything

```bash
docklab down
```

This removes:

- all lab containers
- Docker network

---

# Services

After starting the lab:

| Application | URL |
|------------|------|
| DVWA | http://localhost:8080 |
| Juice Shop | http://localhost:3000 |
| WebGoat | http://localhost:8081/WebGoat |

---

# Network

DockLab creates an isolated Docker network.

Example:

```
cyberlab
```

Every container receives its own private IP.

Example:

```
Pentest Toolkit
172.19.0.3

DVWA
172.19.0.2

Juice Shop
172.19.0.4

WebGoat
172.19.0.5
```

The toolkit can directly reach every vulnerable machine by hostname or IP.

Example:

```
nmap dvwa
curl http://dvwa
```

---

# Included Tools

The toolkit includes many common offensive security tools.

### Enumeration

- nmap
- netcat
- curl
- wget
- httpie
- jq

### Web

- ffuf
- gobuster
- feroxbuster
- whatweb
- nikto

### Exploitation

- sqlmap
- hydra
- metasploit
- searchsploit

### Utilities

- python3
- git
- vim
- tmux
- tree
- less
- dnsutils
- iproute2
- net-tools

### Wordlists

- SecLists
- rockyou.txt

---

# Example Workflow

Start the lab.

```bash
docklab up --shell
```

Scan DVWA.

```bash
nmap dvwa
```

Directory fuzzing.

```bash
ffuf -u http://dvwa/FUZZ -w /usr/share/seclists/Discovery/Web-Content/common.txt
```

Fingerprint.

```bash
whatweb http://dvwa
```

SQL Injection.

```bash
sqlmap -u "http://dvwa/vulnerabilities/sqli/?id=1&Submit=Submit"
```

Run Nikto.

```bash
nikto -h http://dvwa
```

---

# Project Structure

```
.
├── cmd
│   ├── down.go
│   ├── root.go
│   ├── shell.go
│   └── up.go
├── images
│   └── toolkit
│       └── Dockerfile
├── main.go
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

# Roadmap

- [ ] Metasploitable support
- [ ] Custom lab profiles
- [ ] Plugin system
- [ ] Windows Active Directory lab
- [ ] HTB-style machine support
- [ ] YAML configuration
- [ ] User-defined targets
- [ ] Automatic updates

---

# Contributing

Contributions are welcome.

1. Fork the repository

2. Create a feature branch

```bash
git checkout -b feature/my-feature
```

3. Commit your changes

```bash
git commit -m "Added awesome feature"
```

4. Push

```bash
git push origin feature/my-feature
```

5. Open a Pull Request

---

# Security Notice

DockLab contains intentionally vulnerable applications.

Use only:

- on your own machine
- inside isolated Docker networks
- for educational purposes

Never expose these containers directly to the Internet.

---

# License

MIT License

---

# Acknowledgements

This project makes use of several outstanding open-source projects:

- OWASP DVWA
- OWASP Juice Shop
- OWASP WebGoat
- Docker
- Cobra CLI

Huge thanks to the maintainers of these projects.

---

# Author

**Adarsh**

If you found this project useful, consider giving it a ⭐ on GitHub!
```
