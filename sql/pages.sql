CREATE TABLE pages (
    page_id INT AUTO_INCREMENT
        PRIMARY KEY,
    parent_id INT NULL,
    route_id INT NULL,
    label VARCHAR(255) NOT NULL,
    data TEXT NULL,
    CONSTRAINT pages_pages_page_id_fk
        FOREIGN KEY (parent_id) REFERENCES pages (page_id)
            ON UPDATE CASCADE ON DELETE CASCADE
);

