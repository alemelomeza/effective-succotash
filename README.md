# effective-succotash
### Principios fundamentales de Data-Oriented Programming

Los principios fundamentales de la Programación Orientada a Datos (DOP) son cuatro:

*   **Separar el código de los datos**: El código (comportamiento) debe residir en funciones cuyo comportamiento no dependa de los datos encapsulados en el contexto de la función. Este principio sugiere una separación clara entre las entidades de código y las entidades de datos.
    *   En los lenguajes de POO, el código se agrega en métodos estáticos y los datos en clases sin métodos.
    *   En los lenguajes de FP, se debe evitar ocultar datos en el ámbito léxico de las funciones.
    *   Este principio no tiene que ver con la forma en que se representan los datos.
    *   Los beneficios incluyen que el código se puede reutilizar en diferentes contextos, se puede probar de forma aislada y los sistemas tienden a ser menos complejos.
    *   El coste incluye la falta de control sobre qué código accede a qué datos, la ausencia de empaquetado y la creación de más entidades.

*   **Representar los datos con estructuras de datos genéricas**:  Los datos deben representarse con estructuras de datos genéricas como mapas (diccionarios) y arrays (o listas) en lugar de instanciar datos a través de clases específicas.
    *   Este principio no trata la mutabilidad o inmutabilidad de los datos.
    *   Los beneficios incluyen el uso de funciones genéricas que no se limitan a casos de uso específicos y un modelo de datos flexible.
     *  El coste incluye un ligero impacto en el rendimiento, la necesidad de documentar manualmente la forma de los datos y, en algunos lenguajes con tipado estático, la necesidad de un type casting explícito.

*   **Tratar los datos como inmutables**:  Los cambios en los datos se logran creando nuevas versiones de los datos. La referencia a una variable puede cambiar para que se refiera a una nueva versión de los datos, pero el valor de los datos en sí nunca debe cambiar.
    *   Este principio da como resultado un código predecible incluso en un entorno multiproceso, y las comprobaciones de igualdad son rápidas.
    *   Sin embargo, se requiere un cambio mental y, en la mayoría de los lenguajes de programación, se necesita una biblioteca de terceros para proporcionar una implementación eficiente de estructuras de datos persistentes.

*   **Separar el esquema de datos de la representación de datos**: La forma esperada de los datos se expresa como un esquema de datos que se mantiene separado de los datos en sí.
    *   Esto permite a los desarrolladores decidir qué partes de los datos deben tener un esquema y cuáles no.
    *   Los beneficios incluyen la libertad de elegir qué datos deben ser validados, campos opcionales, condiciones avanzadas de validación de datos y la generación automática de visualización del modelo de datos.
    *   El coste incluye una conexión débil entre los datos y su esquema y un pequeño impacto en el rendimiento.

Estos cuatro principios, cuando se combinan, forman un todo cohesivo. Los sistemas construidos utilizando DOP son más simples y fáciles de entender, lo que mejora significativamente la experiencia del desarrollador.

## Estructura de un proyecto con Arquitectura Hexagonal y DOP en Go

La Arquitectura Hexagonal, también conocida como Puertos y Adaptadores, es un patrón de diseño que busca desacoplar el núcleo de la aplicación (lógica de negocio) de las dependencias externas (bases de datos, APIs, interfaces de usuario). Al combinarla con los principios de la Programación Orientada a Datos (DOP), se obtiene una estructura de proyecto flexible, escalable y fácilmente testable.

A continuación, se propone una estructura de proyecto que integra la Arquitectura Hexagonal y DOP en Go:

```
proyecto-go/
├── internal/
│   ├── core/         // Núcleo de la aplicación (lógica de negocio)
│   │   ├── domain/    // Entidades y casos de uso
│   │   │   ├── entities/  // Estructuras de datos (DOP)
│   │   │   └── usecases/ // Lógica de negocio (funciones puras)
│   │   └── ports/     // Interfaces (puertos) para interactuar con el exterior
│   └── infrastructure/ // Implementaciones de los puertos (adaptadores)
│   │   ├── database/   // Adaptadores para bases de datos
│   │   ├── api/        // Adaptadores para APIs
│   │   └── external/   // Adaptadores para servicios externos
├── cmd/
│   └── main.go       // Punto de entrada de la aplicación
├── pkg/              // Librerías reutilizables
├── tests/            // Pruebas unitarias y de integración
└── go.mod            // Archivo de dependencias
```

### Capas de la Arquitectura Hexagonal

1. **Core (Núcleo):**
   - Contiene la lógica de negocio y las entidades del dominio.
   - Define interfaces (puertos) para interactuar con el exterior.
   - No depende de ninguna otra capa.

2. **Infrastructure (Infraestructura):**
   - Contiene las implementaciones concretas de los puertos definidos en el núcleo (adaptadores).
   - Incluye código para bases de datos, APIs, servicios externos, etc.
   - Depende del núcleo.

### Principios de DOP en cada capa

- **Core:** Se aplican los principios de separación de código y datos, representación de datos con estructuras genéricas y, preferiblemente, inmutabilidad de los datos.
- **Infrastructure:** Se aplican los principios de separación de código y datos y separación del esquema de datos de la representación.

### Beneficios de esta estructura

- **Desacoplamiento**: El núcleo de la aplicación es independiente de las dependencias externas.
- **Testabilidad**: Se pueden probar los casos de uso sin necesidad de mocks para las dependencias externas.
- **Flexibilidad**: Se pueden cambiar las implementaciones de los adaptadores sin afectar el núcleo.
- **Mantenibilidad**: La separación de responsabilidades facilita la comprensión y modificación del código.

Esta estructura de proyecto proporciona una base sólida para construir aplicaciones Go robustas y escalables utilizando la Arquitectura Hexagonal y DOP. Recuerda que esta es solo una propuesta y se puede adaptar a las necesidades específicas de cada proyecto.
