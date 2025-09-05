-- exercise 1:  基本CRUD操作

INSERT INTO students (name, age, grade) VALUES('张三', 20, '三年级');

SELECT * FROM students WHERE age > 18;

UPDATE students SET grade = '四年级' Where name = '张三';

DELETE FROM students where age < 15;

-- exercise 2:  事务语句

BEGIN;

SELECT balance INTO temp_balance FROM accounts WHERE id = 'A';

IF balance < 100 THEN
    -- ROLLBACK;
    RAISE EXCEPTION '余额不足';
    -- 在事务内，一旦 RAISE EXCEPTION 被触发，整个事务将进入异常状态，且所有后续的 SQL 语句都不会再执行。
    -- RAISE EXCEPTION 会立即引发一个错误，并且会导致当前事务的回滚
END IF;

UPDATE accounts SET balance = balance - 100 WHERE id = 'A';

UPDATE accounts SET balance = balance + 100 WHERE id = 'B';

INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES('A', 'B', 100);

COMMIT;
