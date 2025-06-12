package main

import (
	"C"
	"github.com/joho/godotenv"
	"os"
)

//export LoadEnvFile
func LoadEnvFile() C.int {
	err := godotenv.Load()
	if err != nil {
		return 1 // Error al cargar el archivo
	}
	return 0 // Éxito
}

//export GetEnvVar
func GetEnvVar(key *C.char) *C.char {
	keyStr := C.GoString(key)
	value := os.Getenv(keyStr)
	return C.CString(value)
}

func main(){}