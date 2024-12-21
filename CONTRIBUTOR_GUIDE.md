# **Contributor Guide**

Welcome to the MightDS project! This guide provides the standards and best practices for contributing to ensure consistency and maintainability across the codebase. Please read through this document before making changes or submitting pull requests.

## **Table of Contents**

1. [Code Standards](#code-standards)  
2. [Naming Conventions](#naming-conventions)  
3. [Logging Guidelines](#logging-guidelines)  
4. [Git Commit Message Guidelines](#git-commit-message-guidelines)  
5. [Pull Request (PR) Process](#pull-request-pr-process)  
6. [Documentation and Comments](#documentation-and-comments)

## **Code Standards**

We follow industry-standard Go coding practices to ensure consistency and readability. The following guidelines should be adhered to:

1. **Indentation:**
   - Use **tabs** for indentation.  
   - One tab per indentation level.
   - Nested logic must be consistently indented.

   **Example:**
   ```go
   func ExampleFunction() {
       if condition {
           DoSomething() // Indented with a single tab
       }
   }
   ```

2. **Line Length:** Limit lines to 80 characters.
3. **Error Handling (Go Built-in Error Style):**
   - Always return errors using Go’s `error` type.  
   - Format error messages clearly, providing enough context for debugging.

   **Example:**
   ```go
   return fmt.Errorf("failed to push item %v onto stack: %w", item, err)
   ```


## **Naming Conventions**

We aim to keep naming clear and descriptive to improve maintainability. 

1. **Packages and Files:**
   - Use **lowercase** names for packages (e.g., `stack`, `logger`, `error`).
   - For files with multiple words, use **underscores** (`_`) instead of hyphens (`-`).
   
   **Good:**
   - `stack.go`, `bounded_stack.go`
   - `logger.go`, `stack_logger.go`

   **Bad:**
   - `bounded-stack.go`, `stack-logger.go`
  
2. **Variables and Functions:**
   - Use **camelCase** for private variables and methods (e.g., `maxSize`, `isEmpty`).
   - Use **TitleCase** for exported variables and methods (e.g., `Push()`, `Pop()`).

## **Logging Guidelines**

1. **Log Levels:**
   - **DEBUG:** For detailed debugging information.  
   - **INFO:** For general system updates or milestones.  
   - **WARN:** For potential issues that might need attention.  
   - **ERROR:** For recoverable issues affecting functionality.  
   - **CRITICAL:** For severe issues requiring immediate attention.

2. **Best Practices:**
   - Always provide meaningful log messages.
   - Avoid logging sensitive data like credentials or PII.

## **Git Commit Message Guidelines**

We follow a simplified Git commit message format to ensure clarity in the project history. 

#### **Commit Message Format:**

```text
<type>: <short description>
```

- **Type:** Use one of the following types:
  - `feat`: Adding a new feature.
  - `fix`: Fixing a bug.
  - `docs`: Updates or additions to documentation.
  - `style`: Code style changes (e.g., formatting, minor refactoring).
  - `refactor`: Refactoring code without adding features or fixing bugs.
  - `test`: Adding or updating tests.
  - `chore`: Miscellaneous changes (e.g., updating dependencies).

- **Description:** A concise and clear description of the change. Keep 
it under 72 characters.

- **Verb Usage:**  
   Use present tense and active voice:
   - **Correct:** `Add stack size validation`
   - **Incorrect:** `Added stack size validation`

- **Examples:**

1. **Adding a new feature:**
   ```text
   feat: Add support for bounded stack

   Implemented the BoundedStack data structure, which limits the number of elements stored.
   ```

2. **Fixing a bug:**
   ```text
   fix: Correct stack overflow handling in Push()

   Fixed the bug where the Push method was allowing elements to be added to a full stack.
   ```

3. **Documentation update:**
   ```text
   docs: Update README with new stack usage examples

   Added examples of using the BoundedStack and Stack data structures with sample code.
   ```

4. **Refactoring:**
   ```text
   refactor: Simplify stack initialization logic

   Refactored the stack initialization process to reduce redundant checks and improve clarity.
   ```

#### **Git Commit Best Practices:**
- **Commit frequency:** Commit frequently but with meaningful changes.
- **Commit only related changes:** Avoid committing unrelated changes together in a single commit. Each commit should reflect a single logical change.
  

## **Pull Request (PR) Process**

When creating a Pull Request (PR), please follow these steps:

1. **Fork the repository** and **create a new branch** for your work.
2. **Make changes** in your branch. Make sure to only include files relevant to the changes you're making.
3. **Commit your changes** with a meaningful message (see [Commit Message Guidelines](#git-commit-message-guidelines)).
4. **Push your branch** to your fork.
5. **Create a PR** from your branch to the main repository.
6. **PR Description:** In the PR description, include the following:
   - A brief summary of what the PR does.
   - Any relevant context or information about why the changes were made.
   - Any issues or bugs being fixed, if applicable.

#### **Example PR Message:**

```text
Add BoundedStack support

This PR adds the BoundedStack data structure, which allows users to define a maximum size for the stack. It includes methods for pushing, popping, and checking if the stack is full.
```

7. **Review:** A reviewer will assess the changes. Respond to any feedback and address necessary changes.
8. **Merge:** Once approved, your PR will be merged into the main branch.

## **Documentation and Comments**

Proper documentation is crucial for maintainability. Developers often add comments to explain what the code does, but **comments should focus on why** the code is there, rather than just describing it.

1. **Function Comments:** Every exported function should have a comment explaining its purpose, the inputs, and expected outputs.

   **Good:**
   ```go
   // Push adds an element to the stack.
   // Returns an error if the stack is full.
   func Push(item Item) error {
       if isFull() {
           return fmt.Errorf("stack overflow")
       }
       stack = append(stack, item)
       return nil
   }
   ```

   **Bad:**
   ```go
   // Push adds an element to the stack
   func Push(item Item) {
       stack = append(stack, item)
   }
   ```

2. **Struct Comments:** For structs, explain what the struct represents and any relevant details for its usage.

   **Good:**
   ```go
   // BoundedStack represents a stack with a defined maximum size.
   // It provides methods to push, pop, and check if the stack is full.
   type BoundedStack struct {
       maxSize int
       stack   []Item
   }
   ```

   **Bad:**
   ```go
   // BoundedStack is a stack.
   type BoundedStack struct {
       maxSize int
       stack   []Item
   }
   ```

3. **Inline Comments:** Use inline comments sparingly. They should clarify complex code but not explain simple code.

   **Good:**
   ```go
   stack = append(stack, item) // Append item to the stack
   ```

   **Bad:**
   ```go
   stack = append(stack, item) // This line appends an item to the stack
   ```

By following these guidelines, we ensure the code is maintainable, readable, and easy for new contributors to understand and extend.