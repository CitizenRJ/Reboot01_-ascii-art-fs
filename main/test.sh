#!/bin/bash
echo go run . "hello" standard 
echo go run . "hello world" shadow 
echo go run . "nice 2 meet you" thinkertoy 
echo go run . "you & me" standard 
echo go run . "123" shadow 
echo go run . "/(\")" thinkertoy 
echo go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow 
echo go run . "\"#$%&/()*+,-./" thinkertoy 
echo go run . "It's Working" thinkertoy 
