##### DATABASE SETUP

```bash
mysql -u <username> -p
```

```sql
CREATE DATABASE inventory;

USE inventory;

CREATE TABLE products(
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  quantity INT,
  price FLOAT(10, 7),
  PRIMARY KEY(id)
);

INSERT INTO products VALUES(
  1,
  "chair",
  100,
  200.00
);
```
