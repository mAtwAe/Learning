package com.example.demo;

import org.springframework.stereotype.Service;

import java.util.Objects;
import java.util.function.Function;

@Service
public class UserRequestMapper implements Mapper<User, UserRequest> {

    @Override
    public UserRequest apply(User user){
        return new UserRequest(
                user.getName()
        );
    }

    @Override
    public User reverse(UserRequest userRequest){
        return new User(
                userRequest.name(),
                null
        );
    }
}
