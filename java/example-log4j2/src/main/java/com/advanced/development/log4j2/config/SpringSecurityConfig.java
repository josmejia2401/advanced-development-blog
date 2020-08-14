package com.advanced.development.log4j2.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;
import org.springframework.context.annotation.PropertySource;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;


@Configuration
@PropertySource({"classpath:application.properties"})
@Profile("default")
public class SpringSecurityConfig implements WebMvcConfigurer {


}