SELECT *
FROM anime;

SELECT *
FROM manga;

SELECT *
FROM genres;

DROP TABLE IF EXISTS anime;

DROP TABLE IF EXISTS manga;

DROP TABLE IF EXISTS genres;

CREATE TABLE genres
(
    id   SERIAL PRIMARY KEY UNIQUE NOT NULL,
    name TEXT                      NOT NULL
);

CREATE TABLE anime
(
    id            SERIAL PRIMARY KEY UNIQUE      NOT NULL,
    genre_id      INTEGER REFERENCES genres (id) NOT NULL,
    title         TEXT                           NOT NULL,
    titlejapanese TEXT                           NOT NULL,
    source        TEXT                           NOT NULL,
    episodes      INTEGER,
    kind          TEXT                           NOT NULL,
    score         DOUBLE PRECISION               NOT NULL,
    status        TEXT                           NOT NULL,
    synopsis      TEXT                           NOT NULL
);


CREATE TABLE manga
(
    id            SERIAL PRIMARY KEY UNIQUE      NOT NULL,
    genre_id      INTEGER REFERENCES genres (id) NOT NULL,
    title         TEXT                           NOT NULL,
    titlejapanese TEXT                           NOT NULL,
    volumes       INTEGER                        NOT NULL,
    chapters      INTEGER                        NOT NULL,
    score         DOUBLE PRECISION               NOT NULL,
    status        TEXT                           NOT NULL,
    synopsis      TEXT                           NOT NULL
);

INSERT INTO genres
VALUES (DEFAULT, 'Action'),
       (DEFAULT, 'Comedy'),
       (DEFAULT, 'Drama'),
       (DEFAULT, 'Psychological'),
       (DEFAULT, 'Horror'),
       (DEFAULT, 'Supernatural'),
       (DEFAULT, 'Romance'),
       (DEFAULT, 'Adventure');



INSERT INTO anime
VALUES (DEFAULT, 3, 'Angel Beats!', 'エンジェルビーツ!', 'Light Novel', 13, 'TV', 8.8, 'RELEASED',
        'Yuzuru Otonashi wakes up in an unfamiliar place, not remembering who he is or what happened to him. Some abnormal schoolgirl named Yuri Nakamura claims that Yuzuru is dead and encourages him to join the Underworld Front.'),
       (DEFAULT, 8, 'KonoSuba', 'この素晴らしい世界に祝福を', 'Light Novel', 10, 'TV', 8.13, 'RELEASED',
        'It just so happened that the life of Kazuma Sato, a hikikomori addicted to video games, was cut short by a traffic accident. Such is the harsh price to pay for a good deed, for the salvation of a person. However, waking up, the guy suddenly realized that he was alive and well, and in front of him he saw a pretty girl, Aqua, who introduced herself as a goddess.'),
       (DEFAULT, 3, 'Steins Gate', ' シュタインズ ゲート', 'Manga', 24, 'TV', 9.09, 'RELEASED',
        'Having rented an apartment in Akihabara, the self-proclaimed mad scientist Okabe Rintaro set up a laboratory there and in the company of his childhood friend Sina Mayuri and otaku hacker Hasida Itaru invents gadgets of the future.'),
       (DEFAULT, 2, 'Non Non Biyori', 'のんのんびより', 'Manga', 12, 'TV', 7.95, 'RELEASED',
        'Hotaru Ichidzho, a fifth grade elementary school student, moved with her parents from Tokyo to the countryside. And I was extremely surprised that local children know nothing about the big city. But she was even more surprised when she learned that the only school in the whole village was attended by only four students.'),
       (DEFAULT, 2, 'Rent a Girlfriend', '彼女、お借りします', 'Manga', 12, 'TV', 7.31, 'RELEASED',
        'Kazuya Kinoshita is an ordinary university student who leads a modest life and has been dating a girl for a month now ... But happiness does not last long, and a friend informs the guy that she is leaving him for another.'),
       (DEFAULT, 4, 'Death Note', 'デスノート', 'Manga', 37, 'TV', 8.63, 'RELEASED',
        'A bored Shinigami Ryuk throws one of his Death Notes into the human world. Just for fun, see what comes of it.'),
       (DEFAULT, 1, 'Fullmetal Alchemist: Brotherhood', '鋼の錬金術師', 'Manga', 64, 'TV', 9.15, 'RELEASED',
        'There are alchemists in this world - people who possess the art of alchemy, the ability to manipulate matter and transform matter. All of them are limited by the basic Law of alchemy: you cannot get something by alchemical means without sacrificing something equivalent to what you received.'),
       (DEFAULT, 5, 'Another', 'アナザー', 'Light Novel', 12, 'TV', 7.5, 'RELEASED',
        '26 years ago, in one of the third grade of high school, there was a student named Misaki. He was good at school and in sports, handsome in appearance and was popular in his class. However, he died suddenly. After that, his classmates decided to act as if he was still around until graduation.'),
       (DEFAULT, 6, 'Bakemonogatari', '化物語', 'Light Novel', 15, 'TV', 8.35, 'RELEASED',
        'The story of Koyomi Araragi, a third grade high school student and former vampire. One day he was walking up the school stairs, and a girl fell on top of him. However, she fell slowly, and her weight turned out to be strange, only a few kilograms.'),
       (DEFAULT, 7, 'Clannad', 'CLANNAD', 'PC Game', 23, 'TV', 8.03, 'RELEASED',
        'Tomoya Okazaki is a bum, confident that life is boring, and he himself is worthless. He hates his city. Together with their friend Sunohara, they constantly skip school and do whatever they please.'),
       (DEFAULT, 7, '5 Centimeters Per Second', '秒速５センチメートル', 'Original', 3, 'Movie', 8.0, 'RELEASED',
        'Five Centimeters Per Second is a romantic drama that focuses on the mundane and harsh reality of long-distance relationships.'),
       (DEFAULT, 7, 'My Senpai is Annoying', '先輩がうざい後輩の話', 'Manga', 12, 'TV', 7.86, 'ONGOING',
        'Futaba Igarashi''s new work may have been amazing, but the picture is spoiled by one significant flaw - the annoying Harumi Takeda.');

