package com.advanced.development.log4j2.services;

import java.io.File;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.core.LoggerContext;

public class Logger2 {
	private org.apache.logging.log4j.Logger logger;

	private Logger2(Class<? extends Object> clazz) {
		this.logger = LogManager.getLogger(clazz);
	}

	private Logger2(String clazz) {
		this.logger = LogManager.getLogger(clazz);
	}

	public static Logger2 getLogger(Class<? extends Object> clazz) {
		return new Logger2(clazz);
	}

	public static Logger2 getLogger(String clazz) {
		return new Logger2(clazz);
	}

	public void info(Object message) {
		this.logger.info(message);
	}

	public void debug(Object message) {
		this.logger.debug(message);
	}

	public void debug(Object message, Throwable throwable) {
		this.logger.debug(message, throwable);
	}

	public void error(Object message) {
		this.logger.error(message);
	}

	public void error(Object message, Throwable throwable) {
		this.logger.error(message, throwable);
	}

	public void warn(Object message) {
		this.logger.warn(message);
	}

	public void warn(Object message, Throwable throwable) {
		this.logger.warn(message, throwable);
	}

	public static void setConfigFile(String logConfigFile) {
		File file = new File(logConfigFile);
		LoggerContext context = (org.apache.logging.log4j.core.LoggerContext) LogManager.getContext(false);
		context.setConfigLocation(file.toURI());
	}
}
