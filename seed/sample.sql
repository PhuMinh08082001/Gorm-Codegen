
create table users (
                       id bigserial primary key,
                       name varchar(128)
);

INSERT INTO public.users ("name")
VALUES('Peter'),('Messi'),('Hana'),('Harry');

create table ordertabs (
                           order_id bigserial,
                           user_id bigint,
                           primary key(order_id),
                           FOREIGN KEY (user_id)
                               REFERENCES users(id)
);

INSERT INTO public.ordertabs
(user_id)
VALUES(1), (2), (3), (1);


CREATE TABLE products (
                          product_id bigserial primary key,
                          product_name varchar(128) ,
                          price int,
                          num_inventory int
);

CREATE TABLE categories (
                            category_id bigserial primary key,
                            category_name varchar(128)
);

CREATE TABLE product_category (
                                  product_id bigint,
                                  category_id bigint,
                                  primary key(product_id, category_id),
                                  FOREIGN KEY (product_id)
                                      REFERENCES products(product_id),
                                  FOREIGN KEY (category_id)
                                      REFERENCES categories(category_id)
);

CREATE TABLE order_detail (
                              order_id bigint,
                              product_id bigint,
                              primary key(order_id, product_id),
                              FOREIGN KEY (product_id)
                                  REFERENCES products(product_id),
                              FOREIGN KEY (order_id)
                                  REFERENCES ordertabs(order_id)
);











