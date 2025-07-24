# Instalaciones

01. Instalamos "Postgre-SQL" para lo cual hacemos lo siguiente:
	01.01 Accedemos al siguiente link:
		01.01.01 https://www.postgresql.org/download/
		01.01.02 Seleccionamos nuestro SO.
			01.01.02.01 Windows
		01.01.03 Damos click en el hiperbinculo "Download the installer"
		01.01.04 Seleccionamos la vercion que mas se adecue a nuestro SO.
			01.01.04.01 En este caso selaccionamos la vercion 17.5 para el SO Windows.
		01.01.05 Ejecutamos el .exe "postgresql-17.5-3-windows-x64.exe".

02. Procedemos a instalar "Docker" para lo cual hacemos lo siguiente:
	02.01 Accedemos al siguiente link:
		02.01.01 https://www.docker.com/
	02.02 Damos click en el boton "Download Docker Desktop".
	02.03 Seleccionamos la opcion de nuestro sistema operativo.
	02.04 Ejecutamos el .exe "Docker Desktop Installer.exe".
	02.05 Damos click en el boton "Ok".
	02.06 Damos click en el boton "Close and restart".
	02.07 Iniciamos la aplicacion.
	02.08 Si nos pide restart, copiamos el comando que nos muestra en pantalla y lo ejecutamos en power shell
	02.09 Procedemos a ejecutar Docker.

03. Instalamos "Go" para lo cual hacemos lo siguiente:
	03.01 Accedemos al siguiente link :
			03.01.01 https://go.dev/
		03.03 Damos click en el boton "Download".
		03.03 Seleccionamos el sistema operativo en el que queremos instalar Go.
			03.03.01 En este caso "Microsoft Windows".
		03.04 Ejecutamos el instalador "go1.24.5.windows-amd64.msi".
		03.05 Damos click en el boton "Siguiente".
		03.06 Damos click en el boton "Siguiente".
		03.07 Damos click en el boton "Siguiente".
		03.08 Damos click en el boton "Install".

04. Para saber si tenemos Go instalado hacemos lo siguiente:
	04.01 Abrimos el "cmd".
	04.02 Ejecutamos el siguiente comando:
		04.02.01 go version

05. Procedemos a configurar las variables de entorno para lo cual hacemos lo siguiente:
	05.01 Abrimos el explorador de carpetas.
	05.02 Buscamos "Este equipo".
	05.03 Damos click derecho.
	05.04 Propiedades.
	05.05 Buscamos "Configuración avanzada del sistema".
	05.06 Seleccionamos la opcion "Ver la configuracion avanzada del sistema".
	05.07 Damos click en el boton "Variables de entorno".
	05.08 Buscamos la variable de entorno "GOPATH" que nos indica donde estara ubicada la carpeta de trabajo.
		05.08.01 Si no existe la variable de entorno "GOPATH" la creamos.
		05.08.02 Indicamos la ubicacion de la variable "GOPATH" en este caso:
			05.08.02.01 C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src
	05.09 Abrimos "cmd" y accedemos a la ubicacion "C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go" y agregamos los directorios "bin", "pkg" y "src" utilizando el siguiente comando:
		05.09.01 mkdir bin
		05.09.02 mkdir pkg
		05.09.03 mkdir src
	05.10 Abrimos "cmd" y accedemos a la ubicacion "C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src" y agregamos el directorio "github.com" con el siguiente comando:
		05.10.01 mkdir github.com
	05.11 Abrimos "cmd" y accedemos a la ubicacion "C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com" y agregamos el directorio "Jucarios1987" con el siguiente comando:
		05.11.01 La carpeta creada debe tener el mismo nombre que mi usuario de "github".
		05.11.02 mkdir Jucarios1987

