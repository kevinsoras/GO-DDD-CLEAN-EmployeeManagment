# Employee Management

Este proyecto es un ejemplo de arquitectura avanzada en Go para la gestión de empleados, aplicando los principios de DDD (Domain-Driven Design), Clean Architecture y SOLID.

## Estructura de carpetas

- `cmd/` — Punto de entrada de la aplicación (servidor HTTP, inyección de dependencias).
- `contexts/employee/` — Contexto de dominio Employee, siguiendo DDD y Clean Architecture:
  - `application/` — Casos de uso y DTOs (Application Layer)
  - `domain/` — Entidades, repositorios, servicios de dominio, value objects (Domain Layer)
  - `infrastructure/` — Implementaciones concretas de repositorios y datasources (Infrastructure Layer)
  - `interfaces/` — Controladores HTTP (Interface Layer)
- `shared/` — Código compartido entre contextos (personas, validaciones, factories, value objects, etc.)

## Endpoints de la API

### POST /employee

**Descripción:** Registra un nuevo empleado en el sistema, incluyendo tanto sus datos personales como los detalles de su empleo. Esta operación es atómica y asegura la consistencia de los datos.

**Método:** `POST`

**URL:** `/employee`

**Cuerpo de la Solicitud (Request Body):**

```json
{
  "person": {
    "type": "NATURAL",
    "email": "juan.perez@empresa.com",
    "phone": "+51987651324",
    "address": "Av. Lima 123, Lima, Perú",
    "country": "Perú",
    "documentNumber": "25312026",
    "firstName": "Juan",
    "lastNamePaternal": "Pérez",
    "lastNameMaternal": "Gómez",
    "birthDate": "1990-05-15T00:00:00Z",
    "gender": "M"
  },
  "employment": {
    "salary": 4500.00,
    "contractType": "INDEFINIDO",
    "startDate": "2024-01-15T00:00:00Z",
    "position": "Desarrollador Senior",
    "workSchedule": "Lunes a Viernes 9:00-18:00",
    "department": "Tecnología",
    "workLocation": "Oficina Central",
    "bankAccount": "0011-0234-56789012",
    "afp": "Integra",
    "eps": "Rímac",
    "hasCTS": true,
    "hasGratification": true,
    "hasVacation": true
  }
}
```

**Respuestas (Responses):**

*   `201 Created`: Empleado registrado exitosamente.
*   `400 Bad Request`: Datos de entrada inválidos (ej. validación fallida).
*   `409 Conflict`: La persona con el documento o email ya existe.
*   `500 Internal Server Error`: Error inesperado en el servidor.

### Documentación de la API (Swagger)

