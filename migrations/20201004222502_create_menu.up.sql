CREATE TABLE menu (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    organization_id BIGSERIAL NOT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE
);