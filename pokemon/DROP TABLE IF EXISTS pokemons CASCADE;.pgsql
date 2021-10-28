DROP TABLE IF EXISTS pokemons CASCADE;

CREATE TABLE pokemons (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    weight INTEGER NOT NULL
);

INSERT INTO pokemons (name, type, weight)
VALUES('Bulbasaur','grass', 69);
INSERT INTO pokemons (name, type, weight)
VALUES('Charmander','fire', 85);
INSERT INTO pokemons (name, type, weight)
VALUES('Squirtle','water', 90);
INSERT INTO pokemons (name, type, weight)
VALUES('Caterpie','bug', 29);
INSERT INTO pokemons (name, type, weight)
VALUES('Pidgey','flying', 18);
INSERT INTO pokemons (name, type, weight)
VALUES('Rattata','normal', 35);
INSERT INTO pokemons (name, type, weight)
VALUES('Articuno','ice', 554);
INSERT INTO pokemons (name, type, weight)
VALUES('Zapdos','electric', 526);
INSERT INTO pokemons (name, type, weight)
VALUES('Moltres','fire', 600);
INSERT INTO pokemons (name, type, weight)
VALUES('Mewtwo','psychic', 1220);