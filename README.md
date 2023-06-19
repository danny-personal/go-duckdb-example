# go-duckdb-example
Playing with DuckDB.

## Installation.
```bash
pip install duckdb==0.8.1
vscode ➜ /usr/bin $ export PATH="/usr/bin:$PATH"
vscode ➜ /usr/bin $ source ~/.bashrc
```

## Usage.
```bash
vscode ➜ /usr/bin $ pwd
/usr/bin
vscode ➜ /usr/bin $ ./duckdb 
v0.8.1 6536a77232
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
D select 43;
┌───────┐
│  43   │
│ int32 │
├───────┤
│    43 │
└───────┘
D .exit
vscode ➜ /usr/bin $ 
```
