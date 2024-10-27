## Work in progress

v0.0.2 (27.10.2024)
* added graceful lib for graceful shutdown
* http-adapter is ready
* main is refactored. Initialization and runtime parts are separated
* configs now are parts of their packages
* consistent package names (all snake_case)
* general refactoring. Refactored names of: packages, variables, functions etc.



v0.0.1
* initial version

#### TODO
1. Get rid of config in domain (+)
2. Rethink the structure of http-adapter (+)
3. Implement graceful shutdown (+)
4. Consider making a branch with DI
5. Separate interface adapters layer and infrastructure layer
6. Think about fatals in adapters constructors (+)
7. Validate how context is propagated in adapters


Notes:
1. snake_case in package names is opinionated
2. adapters implementations are not production ready. They are just examples (except http-adapter)
