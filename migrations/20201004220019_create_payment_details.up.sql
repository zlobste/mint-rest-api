CREATE TABLE payment_details (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    bank VARCHAR NOT NULL,
    account VARCHAR NOT NULL UNIQUE,
    organization_id BIGSERIAL NOT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE
);


