# Ejemplo Log4j 2 en JAVA | Log4j 2 en Springboot | Configuración Log4j 2 | Log4j 2 in SpringBoot| Example Log4j 2 in SpringBoot | Configuring Log4j 2

Apache Log4j 2 o 2.x es una actualización muy importante de Log4j, la cual proporciona innumerables mejoras con respecto a las versiones anteriores, entre estas mejoras: Logback, manejo asíncrono, etc. 

# Nuevas caracteristicas
  - Ninguna
# Construcción
Construido en JAVA a través del IDE Eclipse.
# Dependencias
  
  <groupId>org.springframework.boot</groupId>
	<artifactId>spring-boot-starter-parent</artifactId>
	<version>2.2.4.RELEASE</version>

  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
    <exclusions>
      <exclusion>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-logging</artifactId>
      </exclusion>
    </exclusions>
  </dependency>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-log4j2</artifactId>
  </dependency>
  <dependency>
    <groupId>com.lmax</groupId>
    <artifactId>disruptor</artifactId>
    <version>3.4.2</version>
  </dependency>

# Ejecución y Compilación
```
mvn clean install
```
# Métodos permitidos

# DETALLE TECNICO

Tecnologías usadas:

* [JAVA] - JAVA 11
* [MAVEN] - Maven 

# PARA MÁS DETALLE
- Visitar: [a DETALLE](https://soursop-dev.blogspot.com/2020/08/golang-con-dockergolang-with-docker.html)