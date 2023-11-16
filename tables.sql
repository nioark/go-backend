CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    password VARCHAR(255)
);

-- Create pedidos table
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    quantidade INT,
    usuario_id INT REFERENCES usuarios(id) ON DELETE CASCADE
);