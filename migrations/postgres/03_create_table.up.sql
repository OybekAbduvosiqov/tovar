
CREATE TABLE "order" (
    id UUID PRIMARY KEY,
    total_price NUMERIC NOT NULL,
    branch_id UUID NOT NULL REFERENCES branch(id),
    client_id UUID NOT NULL REFERENCES client(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE bucket (
    id UUID PRIMARY KEY,
    branch_id UUID NOT NULL REFERENCES branch(id),
    client_id UUID NOT NULL REFERENCES client(id),
    product_id UUID NOT NULL REFERENCES product(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


SELECT
    b.id,
    b.branch_id,
    b.client_id,
    b.created_at,
    b.updated_at,
    p.price
FROM branch AS b
LEFT JOIN product AS p ON p.id = b.product_id