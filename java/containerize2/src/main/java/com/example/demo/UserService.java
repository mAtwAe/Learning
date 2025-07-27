package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class UserService {

    private final UserRequestMapper userRequestMapper;
    private final UserDTOMapper userDTOMapper;
    private final UserRepository userRepository;

    public UserService(UserRequestMapper userRequestMapper, UserDTOMapper userDTOMapper, UserRepository userRepository) {
        this.userRequestMapper = userRequestMapper;
        this.userDTOMapper = userDTOMapper;
        this.userRepository = userRepository;
    }


    public List<UserDto> getUsers(){
        return userRepository.findAll()
                .stream()
                .map(userDTOMapper)
                .collect(Collectors.toList());
    }

    public void addUsers(UserRequest userRequest){
        if(!userRequest.name().isBlank()){
            userRepository.save(userRequestMapper.reverse(userRequest));
        }
    }

    public User getUserServiceById(Integer id) {
        return userRepository.findById(id)
                .orElseThrow(() -> new IllegalStateException(
                        id + " not found"
                ));
    }
}
