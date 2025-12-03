package common

func (e *ErrorInfo) Error() string {
	return e.Message
}

const (
	// CÓDIGOS DE ÉXITO
	HTTP_OK         = "200" // GET exitoso
	HTTP_CREATED    = "201" // POST exitoso (recurso creado)
	HTTP_NO_CONTENT = "204" // DELETE exitoso

	// CÓDIGOS DE ERROR DEL CLIENTE
	HTTP_BAD_REQUEST  = "400" // Datos mal formados
	HTTP_UNAUTHORIZED = "401" // Sin login/token
	HTTP_FORBIDDEN    = "403" // Sin permisos
	HTTP_NOT_FOUND    = "404" // Recurso no existe
	HTTP_CONFLICT     = "409" // Recurso duplicado
	HTTP_VALIDATION   = "422" // Datos inválidos pero JSON correcto

	// CÓDIGOS DE ERROR DEL SERVIDOR
	HTTP_SERVER_ERROR = "500" // Error interno
	HTTP_UNAVAILABLE  = "503" // Servicio no disponible
)

const (
	// ERRORES DE VALIDACIÓN - Usar cuando los datos del cliente están mal
	ERR_VALIDATION     = "VALIDATION_ERROR" // Campo inválido
	ERR_REQUIRED_FIELD = "REQUIRED_FIELD"   // Campo obligatorio faltante
	ERR_INVALID_FORMAT = "INVALID_FORMAT"   // Email, teléfono, fecha mal formada
	ERR_INVALID_JSON   = "INVALID_JSON"     // JSON syntax error

	// ERRORES DE AUTENTICACIÓN - Usar en login/registro
	ERR_UNAUTHORIZED  = "UNAUTHORIZED"        // Sin token o credenciales
	ERR_FORBIDDEN     = "FORBIDDEN"           // Token válido pero sin permisos
	ERR_INVALID_TOKEN = "INVALID_TOKEN"       // JWT inválido o malformado
	ERR_EXPIRED_TOKEN = "EXPIRED_TOKEN"       // JWT expirado
	ERR_INVALID_LOGIN = "INVALID_CREDENTIALS" // Usuario/password incorrectos

	// ERRORES DE BASE DE DATOS - Usar en operaciones CRUD
	ERR_NOT_FOUND      = "NOT_FOUND"      // Registro no existe
	ERR_DUPLICATE      = "DUPLICATE"      // Recurso ya existe (email, username)
	ERR_DATABASE_ERROR = "DATABASE_ERROR" // Error de conexión o consulta

	// ERRORES DE NEGOCIO - Usar en lógica específica del dominio
	ERR_INSUFFICIENT_FUNDS = "INSUFFICIENT_FUNDS" // Saldo insuficiente
	ERR_OUT_OF_STOCK       = "OUT_OF_STOCK"       // Producto sin inventario
	ERR_OPERATION_DENIED   = "OPERATION_DENIED"   // Operación no permitida

	// ERRORES DE SISTEMA - Usar en problemas técnicos
	ERR_INTERNAL_ERROR = "INTERNAL_ERROR" // Error inesperado del servidor
	ERR_SERVICE_ERROR  = "SERVICE_ERROR"  // Servicio externo falló
	ERR_TIMEOUT        = "TIMEOUT"        // Operación tardó demasiado
	ERR_RATE_LIMIT     = "RATE_LIMIT"     // Demasiadas peticiones

	//ERROR DE LOGIN - Uso en auth

)

const (
	// OPERACIONES CRUD
	SUCCESS_RETRIEVED = "SUCCESS_RETRIEVED" // GET exitoso
	SUCCESS_CREATED   = "SUCCESS_CREATED"   // POST exitoso
	SUCCESS_UPDATED   = "SUCCESS_UPDATED"   // PUT/PATCH exitoso
	SUCCESS_DELETED   = "SUCCESS_DELETED"   // DELETE exitoso

	// OPERACIONES DE AUTENTICACIÓN
	SUCCESS_LOGIN    = "SUCCESS_LOGIN"    // Login exitoso
	SUCCESS_LOGOUT   = "SUCCESS_LOGOUT"   // Logout exitoso
	SUCCESS_REGISTER = "SUCCESS_REGISTER" // Registro exitoso

	// OPERACIONES DE NEGOCIO
	SUCCESS_PAYMENT  = "SUCCESS_PAYMENT"  // Pago procesado
	SUCCESS_ORDER    = "SUCCESS_ORDER"    // Pedido creado
	SUCCESS_TRANSFER = "SUCCESS_TRANSFER" // Transferencia completada
)
