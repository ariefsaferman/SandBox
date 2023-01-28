--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Homebrew)
-- Dumped by pg_dump version 15.0

-- Started on 2022-12-22 12:46:43 WIB

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

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: nata.nael
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO "nata.nael";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 17244)
-- Name: trainers; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.trainers (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.trainers OWNER TO admin;

--
-- TOC entry 210 (class 1259 OID 17251)
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    email text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- TOC entry 211 (class 1259 OID 17262)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3593 (class 0 OID 17244)
-- Dependencies: 209
-- Data for Name: trainers; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.trainers (id, name) FROM stdin;
2002	trainer 2
2003	trainer 3
\.


--
-- TOC entry 3594 (class 0 OID 17251)
-- Dependencies: 210
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.users (id, email, password) FROM stdin;
1	user@email.com	$2a$10$xlWjyFziWHGsn3CoTyLSveIAAf7q4ivGvLfAW.QjDp1Xf5AdnA9u2
\.


--
-- TOC entry 3602 (class 0 OID 0)
-- Dependencies: 211
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.users_id_seq', 6, true);


--
-- TOC entry 3449 (class 2606 OID 17250)
-- Name: trainers trainers_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.trainers
    ADD CONSTRAINT trainers_pkey PRIMARY KEY (id);


--
-- TOC entry 3451 (class 2606 OID 17261)
-- Name: users unique_email; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_email UNIQUE (email);


--
-- TOC entry 3453 (class 2606 OID 17257)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3601 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: nata.nael
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2022-12-22 12:46:43 WIB

--
-- PostgreSQL database dump complete
--

