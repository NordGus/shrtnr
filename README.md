# Shrtnr

A URL Shortener service to practice Message Driven Software Architecture and distributed systems in Go with the ROM Stack

## Design constrains

The service should only hold a maximum of 2500 short urls. When a new url is added and the service can't store more, the oldest is drop to make space.

> Let's implement some queues

---
Built with the ROM Stack
