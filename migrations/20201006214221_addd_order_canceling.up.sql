ALTER TABLE orders
    ADD canceled BOOLEAN NOT NULL DEFAULT FALSE,
    ADD in_progress BOOLEAN NOT NULL DEFAULT FALSE;