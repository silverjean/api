INSERT INTO users (name, nick, email, password)
values
("Jean", "none", "jean@gmail.com", "$2a$10$8Pb7cUA.JrGtqMHLmkS/mOhwc.ZgG/LeXElUABvZJ7Oajkc1YMWGu")
("Gabriel", "ilha", "gabriel@gmail.com", "$2a$10$8Pb7cUA.JrGtqMHLmkS/mOhwc.ZgG/LeXElUABvZJ7Oajkc1YMWGu")
("Bruna", "bru", "bru@gmail.com", "$2a$10$8Pb7cUA.JrGtqMHLmkS/mOhwc.ZgG/LeXElUABvZJ7Oajkc1YMWGu")
("Gabriela", "bibi", "bibi@gmail.com", "$2a$10$8Pb7cUA.JrGtqMHLmkS/mOhwc.ZgG/LeXElUABvZJ7Oajkc1YMWGu")
("Victoria", "bic", "bic@gmail.com", "$2a$10$8Pb7cUA.JrGtqMHLmkS/mOhwc.ZgG/LeXElUABvZJ7Oajkc1YMWGu")

INSERT INTO followers (user_id, follower_id)
VALUES
(1,2)
(1,4)
(2,1)
