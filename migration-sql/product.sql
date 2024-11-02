CREATE TABLE public.products (
	id serial4 NOT NULL,
	sku text NOT NULL,
	"name" text NOT NULL,
	price numeric(10, 2) NOT NULL,
	inventory_qty int8 NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

