package com.example.gvm;

public class CircleCalculator {
    public static float circumference(float r) {
        float pi = 3.14f;
        float circumference = 2 * pi * r;
        return circumference;
    }

    public static void main(String[] args) {
        System.out.println(circumference(1f));
    }

}
