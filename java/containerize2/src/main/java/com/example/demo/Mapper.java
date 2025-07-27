package com.example.demo;

import java.util.function.Function;

public interface Mapper<S, T> extends Function<S, T> {
    S reverse(T target);
}
