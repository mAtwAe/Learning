package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@SpringBootApplication
@RestController
public class DemoApplication {

	private final UserService userService;

	@Autowired
	public DemoApplication(UserService userService){
		this.userService = userService;
	}

	@RequestMapping("/api/v1/home")
	public String home() {
		return "Hello Docker World";
	}

	@GetMapping
	public List<User> getUser(){
		return userService.getUsers();
	}

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}
}
