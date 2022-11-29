-- SELECT (Column) FROM (TABLE NAME) WHERE (IF)

-- * == ALL 

SELECT * FROM ngletutor.helloworld;

SELECT * FROM ngletutor.category;

SELECT * FROM ngletutor.post;


SELECT * FROM ngletutor.post WHERE category_id = 2 -- category_id is foreign key, category table PRIMARY KEY (id) 

SELECT id, create_at FROM ngletutor.post WHERE id = 1

SELECT id, username FROM ngletutor.user