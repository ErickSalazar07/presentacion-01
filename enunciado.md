# Presentación 1: Exploración de Estilos Arquitectónicos y Stack Tecnológico

## Objetivo

- Capacitar al estudiante en la selección, modelado e implementación de soluciones de software mediante el análisis de estilos arquitectónicos y stacks
tecnológicos reales, utilizando el modelo C4 y criterios técnicos para sustentar la toma de decisiones.

La actividad se centra en:
- Comprensión de estilos arquitectónicos y patrones asociados.
- Uso de stacks tecnológicos reales. Aplicación de patrones de diseño. Modelado de soluciones mediante C4 Model.
- Implementación de arquitectura.
- Análisis de decisiones arquitectónicas asociadas al estilo y al stack.

1. Requerimientos del Stack Tecnológico

Cada grupo deberá definir e implementar un stack tecnológico único, compuesto obligatoriamente por:

- __Frontend:__ Una tecnología de interfaz de usuario.
- __Backend:__ Un framework y runtime de servidor. Persistencia: Una base de datos (relacional o no relacional).
- __Protocolo de Integración:__ Definición clara de comunicación (ej. REST, GraphQL, gRPC, WebSockets).
- __Componentes Adicionales:__ El grupo podrá seleccionar herramientas complementarias según las necesidades de su diseño.

2. Alcance del Caso Práctico

La solución desarrollada debe ser funcional y técnicamente consistente con el estilo arquitectónico elegido, cumpliendo con los siguientes mínimos:

- __Modelado:__ Mínimo 3 entidades de negocio interrelacionadas.
- __Flujo Funcional:__ Implementación de un caso de uso completo de extremo a extremo (end-to-end).
- __Persistencia:__ Gestión real de datos en el motor seleccionado.
- __Robustez:__ Implementación de manejo de errores y validaciones.
- __Justificación:__ Todas las decisiones técnicas deben estar sustentadas en las características del estilo y el stack.

## Hacer: Investigación del estilo y stack

### Para el estilo arquitectónico, el grupo deberá desarrollar:

> [!IMPORTANT]
> Revisar documento Word

- Definición clara (qué es y qué no es).
- Clasificacion del estilo.
- Características principales.
- Historia y evolución.
- Ventajas y desventajas.
- Problemas comunes que se presentan y patrones.
- Identificar patrones aplicables y cuando usarlos.
- Casos de uso (cuándo usarlo y cuándo no).
- Casos de aplicación (ejemplos reales en la industria).

### Para cada una de las tecnologías del stack asignado, el grupo deberá desarrollar:

> [!IMPORTANT]
> Revisar documento Word

- Definición clara (qué es y qué no es).
- Características principales.
- Historia y evolución.
- Ventajas y desventajas.
- Casos de uso (cuándo usarlo y cuándo no).
- Casos de aplicación (ejemplos reales en la industria).
- Relación entre el estilo y las tecnologías seleccionadas.
- Que tan común es el stack designado (relacion entre tecnologias).

## Análisis arquitectónico

> [!IMPORTANT]
> Hacer en el word

realizar matrices de analisis arquitectonicas

- __Matriz de atributos de calidad vs estilo:__ Identificar cómo el estilo soporta o limita cada atributo de calidad.
- __Matriz de análisis de Principios vs estilo:__ Identificar como el estilo cumple los principios SOLID, KISS, DRY, YAGNI, PoLA, Ley de Demeter, STUPID y
  Composicion sobre Herencia.
- __Matriz de análisis de tácticas vs estilo y stack:__ Identificar como un ADR puede ser justificado para el estilo y stack asignado.
- __Matriz de análisis de mercado Laboral vs estilo y stack:__ Identificar como esta el mercado actual para el estilo y el estack definido, con salarios
  mensuales o anueles y la proyeccion que tienen en el mercado laboral. (posibles fuentes: encuesta anual de GitHub, JetBrains o StackOverflow, TIOBE
  Index, PYPL, IEEE Spectrum, Gartner, TechnoVision  HackerRank Developer Skills Report, Hired / LinkedIn Economic Graph)

