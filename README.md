# Instalaciones

## 1. Instalación de PostgreSQL
1. Acceder al siguiente link:  
   https://www.postgresql.org/download/
2. Seleccionar nuestro SO:
   1. Windows
3. Dar click en el hipervínculo "Download the installer"
4. Seleccionar la versión que más se adecue a nuestro SO:
   1. En este caso seleccionamos la versión 17.5 para el SO Windows
5. Ejecutar el .exe "postgresql-17.5-3-windows-x64.exe"

## 2. Instalación de Docker
1. Acceder al siguiente link:  
   https://www.docker.com/
2. Dar click en el botón "Download Docker Desktop"
3. Seleccionar la opción de nuestro sistema operativo
4. Ejecutar el .exe "Docker Desktop Installer.exe"
5. Dar click en el botón "Ok"
6. Dar click en el botón "Close and restart"
7. Iniciar la aplicación
8. Si nos pide restart, copiar el comando que nos muestra en pantalla y ejecutarlo en power shell
9. Proceder a ejecutar Docker

## 3. Instalación de Go
1. Acceder al siguiente link:  
   https://go.dev/
2. Dar click en el botón "Download"
3. Seleccionar el sistema operativo:
   1. En este caso "Microsoft Windows"
4. Ejecutar el instalador "go1.24.5.windows-amd64.msi"
5. Dar click en el botón "Siguiente"
6. Dar click en el botón "Siguiente"
7. Dar click en el botón "Siguiente"
8. Dar click en el botón "Install"

## 4. Verificar instalación de Go
1. Abrir el "cmd"
2. Ejecutar el siguiente comando:  
   `go version`

## 5. Configuración de variables de entorno
1. Abrir el explorador de carpetas
2. Buscar "Este equipo"
3. Dar click derecho
4. Propiedades
5. Buscar "Configuración avanzada del sistema"
6. Seleccionar la opción "Ver la configuración avanzada del sistema"
7. Dar click en el botón "Variables de entorno"
8. Buscar la variable de entorno "GOPATH":
   1. Si no existe, crearla
   2. Indicar la ubicación:  
      `C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src`
9. En "cmd", acceder a la ubicación y crear directorios:
   1. `mkdir bin`
   2. `mkdir pkg`
   3. `mkdir src`
10. Crear directorio github.com:
    1. `mkdir github.com`
11. Crear directorio con nombre de usuario GitHub:
    1. `mkdir Jucarios1987`

## 6. Instalación de Git
1. Acceder al siguiente link:  
   https://git-scm.com/
2. Dar click en el botón "Download for Windows"
3. Seleccionar la opción de descarga para Windows a 64 bits:
   1. Git for Windows/x64 Setup
4. Guardar el .exe en la carpeta de descargas:
   1. Git-2.50.1-64-bit.exe
5. Ejecutar el .exe "Git-2.50.1-64-bit.exe"
6. Dar click en el botón "Install"

## 7. Configuración de GitHub
1. Acceder a:  
   https://github.com/
2. Crear una cuenta e ingresar
3. Crear nuevo repositorio:
   1. Nombre: transaccionesBancarias
   2. Descripción: Repositorio microservicio que se encarga de registrar y consultar transacciones bancarias
   3. Configurar como "Público"
   4. Dar click en "Crear repositorio"

## 8. Clonar repositorio
1. Acceder al repositorio en GitHub
2. Copiar el link de la sección "Quick setup"
   1. `https://github.com/Jucarios1987/transaccionesBancarias.git`
3. Abrir "cmd"
4. Navegar a la carpeta del usuario GitHub:
   1. `cd C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com\Jucarios1987`
5. Clonar repositorio:
   1. `git clone https://github.com/Jucarios1987/transaccionesBancarias.git`
6. Verificar contenido:
   1. `dir`
7. Verificar estado:
   1. `git status`

## 9. Instalación de Visual Studio Code
1. Acceder al siguiente link:  
   https://code.visualstudio.com/
2. Dar click en el botón "Download for Windows"

## 10. Configuración de Visual Studio Code
1. Abrir VS Code
2. Click en "Archivo" > "Open File"
3. Abrir la carpeta del repositorio local

## 11. Instalación de extensiones
1. Go
2. Go Nightly
3. Go Doc
4. Go Test Explorer
5. Go Snippets
6. Go Grammar
7. Go Syntax

## 12. Uso de terminal en VS Code
1. Abrir terminal
2. Ejecutar:  
   `git status`

# Crear API Rest con Go

## 1. Creación de proyecto
1. Abrir VS Code en la ubicación del proyecto
2. Ejecutar en consola:  
   `go mod init github.com/Jucarios1987/transaccionesBancarias`
3. Se crea archivo "go.mod"
4. Crear archivo "main.go" en la raíz
5. Crear carpeta "variables":
   1. Crear archivo "enteras.go"

## 2. Configuración básica
1. En "main.go", agregar funcionalidad para imprimir mensaje
2. Ejecutar programa:  
   `go run main.go`

## 3. Estructura de módulos
1. Definir package
2. Definir imports
3. Agregar funcionalidades

## 4. Uso de gorilla/mux
1. Acceder a:  
   https://github.com/gorilla/mux
2. Instalar:  
   `go get -u github.com/gorilla/mux`

## 5. Configuración de rutas
1. Crear carpeta "routes"
   1. Crear archivo "index.routes.go"
      1. Importar módulo HTTP
      2. Crear función "HomeHandler"
2. Configurar router en "main.go"

[... continuación con el resto de secciones ...]

# Estructura de Proyectos Go

## 1. Archivos básicos
1. `go.mod` (creado con `go mod init`)
2. `main.go` (punto de entrada)

## 2. Estructura de carpetas
1. `routes/` - Manejo de rutas
   1. `index.routes.go`
   2. `accounts.routes.go`
   3. `transactions.routes.go`
2. `db/` - Conexiones a base de datos
   1. `connection.go`
3. `models/` - Modelos de datos
   1. `accounts.go`
   2. `transactions.go`

## 3. Flujo de trabajo
1. Configurar conexión a DB
2. Definir modelos
3. Configurar migraciones
4. Definir rutas y handlers
5. Probar con Thunder Client