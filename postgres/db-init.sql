create schema if not exists pixelpay;

CREATE TABLE IF NOT EXISTS pixelpay.person (
      id UUID PRIMARY KEY NOT NULL,
      name VARCHAR(255) NOT NULL,
      document VARCHAR(255) NOT NULL,
      email VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS pixelpay_person_document_idx ON pixelpay.person (document);

CREATE TABLE if NOT EXISTS pixelpay.transaction (
      id UUID PRIMARY KEY NOT NULL,
      payer UUID NOT NULL, -- create who pays
      payee UUID NOT NULL, -- create who receives
      value DECIMAL(10,2) NOT NULL,
      created_at TIMESTAMP NOT NULL,

      CONSTRAINT fk_payer FOREIGN KEY (payer) REFERENCES pixelpay.person(id),
      CONSTRAINT fk_payee FOREIGN KEY (payee) REFERENCES pixelpay.person(id)
)
