# Experiencia en el proyecto de concurrencia

## Tabla de contenidos
- Introducción
- Tecnologías utilizadas
- Concurrencia
- Estructura del proyecto
- Desafíos y soluciones
- Conclusiones

## Introducción

El desarrollo de este proyecto que consiste en un simulador de un estacionamiento con semáforos fue una experiencia desafiante e interesante que me permitió combinar el uso de estructuras de datos como los canales para aplicar la concurrencia en Go y el diseño de una interfaz gráfica para representar visualmente el estado de un estacionamiento en tiempo real. En este artículo, quiero compartir la experiencia, los desafíos técnicos que enfrenté, las soluciones implementadas y las lecciones que aprendí durante el desarrollo. El objetivo principal de este proyecto era construir un sistema que pudiera simular la entrada y salida de vehículos en un estacionamiento, así como la disponibilidad de los espacios para los vehículos reflejando en tiempo real el estado de ocupación de cada espacio.

## Tecnologías

- Go

El lenguaje principal del proyecto. La capacidad de Go para manejar concurrencia con goroutines fue crucial para simular múltiples eventos de entrada y salida de vehículos, así como la disponibilidad de los espacios del estacionamiento.

- Fyne

Se utilizó este framework de Go para interfaces gráficas, con lo que me permitió construir una interfaz visualmente intuitiva y reactiva para representar el estacionamiento. Se utilizó un semáforo sync.Mutex para manejar la exclusión mutua entre los vehiculos al momento de entrar y salir del estacionamiento.

# Concurrencia

El uso de goroutines para manejar el acceso de los vehículos facilitó la gestión de eventos en tiempo real. Sin embargo, asegurar que las actualizaciones de la interfaz fueran consistentes requirió un manejo cuidadoso de concurrencia.

# Estructura del proyecto

- models:

Contiene las entidades principales del sistema, como Parking y Vehicle. Este paquete
se encarga de definir la lógica básica y el estado de los componentes clave.

- scenes:

Se encarga de la lógica de simulación, como el flujo de vehículos entrando y saliendo
del estacionamiento. Este módulo también maneja la sincronización y el control del estado de
los espacios de estacionamiento.

- views:

Administra la interfaz gráfica de usuario (GUI) usando el framework fyne. Incluye
funciones para inicializar la ventana principal, mostrar el estado del estacionamiento, y
actualizar la visualización cada vez que un vehículo entra o sale.

# Desafíos y soluciones

Uno de los aspectos más desafiantes durante el desarrollo del proyecto
fue la instalación e implementación del paquete fyne, debido al poco conocimiento sobre este framwork y de como saber optimizar la interfaz gráfica. Inicialmente, la interfaz intentaba actualizar todos los espacios en cada evento, lo cual afectaba el rendimiento. Para solucionar esto, se trató  de optimizar el código para que solo los espacios que cambian de estado se actualicen, mejorando significativamente la eficiencia.

# Conclusión

El desarrollo de este simulador de estacionamiento fue una experiencia enriquecedora que me permitió experimentar con concurrencia en Go y explorar el diseño de interfaces gráficas con Fyne. Logré crear una simulación funcional que no solo refleja el estado en tiempo real, sino que también permite visualizar el funcionamiento de los semáforos en la ocupación de espacios de estacionamiento, así como en la entrada y salida de estos. Este proyecto me ha motivado a seguir explorando el potencial de Go y a profundizar en la creación de aplicaciones interactivas con interfaces gráficas.