06. Procedemos a instalar git para lo cual hacemos lo siguiente:
	06.01 Accedemos al siguiente link:
		06.01.01 https://git-scm.com/
	06.02 Damos click en el boton "Download for Windows".
	06.03 Seleccionamos la opcion de descarga para Windows a 62 bits.
		06.03.01 Git for Windows/x64 Setup.
	06.04 Guardamos el .exe en la carpeta de descargas:
		06.04.01 Git-2.50.1-64-bit.exe
	06.05 Ejecutamos el .exe "Git-2.50.1-64-bit.exe".
	06.06 Damos click en el boton "Install".

07. Accedemos a la pagina de Git Hub "https://github.com/" y hacemos lo siguiente:
	07.01 Creamos una cuenta e ingresamos a esta.
	07.02 Creamos un nuevo repositorio.
		07.02.01 Nombre: transaccionesBancarias
		07.02.02 Descripcion: Repositorio microservicio que se encarga de registrar y consultar transacciones bancarias.
		07.02.03 Configuramos el repositorio para que sea "Publico".
		07.02.04 Damos click en el boton "Crear repositorio".

08. Procedemos a clonar el repositorio "transaccionesBancarias" en nuestra PC paralo cual hacemos lo siguiente:	
	08.01 Accedemos al repositorio "transaccionesBancarias" en GitHub.
	08.02 Copiamos el link de la seccion "Quick setup — if you’ve done this kind of thing before":
		08.02.01 https://github.com/Jucarios1987/transaccionesBancarias.git
	08.03 Abrimos el "cmd".
	08.04 hacemos un cd a la carpeta que creamos con el nombre de nuestro usuario de GitHub "Jucarios1987".
		08.04.01 cd C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com\Jucarios1987
	08.05 Introducimos el siguiente comando "cmd" en la ruta "C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com\Jucarios1987"
		08.05.01 git clone https://github.com/Jucarios1987/transaccionesBancarias.git
		08.05.02 Presionamos enter.
	08.06 Ejecutamos el siguiente comando en el "cmd" en la siguiente ubicacion "C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com\Jucarios1987".
		08.06.01 dir
	08.07 Nos ubicamos en la carpeta "C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com\Jucarios1987\transaccionesBancarias" en el "cmd" y ejecutamos el siguiente comando:
		08.07.01 git status
		08.07.02 Obtenemos la siguiente informacion:
			08.07.02.01 Nombre de nuestra rama principal.
			08.07.02.02 Informacion de  cambios pendientes.

09. Procedemos a instalar Visual Studio Code para lo cual hacemos lo siguiente:
	09.01 Accedemos al siguiente link:
		09.01.01 https://code.visualstudio.com/
	09.02 Damos click en el boton "Download for Windows".

10. Procedemos a abrir "Visual Studio Code" y hacemos lo siguiente:
	10.01 Click en el boton "Archivo".
	10.02 Click "Open FIle"
	10.03 Abrimos la ubicacion de la carpeta del repositorio local "transaccionesBancarias"

11. Procedemos a instalar las siguientes extenciones:
	11.01 Go
	11.02 Go Nightly
	11.03 Go Doc
	11.04 Go Test Explorer
	11.05 Go Snippets
	11.06 Go Grammar
	11.07 Go Syntax

12. Procedemos a abrir una terminal dentro de visual studio code y hacemos lo siguiente:
	12.01 Ejecutamos el siguiente comando:
		12.01.01 git status

# Crear-API-Rest-Con-Go

01. Para crear un nuevo proyecto de Go hacemos lo siguiente:
	01.01 Abrimos visual studio code en la siguiente ubicacion:
		01.01.01 C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\github.com\Jucarios1987\transaccionesBancarias> 
	01.02 Abrimos la consola.
	01.03 Ejecutamos el siguiente comando. La ruta que adicionamos al comando hace referencia a la ruta apartir de la carpeta github.com hasta la carpeta del repositorio destino:
		01.03.01 go mod init github.com/Jucarios1987/transaccionesBancarias
	01.04 Se crea el archivo "go.mod"
	01.05 Procedemos a agregar el archivo "main.go" a la raiz del proyecto justo en el mismo nivel del archivo "go.mod"
	01.06 Creamos la carpeta "variables" en la raiz de nuestro proyecto "transaccionesBancarias".
		01.06.01 Creamos el archivo "enteras.go" el cual contendra todas la variables de este tipo.
		01.06.02 Definimos el package del archivo creado.
		01.06.03 Si el archivo ".go" creado esta en la raiz del proyecto el "package" se nombra "package main".
		01.03.04 Si el archivo ".go" creado esta en una sub carpeta del proyecto raiz el "package" se nombra como la carpeta contenedora "package variables".

