BEGIN;

INSERT INTO customers (id, first_name, last_name, email) VALUES
    ('6aa6cb6c-6054-4943-a0a7-f279cf6ceabd', 'John', 'Doe', 'john.doe@example.com');

INSERT INTO accounts (id, customer_id, account_type, account_number, status) VALUES
    ('d7ee4877-7645-461a-b2cc-2f2c8f6a7284', '6aa6cb6c-6054-4943-a0a7-f279cf6ceabd', 'ISA', '1234567890', 'ACTIVE');

INSERT INTO funds (id, name, category, currency, risk_return) VALUES
    ('cb91e975-d8bc-423b-bc99-fa6f396c2eaf', 'Cushon Equities Fund', 'EQUITY', 'GBP', 'LOW');

INSERT INTO investments (id, account_id, amount, status) VALUES
    ('f9d7338c-62e9-4f60-8d3a-9082f1616a23', 'd7ee4877-7645-461a-b2cc-2f2c8f6a7284', '20000', 'PENDING');

INSERT INTO investment_funds (id, investment_id, fund_id, amount) VALUES
    ('1c570240-b164-4111-a321-25d0467eb5ce', 'f9d7338c-62e9-4f60-8d3a-9082f1616a23', 'cb91e975-d8bc-423b-bc99-fa6f396c2eaf', '20000');

COMMIT;