# FizzBuzz Generator

A concurrent Go program that generates numbers, maps them to FizzBuzz values, and prints the results in numeric order.

## Run

```sh
go run .
```

The program processes the numbers from 1 through 10,000,000 using 100 workers. Each number is printed as `number -> value`:

- Multiples of both 3 and 5 become `FizzBuzz`.
- Multiples of 3 become `Fizz`.
- Multiples of 5 become `Buzz`.
- All other numbers remain their decimal representation.

## Workflow

```mermaid
flowchart TD
    A[Start] --> B[Gen emits 1..10,000,000]
    B --> C[Shared input channel]
    C --> D[100 concurrent workers]
    D --> E{Classify each number}
    E -->|Divisible by 3 and 5| F[FizzBuzz]
    E -->|Divisible by 3| G[Fizz]
    E -->|Divisible by 5| H[Buzz]
    E -->|Otherwise| I[Number as text]
    F --> J[Response channel]
    G --> J
    H --> J
    I --> J
    J --> K[SortedBuffer restores numeric order]
    K --> L[Print number -> value]
    D -->|All workers finish| M[Close response channel]
    M --> K
    L --> N[End]
```
