# Postgres-Database-Transaction-in-Golang
Demonstrate usage of transaction in Postgres Database using Golang

# Table schema
postgres=# \d sales.employee
              Table "sales.employee"
 Column |  Type   | Collation | Nullable | Default 
--------+---------+-----------+----------+---------
 id     | integer |           | not null | 
 name   | text    |           |          | 
 role   | text    |           |          | 
 salary | integer |           |          | 
Indexes:
    "employee_pkey" PRIMARY KEY, btree (id)
    
    
# For Single insert db query execution
C02YM0ADJG5J:plsql pkuma679$ go run main.go 
Executing with transaction
Total time taken1 :  0.520
Executing without transaction
Total time taken1 :  0.788
Executing with transaction
Total time taken1 :  0.620
Executing without transaction
Total time taken1 :  0.551
Executing with transaction
Total time taken1 :  0.516
Executing without transaction
Total time taken1 :  0.556
Executing with transaction
Total time taken1 :  0.798
Executing without transaction
Total time taken1 :  0.512
Executing with transaction
Total time taken1 :  0.518
Executing without transaction
Total time taken1 :  1.125

# For 500 insert db query execution
C02YM0ADJG5J:plsql pkuma679$ go run main.go 
Executing with transaction
Total time taken1 :  13.415
Executing without transaction
Total time taken1 :  10.295
Executing with transaction
Total time taken1 :  11.772
Executing without transaction
Total time taken1 :  10.107
Executing with transaction
Total time taken1 :  12.091
Executing without transaction
Total time taken1 :  10.308
Executing with transaction
Total time taken1 :  11.933
Executing without transaction
Total time taken1 :  10.245
Executing with transaction
Total time taken1 :  11.944
Executing without transaction
Total time taken1 :  10.191
