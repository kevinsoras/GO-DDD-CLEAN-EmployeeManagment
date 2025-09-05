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

## Ruta principal

- POST `/employee` — Registra un empleado y su persona asociada.

## Principios aplicados

### DDD (Domain-Driven Design)
- **Contextos**: El código está organizado por contexto de dominio (`contexts/employee`, `shared`).
- **Entidades y Value Objects**: En `domain/entities` y `domain/value_objects`.
- **Aggregates y Factories**: En `shared/domain/aggregates` y `shared/domain/factories`.
- **Servicios de Dominio**: Ejemplo: `peruvian_labor_service.go`.

### Clean Architecture
- **Separación de capas**: Application, Domain, Infrastructure, Interfaces.
- **Inversión de dependencias**: Los casos de uso dependen de interfaces, no de implementaciones concretas.
- **DTOs**: Separados de las entidades de dominio.

### SOLID
- **S (Single Responsibility)**: Cada archivo/clase tiene una responsabilidad clara (ej: un caso de uso, un repositorio, un factory).
- **O (Open/Closed)**: Factories y repositorios pueden extenderse sin modificar el código existente.
- **L (Liskov Substitution)**: Las implementaciones de repositorios y datasources cumplen las interfaces del dominio.
- **I (Interface Segregation)**: Las interfaces son específicas y no fuerzan métodos innecesarios.
- **D (Dependency Inversion)**: Los casos de uso y controladores dependen de abstracciones (interfaces), no de detalles.

## Observaciones y sugerencias

- **Bien aplicado**: La estructura respeta DDD y Clean Architecture. Los casos de uso no dependen de infraestructura. Los value objects y entidades están bien separados.
- **SOLID**: Se cumple en la mayoría de los casos, especialmente con el uso de interfaces y factories.
- **Mejoras posibles**:
  - Persistencia de la persona: Actualmente el factory crea la entidad, pero no la persiste en la base de datos. Deberías agregar un repositorio para personas y persistirla antes de crear el empleado.
  - Validaciones: Ya usas validaciones avanzadas, pero podrías centralizar aún más la lógica de validación de negocio en servicios de dominio.
  - Tests: Agrega pruebas unitarias para casos de uso y servicios de dominio.

---

> Si necesitas ejemplos de cómo mejorar la persistencia o agregar tests, ¡avísame!