## Diseño: Ejemplo práctico y funcional

> [!IMPORTANT]
> Hacer en el word

El sistema debe ser modelado utilizando HLD +  C4 Model:

__Obligatorio:__
- Diagrama de alto nivel (HLD)
- Diagrama de Contexto (Nivel 1)
- Diagrama de Contenedores (Nivel 2)
- Diagrama Dinámico (flujo principal)
- Diagrama de Despliegue

__Opcional:__
- Diagrama de componentes (Nivel 3)
- Diagrama de Codigo (Nivel 3)
- Diagrama de System landscape
- Modelo de datos.

## Implementacion: Ejemplo práctico y funcional

> [!IMPORTANT]
> Usar VMs de los integrantes

- Código funcional en repositorio Git público con TAG
- Uso obligatorio de contenedores: Docker o Podman
- Implementación del flujo principal Integración entre componentes

## Entregables

> [!IMPORTANT]
> Tener los siguientes apartados

- Documento técnico
- Investigación
- Análisis arquitectónico
- Diseño ejemplo practico
- Lecciones aprendidas
- Repositorio Git público
- Código fuente
- Creacion de TAG y Release
- README con:
    + Descripción del sistema
    + Tecnologías usadas
    + Pasos para despliegue
- Presentación
- Uso de diagramas e infografías
- Explicación clara de los temas
- Estructura fluida
- Uso adecuado de colores

## Sustentación

Exposición de 10 a 20 minutos por grupo Todos los integrantes deben participar


# TEMAS A SELECCIONAR

## 1. CATÁLOGO DE ESTILOS ARQUITECTÓNICOS

Se refiere a la organización macroscópica del sistema. A. Estilos Estructurales (Lógica y Organización)

- __Monolítico:__ Monolito Modular, Microkernel (Plug-in), Monolito Monolítico.
- __Capas (Layered Architecture):__ N-Tiers, Local o Distribuido, Capas Cerradas vs Abiertas.
- __Arquitecturas de Dominio / "Círculos":__ Arquitectura Hexagonal (Ports & Adapters). Onion Architecture. Clean Architecture. Screaming Architecture.

### B. Estilos Distribuidos y Cloud-Native

- __Cliente-Servidor:__ Frontend / Backend (SPA + API) y P2P (Peer-to-Peer).
- __SOA (Service-Oriented Architecture):__ ESB-based.

- __Microservicios:__ Orquestados vs Coreografiados.
- __Serverless / FaaS (Function as a Service):__ Arquitecturas efímeras y escalado a cero.
- __Space-Based Architecture:__ Para alta escalabilidad basada en memoria compartida/tuplas. Event-Driven

- __Architecture (EDA):__ Broker, Mediator, Streaming.
- __Mensajería Asíncrona:__ Publicador-Suscriptor (Pub-Sub) y Colas de Mensajes.

### C. Estilos de Datos y Mensajería

- __Arquitectura Centrada en Datos:__ Data-Centric, blackboard.
- __Streaming:__ Procesamiento en tiempo real (Streaming).
- __Tuberías y Filtros (Pipes &amp; Filters):__ Procesamiento secuencial.
- __Kappa & Lambda Architectures:__ Fundamentales para Big Data y analítica en tiempo real.
- __Data Mesh:__ Arquitectura de datos distribuida por dominios. Data Lake y Data WareHouse: Arquitectura de datos

### D. Patrones de Presentación (Frontend/UI)

MVC, MVP, MVVM. Flux / Redux / MVI (Model-View-Intent).
- __Island Architecture:__ (Renderizado selectivo moderno).
- __Micro-frontends:__ (Extensión del concepto de microservicios a la UI).

## 2. PATRONES ASOCIADOS

