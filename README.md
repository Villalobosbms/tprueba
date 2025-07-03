#  QUE APRENDI Y QUE PUEDO MEJORAR

Aprendi mucho de esta prueba tecnica ya es la primera vez que presento una prueba tecnica en GO ya venia estudiandolo pero los conceptos basicos y poder realizar una api rest sencilla me sirvio mucho.

- Mejoraria el la conexion a la base de datos con un archivo .env o algo parecido para que la conexion a la base de datos fuera mas segura 
- Mejoraria la conpresion de los punteros en GO para asi facilitar su implementacion esta mejora es para mi no para el programa
  
#  API REST de Gestión de Usuarios

Esta es una API REST construida en **Go** utilizando **Gorilla Mux** para enrutamiento, **GORM** como ORM y **PostgreSQL** como base de datos. La API permite crear, leer, actualizar y eliminar usuarios.

---

##  Requisitos

Antes de ejecutar este proyecto, asegúrate de tener instalado:

- [Go](https://golang.org/dl/) (v1.18 o superior)
- [PostgreSQL](https://www.postgresql.org/)
- [Git](https://git-scm.com/)

---

## ⚙️ Instalación y configuración

### 1. Inicializar el módulo Go (solo una vez)
```bash
go mod init github.com/villalobosbms/tprueba
```

### 2. Instalar dependencias

#### Gorilla Mux (para rutas y servidor HTTP)
```bash
go get -u github.com/gorilla/mux
```

#### GORM (ORM para manejar la base de datos)
```bash
go get -u gorm.io/gorm
```

#### Driver de PostgreSQL para GORM
```bash
go get -u gorm.io/driver/postgres
```

#### Air (opcional, para recarga automática en desarrollo)
```bash
go install github.com/air-verse/air@latest
```

> Si usas Air por primera vez, ejecuta:
```bash
air init
```

Esto generará un archivo `.air.toml` para configuración.

---

## 🗄️ Configuración de la Base de Datos

Asegúrate de tener una base de datos PostgreSQL creada y configura la cadena de conexión directamente en el fichero que se encuentra en la ruta 
../db/connection en la variable

c:

```go
var DSN = "host=localhost user=tu_usuario password=tu_contraseña dbname=tu_base port=5432"
```

---

## ▶️ Ejecutar el proyecto

### Opción 1: Ejecutar con `go run`
```bash
go run main.go
```

### Opción 2: Ejecutar con recarga automática usando `air`
```bash
air
```

---

## 📬 Endpoints disponibles

- `GET    /user` → Lista todos los usuarios
- `GET    /user/{id}` → Obtiene un usuario por ID
- `POST   /user` → Crea un nuevo usuario
- `PUT    /user/{id}` → Actualiza un usuario por ID
- `DELETE /user/{id}` → Elimina un usuario por ID

---

## 📦 Estructura del JSON para crear usuario

```json
{
  "first_name": "Carlos",
  "second_name": "Andrés",
  "last_name": "Pérez",
  "age": 28,
  "password": "miclave123"
}
```

---

## 🛠 Autor

Desarrollado por [@villalobosbms](https://github.com/villalobosbms)