02. Debemos importar un paquete de go que se utiliza para enviar texto a la consola, para lo cual hacemos lo siguiente:
	02.01 Nos ubicamos en el archivo "main.go"
	02.02 Agregamos la funcionalidad para que se imprima 'Hola Juan Go' en el archivo "main.go".

03. Procedemos a ejecutar nuestro programa para lo cual hacemos lo siguiente:
	03.01 Abrimos la terminal de Visual Studio Code en la hubicacion de nuestra carpeta raiz.
	03.02 Ejecutamos el siguiente comando.
		03.02.01 go run main.go

04. La estructura de todos los modulos o archivos de "go" es la siguiente:
	04.01 Se define el package.
	04.02 Se definen las librerias o modulos que utilizaremos en "import".
	04.03 Adicionamos el resto de funcionalidades al modulo (declaraciones de variables, funciones,..., etc)

05. Utilizamos el modulo "gorilla max" para poder crear un servidor, para lo cual hacemos lo siguiente:
	05.01 Accedemos al siguiente link:
		05.01.01 https://github.com/gorilla/mux
	05.02 Buscamos la seccion "Install" y copiamos el siguiente comando:
		05.02.01 go get -u github.com/gorilla/mux
	05.03 Abrimos nuestro proyecto de "go" en Visual Studio Code.
	05.04 Abrimos una terminal y ejecutamos el comando "go get -u github.com/gorilla/mux".
	
06. Creamos una carpeta llamada "routes" en la raiz de nuestro proyecto.
	06.01 Creamos el archivo "index.routes.go" dentro de la carpeta "routes".
		06.01.01 Importamos el modulo HTTP.
		06.01.02 Creamos la funcion "HomeHandler".
	06.02 Procedemos a crear un router y adicionamos su funcionalidad en el archivo "main.go" de nuestro proyecto.

07. Procedemos a ejecutar y probar nuestra app para lo cual hacemos lo siguiente:
	07.01 Abrimos una nueva terminal desde Visual Studio Code y ejecutamos el siguiente comando:
		07.01.01 go run .
	07.02 Abrimos el siguiente link en el navegador:
		07.02.01 http://localhost:3000/
		07.02.02 El puerto que utilizamos depende del que se definio cuando inicializamos el servidor en el archivo "main.go"

08. Procedemos a intalar el modulo "" que nos permite reiniciar la consola cuando hagamos un cambio para lo cual hacemos lo siguiente:
	08.01 Ejecutamos el siguiente comando en la consola de nuestro proyecto:
		08.01.01 go install github.com/cosmtrek/air@latest
	08.02 Abrimos nuestro proyecto en Visual Studio Code
	08.03 Abrimos una nueva terminal.
	08.04 Ejecutamos el comando "go install github.com/cosmtrek/air@latest"
	08.05 Comprobamos que en nuestras variables de entorno "Path" este definida la siguiente ruta:
		08.05.01 C:\01-Plan_Carrera\08_Golang\04_Codigo\99_Go\src\bin
		08.05.02 Esta ruta cambia dependiendo del lugar en el que se configura la variable de entrono "GOPATH" durante la instalacion y configuracion de "Go".
	08.06 Reiniciamos Visual Studio Code
	08.05 Ejecutamos el comando "air init"
	08.06 Ahora para ejecutar nuestro proyecto solo debemos abrir una nueva terminal y ejecutar el siguiente comando:
		08.06.01 air
	08.07 Mientras el comando "Air" se este ejecutando podemos hacer cambios en nuestra app y al guardar estos se veran reflejados sin necesidad de parar la ejecucion de esta.

