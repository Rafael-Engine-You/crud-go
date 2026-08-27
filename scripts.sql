CREATE TABLE
    clientes (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(150) NOT NULL,
        email VARCHAR(150) NOT NULL UNIQUE,
        telefone VARCHAR(20),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    categorias (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(100) NOT NULL UNIQUE
    );

CREATE TABLE
    produtos (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(150) UNIQUE NOT NULL,
        preco NUMERIC(10, 2) NOT NULL,
        categoria_id INTEGER NOT NULL,
        CONSTRAINT fk_categoria FOREIGN KEY (categoria_id) REFERENCES categorias (id) ON UPDATE NO ACTION ON DELETE CASCADE
    );