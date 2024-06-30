# Instruction

## Setup infrastructure

Prepare the environment for running tests
```bash
make prepare_test_environment
```

## Run tests

Run unit tests (only possible inside a container)
```bash
make unit_test
```

Run e2e tests outside the container
```bash
make e2e_test
```