INSERT INTO manga
VALUES (DEFAULT, 2, 'Chainsaw Man', 'チェンソーマン', 11, 104, 9.83, 'RELEASED',
        'I always dreamed of living an ordinary life: sleeping in a warm bed, eating toast with jam in the morning, going on dates with my girlfriend and smiling every day.'),
       (DEFAULT, 1, 'Kimetsu no Yaiba', '鬼滅の刃', 23, 235, 9.78, 'RELEASED',
        'Tanjiro is the eldest son of a family that has lost a breadwinner. One day, he leaves for another city to sell charcoal, but ends up staying overnight in someone else''s house instead of going home.'),
       (DEFAULT, 7, 'Ijiranaide, Nagatoro-san', 'イジらないで、長瀞さん', 11, 109, 8.78, 'ONGOING',
        'Nagatoro is a transferred high school student who enjoys teasing her senpai. But he will endure everything. And even if he has to go through all kinds of bullying - he will forgive everything because he is in love.'),
       (DEFAULT, 3, 'Three Days of Happiness', '寿命を買い取ってもらった。一年につき、一万円で。', 2, 18, 9.91, 'RELEASED',
        'It looks like nothing good was in store for me in the future. Therefore, the cost of my life was only 10,000 yen per year.'),
       (DEFAULT, 5, 'Mieru Ko-chan', '見える子ちゃん', 7, 47, 8.86, 'ONGOING',
        'Miko is a high school student cursed with a sixth sense: she has the ability to see ghosts. She tries to live a normal life, but it''s hard when the ghosts continue to haunt her.'),
       (DEFAULT, 4, 'Welcome to the NHK!', 'NHKにようこそ!', 8, 40, 7.61, 'RELEASED',
        'Conspiracies. Everywhere. Nippon Hikikomori Kyokai. Welcome to Loneliness. Welcome to the Ranks of the Society Forsaken. You are a hikikomori and your life is insignificant.'),
       (DEFAULT, 6, 'Noragami', 'ノラガミ', 24, 123, 8.9, 'ONGOING',
        'God who can overturn a sentence with a sword. In the human world there is a God who can penetrate into the very heart of a person. When you are sad, when you are tired, if you look up to heaven, you will see a phone number.'),
       (DEFAULT, 8, 'No Game No Life', 'ノーゲーム・ノーライフ', 11, 66, 8.76, 'ONGOING',
        'The manga is tied to 18-year-old Sora and 11-year-old Shiro, siblings whose reputation for impeccable NEET, hikikomori and gamers has spawned legends all over the internet.');