Nota para el estudiante:
No es necesario usar todos los patrones. El objetivo es que seleccionen los que realmente necesita su estilo arquitectónico. Por ejemplo, si eligen
Microservicios, es obligatorio hablar de API Gateway y Circuit Breaker, pero si eligen un Monolito Modular, estos últimos podrían no ser necesarios. A.
Persistencia y Gestión de Estado

Patrones de Acceso:
DAO, Repository, Unit of Work, Data Mapper, Active Record.
- __Estrategias de Datos:__ Database per Service, Shared Database (Anti-patrón en ciertos contextos).
- __Consistencia y Consulta:__ CQRS, Event Sourcing, Materialized View.
- __Bajo Nivel:__ Write-Ahead Log (WAL).

### B. Comunicación e Integración

Estructuras de Red: API Gateway, BFF (Backend for Frontend), Service Mesh, Load Balancer, Proxy Inverso. Componentes: Sidecar / Ambassador, Gateway
Routing, Aggregator. Migración y Desacoplamiento: Strangler Fig (Higo Estrangulador), ACL (Anti-Corruption Layer), Facade. Patrones de Eventos:
Publish–Subscribe, Event Streaming, Message Queue, Dead Letter Queue (DLQ), Event Notification, Event Replay.

### C. Estrategias de Renderizado y Composición (Tácticas de UI)

CSR (Client-Side Rendering): Renderizado en el navegador (SPAs). SSR (Server-Side Rendering): Renderizado en el servidor (SEO-friendly). SSG (Static
Site Generation): Generación de sitios estáticos. ISR (Incremental Static Regeneration): Regeneración híbrida.

### D. Resiliencia, Estabilidad y Consistencia

Estabilidad y Resiliencia:
Circuit Breaker, Retry, Timeout, Bulkhead, Rate Limiting, Throttling.

Consistencia Distribuida:
Saga (Orquestación/Coreografía), Two-Phase Commit (2PC), TCC (Try-Confirm-Cancel), Idempotency, Eventual Consistency, Compensating Transactions.

Observabilidad:
Distributed Tracing, Structured Logging, Health Checks, Logs Centralizados.

### E. Seguridad y Gestión de Acceso

Identidad:
OAuth 2.0, OpenID Connect, SAML, JWT, IdP (Identity Provider), SSO (Single Sign-On), Federations.

Arquitectura de Seguridad:
Zero Trust, Multitenancy, Sidecar Security, Secrets Management.

Control de Acceso:
RBAC (Roles), ABAC (Atributos), PBAC (Políticas).

Protección de Datos:
API Key Management, Encryption at Rest / in Transit.

### F. Diseño de Código (GoF + Modernos)

Creacionales: Singleton, Factory, Abstract Factory, Builder, Prototype. Estructurales: Adapter, Bridge, Composite, Decorator, Flyweight, Proxy.
Comportamiento: Observer, Strategy, Command, Chain of Responsibility, Mediator, State, Template Method, Visitor, Memento.

## 3. CATÁLOGO DE ANTIPATRONES (Riesgos y Errores)

Arquitectónicos: Distributed Monolith (el más peligroso hoy), Big Ball of Mud, Spaghetti Architecture, God Service, Acoplamiento Estrecho (Tight
Coupling), Vendor Lock-in. De Diseño: Golden Hammer (usar la misma herramienta para todo), Lava Flow (código muerto), Optimización Prematura,
Reinventar la Rueda, Tight Coupling. De Implementación: Anemic Domain Model (clases de datos sin lógica), Fat Models / Skinny Controllers, Hardcoding,
God Object, Chatty Microservices (exceso de llamadas), Overengineering. De Datos: Data Silos, Lack of Versioning, Database-as-IPC (usar la DB para
comunicar servicios), Duplicación sin control, falta de normalización/sobre-normalización, ausencia de versionado de esquemas. Operacionales: ClickOps
(configurar todo manual), Works on My Machine, Lack of Observability, Ausencia de CI/CD, despliegue manual, falta de monitoreo y ausencia de planes de
Rollback.

