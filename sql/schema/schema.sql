BEGIN;

CREATE TABLE customers (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customers(id),
    account_type VARCHAR(20) NOT NULL, -- ISA
    account_number VARCHAR(20) UNIQUE NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE funds (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(20) NOT NULL, -- EQUITY, BOND, etc.
    currency VARCHAR(3) NOT NULL DEFAULT 'GBP',
    risk_return VARCHAR(10) -- LOW, MEDIUM, HIGH, VERY_HIGH
);

CREATE TABLE investments (
     id UUID PRIMARY KEY,
     account_id UUID NOT NULL REFERENCES accounts(id),
     amount DECIMAL(7, 2) NOT NULL,
     status VARCHAR(10) NOT NULL DEFAULT 'PENDING',
     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_investments_account_id ON investments(account_id);

CREATE TABLE investment_funds (
    id UUID PRIMARY KEY,
    investment_id UUID NOT NULL REFERENCES investments(id) ON DELETE CASCADE,
    fund_id UUID NOT NULL REFERENCES funds(id),
    amount DECIMAL(7, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(investment_id, fund_id)
);

CREATE INDEX idx_investment_funds_investment_id ON investment_funds(investment_id);
CREATE INDEX idx_investment_funds_fund_id ON investment_funds(fund_id);


COMMIT;