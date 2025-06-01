BEGIN;

-- Insert test customers
INSERT INTO customers (id, first_name, last_name, email) VALUES
    ('test-id', 'Test', 'User', 'test.user@example.com');
COMMIT;
