--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Debian 14.1-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-07-15 10:15:12

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
-- TOC entry 3385 (class 1262 OID 16641)
-- Name: e_user; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE e_user WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


\connect e_user

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
-- TOC entry 228 (class 1259 OID 25380)
-- Name: resource_resources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.resource_resources (
    id bigint NOT NULL,
    resource_item_id bigint,
    key text,
    type integer,
    index boolean,
    path text,
    name text,
    menu boolean,
    icon text,
    "desc" text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    sort bigint
);


--
-- TOC entry 227 (class 1259 OID 25379)
-- Name: resources_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.resources_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3386 (class 0 OID 0)
-- Dependencies: 227
-- Name: resources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.resources_id_seq OWNED BY public.resource_resources.id;


--
-- TOC entry 3230 (class 2604 OID 25383)
-- Name: resource_resources id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resource_resources ALTER COLUMN id SET DEFAULT nextval('public.resources_id_seq'::regclass);


--
-- TOC entry 3379 (class 0 OID 25380)
-- Dependencies: 228
-- Data for Name: resource_resources; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (7, NULL, '', 2, false, '', '空资讯', false, '', '', '2022-02-22 10:03:10.542+00', '2022-02-22 10:03:10.542+00', NULL, NULL);
INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (18, 7, 'E_HOME', 0, false, '/', '首页', true, 'home', '', '2022-04-22 14:26:19.424687+00', '2022-05-09 13:55:05.113448+00', NULL, NULL);
INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (24, 19, '1111', 0, false, '', '', false, '', '', '2022-05-09 14:02:29.605442+00', '2022-05-09 14:02:29.605442+00', '2022-05-09 14:02:37.103792+00', NULL);
INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (11, 7, 'E_SYSTEM', 2, true, '/system', '系统', true, 'system', '', '2022-02-22 10:17:50.201835+00', '2022-06-13 13:34:14.187887+00', NULL, NULL);
INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (19, 11, 'SYSTEM_RESOURCE', 0, false, '/system/resource', '资源', true, 'resource', '', '2022-04-22 14:45:25.271679+00', '2022-06-13 13:34:53.248392+00', NULL, NULL);
INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (22, 11, 'SYSTEM_ROLE', 0, false, '/system/role', '角色', true, 'role', '', '2022-04-25 09:28:25.759163+00', '2022-06-13 13:35:13.034152+00', NULL, NULL);
INSERT INTO public.resource_resources (id, resource_item_id, key, type, index, path, name, menu, icon, "desc", created_at, updated_at, deleted_at, sort) VALUES (23, 11, 'SYSTEM_ACCOUNT', 0, false, '/system/account', '账号', true, 'account', '', '2022-05-09 14:01:59.767487+00', '2022-06-13 13:35:25.88377+00', NULL, NULL);


--
-- TOC entry 3387 (class 0 OID 0)
-- Dependencies: 227
-- Name: resources_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.resources_id_seq', 1, false);


--
-- TOC entry 3234 (class 2606 OID 25389)
-- Name: resource_resources resources_key_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resource_resources
    ADD CONSTRAINT resources_key_key UNIQUE (key);


--
-- TOC entry 3236 (class 2606 OID 25387)
-- Name: resource_resources resources_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resource_resources
    ADD CONSTRAINT resources_pkey PRIMARY KEY (id);


--
-- TOC entry 3231 (class 1259 OID 25523)
-- Name: idx_resource_resources_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_resource_resources_deleted_at ON public.resource_resources USING btree (deleted_at);


--
-- TOC entry 3232 (class 1259 OID 25395)
-- Name: idx_resources_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_resources_deleted_at ON public.resource_resources USING btree (deleted_at);


--
-- TOC entry 3238 (class 2606 OID 25768)
-- Name: resource_resources fk_resource_resources_children; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resource_resources
    ADD CONSTRAINT fk_resource_resources_children FOREIGN KEY (resource_item_id) REFERENCES public.resource_resources(id);


--
-- TOC entry 3237 (class 2606 OID 25763)
-- Name: resource_resources fk_resources_children; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resource_resources
    ADD CONSTRAINT fk_resources_children FOREIGN KEY (resource_item_id) REFERENCES public.resource_resources(id);


-- Completed on 2022-07-15 10:15:15

--
-- PostgreSQL database dump complete
--

