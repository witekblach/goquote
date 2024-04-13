CREATE TABLE IF NOT EXISTS quotes
(
    text   text NOT NULL,
    viewed int  NOT NULL DEFAULT 0
) STRICT;

-- some sample data
INSERT INTO quotes (TEXT, VIEWED)
VALUES ("Mało jem, mało śpię, mało mówię, dużo piszę. Duże sumy, luźne ciuchy, zakoduje sobie życie - otsochodzi", 0),
       ("I know it hurts sometimes but you'll get over it; You'll find another life to live - lil uzi vert", 0);

