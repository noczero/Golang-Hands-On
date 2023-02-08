MySQL Database Implementation
===

This module implements hexagonal (port-adapter) architecture based on repository layered design pattern.

There 2 layered for database :
1. Entity layer represent database schema
2. Repository layer represent database operation including insert, find by id, and find all.

In repository layer there is interface and implementation of its interface.

Then it tested and invoke the implementation in _test.go

Also, it provide mysql container that can be running with ```docker-compose up -d``` and all tables will create automatically.
