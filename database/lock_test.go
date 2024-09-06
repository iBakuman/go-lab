package database

// When using `SELECT ... FOR UPDATE` in a transaction, the purpose is to lock the rows that are being selected so that
// no other transactions can modify or lock them until the current transaction is complete. This locking mechanism is
// designed to prevent other transactions from reading or writing to the locked rows, ensuring data consistency and
// preventing race conditions.
//
// ### Why Thread B Cannot Execute the Same `SELECT ... FOR UPDATE`
//
// Here’s a detailed explanation of why thread B cannot execute the same `SELECT ... FOR UPDATE` while thread A is
// already executing it:
//
// 1. **Row Locking Mechanism**: When thread A executes `SELECT * FROM test LIMIT 1 FOR UPDATE;` within a transaction,
// it places a lock on the rows being selected. The lock type is typically an exclusive lock (depending on the
// database's locking strategy). This lock means that thread A has exclusive access to modify the selected rows.
//
// 2. **Lock Contention**: If thread B tries to execute the same `SELECT * FROM test LIMIT 1 FOR UPDATE;`, it attempts
// to acquire an exclusive lock on the same rows. However, because thread A already holds the lock, thread B cannot
// acquire the same lock. As a result, thread B will be blocked, waiting for thread A to release the lock.
//
// 3. **Transaction Isolation**: The isolation level of the transaction plays a role here. By default, `FOR UPDATE`
// applies row-level locking that ensures any changes made by one transaction are isolated from other transactions.
// Until thread A completes its transaction (either commits or rolls back), other transactions (like thread B’s) are
// prevented from acquiring locks on the same rows.
//
// 4. **Deadlock Prevention and Consistency**: This behavior prevents deadlocks and ensures consistency. If thread B
// were allowed to lock the same rows simultaneously, it could lead to inconsistent reads, updates, or even deadlocks
// where neither thread can proceed.
//
// ### Example Scenario
//
// Assume you have a table `test` with multiple rows. Here’s what happens:
//
// - **Thread A**: Starts a transaction and executes `SELECT * FROM test LIMIT 1 FOR UPDATE;`. This locks the first
// row(s) of the `test` table based on the database's row selection order.
//
// - **Thread B**: Starts a transaction and attempts to execute the same query. However, because thread A holds a lock
// on the selected rows, thread B’s query is blocked. It cannot proceed until thread A completes its transaction.
//
// ### Behavior Based on Database
//
// Different databases handle `LIMIT` with `FOR UPDATE` differently, but in general:
//
// - **MySQL**: The `LIMIT` clause with `FOR UPDATE` will lock the first `n` rows that match the criteria. If another
// transaction tries to `SELECT ... FOR UPDATE` the same rows, it will be blocked until the first transaction releases
// the locks.
//
// - **PostgreSQL**: Similar to MySQL, using `LIMIT` with `FOR UPDATE` locks the specified number of rows. Other
// transactions attempting to `SELECT ... FOR UPDATE` the same rows will wait for the lock to be released.
//
// ### Handling Concurrent Access
//
// To manage concurrent access and avoid blocking:
//
// 1. **Use SKIP LOCKED**: Some databases, like PostgreSQL, support `FOR UPDATE SKIP LOCKED`, which allows other
// transactions to skip over rows that are currently locked by another transaction. This way, thread B will skip the
// locked rows and try to lock other rows:
//
// ```sql SELECT * FROM test LIMIT 1 FOR UPDATE SKIP LOCKED; ```
//
// 2. **Reduce Lock Scope**: Design your queries and application logic to minimize the scope and duration of locks. This
// can be achieved by processing smaller chunks of data at a time.
//
// 3. **Use Queues or Job Systems**: As mentioned earlier, using a job queue where each job processes a single user can
// ensure no two threads/processes work on the same data concurrently.
//
// ### Summary
//
// - `SELECT ... FOR UPDATE` locks rows to prevent other transactions from accessing or modifying them concurrently,
// ensuring data consistency and preventing race conditions. - Thread B cannot execute the same `SELECT ... FOR UPDATE`
// because it tries to acquire the same locks already held by thread A. - Using mechanisms like `SKIP LOCKED` or
// carefully managing transaction scopes can help mitigate blocking issues in high-concurrency environments.
