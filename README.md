# X-Twitter (GO)

---

## Índice
- [Descripción General](#descripcion-general)
- [Arquitectura](#arquitectura)
- [Componentes](#componentes)
- [Tecnologías Utilizadas](#tecnologias-utilizadas)
- [Instrucciones de Uso](#instrucciones-de-uso)
- [Escalabilidad](#escalabilidad)
- [Business Assumptions](#business-assumptions)

---
## Descripción General
X-Twitter es una plataforma de microblogging inspirada en Twitter que permite a los usuarios publicar mensajes cortos, seguir a otros usuarios y ver un timeline con los tweets de las personas a las que siguen.

## Arquitectura
La aplicación sigue una arquitectura limpia (Clean Architecture), que se centra en separar la lógica y mantener la independencia entre capas. A continuación se detallan los principales componentes y capas de la arquitectura:

### Capas Principales
1. **Controllers:** Esta capa maneja las solicitudes HTTP y la interacción con el usuario. Los controladores de usuario y tweets están ubicados aquí.
2. **Servicios (Services):** Contiene la lógica de negocio de la aplicación. Los servicios de usuario y tweet interactúan con la capa de acceso a datos y realizan operaciones específicas.
3. **Acceso a Datos (DB):** Gestiona la interacción con la base de datos. Incluye las interfaces para usuarios y tweets, así como implementaciones concretas.
4. **Base de Datos en Memoria:** Implementación específica de la base de datos que almacena usuarios y tweets en memoria.

### Flujo de Datos
1. **Creación de Usuarios y Tweets:** Describe cómo se crean y almacenan nuevos usuarios y tweets en la base de datos.
2. **Relaciones de Usuario:** Explica cómo se gestionan las relaciones de seguir y ser seguido entre usuarios.
3. **Publicación y Obtención de Tweets:** Detalla el flujo para publicar nuevos tweets y obtener el timeline de un usuario.

### Escalabilidad
La arquitectura está diseñada para escalar a medida que aumenta la cantidad de usuarios. La base de datos en memoria puede ser reemplazada por una solución más robusta en un entorno de producción.

### Consideraciones de Implementación
Esta implementación se ha diseñado para abordar los requisitos fundamentales del challenge, centrándose en la funcionalidad esencial de un microblogging. Sin embargo, es importante tener en cuenta las siguientes consideraciones y limitaciones:

1. **Seguridad:** La solución actual no aborda aspectos críticos de seguridad como la autenticación de usuarios, autorización, y gestión de sesiones. En un entorno de producción, se requeriría una capa de seguridad robusta.
2. **Logueo de Usuarios:** La aplicación actual no cuenta con un sistema de logueo de usuarios. Todos los usuarios son considerados válidos sin autenticación adicional.
3. **Escalabilidad:** Aunque se ha diseñado para ser escalable, la implementación en memoria puede no ser adecuada para entornos de producción con millones de usuarios. En un escenario de producción real, se recomendaría utilizar una base de datos
   persistente y técnicas de escalabilidad.
4. **Manejo de Errores:** La gestión de errores en esta implementación es simplificada y no contempla situaciones excepcionales o errores inesperados.

Estas consideraciones son importantes para futuras mejoras y desarrollo en un entorno de producción real.

## Componentes
La arquitectura de X-Twitter está compuesta por varios componentes que trabajan de manera conjunta para proporcionar la funcionalidad completa de la plataforma de microblogging. A continuación, se describen los componentes clave:

### 1. Servicios

#### Tweet Service
- **Descripción:** Encargado de gestionar la publicación y recuperación de tweets.
- **Funcionalidades:**
  - Publicación de tweets.
  - Obtención de tweets por usuario.
  - Gestión de interacciones de usuarios con tweets.

#### User Service
- **Descripción:** Maneja las operaciones relacionadas con los usuarios, como la creación de usuarios, seguimiento y obtención de seguidores.
- **Funcionalidades:**
  - Creación de usuarios.
  - Seguimiento de usuarios.
  - Obtención de seguidores.

### 2. Controladores

#### Tweet Controller
- **Descripción:** Gestiona las solicitudes HTTP relacionadas con tweets.
- **Funcionalidades:**
  - Manejo de la publicación de tweets.
  - Obtención de tweets por usuario.

#### User Controller
- **Descripción:** Controla las solicitudes HTTP relacionadas con usuarios.
- **Funcionalidades:**
  - Creación de usuarios.
  - Seguimiento de usuarios.
  - Obtención de seguidores.

### 3. Base de Datos (In-Memory)
- **Descripción:** Almacena usuarios y tweets de manera temporal en la memoria.
- **Consideraciones:**
  - Implementación básica para simplificar el desarrollo.
  - No apta para entornos de producción a gran escala.

### 4. Router (Gin)
- **Descripción:** Enrutador HTTP que gestiona las solicitudes y dirige el tráfico a los controladores correspondientes.

Estos componentes trabajan juntos para proporcionar una plataforma funcional de microblogging. En futuras iteraciones, estos componentes podrían escalarse o reemplazarse con implementaciones más robustas según sea necesario.

## Tecnologías Utilizadas
X-Twitter está construido utilizando varias tecnologías que se han seleccionado para proporcionar una base sólida y eficiente para la plataforma. A continuación, se detallan las tecnologías clave utilizadas:

### 1. Lenguaje de Programación: Go
- **Descripción:** Go (o Golang) se eligió como el lenguaje principal de desarrollo debido a su rendimiento, simplicidad y concurrencia integrada. La eficiencia en el uso de recursos y la facilidad de desarrollo en entornos concurrentes hacen que
- Go sea una elección sólida para construir servicios backend.

### 2. Framework Web: Gin
- **Descripción:** Gin es un framework web para Go que se utiliza para gestionar las rutas HTTP y proporcionar un manejo eficiente de solicitudes y respuestas. Su enfoque minimalista y su rendimiento lo hacen adecuado para aplicaciones web simples y rápidas.

### 3. Base de Datos: In-Memory
- **Descripción:** Se ha implementado una base de datos en memoria para simplificar el desarrollo inicial. Aunque es una solución básica y no apta para entornos de producción a gran escala, cumple con los requisitos del challenge y proporciona
- un almacenamiento temporal para usuarios y tweets.

Estas tecnologías fueron seleccionadas para facilitar el desarrollo, mantener la eficiencia y cumplir con los requisitos específicos del challenge. En iteraciones futuras, podría considerarse la adopción de tecnologías más robustas según sea necesario.

## Componentes
La arquitectura de X-Twitter se organiza en varios componentes que trabajan juntos para proporcionar la funcionalidad completa de la plataforma de microblogging. A continuación, se describen los componentes principales:

### 1. Controllers
- **Descripción:** Los controladores manejan las solicitudes HTTP, interpretan los parámetros de la solicitud y llaman a los servicios correspondientes para ejecutar la lógica de negocio. En X-Twitter, hay controladores para gestionar usuarios y tweets.

### 2. Services
- **Descripción:** Los servicios contienen la lógica de negocio de la aplicación. En X-Twitter, existen servicios separados para usuarios y tweets. Estos servicios interactúan con la base de datos y aplican reglas de negocio, como la validación de tweets
- y la gestión de seguidores.

### 3. Base de Datos en Memoria
- **Descripción:** Para simplificar el desarrollo inicial, se ha implementado una base de datos en memoria. Almacena información sobre usuarios y tweets. A tener en cuenta, esta implementación es básica y no está diseñada para entornos de producción a
- gran escala.

### 4. Gin (Framework Web)
- **Descripción:** Gin se utiliza como el framework web principal para manejar las rutas HTTP. Interactúa con los controladores y facilita la gestión de solicitudes y respuestas HTTP.

Esta estructura modular permite una fácil extensión y mantenimiento. Cada componente tiene responsabilidades específicas y se comunica con otros componentes según sea necesario.

## Instrucciones de Uso
A continuación, se detallan los pasos para ejecutar X-Twitter localmente:

### Prerequisitos
Asegúrate de tener instalado Go en tu sistema. Puedes descargarlo [aquí](https://golang.org/dl/).

### Pasos
1. Clona este repositorio:
2. Navega al directorio del proyecto Ej: cd X-Twitter
3. Ejecuta la aplicación Ej: cd cmd
   go run main.go
   La aplicación estará disponible en http://localhost:8080.
   Mediante Postman podrás realizar las peticiones necesarias para probar la funcionalidad.

## Instrucciones de Uso

1. **Pruebas en Postman:**
   - Para facilitar las pruebas, he creado una colección de Postman que puedes descargar [aquí](https://universal-spaceship-200633.postman.co/workspace/New-Team-Workspace~8c8f3f5b-5dbf-423d-8ed2-d8eb314bf076/folder/15549609-d9eaf677-91d6-4b5c-a100-aa7faf896647?ctx=documentation).

   - Importa la colección en tu aplicación Postman y encontrarás ejemplos de solicitudes para probar las diferentes funcionalidades de X-Twitter.

2. **Ejecución Local:**
   - Si prefieres ejecutar la aplicación localmente, sigue estos pasos:

     # Clona el repositorio

     # Navega al directorio del proyecto
     cd X-Twitter

     # Ejecuta la aplicación
     go run main.go

   - La aplicación estará disponible en `http://localhost:8080`.
   - Endpoints
   - Creación de usuarios:
   - http://localhost:8080/create_user
   - body: {"user_id": 3, "username": "user3"}
   - Publicar tweeter:
   - http://localhost:8080/publish_tweet
   - body: {"user_id": 3, "content": "Este es un tweet del user3"}
   - Seguir a un usuario:
   - http://localhost:8080/follow_user
   - body: {"user_id": 3, "follower_id": 1}
   - Obtener tweets:
   - http://localhost:8080/get_tweets/3
   - Obtener seguidores:
   - http://localhost:8080/get_followers/3


## Escalabilidad
X-Twitter ha sido diseñada con la escalabilidad en mente para manejar un crecimiento significativo de usuarios y tweets. Algunas estrategias para la escalabilidad podrían ser:

1. **Base de Datos Distribuida:** Se podría utilizar una base de datos distribuida que permita escalar horizontalmente a medida que aumenta la cantidad de datos y usuarios.
2. **Caché:** Se puede implementar un sistema de caché para reducir la carga en la base de datos y mejorar los tiempos de respuesta.
3. **Balanceo de Carga:** Para distribuir equitativamente las solicitudes, podría emplearse un balanceador de carga que redirija el tráfico entre varios nodos del servidor.
4. **Microservicios:** La arquitectura de microservicios permite escalar componentes específicos de la aplicación de manera independiente, por lo que sería una gran elección.

Estas estrategias aseguran que X-Twitter pueda manejar un alto número de usuarios concurrentes y mantener un rendimiento óptimo de ser necesario.

## Business Assumptions
- **Optimización para Operaciones de Lectura:** La solución actual, aunque funcional, no se considera optimizada para un entorno de producción con un gran número de usuarios y tweets. Se asume que, para alcanzar un rendimiento óptimo en operaciones
- de lectura, se requerirán mejoras significativas, como:
  - **Uso de Caching:** Implementación de un sistema de caché para reducir la latencia en la recuperación de tweets frecuentemente consultados.
  - **Índices Eficientes en la Base de Datos:** Creación de índices eficientes para acelerar las consultas de búsqueda, especialmente en operaciones relacionadas con el timeline y seguidores.
  - **Escalabilidad Horizontal:** Consideración de estrategias de escalabilidad horizontal para distribuir la carga de lectura entre múltiples nodos o servidores.
  - **Almacenamiento de Datos Denormalizado:** Exploración de enfoques de almacenamiento denormalizado para reducir la complejidad en las consultas de lectura.

Estas mejoras deben abordarse en fases posteriores del desarrollo para garantizar un rendimiento óptimo y escalabilidad a medida que la plataforma crece en usuarios y actividades.
