The test assignment consists of two independent parts. In the first part you need to implement a
CLI tool in go language. In the second part you need to design the architecture of the system
that will run in Kubernetes. On each task you can spend no more than 2 hours, so 4 hours in
total.

CLI tool
You need to implement a CLI tool that takes a list of URLs as input, visits each URL, and prints
the list of pairs to the output: URL and response body size. The output must be sorted by the
size of the response body. It should be implemented in the most efficient way.
Time limit: 2 hours
Programming Language: go (only)

k8s app architecture
You need to design the architecture of a simple system that will run in Kubernetes. You can
share the result in any desired format (text, UML diagrams, k8s manifests, hand drawings, etc.).
The system has the following components:
● SPA frontend;
● API backend;
● Postgres cluster;
● S3 bucket (to store some files within the application);
● 3rd party service (integration with external systems via API, like DocuSign for signing
documents);
● SQL script for initializing the database;
● some script that fills the data for ephemeral environments (like for testing purpose).
Time limit: 2 hours