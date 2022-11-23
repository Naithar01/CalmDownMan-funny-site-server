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
    category_id INT NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,

    PRIMARY KEY(id),
    FOREIGN KEY(category_id) REFERENCES ngletutor.category(id) ON DELETE CASCADE
)