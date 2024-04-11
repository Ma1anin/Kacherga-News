use newsApp

CREATE TABLE IF NOT EXISTS [Role] (
    ID SERIALIZABLE,
    [name] VARCHAR(20),
    PRIMARY KEY (ID)
)

CREATE TABLE IF NOT EXISTS User (
    login VARCHAR(20),
    fullName VARCHAR(50) NOT NULL,
    roleID INT NOT NULL,
    avatarUrl VARCHAR(255),
    PRIMARY KEY (login),
    FOREIGN KEY (roleID) REFERENCES [Role] (ID)
)

CREATE TABLE IF NOT EXISTS News (
    ID SERIALIZABLE,
    title VARCHAR(100) NOT NULL,
    content VARCHAR(255) NOT NULL,
    imageUrl VARCHAR(255),
    [createAt] TIMESTAMP DEFAULT NOW(),
    authorID INT NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (authorID) REFERENCES User (ID)
)

CREATE TABLE IF NOT EXISTS [Event] (
    ID SERIALIZABLE,
    title VARCHAR(100) NOT NULL,
    content VARCHAR(255) NOT NULL,
    [createAt] TIMESTAMP DEFAULT NOW(),
    authorID INT NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (authorID) REFERENCES User (ID)
)

INSERT INTO Role ([name])
VALUES ('admin'),
       ('user')