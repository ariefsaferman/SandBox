--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)

-- Started on 2022-12-13 16:00:44 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16589)
-- Name: authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.authors OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16588)
-- Name: authors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_id_seq OWNER TO postgres;

--
-- TOC entry 3411 (class 0 OID 0)
-- Dependencies: 214
-- Name: authors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_id_seq OWNED BY public.authors.id;


--
-- TOC entry 217 (class 1259 OID 16602)
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying NOT NULL,
    description character varying NOT NULL,
    quantity integer NOT NULL,
    cover character varying,
    author_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.books OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16601)
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO postgres;

--
-- TOC entry 3412 (class 0 OID 0)
-- Dependencies: 216
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- TOC entry 221 (class 1259 OID 16637)
-- Name: borrowing_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.borrowing_records (
    id integer NOT NULL,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    status character varying NOT NULL,
    borrowing_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    returning_date timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.borrowing_records OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16636)
-- Name: borrowing_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.borrowing_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.borrowing_records_id_seq OWNER TO postgres;

--
-- TOC entry 3413 (class 0 OID 0)
-- Dependencies: 220
-- Name: borrowing_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.borrowing_records_id_seq OWNED BY public.borrowing_records.id;


--
-- TOC entry 219 (class 1259 OID 16626)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    phone character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    password character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16625)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3414 (class 0 OID 0)
-- Dependencies: 218
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3231 (class 2604 OID 16592)
-- Name: authors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors ALTER COLUMN id SET DEFAULT nextval('public.authors_id_seq'::regclass);


--
-- TOC entry 3234 (class 2604 OID 16605)
-- Name: books id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- TOC entry 3240 (class 2604 OID 16640)
-- Name: borrowing_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.borrowing_records ALTER COLUMN id SET DEFAULT nextval('public.borrowing_records_id_seq'::regclass);


--
-- TOC entry 3237 (class 2604 OID 16629)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3399 (class 0 OID 16589)
-- Dependencies: 215
-- Data for Name: authors; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.authors (id, name, created_at, updated_at, deleted_at) VALUES (1, 'arief', '2022-12-08 18:14:49.365623', '2022-12-08 18:14:49.365623', NULL);
INSERT INTO public.authors (id, name, created_at, updated_at, deleted_at) VALUES (2, 'dafur', '2022-12-08 18:14:49.365623', '2022-12-08 18:14:49.365623', NULL);