## 4. FRAMEWORKS BACKEND

Esta lista clasifica las herramientas por su arquitectura interna y el lenguaje que las sustenta: A. Frameworks Full-Stack & Enterprise

Son robustos, incluyen todo lo necesario (ORM, Seguridad, Inyección de Dependencias) y fuerzan una estructura clara. Ideales para Arquitectura en
Capas y Monolitos Modulares.

Jakarta EE (Anteriormente Java EE): Java, Kotlin. Es el estándar de la industria sobre el cual se construyen muchos otros frameworks. Ideal para
grandes corporaciones que requieren estabilidad a muy largo plazo. Spring Boot: Java, Kotlin, Groovy. (El rey del sector bancario/seguros). .NET: C#,
F#. (gran ecosistema empresarial). Django: Python. (Enfoque en seguridad y desarrollo rápido). NestJS: TypeScript, JavaScript. (Arquitectura inspirada
en Angular para Node.js). Laravel: PHP. (Elegancia y herramientas de productividad masiva). Symfony: PHP. (Basado en componentes desacoplados; es la
base de muchos otros sistemas). Ruby on Rails: Ruby. (El pionero de la convención sobre configuración). Odoo Framework: Python. (Específico para
aplicaciones de gestión empresarial/ERP).

### B. Frameworks Minimalistas y Micro-frameworks

Enfocados en la velocidad y la baja latencia. Son la base preferida para Microservicios, Serverless y Sidecars.

FastAPI: Python. (Basado en estándares OpenAPI y tipos de Python). Express.js: JavaScript, TypeScript. (El
estándar minimalista de Node.js). Koa.js: JavaScript, TypeScript. (Koa es el sucesor espiritual de Express, más
ligero). Gin / Echo: Go (Golang). (Diseñados para alta concurrencia y simplicidad). Flask: Python. (Versión ligera
para micro-herramientas). Fastify: JavaScript, TypeScript. (Enfoque extremo en rendimiento de JSON). Fiber: Go
(Golang). (Sintaxis similar a Express pero con la velocidad de Go). Echo: Go (Golang). Hapi.js: JavaScript,
TypeScript. (Enfoque extremo en seguridad y configuración, muy usado en servicios críticos). Sanic: Python.
(Framework asíncrono diseñado para ser muy rápido, similar a FastAPI).

### C. Frameworks Cloud-Native & Reactive (Alta Eficiencia)

Diseñados específicamente para entornos de contenedores (Docker/Kubernetes) con tiempos de arranque casi instantáneos y bajo consumo de RAM.

Quarkus: Java, Kotlin. (Optimizado para GraalVM y Kubernetes). Micronaut: Java, Kotlin, Groovy. (Inyección de dependencias en tiempo de compilación).
Actix-web / Axum / Rocket : Rust. (Seguridad de memoria y el rendimiento más alto del mercado). Ktor: Kotlin. (Creado por JetBrains, puramente
asíncrono). Go (Standard Library): Soporta Go. (Muchos arquitectos prefieren no usar frameworks en Go por su potente librería nativa). Helidon: Java.
(La respuesta de Oracle para microservicios ligeros en Java).

### D. Frameworks de Alta Concurrencia y Tiempo Real

Utilizan modelos de actores o máquinas virtuales especializadas para manejar millones de conexiones simultáneas. Ideales para
EDA (Event-Driven Architecture).

Phoenix: Elixir. (Corre sobre la BEAM, la máquina virtual de Erlang; imbatible en chat/streaming). Akka HTTP / Play Framework :
Scala, Java. (Basado en el modelo de actores para sistemas distribuidos). Vert.x: Java, Kotlin, JavaScript, Groovy. (Un toolkit
políglota y reactivo para construir sistemas distribuidos de alto rendimiento). Tornado: Python. (Históricamente importante
para WebSockets y conexiones de larga duración).

