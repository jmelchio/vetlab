--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3 (Debian 10.3-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: diagnostic_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.diagnostic_report (
    id integer NOT NULL,
    request_id integer NOT NULL,
    org_id integer NOT NULL,
    customer_id integer NOT NULL,
    user_id integer NOT NULL,
    date timestamp without time zone NOT NULL,
    report_body character varying,
    report_file character varying(100)
);


ALTER TABLE public.diagnostic_report OWNER TO postgres;

--
-- Name: TABLE diagnostic_report; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.diagnostic_report IS 'Table to store information on completed diagnostic reports in the vetlab system';


--
-- Name: diagnostic_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.diagnostic_report_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.diagnostic_report_id_seq OWNER TO postgres;

--
-- Name: diagnostic_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.diagnostic_report_id_seq OWNED BY public.diagnostic_report.id;


--
-- Name: diagnostic_request; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.diagnostic_request (
    id integer NOT NULL,
    org_id integer NOT NULL,
    customer_id integer NOT NULL,
    user_id integer NOT NULL,
    date timestamp without time zone NOT NULL,
    description character varying
);


ALTER TABLE public.diagnostic_request OWNER TO postgres;

--
-- Name: TABLE diagnostic_request; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.diagnostic_request IS 'Table to store information for requests for labwork in the vetlab system';


--
-- Name: diagnostic_request_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.diagnostic_request_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.diagnostic_request_id_seq OWNER TO postgres;

--
-- Name: diagnostic_request_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.diagnostic_request_id_seq OWNED BY public.diagnostic_request.id;


--
-- Name: vetlab_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vetlab_user (
    id integer NOT NULL,
    user_name character varying(20) NOT NULL,
    first_name character varying(25),
    last_name character varying(25),
    email character varying(40) NOT NULL,
    password_hash character varying(60) NOT NULL,
    org_id integer,
    admin_user boolean DEFAULT false NOT NULL
);


ALTER TABLE public.vetlab_user OWNER TO postgres;

--
-- Name: TABLE vetlab_user; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.vetlab_user IS 'stores the user in the vetlab system';


--
-- Name: user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_user_id_seq OWNER TO postgres;

--
-- Name: user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_user_id_seq OWNED BY public.vetlab_user.id;


--
-- Name: vet_org; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vet_org (
    id integer NOT NULL,
    org_name character varying(30) NOT NULL,
    street character varying(40) NOT NULL,
    house_number character varying(10) NOT NULL,
    city character varying(40) NOT NULL,
    province character varying(30) NOT NULL,
    country character varying(40) NOT NULL,
    postal_code character varying(10) NOT NULL,
    email character varying(40) NOT NULL,
    phone character varying(20) NOT NULL,
    fax character varying(20)
);


ALTER TABLE public.vet_org OWNER TO postgres;

--
-- Name: TABLE vet_org; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.vet_org IS 'Table to store veterinary practice information for the vetlab system';


--
-- Name: vet_org_org_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vet_org_org_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vet_org_org_id_seq OWNER TO postgres;

--
-- Name: vet_org_org_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vet_org_org_id_seq OWNED BY public.vet_org.id;


--
-- Name: diagnostic_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_report ALTER COLUMN id SET DEFAULT nextval('public.diagnostic_report_id_seq'::regclass);


--
-- Name: diagnostic_request id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_request ALTER COLUMN id SET DEFAULT nextval('public.diagnostic_request_id_seq'::regclass);


--
-- Name: vet_org id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vet_org ALTER COLUMN id SET DEFAULT nextval('public.vet_org_org_id_seq'::regclass);


--
-- Name: vetlab_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vetlab_user ALTER COLUMN id SET DEFAULT nextval('public.user_user_id_seq'::regclass);


--
-- Name: diagnostic_report diagnostic_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_report
    ADD CONSTRAINT diagnostic_report_pkey PRIMARY KEY (id);


--
-- Name: diagnostic_request diagnostic_request_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_request
    ADD CONSTRAINT diagnostic_request_pkey PRIMARY KEY (id);


--
-- Name: vetlab_user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vetlab_user
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: vet_org vet_org_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vet_org
    ADD CONSTRAINT vet_org_pkey PRIMARY KEY (id);


--
-- Name: diagnostic_report_org_id_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX diagnostic_report_org_id_index ON public.diagnostic_report USING btree (org_id DESC);


--
-- Name: user_user_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_user_name_uindex ON public.vetlab_user USING btree (user_name);


--
-- Name: diagnostic_report diagnostic_report_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_report
    ADD CONSTRAINT diagnostic_report_user_id_fk FOREIGN KEY (user_id) REFERENCES public.vetlab_user(id);


--
-- Name: diagnostic_report diagnostic_report_user_id_fk_2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_report
    ADD CONSTRAINT diagnostic_report_user_id_fk_2 FOREIGN KEY (customer_id) REFERENCES public.vetlab_user(id);


--
-- Name: diagnostic_report diagnostic_report_vet_org_org_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_report
    ADD CONSTRAINT diagnostic_report_vet_org_org_id_fk FOREIGN KEY (org_id) REFERENCES public.vet_org(id);


--
-- Name: diagnostic_request diagnostic_request_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_request
    ADD CONSTRAINT diagnostic_request_user_id_fk FOREIGN KEY (customer_id) REFERENCES public.vetlab_user(id);


--
-- Name: diagnostic_request diagnostic_request_user_id_fk_2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_request
    ADD CONSTRAINT diagnostic_request_user_id_fk_2 FOREIGN KEY (user_id) REFERENCES public.vetlab_user(id);


--
-- Name: diagnostic_request diagnostic_request_vet_org_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.diagnostic_request
    ADD CONSTRAINT diagnostic_request_vet_org_id_fk FOREIGN KEY (org_id) REFERENCES public.vet_org(id);


--
-- Name: vetlab_user user_vet_org_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vetlab_user
    ADD CONSTRAINT user_vet_org_id_fk FOREIGN KEY (org_id) REFERENCES public.vet_org(id);


--
-- PostgreSQL database dump complete
--