09. Adicionamos el archivo "cuentas.routes.go" encargado de manejar las rutas para el CRUD de cuentas bancarias a la carpeta "routes"

10. Buscar la libreria "gorm" de "Go" que nos permite conectarnos a dibersas bases de datos es decir es un "ORM" para lo cual accedemos al siguiente link:
	10.01 https://gorm.io/index.html
	10.02 Ingresamos a la documentacion dando click en el boton "->" lo cual nos lleva a la siguiente pagina:
		10.02.01 https://gorm.io/docs/
	10.03 Buscamos la seccion "Install" y ejecutamos el comando para instalacion de "gorm":
		10.03.01 go get -u gorm.io/gorm
	10.04 Despues procedemos a instalar el driver de nuesto sistema de gestion de base de datos para lo cual hacemos lo siguiente:
		10.04.01 Buscamos la seccion "Getting Started" del menu verticla izquierdo.
		10.04.02 Damos click en la opcion "Connecting to Database" lo cual nos abre la siguiente ubicacion:
			10.04.02.01 https://gorm.io/docs/connecting_to_the_database.html
		10.04.03 Buscamos la seccion de la base de datos a la cual nos conectaremos en este caso "PostgreSQL".
		10.04.04 Copiamos el nombre del modulo que nos interesa:
			10.04.04.01 "gorm.io/driver/postgres"
		10.04.05 Ejecutamos el comando de instalacion del drive segun el nombre del mudulo que obtubimos en el numeral [10.04.04.01] en la terminal de nuestra app de go:
			10.04.05.01 go get -u gorm.io/driver/postgres

11. Creamos una instancia de "Postgre-SQL" en docker, para lo cual hacemos lo siguiente:
	11.01 Abrimos nuestro proyecto de "Go" en Visual Studio Code.
	11.02 Abrimos una nueva terminal.
	11.03 Ejecutamos el siguiente comando:
		11.03.01 docker run --name some-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1087989233 -p 5432:5432 -d postgres
	11.04 Para comprobar la instalacion del contenedor ejecutamos el siguiente comando:
		11.04.01 docker ps

12. Comandos utiles para manipular docker:
	12.01 Para detener un contenedor:
		12.01.01 docker stop [NameContenedor]
	12.02 Eliminar un contenedor:
		12.02.01 docker rm [NameContenedor]
	12.03 Crear un contenedor para instanciar base de dator Postgres:
		12.03.01 docker run --name some-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1087989233 -p 5432:5432 -d postgres
	12.04 Listar los contenedores  existentes:
		01.04.01 docker ps

13. Creamos una carpeta llamada "db" para el manejo de las conexiones a la base de datos en la raiz de nuestro proyecto.
	13.01 Dentro de la carpeta "db" creamos el archivo "connection.go".
	13.02 Creamos la funcion "DBConnection" que nos permitira conectarnos a la DB dentro del archivo "connection.go".

14. Procedmos a crear la base de datos en el contenedor "some-postgres" de docker para lo cual hacemos lo siguiente:
	14.01 Abrimos nuestro proyecto de Go en Visual Studio Code.
	14.02 Abrimos una nueva terminal y ejecutamos el siguiente comando pasa obtener el nombre del contenedor:
		14.02.01 docker ps
	14.03 Ejecutamos el contenedor "some-postgres" en modo interactivo para lo cual ejecutamos el siguiente comando:
		14.03.01 docker exec -it some-postgres bash
	14.04 Nos conectamos a postgres ejecutando el siguiente comando:
		14.04.01 psql -U postgres --password
		14.04.02 Ingresamos la contraseña: 1087989233
	14.05 Para ver que bases de datos tengo en el contenedor utilizamos el siguiente comando:
		14.05.01 \l
	14.06 Procedemos a crear una nueva base de datos para lo cual utilizamos el siguiente comando:
		14.06.01 CREATE DATABASE cuentasbancarias;
	14.07 Procedemos a conectarnos a la base de datos para lo cual ejecutamos el siguiente comado:
		14.07.01 \c cuentasbancarias
		14.07.02 Ingresamos la contraseña: 1087989233

