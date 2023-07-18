CREATE TABLE cards_transaction (
  card_number TEXT,
  description TEXT,
  amount NUMERIC,
  currency TEXT,
  date TIMESTAMP
);

INSERT INTO cards_transaction (
  card_number,
  description,
  amount,
  currency,
  date
) VALUES ( 
  '1234',
  'Mercado Livre',
  100,
  'BRL',
  '2023-07-01T10:00:00'
);

INSERT INTO cards_transaction (
  card_number,
  description,
  amount,
  currency,
  date
) VALUES ( 
  '1234',
  'Amazon',
  300,
  'USD',
  '2023-07-01T10:00:00'
);

INSERT INTO cards_transaction (
  card_number,
  description,
  amount,
  currency,
  date
) VALUES ( 
  '1234',
  'Submarino',
  50,
  'BRL',
  '2023-07-01T10:00:00'
);

INSERT INTO cards_transaction (
  card_number,
  description,
  amount,
  currency,
  date
) VALUES ( 
  '1234',
  'Extra',
  1000,
  'BRL',
  '2023-06-01T10:00:00'
);

INSERT INTO cards_transaction (
  card_number,
  description,
  amount,
  currency,
  date
) VALUES ( 
  '1234',
  'Google',
  50,
  'USD',
  '2023-06-01T10:00:00'
);
