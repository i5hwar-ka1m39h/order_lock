-- name: CreateEmailDelivery :one
INSERT INTO email_delivery (
	recipient,
	sender,
	cc,
	bcc,
	subject,
	body,
	delivery_status,
	is_opened,
	opened_at,
	sent_at,
	failed_at,
	error_message,
	retry_count,
	max_retries,
	metadata,
	external_id
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15::jsonb, $16
) RETURNING *;


-- name: GetEmailDeliveries :many
SELECT * FROM email_delivery;

-- name: SelectEmailDelivery :one
SELECT * FROM email_delivery WHERE id=$1;