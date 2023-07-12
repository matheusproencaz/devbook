INSERT INTO users(name, nick, email, password) VALUES
('Usuario 1', 'user1', 'user1@gmail.com', '$2a$10$SLU6BehOH8EhOjZanZ2fC.4z7QRynJzw5.6.OuCH0ypz2EnniOgiq'),
('Usuario 2', 'user2', 'user2@gmail.com', '$2a$10$SLU6BehOH8EhOjZanZ2fC.4z7QRynJzw5.6.OuCH0ypz2EnniOgiq'),
('Usuario 3', 'user3', 'user3@gmail.com', '$2a$10$SLU6BehOH8EhOjZanZ2fC.4z7QRynJzw5.6.OuCH0ypz2EnniOgiq'),
('Usuario 4', 'user4', 'user4@gmail.com', '$2a$10$SLU6BehOH8EhOjZanZ2fC.4z7QRynJzw5.6.OuCH0ypz2EnniOgiq');

INSERT INTO followers(user_id, follower_id) VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO posts(title, content, author_id) VALUES
("Title Post User 1", "Content Post User 1", 1),
("Title Post User 2", "Content Post User 2", 2),
("Title Post User 3", "Content Post User 3", 3);