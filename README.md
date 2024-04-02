# Token-based pagination

Token-based pagination offers several advantages over traditional offset-based pagination:

1. Consistency in Results: Token-based pagination ensures consistent results between paginated requests, even if data changes occur between requests. This consistency is crucial, especially in dynamic applications where data can be added, modified, or deleted frequently.


2. Improved Performance: With token-based pagination, the database doesn't need to calculate offsets for each page, which can be resource-intensive, especially with large datasets. Instead, it relies on predefined cursors or tokens, resulting in faster and more predictable performance.


3. Scalability: Token-based pagination scales better with large datasets because it avoids the performance degradation associated with offset-based pagination when fetching pages deep into the dataset. Since tokens are typically based on unique identifiers, the performance remains consistent regardless of the dataset's size.


4. Reduced Server Load: Offset-based pagination can put a significant load on the server, especially when dealing with large datasets, as it requires the server to compute offsets for each page. Token-based pagination offloads this computation to the client side, reducing the server load and improving scalability.


5. Flexibility: Tokens can carry additional metadata or contextual information, allowing for more flexible pagination strategies. For example, tokens can include sorting criteria, filters, or other parameters, enabling more precise and customizable data retrieval.


6. Bookmarking: Tokens serve as bookmarks to specific points in the dataset, allowing users to bookmark and revisit specific pages easily. This feature enhances user experience by providing a seamless way to navigate through paginated data.


7. Support for Real-time Updates: Token-based pagination can seamlessly support real-time updates and live data feeds. Since tokens are based on unique identifiers, new data can be inserted into the dataset without affecting the pagination logic, ensuring a smooth user experience.


Overall, token-based pagination offers better performance, scalability, and flexibility compared to offset-based pagination, making it a preferred choice for paginating large datasets in modern web applications.

## Initialize DB - this will create `data.db` in project's root directory
```shell
go run cmd/init/init_db.go
```

## Run HTTP server
```shell
go run cmd/http/server.go
```