### E. Frameworks de Orquestación y APIs

Apollo Server: JavaScript, TypeScript. (El estándar para arquitecturas basadas en GraphQL). Hasura / PostgREST: (Herramientas
que transforman la base de datos directamente en una API, reduciendo el código de backend).

## 5. FRAMEWORKS Y LIBRERÍAS FRONTEND A. Los Principales (The Big Three + 1)

Son los pilares del mercado. Elegir uno de estos garantiza soporte a largo plazo y gran cantidad de talento humano.

React: (Librería). Basada en componentes y Virtual DOM. Soporta: JavaScript, TypeScript, JSX. Angular: Framework integral (Enterprise). Incluye RxJS
para programación reactiva. Soporta: TypeScript. Vue.js: Framework progresivo. Muy flexible y equilibrado. Soporta: JavaScript, TypeScript. Svelte:
(Compilador). Transforma el código en JS puro en tiempo de compilación. Soporta: JavaScript, TypeScript.

### B. Meta-frameworks (Arquitectura de Renderizado)

Son los que realmente definen la arquitectura de la solución (SSR, SSG, ISR). Indispensables para producción.

Next.js: El estándar para React. Soporta: JavaScript, TypeScript. Nuxt.js: El estándar para Vue. Soporta: JavaScript, TypeScript. SvelteKit: El
meta-framework para Svelte. Soporta: JavaScript, TypeScript. Angular Universal / Analog: Soluciones de SSR para Angular. Soporta: TypeScript. Remix:
Enfocado en estándares web y mutaciones de datos. Soporta: JavaScript, TypeScript. SolidStart: Meta-framework para SolidJS (alto rendimiento).
Soporta: JavaScript, TypeScript.

### C. Alternativos Modernos y Alto Rendimiento (Islands & Resumability)

Herramientas que buscan eliminar el costo de la "hidratación" de JavaScript en el navegador.

Astro: Usa Island Architecture (envía HTML puro y solo carga JS donde hay interacción). Soporta: Markdown, MDX, JS, TS (y componentes de
React/Vue/Svelte). Qwik: Basado en Resumability (carga instantánea sin ejecución de JS inicial). Soporta: TypeScript, JSX. SolidJS: Reactividad fina
sin Virtual DOM. Soporta: JavaScript, TypeScript, JSX. Preact: Versión ultraligera de React (3kb). Soporta: JavaScript, TypeScript. Lit: Basado en Web
Components nativos. Soporta: JavaScript, TypeScript.

## 6. DESARROLLO MÓVIL, ESCRITORIO E HÍBRIDO A. Multiplataforma Nativo (Compilados)

Flutter: Usa el lenguaje Dart. Renderiza su propia UI (Motor Skia/Impeller). React Native: Usa JavaScript, TypeScript. Componentes JS que se mapean a
nativos. Kotlin Multiplatform (KMP): Usa Kotlin. Permite compartir lógica de negocio pura entre iOS, Android y Desktop. SwiftUI / Jetpack Compose:
(Nativos puros). Soporta: Swift (iOS) y Kotlin (Android).

### B. Híbridos y Web-Native (WebView/Bridge)

Ionic: Usa JavaScript, TypeScript (con React, Angular o Vue). Basado en Capacitor. Quasar: Framework basado en Vue para generar SPA, PWA, Mobile y
Desktop. Soporta: JavaScript, TypeScript. Tauri: El sucesor moderno de Electron. Usa Rust para el backend y HTML/JS para el frontend. Mucho más
ligero. Electron: Usa JavaScript, TypeScript. Para aplicaciones de escritorio potentes (ej. VS Code).

## 7. LIBRERÍAS Y HERRAMIENTAS TRANSVERSALES A. Diseño y Estilizado (UI/UX Systems)

Definen cómo se construye la capa visual y la consistencia del diseño.

