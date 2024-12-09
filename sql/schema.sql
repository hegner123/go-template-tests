CREATE TABLE IF NOT EXISTS pages
(
    page_id   int SERIAL DEFAULT VALUE null
        primary key,
    parent_id int                      null,
    route_id  int                      null,
    label     VARCHAR(255)             not null,
    data      TEXT                     null,
    constraint pages_pages_page_id_fk
        foreign key (parent_id) references pages (page_id)
            on update cascade on delete cascade
);

CREATE TABLE IF NOT EXISTS routes (
    route_id INT NULL
        PRIMARY KEY,
    slug TEXT NULL
);
