BEGIN;

INSERT INTO customers (id, first_name, last_name, email) VALUES
    ('6aa6cb6c-6054-4943-a0a7-f279cf6ceabd', 'John', 'Doe', 'john.doe@example.com');

INSERT INTO accounts (id, customer_id, account_type, account_number, status) VALUES
    ('d7ee4877-7645-461a-b2cc-2f2c8f6a7284', '6aa6cb6c-6054-4943-a0a7-f279cf6ceabd', 'ISA', '1234567890', 'ACTIVE');

INSERT INTO funds (id, name, category, currency, risk_return) VALUES
    ('cb91e975-d8bc-423b-bc99-fa6f396c2eaf', 'Cushon Equities Fund', 'EQUITY', 'GBP', 'LOW');

COMMIT;