--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-03-05 22:20:47

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
-- TOC entry 202 (class 1259 OID 16414)
-- Name: Parking; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Parking" (
    "CarId" text,
    "OwnerName" text,
    "PakingId" integer NOT NULL
);


ALTER TABLE public."Parking" OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16397)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    userid integer NOT NULL,
    name text,
    age integer,
    location text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 16395)
-- Name: users_userid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_userid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_userid_seq OWNER TO postgres;

--
-- TOC entry 2999 (class 0 OID 0)
-- Dependencies: 200
-- Name: users_userid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_userid_seq OWNED BY public.users.userid;


--
-- TOC entry 2856 (class 2604 OID 16400)
-- Name: users userid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN userid SET DEFAULT nextval('public.users_userid_seq'::regclass);


--
-- TOC entry 2993 (class 0 OID 16414)
-- Dependencies: 202
-- Data for Name: Parking; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Parking" ("CarId", "OwnerName", "PakingId") FROM stdin;
		4
		5
		3
		8
		9
		10
		1
		6
		7
		2
\.


--
-- TOC entry 2992 (class 0 OID 16397)
-- Dependencies: 201
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (userid, name, age, location) FROM stdin;
1	gopher	25	India
\.


--
-- TOC entry 3000 (class 0 OID 0)
-- Dependencies: 200
-- Name: users_userid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_userid_seq', 1, true);


--
-- TOC entry 2860 (class 2606 OID 16421)
-- Name: Parking Parking_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Parking"
    ADD CONSTRAINT "Parking_pkey" PRIMARY KEY ("PakingId");


--
-- TOC entry 2858 (class 2606 OID 16405)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (userid);


-- Completed on 2021-03-05 22:20:48

--
-- PostgreSQL database dump complete
--

