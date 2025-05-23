# 📁 Distributed File System (DFS) - Go Implementation

> A fault-tolerant, scalable **Distributed File System** built with **Go**, supporting file uploads, downloads, and multi-node replication. This project simulates a simplified version of cloud storage services using a Master Tracker and multiple Data Keeper nodes.

---

## 📋 Table of Contents

- [🌟 Project Overview](#project-description)
- [🏗️ System Architecture](#system-architecture)
- [📂 Project Structure](#project-structure)
- [🚀 Getting Started](#how-to-run)
- [🤝 Contributors](#contributors)

---

## 🌟 Project Overview <a name="project-description"></a>

As distributed systems grow more prevalent, the ability to design and implement reliable file systems that span multiple machines becomes crucial. This DFS project implements a system in which:

- The **Master Tracker** manages metadata and oversees the system.
- Multiple **Data Keeper Nodes** store file data and send heartbeats.
- All **Data Keepers connect directly to the Master Tracker**.
- A **Client connects to the Master Tracker** to upload/download files.
- Files are automatically **replicated across 3 nodes** for fault tolerance.
- Nodes communicate using **gRPC** for metadata and **file transfers**.
- The system supports **parallel file downloads** from multiple nodes.
- All services are **multi-threaded** using Go’s goroutines.

---

## 🏗️ System Architecture <a name="system-architecture"></a>

```text
                           +------------------+
                           |  Master Tracker  |  <----------------------------------+
                           +------------------+                                     │
                            ▲    ▲    ▲    ▲                                        │
         gRPC (control)     │    │    │    │     gRPC (control)                     │
                            │    │    │    │                                        │
                            │    │    │    │                                        │
                            │    │    │    │                                        │
       +--------------------+    │    +---------------------+                       │
       |                         |                          |                       │
       v                         v                          v                       │
+----------------+     +----------------+         +----------------+                │
| Data Keeper 1  |     | Data Keeper 2  |         | Data Keeper 3  |                │
+----------------+     +----------------+         +----------------+                │
                              ^                                                     │
           ^                  |                           ^                         │
           |     gRPC Streams (File Transfer)             |                         │
           +----------------------------------------------+                         │
                              |                                                     │
                              v                                                     │
                       +-------------+                     gRPC (metadata)          │
                       |   Client    | <--------------------------------------------+
                       +-------------+
```

---

## 📂 Project Structure <a name="project-structure"></a>

```text
src/
├── client/                          # Client application
│   ├── config/
│   │   └── config.json              # Client configuration (e.g., master IP and port)
│   ├── download/
│   │   └── download.go              # Logic for downloading files from data keepers
│   ├── upload/
│   │   └── upload.go                # Logic for uploading files to data keepers
│   ├── Interface/
│   │   └── interface.go             # Client-side interface definitions
│   └── main.go                      # Entry point for the Client
│
├── master/                          # Master Tracker service
│   ├── download/
│   │   └── download.go              # Handles download coordination
│   ├── heartbeat/
│   │   ├── heartbeat.go             # Receives heartbeats from data keepers
│   │   ├── FIleInstancesCheck.go    # Checks if file instances meet redundancy
│   │   └── Isidle.go                # Detects idle data keeper nodes
│   ├── lookup/
│   │   ├── file/
│   │   │   └── lookup.go            # File location lookups
│   │   └── node/
│   │       └── lookup.go            # Node availability lookups
│   ├── register/
│   │   └── register.go              # Handles node registration
│   ├── replicate/
│   │   └── replicate.go             # Replication service logic
│   ├── upload/
│   │   └── upload.go                # File upload handler
│   └── main.go                      # Entry point for the Master Tracker
│
├── node/                            # Data Keeper node logic
│   ├── config/
│   │   └── config.json              # Node configuration (e.g., master IP, ports)
│   ├── download/
│   │   └── download.go              # Responds to client download requests
│   ├── heartbeat/
│   │   └── heartbeat.go             # Sends heartbeats to the master
│   ├── register/
│   │   └── register.go              # Registers node with the master
│   ├── replicate/
│   │   └── replicate.go             # Handles replication triggered by the master
│   ├── upload/
│   │   └── upload.go                # Receives uploaded files from client
│   └── main.go                      # Entry point for the Data Keeper
│
├── uploads/                         # Stores uploaded files
├── downloads/                       # Stores downloaded files
├── schema/                          # Protocol buffer definitions for gRPC services
│   ├── download/
│   ├── heartbeat/
│   ├── register/
│   ├── register/
│   ├── replicate/
│   ├── upload/
│   │   
│
├── NodeId.bat                       # Script for assigning or managing Node IDs
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
```

---

## 🚀 Getting Started <a name="how-to-run"></a>

To get started with the Distributed File System, follow these steps for configuration and execution:

### 🛠️ 1. Prerequisites

Ensure you have Go installed:

```bash
go version
```

Install required dependencies:

```bash
go mod tidy
```

### 🔧 2. Configuration

#### Client Configuration (`src/client/config/config.json`)

Update the client configuration with the Master Tracker's IP and port:

```json
{
    "serverIP": ,
    "serverPort": 
}
```

#### Node Configuration (`src/node/config/config.json`)

Set the configuration for each Data Keeper node:

```json
{
    "nodeID": -1,
    "serverIP": ,
    "serverPort": 
}
```

### ▶️ 3. Running the Services

Each service should be launched from its respective directory.

#### 🧠 Master Tracker

Only **one instance** of the Master Tracker should be running:

```bash
cd src/master
go run main.go
```

#### 💾 Data Keeper Nodes

Run multiple nodes:

```bash
cd src/node
go run main.go
```

#### 👤 Client

Run multiple clients to **upload or download files**:

```bash
cd src/client
go run main.go
```

#### 📤 Uploading Files

Use the client interface to **upload a file**.  
It will be **automatically replicated across 3 nodes**.

#### 📥 Downloading Files

Download files using the client.  
The download occurs in **parallel from multiple nodes** for improved performance.

---

## 🤝 Contributors <a name="contributors"></a>

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Omar-Said-4" target="_black">
        <img src="https://avatars.githubusercontent.com/u/87082462?v=4" alt="Omar Said"/>
        <br />
        <sub><b>Omar Said</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/MostafaMagdyy" target="_black">
        <img src="https://avatars.githubusercontent.com/u/97239596?v=4" alt="Mostafa Magdy"/>
        <br />
        <sub><b>Mostafa Magdy</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/nouraymanh" target="_black">
        <img src="https://avatars.githubusercontent.com/u/102790603?v=4" alt="Nour Ayman"/>
        <br />
        <sub><b>Nour Ayman</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/3abqreno" target="_black">
        <img src="https://avatars.githubusercontent.com/u/102177769?v=4" alt="Abdelrahman Mohamed"/>
        <br />
        <sub><b>Abdelrahman Mohamed</b></sub>
      </a>
    </td>
  </tr>
</table>