15. Nos ubicamos en el archivo "main.go" y adicionamos la funcionalidad para que inicie la coneccion a la base de datos.

16. Ejecutamos el siguiente comando para iniciar nuestro proyecto:
	16.01 go run .

17. Si tenemos problema con los puertos hacemos lo siguiente:
	17.01 Abrimos nuestro proyecto en Visual Studio Code.
	17.02 Abrimos una nueva terminal.
	17.03 Ejecutamos los siguientes comados:
		17.03.01 docker stop some-postgres
		17.03.02 docker rm some-postgres
		17.03.03 docker run --name some-postgres -e POSTGRES_PASSWORD=1087989233 -e POSTGRES_DB=cuentasbancarias -p 5433:5432 -d postgres

18. Creamos la carpeta "models" en la raiz del proyecto.
	18.01 Procedemos a crear el archivo "accounts.go" dentro de la carpeta "models".
		18.01.01 Procedemos a agregar la funcionalidad en nuestro modelo "accounts.go".
	18.02 Procedemos a crear el archivo "transactions.go" dentro de la carpeta "models".
		18.02.01 Procedemos a agregar la funcionalidad en nuestro modelo "transactions.go".

19. Nos ubicamos en el archivo "main.go" y configuramos la ejecucion de las migraciones para crear las tablas en la DB.

20. Procedemos a crear los archivos para el manejo de las rutas para las peticiones http de nuestra app "accounts.routes.go" y "transactions.routes.go" en la carpeta "routes".
	20.01 Agregamos las funciones del CRUD.

21. Inicializamos las rutas para menejo de cuentras y transacciones en nuestro archivo "main.go"

22. Instalamos la extencion "Thunder Client" en visual studio code.
	22.01 Podemos probar nuestro sistema de rutas de la siguiente manera:
		22.01.01 Damos click en el boton "Thunder Client" del menu lateral izquierdo.
		22.01.02 Damos click en el boton "New Request"
		22.01.03 Seleccionamos el tipo de peticion "GET", "POST" o "DELETE"
		22.01.04 Agregamos la url:
			22.01.04.01 http://localhost:3000/accounts
			22.01.04.02 http://localhost:3000/transactions
		22.01.05 Damos click en el boton "Send".

23. Procedemos a agregar la funcionalidad de las peticiones HTTP asi:
	23.01 Nos ubicamos en el archivo "accounts.routes.go" que esta dentro de la carpeta "routes" y agregamos la funcionalidad de los metodos.
	23.02 Nos ubicamos en el archivo "transactions.routes.go" que esta dentro de la carpeta "routes" y agregamos la funcionalidad de los metodos.

24. Para probar las solisitudes HTTP hacemos lo siguiente:
	24.01 Damos click en el boton "Thunder Client" del menu lateral izquierdo.
	24.02 Damos click en el boton "New Request"
	24.03 Agregamos la url y seleccionamos tipo de solicitud "POST":
		24.03.01 http://localhost:3000/accounts
			24.03.01.01 Damos click en "body"
			24.03.01.02 Adicionamos el JSON de la solicitud:
						{
						  "name" : "MarcelaAccount"
						}
			24.03.01.03 Damos click en el boton "Send".
		24.03.02 http://localhost:3000/transactions
			24.03.02.01 Damos click en "body"
			24.03.02.02 Adicionamos el JSON de la solicitud:
						{
						  "account_id" : 36,
						  "amount" : 731.3,
						  "currency" : "EUR",
						  "type" : "deposit",
						  "description" : "Deposito"
						}
			24.03.02.03 Damos click en el boton "Send".

# Estructura-Proyectos-Go

01. Al crear el proyecto con el comando "go mod init github.com/Jucarios1987/transaccionesBancarias" se crea el archivo "go.mod"

02. Despues de crear el programa debemos crear un archivo llamado "main.go" que sera el encargado de inicializar nuestra app.