Tailwind CSS: Framework de utilidades CSS (Utility-first). El estándar moderno para flexibilidad total. Material UI (MUI) / Joy UI: Implementación de
Google Material Design para React. Shadcn/ui: No es una librería de componentes tradicional, sino una colección de componentes reutilizables que
copias y pegas (basado en Tailwind y Radix). Bootstrap: El veterano basado en clases predefinidas. Útil para prototipado rápido y sistemas internos.
DaisyUI / Headless UI: Componentes sin estilos predefinidos para máxima personalización.

### B. Gestión de Datos y Persistencia (ORM/ODM)

Mapean los objetos del lenguaje a la estructura de la base de datos.

Prisma: ORM moderno para Node.js/TypeScript con seguridad de tipos automática. Mongoose: ODM estándar para modelado de datos en MongoDB. Spring Data
JPA / Hibernate: El estándar para persistencia en el ecosistema Java/Jakarta y Spring. Entity Framework (EF) Core: El ORM potente y oficial para .NET.
SQLAlchemy: El ORM más flexible para Python. Eloquent: El ORM intuitivo de Laravel (PHP). Drizzle ORM: Alternativa ultraligera y rápida para
TypeScript.

### C. Comunicación y Consumo de APIs

Axios: Cliente HTTP basado en promesas para el navegador y Node.js. TanStack Query (React Query): Gestión de estado asíncrono, caché y sincronización
de datos en el frontend. Apollo Client / Relay: Clientes especializados para consumir GraphQL. gRPC-web / Connect RPC: Para comunicación binaria de
alto rendimiento entre frontend y backend.

### D. Validación y Utilidades de Tipado

Zod: Librería de declaración y validación de esquemas con inferencia de tipos (TypeScript). Pydantic: Validación de
datos para Python (usada internamente por FastAPI). Zustand / Redux Toolkit: Gestión de estado global en el frontend
(Zustand es la tendencia por su simplicidad). Lodash: Utilidades de manipulación de arreglos, objetos y funciones.

### E. Pruebas (Testing Stack)

Jest / Vitest: Frameworks de ejecución de pruebas unitarias. Cypress / Playwright: Herramientas para pruebas End-to-End (E2E). (Playwright es el líder
actual por velocidad). Testing Library (RTL): Para probar componentes desde la perspectiva del usuario. Test Runners (Motores): JUnit (Java), Pytest
(Python), Jest (JS), xUnit (.NET). Mocking (Simulación): Mockito, Moq, Mockery. Moq / NSubstitute, unittest.mock / Pytest-mock, GoMock / Testify,
Mockery / PHPUnit Mocks. Assertions (Validación): AssertJ, Hamcrest, Chai. Contenedores de Prueba: Testcontainers Librería vital en 2026 para levantar
bases de datos reales en Docker temporalmente durante los tests.

## 8. MOTORES DE BASES DE DATOS A. Relacionales (SQL / ACID)

Garantizan consistencia total y relaciones complejas.

PostgreSQL: El estándar de la industria (Extensible con JSONB y Vectores). MySQL / MariaDB: El motor más popular para web y ecosistemas de código
abierto. SQL Server: Líder en integración empresarial con el stack de Microsoft. Oracle Database: Para sistemas críticos de escala masiva y alta
transaccionalidad. IBM Db2: Motor histórico para entornos corporativos de gran escala (Legacy). SQLite / H2: Motores embebidos. SQLite es el estándar
para móviles/Edge; H2 se usa principalmente para Testing en Java.

### B. NoSQL: Documentales y Columnares (Agilidad & Big Data)

Para datos semi-estructurados y escalabilidad horizontal masiva.

MongoDB: Líder en documentos JSON para desarrollo ágil. Couchbase / CouchDB: Enfoque en alta disponibilidad y sincronización móvil/offline. Cassandra
/ HBase: Motores de Wide-Column. Diseñados para escrituras masivas distribuidas (Big Data).

### C. Clave-Valor e In-Memory (Caché & Performance)