La documentación interactiva de la API se genera automáticamente usando [Swag](https://github.com/swaggo/swag).

1.  **Generar la Documentación:**
    Desde la raíz del proyecto, ejecuta:
    ```bash
    make swagger-docs
    ```
    Esto creará la carpeta `docs/` con los archivos de especificación OpenAPI.

2.  **Acceder a la Interfaz de Usuario de Swagger:**
    Una vez que la aplicación esté corriendo (`go run cmd/main.go`), abre tu navegador y navega a:
    ```
    http://localhost:3000/swagger/index.html
    ```
    Aquí podrás ver y probar todos los endpoints documentados.

### Herramientas de Calidad de Código

El proyecto utiliza las siguientes herramientas para asegurar la calidad y consistencia del código:

-   **`gofmt`**: Formateador estándar de Go.
-   **`goimports`**: Gestiona automáticamente las importaciones y formatea el código.
-   **`golangci-lint`**: Un linter rápido y potente que combina múltiples linters.

**Uso:**

-   **Formatear el código:**
    ```bash
    make format
    ```

-   **Ejecutar linters:**
    ```bash
    make lint
    ```

-   **Ejecutar todas las comprobaciones de calidad (formato + lint):**
    ```bash
    make check
    ```

## Configuración y Ejecución

Para arrancar la aplicación, sigue estos pasos:

1.  **Variables de Entorno:**
    Crea un archivo `.env` en la raíz del proyecto con la siguiente configuración. Asegúrate de ajustar `DB_URL` a tu configuración de PostgreSQL.

    ```dotenv
    DB_URL=postgres://user:password@host:port/database?sslmode=disable

    # Configuración de Logging
    # LOG_LEVEL: DEBUG, INFO, WARN, ERROR
    LOG_LEVEL=DEBUG
    # LOG_FORMAT: TEXT, JSON
    LOG_FORMAT=TEXT
    # LOG_OUTPUTS: Lista separada por comas de destinos (ej: stdout,file)
    LOG_OUTPUTS=stdout,file
    # LOG_FILE_PATH: Ruta al archivo de log (ej: app.log o logs/app.log)
    LOG_FILE_PATH=app.log
    ```

2.  **Ejecutar la Aplicación:**
    Asegúrate de tener Go instalado (versión 1.24.0 o superior).
    Desde la raíz del proyecto, ejecuta:

    ```bash
    go run cmd/main.go
    ```
    La aplicación se iniciará en el puerto 3000.

## Principios aplicados y Mejoras Recientes

### DDD (Domain-Driven Design)
- **Contextos**: El código está organizado por contexto de dominio (`contexts/employee`, `shared`).
- **Entidades y Value Objects**: En `domain/entities` y `domain/value_objects`.
- **Aggregates y Factories**: En `shared/domain/aggregates` y `shared/domain/factories`.
- **Servicios de Dominio**: Ejemplo: `peruvian_labor_service.go`.
- **Patrón Command**: Implementado para encapsular las solicitudes a los casos de uso, mejorando el desacoplamiento y la extensibilidad.

### Clean Architecture
- **Separación de capas**: Application, Domain, Infrastructure, Interfaces.
- **Inversión de dependencias**: Los casos de uso dependen de interfaces, no de implementaciones concretas.
- **DTOs**: Separados de las entidades de dominio.
- **Unit of Work (UoW)**: Implementado para gestionar transacciones atómicas a nivel de caso de uso, asegurando la consistencia de los datos.
- **Manejo de Errores Enriquecido**: Los errores de infraestructura se traducen a errores de dominio "ricos" en la capa de infraestructura, y el controlador los maneja de forma centralizada y genérica.
- **Ensamblaje de Dependencias**: La lógica de inyección de dependencias se ha centralizado en un paquete `app/` para una configuración más limpia y mantenible.

### SOLID
- **S (Single Responsibility)**: Cada archivo/clase tiene una responsabilidad clara (ej: un caso de uso, un repositorio, un factory).
- **O (Open/Closed)**: Factories y repositorios pueden extenderse sin modificar el código existente.
- **L (Liskov Substitution)**: Las implementaciones de repositorios y datasources cumplen las interfaces del dominio.
- **I (Interface Segregation)**: Las interfaces son específicas y no fuerzan métodos innecesarios.
- **D (Dependency Inversion)**: Los casos de uso y controladores dependen de abstracciones (interfaces), no de detalles.

### Observabilidad (Logging)
- **Logging Estructurado**: Integración de `log/slog` para logs estructurados y con niveles.
- **Configurable por Entorno**: El nivel y formato del log se configuran a través de variables de entorno (`LOG_LEVEL`, `LOG_FORMAT`).
- **Salida Múltiple**: Capacidad de escribir logs simultáneamente en la consola (`stdout`) y en un archivo (`LOG_FILE_PATH`).

### Testing
- **Tests Unitarios**: Ubicados junto al código que prueban, dentro de sus respectivos paquetes.
- **Tests End-to-End (E2E)**: Implementados en la carpeta `tests/e2e/` para verificar el flujo completo de la aplicación, utilizando `testcontainers-go` para entornos de base de datos aislados.

### Migraciones
- **Gestión Centralizada**: La creación de migraciones se gestiona a través de `Makefile`, requiriendo la especificación explícita del contexto (`employee` o `shared`) para asegurar la ubicación correcta de los archivos de migración.
