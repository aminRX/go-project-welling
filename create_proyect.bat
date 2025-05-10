@echo off
set "ROOT=go-microservices"

rem Crear carpetas
mkdir %ROOT%
cd %ROOT%
mkdir cmd\server
mkdir config
mkdir controllers
mkdir db\migrations
mkdir middleware
mkdir models
mkdir routes
mkdir services
mkdir utils

rem Crear archivos vacíos
echo // Entrypoint >> cmd\server\main.go
echo // Carga de configuración >> config\config.go
echo // Handlers HTTP >> controllers\user_controller.go
echo // Lógica de conexión a DB >> db\postgres.go
echo -- SQL Migrations >> db\migrations\001_init.sql
echo // Middlewares >> middleware\auth.go
echo // Modelos de GORM >> models\user.go
echo // Agrupación de rutas >> routes\routes.go
echo // Lógica de negocio >> services\user_service.go
echo // Funciones auxiliares >> utils\hash.go

rem Archivos raíz
echo module go-microservices> go.mod
echo // Main entrypoint >> main.go

echo Proyecto creado exitosamente.
pause
