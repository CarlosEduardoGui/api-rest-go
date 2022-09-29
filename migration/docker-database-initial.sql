CREATE TABLE personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES 
('teste1', 'grande teste1'),
('teste2', 'grande teste2')