CREATE TABLE company_members (
    user_id INT NOT NULL,
    company_id INT NOT NULL,
    position VARCHAR NOT NULL,
    PRIMARY KEY (user_id, company_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
);