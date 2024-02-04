CREATE TABLE IF NOT EXISTS users(
    id serial NOT NULL, 
    lineuserid varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (lineuserid)
);


CREATE TABLE IF NOT EXISTS menus(
    id serial NOT NULL, 
    userid integer NOT NULL,
    menuname varchar(255) NOT NULL,
    imageurl varchar(255),
    ingredients TEXT[] NOT NULL,
    quantities TEXT[] NOT NULL,
    recipes TEXT[] NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT menus_userid_fkey
    FOREIGN KEY (userid)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS likes(
    id serial NOT NULL, 
    userid integer NOT NULL,
    menuid integer NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (userid, menuid),
    CONSTRAINT likes_userid_fkey
    FOREIGN KEY (userid)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
    CONSTRAINT likes_menuid_fkey
    FOREIGN KEY (menuid)
    REFERENCES menus(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
