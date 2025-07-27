package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/api/v2/user")
public class DemoApplication {

	private final UserService userService;

	public DemoApplication(UserService userService){
		this.userService = userService;
	}

	public String home() {
		return "Hello Docker World";
	}

	@GetMapping
	public List<UserDto> getUser(Pageable pages){
		return userService.getUsers();
	}

	@PostMapping
	public ResponseEntity addNewUser(@Validated @RequestBody UserRequest user){
		userService.addUsers(user);
		return ResponseEntity.status(HttpStatus.CREATED).build();
	}

	@GetMapping("{id}")
	public User getUserById(
			@PathVariable Integer id
	){
		return userService.getUserServiceById(id);
	}

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}
}