Redis: El estándar para caché, colas y estructuras de datos en tiempo real. Memcached: Sistema de caché multihilo, simple y puramente volátil.
Dragonfly: Alternativa moderna de ultra alto rendimiento compatible con Redis.

### D. Especializadas (Grafos, Tiempo, Búsqueda y Vectores)

Neo4j: Orientada a Grafos. Relaciones complejas (fraude, redes sociales). InfluxDB / TimescaleDB: Orientadas a Series de Tiempo (IoT, métricas).
Elasticsearch / Meilisearch: Motores de Búsqueda y analítica de texto completo.

### E. MODELOS DE DESPLIEGUE CLOUD

Esta sub-categoría es vital para decidir quién gestiona la base de datos.

DBaaS (Gestionadas): Amazon RDS, MongoDB Atlas, Azure SQL. El proveedor maneja parches y backups. Serverless DB (Escalado a Cero): Amazon DynamoDB,
Amazon Aurora Serverless, Neon (Postgres), PlanetScale (MySQL). Pagas solo por lectura/escritura real. Cloud-Native / Distributed SQL: Google Spanner,
CockroachDB, YugabyteDB. Consistencia fuerte en múltiples regiones geográficas simultáneamente. Edge Databases: Turso (libSQL), Cloudflare D1. Bases
de datos que viven cerca del usuario para latencia mínima. Analytical Cloud (OLAP): Snowflake, Google BigQuery, Amazon Redshift. Para procesamiento de
petabytes y Business Intelligence.

### F. NewSQL / Distribuidas (Escalabilidad Global)

Combinan el lenguaje SQL con la escalabilidad horizontal de NoSQL.

CockroachDB: SQL distribuido diseñado para sobrevivir a caídas de centros de datos completos. Google Spanner: Base de datos con consistencia fuerte a
escala planetaria. BigChainDB: Combina las características de una base de datos (escalabilidad, consultas ricas) con las de blockchain (activos,
inmutabilidad, descentralización). Usa MongoDB por debajo. OrbitDB: Una base de datos distribuida P2P que corre sobre IPFS. Ideal para aplicaciones
totalmente descentralizadas (DApps). Gun DB: Una base de datos gráfica y descentralizada que permite sincronización en tiempo real sin un servidor
central obligatorio.

### G. Vectoriales (IA y Machine Learning)

Fundamentales para aplicaciones de Inteligencia Artificial (LLMs).

Pinecone / Milvus: Especializadas en búsqueda semántica y almacenamiento de vectores. pgvector: Extensión que convierte a PostgreSQL en una base de
datos vectorial.

## 9. PROTOCOLOS Y ESTRATEGIAS DE COMUNICACIÓN

Esta categoría define el "idioma" y las reglas de transporte entre los componentes del sistema. A. Comunicación Síncrona (Request-Response)

REST (HTTP/1.1 & 2): El estándar basado en recursos y verbos HTTP. GraphQL: Consultas flexibles definidas por el cliente. gRPC (Protobuf): RPC binario
de alto rendimiento (Google). SOAP: Protocolo basado en XML para entornos corporativos legacy. HTTP/3 (QUIC): Transporte moderno sobre UDP para baja
latencia en redes inestables.

### B. Comunicación en Tiempo Real y Streaming

WebSockets: Conexión persistente bidireccional (Full-duplex). SSE (Server-Sent Events): Flujo
unidireccional del servidor al cliente. WebRTC: Comunicación P2P para audio y video en tiempo
real. WebTransport: Alternativa moderna a WebSockets sobre QUIC.

### C. Protocolos Especializados (IoT, Industrial y Media)

MQTT: Protocolo ligero de mensajería para IoT. CoAP: "REST" optimizado para dispositivos de baja potencia (UDP). AMQP: Protocolo avanzado de colas
(usado por RabbitMQ). DDS: Bus de datos en tiempo real para sistemas críticos/robótica. RTSP / RTMP / HLS: Protocolos específicos para streaming de
video y cámaras. LWM2M: Gestión de dispositivos y telemetría en IoT.