--
-- TOC entry 3401 (class 0 OID 16602)
-- Dependencies: 217
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (3, 'negeri di ujung tandik', 'scifi-action', 10, '', 1, '2022-12-08 18:24:03.550168', '2022-12-08 18:24:03.550168', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (5, 'buku last', 'scifi-action', 10, '', 1, '2022-12-08 18:26:05.227429', '2022-12-08 18:26:05.227429', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (11, 'bisayukbisa', 'scifi-action', 1, '', 1, '2022-12-09 16:56:32.54253', '2022-12-09 16:56:32.54253', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (12, 'ayoooo', 'scifi-action', 0, '', 1, '2022-12-09 17:02:27.509775', '2022-12-09 17:02:27.509775', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (1, 'hujan', 'scifi-comedy', 5, '', 1, '2022-12-08 18:18:37.488653', '2022-12-09 19:34:39.301795', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (2, 'harry potter', 'scifi-action', 8, '', 2, '2022-12-08 18:21:46.584869', '2022-12-12 18:03:56.047812', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (4, 'neraka', 'scifi-action', 9, '', 1, '2022-12-08 18:25:36.306712', '2022-12-13 11:04:52.994547', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (14, 'Alice in Border Land', 'Book of a wizard', 10, '', 1, '2022-12-13 11:56:07.221387', '2022-12-13 11:56:07.221387', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (13, 'test', 'scifi-action', 5, '', 1, '2022-12-09 17:07:34.897073', '2022-12-13 13:13:45.778664', NULL);
INSERT INTO public.books (id, title, description, quantity, cover, author_id, created_at, updated_at, deleted_at) VALUES (9, 'bisayukbisa2', 'scifi-action', 0, '', 1, '2022-12-09 16:50:17.910074', '2022-12-13 13:25:27.882861', NULL);


--
-- TOC entry 3405 (class 0 OID 16637)
-- Dependencies: 221
-- Data for Name: borrowing_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (1, 1, 1, 'Returned', '2022-12-09 16:22:16.886849', '2022-12-09 19:34:39.300983', '2022-12-09 16:22:16.887338', '2022-12-09 19:34:39.30115', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (2, 2, 2, 'Returned', '2022-12-12 07:00:51.561198', '2022-12-12 07:01:26.457524', '2022-12-12 07:00:51.567798', '2022-12-12 07:01:26.457659', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (3, 2, 2, 'Borrowed', '2022-12-12 07:03:15.714178', '0001-01-01 00:00:00', '2022-12-12 07:03:15.716339', '2022-12-12 07:03:15.716339', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (4, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 16:35:57.240471', '2022-12-12 16:24:57.154561', '2022-12-12 16:35:57.24067', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (5, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 16:41:47.638986', '2022-12-12 16:37:27.82442', '2022-12-12 16:41:47.639109', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (6, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 16:45:24.706547', '2022-12-12 16:44:55.773511', '2022-12-12 16:45:24.706697', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (7, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 16:54:49.884343', '2022-12-12 16:53:34.161329', '2022-12-12 16:54:49.884544', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (8, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:14:05.09974', '2022-12-12 17:03:03.602739', '2022-12-12 17:14:05.099865', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (9, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:15:17.287504', '2022-12-12 17:14:12.258357', '2022-12-12 17:15:17.28759', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (10, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:19:40.840095', '2022-12-12 17:15:22.87564', '2022-12-12 17:19:40.840153', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (11, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:27:29.488731', '2022-12-12 17:27:22.005641', '2022-12-12 17:27:29.488839', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (12, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:29:17.180494', '2022-12-12 17:29:05.015443', '2022-12-12 17:29:17.180555', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (13, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:32:15.671401', '2022-12-12 17:32:08.085403', '2022-12-12 17:32:15.671495', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (14, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:34:27.602504', '2022-12-12 17:32:43.629277', '2022-12-12 17:34:27.602575', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (15, 1, 2, 'Returned', '0001-01-01 00:00:00', '2022-12-12 17:35:03.561666', '2022-12-12 17:34:34.92448', '2022-12-12 17:35:03.561721', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (16, 1, 2, 'Borrowed', '0001-01-01 00:00:00', '0001-01-01 00:00:00', '2022-12-12 18:03:56.047367', '2022-12-12 18:03:56.047367', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (17, 1, 4, 'Borrowed', '0001-01-01 00:00:00', '0001-01-01 00:00:00', '2022-12-13 11:04:52.98618', '2022-12-13 11:04:52.98618', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (19, 1, 13, 'Returned', '0001-01-01 00:00:00', '2022-12-13 13:13:45.777363', '2022-12-13 13:01:41.865412', '2022-12-13 13:13:45.777505', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (18, 1, 9, 'Returned', '0001-01-01 00:00:00', '2022-12-13 13:19:47.378273', '2022-12-13 11:08:13.501352', '2022-12-13 13:19:47.378339', NULL);
INSERT INTO public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) VALUES (20, 1, 9, 'Borrowed', '0001-01-01 00:00:00', '0001-01-01 00:00:00', '2022-12-13 13:25:27.882146', '2022-12-13 13:25:27.882146', NULL);


--
-- TOC entry 3403 (class 0 OID 16626)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, name, email, phone, created_at, updated_at, deleted_at, password) VALUES (2, 'rei', 'rei@shopee.com', '081230', '2022-12-12 06:53:15.516657', '2022-12-12 06:53:15.516657', NULL, NULL);
INSERT INTO public.users (id, name, email, phone, created_at, updated_at, deleted_at, password) VALUES (1, 'arief', 'arief.saferman@shopee.com', '091922', '2022-12-09 16:15:49.616299', '2022-12-09 16:15:49.616299', NULL, '$2a$04$0PNbSOX837qG7kTFnG6o3eBxik0WJDid6.YuNd27bQSnS4Njd3WD6');


--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 214
-- Name: authors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.authors_id_seq', 2, true);


--
-- TOC entry 3416 (class 0 OID 0)
-- Dependencies: 216
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.books_id_seq', 15, true);


--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 220
-- Name: borrowing_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.borrowing_records_id_seq', 20, true);


--
-- TOC entry 3418 (class 0 OID 0)
-- Dependencies: 218
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 3245 (class 2606 OID 16600)
-- Name: authors authors_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_name_key UNIQUE (name);


--
-- TOC entry 3247 (class 2606 OID 16598)
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- TOC entry 3249 (class 2606 OID 16611)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- TOC entry 3251 (class 2606 OID 16613)
-- Name: books books_title_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_title_key UNIQUE (title);


--
-- TOC entry 3255 (class 2606 OID 16645)
-- Name: borrowing_records borrowing_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.borrowing_records
    ADD CONSTRAINT borrowing_records_pkey PRIMARY KEY (id);


--
-- TOC entry 3253 (class 2606 OID 16635)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2022-12-13 16:00:44 WIB

--
-- PostgreSQL database dump complete
--

