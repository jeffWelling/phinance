CREATE TABLE Accounts (
    account_id integer PRIMARY KEY,
    account_name varchar,
    account_type varchar,
    created_at timestamp,
    currency varchar
);

CREATE TABLE Transactions (
    from_account integer,
    to_account integer,
    amount integer,
    transaction_date timestamp
);

CREATE TABLE AccountRelationships (
    parent_account_id integer,
    child_account_id integer
);
