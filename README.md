<h2 align="center">
<img src=".github/images/sqlmustcontext.png" alt="sqlmustcontext logo" width="500">
</h2>
<h2 align="center">
  a <code>go/analysis</code> linter to enforce that the standard library `database/sql` package's context-enabled functions are always called versus their non-context-enabled counterparts
</h2>
<div align="center">

&nbsp;&nbsp;&nbsp;[golangci-lint][golang-ci-lint-link]&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[contributing new linters][contribute-new-linters-link]&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;[linter building tutorial][tutorial-link]

[![Made With Go][made-with-go-badge]][for-the-badge-link]

</div>

---
## Overview
`sqlmustcontext` is a Go linter built using the `go/analysis` framework. Its purpose is to enforce the use of "context-enabled" functions from the standard library's `database/sql` package instead of their non-context counterparts. This ensures that database operations are properly tied to a `context.Context`, promoting better resource management and cancellation handling in Go applications.

For example, the linter will flag the use of functions like `db.Ping()` or `db.Exec()` and suggest using their context-enabled equivalents, such as `db.PingContext(ctx)` or `db.ExecContext(ctx, query)`. Similarly, it enforces the use of context-enabled methods for transactions, such as `tx.ExecContext(ctx, query)` instead of `tx.Exec(query)`.

### Benefits of Using Context-Enabled Functions

Using the context-enabled functions from the `database/sql` package provides several key benefits:

1. **Tracing and Monitoring**  
   Context-enabled functions allow you to pass a `context.Context` object, which can carry metadata such as request IDs or tracing information. This is essential for distributed systems where tracing requests across services is critical for debugging and performance monitoring.

2. **Timeouts and Deadlines**  
   By using a `context.Context`, you can set timeouts or deadlines for database operations. This ensures that long-running queries or operations do not block your application indefinitely, improving overall responsiveness and reliability.

3. **Cancellation Propagation**  
   Context-enabled functions support cancellation, allowing you to cancel database operations when the parent context is canceled. This is particularly useful in scenarios where a user aborts a request or when a service shuts down gracefully.

4. **Resource Management**  
   Proper use of `context.Context` helps in managing resources effectively by ensuring that operations are tied to the lifecycle of the context. This reduces the risk of resource leaks, such as open connections or unclosed transactions.

By enforcing the use of context-enabled functions, `sqlmustcontext` helps you write more robust, maintainable, and production-ready Go applications.

## Development

### Prerequisites

**Go**  
  To install a valid Go toolchain, follow these steps:
1. Download and install Go from the [official Go website](https://go.dev/dl/).
2. Ensure that the Go binary is added to your system's PATH. You can verify this by running `go version` in your terminal.
3. Optionally, install tools like `golangci-lint` for linting by running:
   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

  For more detailed instructions, refer to the [Go installation guide](https://go.dev/doc/install).

### Testing
This repository contains a go linter written in the go/analysis style. It has a test located at [pkg/analyzer/analyzer_test.go](pkg/analyzer/analyzer_test.go), that currently provides 95%+ test coverage.

Should you wish to try this linter on repositories locally, run `make install` followed by `sqlmustcontext <path to lint>`.


<!--

Reference Variables

-->

<!-- Badges -->
[made-with-go-badge]: .github/images/made-with-go.svg

<!-- Links -->
[blank-reference-link]: #
[for-the-badge-link]: https://forthebadge.com
[contribute-new-linters-link]: https://golangci-lint.run/contributing/new-linters/
[golang-ci-lint-link]: https://golangci-lint.run/
[tutorial-link]: https://disaev.me/p/writing-useful-go-analysis-linter/