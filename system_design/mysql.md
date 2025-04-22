MySQL’s architecture is modular in that the storage engine is responsible for the actual data storage and retrieval. There are multiple storage engines, but the most prominent ones are:
* InnoDB (transactional engine, ACID-compliant).
* MyISAM (non-transactional engine, faster in certain read-heavy use cases).


## a. InnoDB Internals
InnoDB is the default storage engine in MySQL, and it’s optimized for performance and reliability. Here’s how it works internally:
* Buffer Pool: InnoDB uses a memory cache called the buffer pool to store frequently accessed data. This minimizes disk I/O, as reading from RAM is much faster than from disk.
* B-tree Indexes: InnoDB uses B-tree indexes to speed up the retrieval of rows. These indexes are self-balancing and allow for efficient searching, inserting, and deleting.
* Transaction Management: InnoDB is known for its support of transactions, which ensures ACID (Atomicity, Consistency, Isolation, Durability) properties. It does this by using undo logs, redo logs, and lock mechanisms.
* Undo logs are used to revert data back to a consistent state if a transaction fails.
* Redo logs are used for crash recovery. After a crash, InnoDB uses these logs to restore the database to the state it was in before the crash.
* Concurrency Control: InnoDB employs a technique called Multi-Version Concurrency Control (MVCC) to allow multiple transactions to happen concurrently without locking the entire table.
