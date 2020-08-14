package com.advanced.development.log4j2.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;
import org.springframework.context.annotation.PropertySource;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
@PropertySource({ "classpath:application-dev.properties" })
@Profile("dev")
public class PropertiesSourceDev implements WebMvcConfigurer {

	@Bean
	public RestTemplate getRestTemplate() {
		return new RestTemplate();
	}
}