

# Key

Redis does not immediately delete a key at the exact moment it expires. Instead, it uses two strategies:

- **Lazy Deletion:** When you try to access a key, Redis checks if it has expired. If it has, the key is deleted at that time.
- **Periodic Deletion:** Redis also runs a background process that periodically scans a subset of keys and removes those that have expired, even if they haven't been accessed.