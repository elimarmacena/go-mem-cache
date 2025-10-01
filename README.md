# go-mem-cache
Personal implementation of in memory cache using golang.

## Introduction
This project aims to create a TTL cache implementation avoiding external lib for the development.

## Operations
The project present the following operations related to the cache database.
* Read
  * Get
* Write
  * Set
  * Delete
  * Clear

Aiming the reponse time, only the write operations presents a thread control.