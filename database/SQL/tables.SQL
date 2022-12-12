CREATE TABLE ngletutor.helloworld (
	id INT NOT NULL AUTO_INCREMENT,
    world VARCHAR(45) NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,

    PRIMARY KEY(id)
)

CREATE TABLE ngletutor.category (
	id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(45) NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,

    PRIMARY KEY(id)
)

CREATE TABLE ngletutor.post (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(30) NOT NULL, 
    content TEXT NOT NULL,
    -- view Default: 0 
    view INT, 
    category_id INT NOT NULL,
    author_id INT NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,

    PRIMARY KEY(id),
    FOREIGN KEY(category_id) REFERENCES ngletutor.category(id) ON DELETE CASCADE,
    FOREIGN KEY(author_id) REFERENCES ngletutor.user(id) ON DELETE CASCADE
)

CREATE TABLE ngletutor.user (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    password TEXT NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,

    PRIMARY KEY(id)
)