## 10. EVENTOS, MENSAJERÍA Y ORQUESTACIÓN (EDA) A. Protocolos de Mensajería (El estándar de comunicación)

Un arquitecto debe saber que el broker es intercambiable si el protocolo es estándar.

AMQP (Advanced Message Queuing Protocol): El estándar para colas seguras y transaccionales (usado por RabbitMQ). MQTT (Message Queuing Telemetry
Transport): El estándar para IoT y redes inestables. STOMP: Protocolo simple basado en texto, ideal para WebSockets y clientes ligeros. CloudEvents:
No es un protocolo de transporte, sino una especificación de la CNCF para estandarizar la descripción de eventos (metadata) entre diferentes nubes.

### B. Message Brokers (Enrutamiento y Colas)

Enfocados en la entrega garantizada y lógica de "Push".

RabbitMQ: El más versátil para enrutamiento complejo (Exchanges). NATS / NATS JetStream: Ultra-rápido, escrito en Go, ideal para microservicios
modernos y Edge computing. ActiveMQ (Classic & Artemis): El estándar histórico en el mundo Java/Enterprise. ZeroMQ: No es un servidor, sino una
librería de sockets inteligentes para comunicación masiva sin broker central (Brokerless).

### C. Event Streaming (Logs distribuidos y Replay)

Enfocados en el almacenamiento masivo de eventos y su procesamiento secuencial.

Apache Kafka: El líder absoluto. Basado en el concepto de Append-only Log. Redpanda: Compatible con la API de Kafka pero sin la complejidad de la JVM
(Java Virtual Machine). Apache Pulsar: Arquitectura que separa el almacenamiento (BookKeeper) del cómputo (Broker). Muy escalable. StreamNative:
Versión gestionada de Pulsar.

### D. Servicios Gestionados (Cloud Native)

Para arquitecturas donde se delega la infraestructura al proveedor.

Amazon SQS (Colas) / SNS (Notificaciones): El combo estándar en AWS. Amazon EventBridge: Un Event Bus moderno que permite filtrar y enrutar eventos
entre servicios de terceros y microservicios propios. Azure Service Bus / Event Grid: Los equivalentes en la nube de Microsoft. Google Cloud Pub/Sub:
Mensajería global con escalabilidad automática en GCP.

### E. Procesamiento de Eventos en Tiempo Real (CEP)

Librerías para analizar los flujos de datos mientras ocurren.

Apache Flink: Para procesamiento de flujos de datos complejos y estados. Kafka Streams: Librería para construir aplicaciones que procesan datos dentro
de Kafka. Spark Streaming: Para procesamiento de micro-batches de datos.

## 11. DEVOPS, INFRAESTRUCTURA Y CLOUD

El entorno donde la arquitectura "cobra vida" y se escala. A. Contenedores y Orquestación

Docker / Podman: Empaquetado y aislamiento de aplicaciones. Kubernetes (K8s): El orquestador estándar de la industria. Nomad: Alternativa simplificada
para orquestación de aplicaciones. Helm: Gestor de paquetes para aplicaciones en Kubernetes.

### B. Infraestructura como Código (IaC) y GitOps

Terraform / OpenTofu: Definición de infraestructura agnóstica a la nube. Pulumi: IaC usando lenguajes de programación (TS, Python,
Go). Ansible / Chef / Puppet: Gestión de configuración de servidores. ArgoCD / Flux: Herramientas para despliegue continuo basado en
GitOps.

### C. Proveedores y Modelos Cloud

Hyperscalers: AWS, Microsoft Azure, Google Cloud (GCP). Especializados: DigitalOcean, Vercel, Netlify, Cloudflare (Edge Computing).
Observabilidad: Prometheus, Grafana, ELK Stack (Elasticsearch, Logstash, Kibana), Datadog.


