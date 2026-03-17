# DOKUMENTASI

### FLOW DETAIL REQUEST
<img width="6250" height="3165" alt="User Creation Request Flow-2026-03-17-133130" src="https://github.com/user-attachments/assets/a5fee317-cf4c-4afa-8b16-f49ecbd5ffb1" />

### CODE
```bash
---
config:
  theme: forest
---
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB
Client->>Middleware: HTTP Request
Middleware->>Handler: forward request
Handler->>Service: CreateUser()
Service->>Repository: SaveUser()
Repository->>DB: INSERT USER
DB-->>Repository: OK
Repository-->>Service: user saved
Service-->>Handler: user response
Handler-->>Client: JSON Response
```
---

### DATABASE

<p align="center">
<img width="1355" height="379" alt="Screenshot (640)" src="https://github.com/user-attachments/assets/437cf9c2-285c-4eaf-9786-afac105be55d" />
<img width="1354" height="548" alt="Screenshot (641)" src="https://github.com/user-attachments/assets/e1b3b9db-98b4-4a9b-bd68-7e28be6aa019" />
<img width="1333" height="555" alt="Screenshot (642)" src="https://github.com/user-attachments/assets/f82680af-25c7-4e90-964d-c33bad282f4d" />

  
</p>

---

### HIT API

<p align="center">
<img width="1366" height="342" alt="Screenshot (635)" src="https://github.com/user-attachments/assets/fd498328-6896-4087-a83d-e04aa0267a07" />
<img width="1366" height="577" alt="Screenshot (636)" src="https://github.com/user-attachments/assets/3ae9faa5-1f0c-43f4-8259-a8edff6f0df6" />
<img width="1366" height="584" alt="Screenshot (637)" src="https://github.com/user-attachments/assets/aff84e51-dfcf-4ef8-9129-48821439d78d" />
<img width="1366" height="582" alt="Screenshot (638)" src="https://github.com/user-attachments/assets/46cd387c-e61f-4d6f-b5c4-bb31a8ca06d4" />
  <img width="1366" height="544" alt="Screenshot (639)" src="https://github.com/user-attachments/assets/fd826a3a-642e-4c75-9b4f-fb575ede6634" />

</p>

