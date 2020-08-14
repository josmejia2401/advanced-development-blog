package com.advanced.development.log4j2;


import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.advanced.development.log4j2.services.Logger2;

@SpringBootApplication
//@ComponentScan(basePackages = { "com.advanced.development.*" })  
@EnableAutoConfiguration
public class StartApplication implements CommandLineRunner {

    private static final Logger2 LOGGER = Logger2.getLogger(StartApplication.class);

    
    public static void main(String[] args) {
        SpringApplication.run(StartApplication.class, args);
    }

    @Override
    public void run(String... args) {
    	LOGGER.info("StartApplication...");
    }

}