03. Para Crear una API Rest debemos hacer lo siguiente:
	03.01 Creamos una carpeta llamada "routes" en la raiz de nuestro proyecto.
		03.01.01 Creamos el archivo "index.routes.go" dentro de la carpeta "routes".
			03.01.01.01 Creamos la funcion "HomeHandler" en el archivo "index.routes.go".
		03.01.02 Dentro de la carpeta "routes" van todos los archivos que necesitemos para definir las rutas que tendra nuestra app.
		03.01.03 En este caso agregaremos el archivo "cuentas.routes.go"
	03.02 Creamos una carpeta llamada "db" para el manejo de las conexiones a la base de datos en la raiz de nuestro proyecto.
		03.02.01 Dentro de la carpeta "db" creamos el archivo "connection.go".
		03.02.02 Agregasmo la funcionalidad para la coneccion a la base de datos.
	03.03 Nos ubicamos en el archivo "main.go" y adicionamos la funcionalidad para que inicie la coneccion a la base de datos.
	03.04 Creamos la carpeta "models" en la raiz del proyecto.
		03.04.01 Procedemos a crear el archivo "Accounts.go" dentro de la carpeta "models".
			03.04.01.01 Procedemos a agregar la funcionalidad en nuestro modelo "Accounts.go".
		03.04.02 Procedemos a crear el archivo "Transactions.go" dentro de la carpeta "models".
			03.04.02.01 Procedemos a agregar la funcionalidad en nuestro modelo "Transactions.go".
	03.05 Nos ubicamos en el archivo "main.go" y configuramos la ejecucion de las migraciones para crear las tablas en la DB.
	03.06 Procedemos a crear los archivos para el manejo de las rutas para las peticiones http de nuestra app "accounts.routes.go" y "transactions.routes.go" en la carpeta "routes".
		03.06.01 Agregamos las funciones del CRUD.
	03.07 Inicializamos las rutas para menejo de cuentras y transacciones en nuestro archivo "main.go"
	03.08 Instalamos la extencion "Thunder Client" en visual studio code.
	03.09 Podemos probar nuestro sistema de rutas de la siguiente manera:
		03.09.01 Damos click en el boton "Thunder Client" del menu lateral izquierdo.
		03.09.02 Damos click en el boton "New Request"
		03.09.03 Seleccionamos el tipo de peticion "GET", "POST" o "DELETE"
		03.09.04 Agregamos la url:
			03.09.04.01 http://localhost:3000/accounts
			03.09.04.02 http://localhost:3000/transactions
		03.09.05 Damos click en el boton "Send".
	03.10 Procedemos a agregar la funcionalidad de las peticiones HTTP asi:
		03.10.01 Nos ubicamos en el archivo "accounts.routes.go" que esta dentro de la carpeta "routes" y agregamos la funcionalidad de los metodos.
		03.10.02 Nos ubicamos en el archivo "transactions.routes.go" que esta dentro de la carpeta "routes" y agregamos la funcionalidad de los metodos.
	03.11 Para probar las solisitudes HTTP hacemos lo siguiente:
		03.11.01 Damos click en el boton "Thunder Client" del menu lateral izquierdo.
		03.11.02 Damos click en el boton "New Request"
		03.11.03 Agregamos la url y seleccionamos tipo de solicitud "POST":
			03.11.03.01 http://localhost:3000/accounts
				03.11.03.01.01 Damos click en "body"
				03.11.03.01.02 Adicionamos el JSON de la solicitud:
							{
							  "name" : "MarcelaAccount"
							}
				03.11.03.01.03 Damos click en el boton "Send".
			03.11.03.02 http://localhost:3000/transactions
				03.11.03.02.01 Damos click en "body"
				03.11.03.02.02 Adicionamos el JSON de la solicitud:
							{
							  "account_id" : 36,
							  "amount" : 731.3,
							  "currency" : "EUR",
							  "type" : "deposit",
							  "description" : "Deposito"
							}
				03.11.03.02.03 Damos click en el boton "Send".