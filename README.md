# env
Una biblioteca ligera para leer .env en C.  
Compilada usando: `go build -o env.dll -buildmode=c-shared env.go`

---

### 📥 Descargar la librería

| Linux | Windows |
| --- | --- |
| `wget https://github.com/IngenieroRicardo/env/releases/download/2.0/env.so` | `Invoke-WebRequest https://github.com/IngenieroRicardo/env/releases/download/2.0/env.dll -OutFile ./env.dll` |
| `wget https://github.com/IngenieroRicardo/env/releases/download/2.0/env.h` | `Invoke-WebRequest https://github.com/IngenieroRicardo/env/releases/download/2.0/env.h -OutFile ./env.h` |

---

### 🛠️ Compilar

| Linux | Windows |
| --- | --- |
| `gcc -o main.bin main.c ./env.so` | `gcc -o main.exe main.c ./env.dll` |
| `x86_64-w64-mingw32-gcc -o main.exe main.c ./env.dll` |  |

---

### 🧪 Ejemplo

```C
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "env.h"

int main() {
    // Cargar el archivo .env
    int loadResult = LoadEnvFile();
    if (loadResult != 0) {
        printf("Error: No se pudo cargar el archivo .env\n");
        return 1;
    }
    
    // Obtener una variable de entorno específica
    char* varName = "basededatos"; // Nombre de la variable a leer
    char* value = GetEnvVar(varName);
    
    if (value != NULL && strlen(value) > 0) {
        printf("El valor de '%s' es: %s\n", varName, value);
    } else {
        printf("La variable '%s' no está definida\n", varName);
    }
    
    // Liberar memoria asignada por GetEnvVar
    free(value);
    
    return 0;
}
```

```.env
variable=String de conexion
basededatos=mysql

```

---


## 📚 Documentación de la API

#### Manejo Básico de env
- `int LoadEnvFile()`: Cargar archivo .env
- `char* GetEnvVar(char* key)`: Obtener dato a partir del key
