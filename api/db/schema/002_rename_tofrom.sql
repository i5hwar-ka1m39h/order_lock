-- +goose up
ALTER TABLE email_delivery
    RENAME COLUMN "to" TO recipient;

ALTER TABLE email_delivery
    RENAME COLUMN "from" TO sender;

-- +goose down
ALTER TABLE email_delivery
    RENAME COLUMN recipient TO "to";

ALTER TABLE email_delivery
    RENAME COLUMN sender TO "from";
