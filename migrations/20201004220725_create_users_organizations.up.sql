CREATE TABLE users_organizations (
    organization_id BIGSERIAL NOT NULL,
    user_id BIGSERIAL NOT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
