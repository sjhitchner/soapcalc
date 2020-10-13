--
-- PostgreSQL database dump
--

-- Dumped from database version 10.6 (Debian 10.6-1.pgdg90+1)
-- Dumped by pg_dump version 12.4

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

--
-- Name: additive; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.additive (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    note text NOT NULL
);


ALTER TABLE public.additive OWNER TO soap;

--
-- Name: additive_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.additive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.additive_id_seq OWNER TO soap;

--
-- Name: additive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.additive_id_seq OWNED BY public.additive.id;


--
-- Name: additive_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.additive_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    additive_id integer NOT NULL,
    supplier_id integer NOT NULL
);


ALTER TABLE public.additive_inventory OWNER TO soap;

--
-- Name: additive_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.additive_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.additive_inventory_id_seq OWNER TO soap;

--
-- Name: additive_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.additive_inventory_id_seq OWNED BY public.additive_inventory.id;


--
-- Name: auth_group; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_group (
    id integer NOT NULL,
    name character varying(150) NOT NULL
);


ALTER TABLE public.auth_group OWNER TO soap;

--
-- Name: auth_group_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_group_id_seq OWNER TO soap;

--
-- Name: auth_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_group_id_seq OWNED BY public.auth_group.id;


--
-- Name: auth_group_permissions; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_group_permissions (
    id integer NOT NULL,
    group_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE public.auth_group_permissions OWNER TO soap;

--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_group_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_group_permissions_id_seq OWNER TO soap;

--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_group_permissions_id_seq OWNED BY public.auth_group_permissions.id;


--
-- Name: auth_permission; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_permission (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    content_type_id integer NOT NULL,
    codename character varying(100) NOT NULL
);


ALTER TABLE public.auth_permission OWNER TO soap;

--
-- Name: auth_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_permission_id_seq OWNER TO soap;

--
-- Name: auth_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_permission_id_seq OWNED BY public.auth_permission.id;


--
-- Name: auth_user; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_user (
    id integer NOT NULL,
    password character varying(128) NOT NULL,
    last_login timestamp with time zone,
    is_superuser boolean NOT NULL,
    username character varying(150) NOT NULL,
    first_name character varying(150) NOT NULL,
    last_name character varying(150) NOT NULL,
    email character varying(254) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    date_joined timestamp with time zone NOT NULL
);


ALTER TABLE public.auth_user OWNER TO soap;

--
-- Name: auth_user_groups; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_user_groups (
    id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer NOT NULL
);


ALTER TABLE public.auth_user_groups OWNER TO soap;

--
-- Name: auth_user_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_user_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_user_groups_id_seq OWNER TO soap;

--
-- Name: auth_user_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_user_groups_id_seq OWNED BY public.auth_user_groups.id;


--
-- Name: auth_user_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_user_id_seq OWNER TO soap;

--
-- Name: auth_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_user_id_seq OWNED BY public.auth_user.id;


--
-- Name: auth_user_user_permissions; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_user_user_permissions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE public.auth_user_user_permissions OWNER TO soap;

--
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_user_user_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_user_user_permissions_id_seq OWNER TO soap;

--
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_user_user_permissions_id_seq OWNED BY public.auth_user_user_permissions.id;


--
-- Name: django_admin_log; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_admin_log (
    id integer NOT NULL,
    action_time timestamp with time zone NOT NULL,
    object_id text,
    object_repr character varying(200) NOT NULL,
    action_flag smallint NOT NULL,
    change_message text NOT NULL,
    content_type_id integer,
    user_id integer NOT NULL,
    CONSTRAINT django_admin_log_action_flag_check CHECK ((action_flag >= 0))
);


ALTER TABLE public.django_admin_log OWNER TO soap;

--
-- Name: django_admin_log_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.django_admin_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.django_admin_log_id_seq OWNER TO soap;

--
-- Name: django_admin_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.django_admin_log_id_seq OWNED BY public.django_admin_log.id;


--
-- Name: django_content_type; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_content_type (
    id integer NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL
);


ALTER TABLE public.django_content_type OWNER TO soap;

--
-- Name: django_content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.django_content_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.django_content_type_id_seq OWNER TO soap;

--
-- Name: django_content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.django_content_type_id_seq OWNED BY public.django_content_type.id;


--
-- Name: django_migrations; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_migrations (
    id integer NOT NULL,
    app character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    applied timestamp with time zone NOT NULL
);


ALTER TABLE public.django_migrations OWNER TO soap;

--
-- Name: django_migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.django_migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.django_migrations_id_seq OWNER TO soap;

--
-- Name: django_migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.django_migrations_id_seq OWNED BY public.django_migrations.id;


--
-- Name: django_session; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_session (
    session_key character varying(40) NOT NULL,
    session_data text NOT NULL,
    expire_date timestamp with time zone NOT NULL
);


ALTER TABLE public.django_session OWNER TO soap;

--
-- Name: fragrance; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.fragrance (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    note text NOT NULL
);


ALTER TABLE public.fragrance OWNER TO soap;

--
-- Name: fragrance_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.fragrance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fragrance_id_seq OWNER TO soap;

--
-- Name: fragrance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.fragrance_id_seq OWNED BY public.fragrance.id;


--
-- Name: fragrance_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.fragrance_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    fragrance_id integer NOT NULL,
    supplier_id integer NOT NULL
);


ALTER TABLE public.fragrance_inventory OWNER TO soap;

--
-- Name: fragrance_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.fragrance_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fragrance_inventory_id_seq OWNER TO soap;

--
-- Name: fragrance_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.fragrance_inventory_id_seq OWNED BY public.fragrance_inventory.id;


--
-- Name: lipid; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lipid (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    lauric integer NOT NULL,
    myristic integer NOT NULL,
    palmitic integer NOT NULL,
    stearic integer NOT NULL,
    ricinoleic integer NOT NULL,
    oleic integer NOT NULL,
    linoleic integer NOT NULL,
    linolenic integer NOT NULL,
    hardness integer NOT NULL,
    cleansing integer NOT NULL,
    conditioning integer NOT NULL,
    bubbly integer NOT NULL,
    creamy integer NOT NULL,
    iodine integer NOT NULL,
    ins integer NOT NULL,
    inci_name character varying(100) NOT NULL,
    family character varying(50) NOT NULL,
    naoh double precision NOT NULL
);


ALTER TABLE public.lipid OWNER TO soap;

--
-- Name: lipid_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lipid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lipid_id_seq OWNER TO soap;

--
-- Name: lipid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lipid_id_seq OWNED BY public.lipid.id;


--
-- Name: lipid_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lipid_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    sap double precision NOT NULL,
    naoh double precision NOT NULL,
    koh double precision NOT NULL,
    grams_per_liter double precision NOT NULL,
    lipid_id integer NOT NULL,
    supplier_id integer NOT NULL
);


ALTER TABLE public.lipid_inventory OWNER TO soap;

--
-- Name: lipid_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lipid_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lipid_inventory_id_seq OWNER TO soap;

--
-- Name: lipid_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lipid_inventory_id_seq OWNED BY public.lipid_inventory.id;


--
-- Name: lye; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lye (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    kind character varying(4) NOT NULL,
    name character varying(100) NOT NULL,
    note text NOT NULL
);


ALTER TABLE public.lye OWNER TO soap;

--
-- Name: lye_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lye_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lye_id_seq OWNER TO soap;

--
-- Name: lye_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lye_id_seq OWNED BY public.lye.id;


--
-- Name: lye_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lye_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    concentration double precision NOT NULL,
    lye_id integer NOT NULL,
    supplier_id integer NOT NULL
);


ALTER TABLE public.lye_inventory OWNER TO soap;

--
-- Name: lye_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lye_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lye_inventory_id_seq OWNER TO soap;

--
-- Name: lye_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lye_inventory_id_seq OWNED BY public.lye_inventory.id;


--
-- Name: recipe; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    note text NOT NULL
);


ALTER TABLE public.recipe OWNER TO soap;

--
-- Name: recipe_additive; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_additive (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    percentage double precision NOT NULL,
    additive_id integer NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.recipe_additive OWNER TO soap;

--
-- Name: recipe_additive_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_additive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_additive_id_seq OWNER TO soap;

--
-- Name: recipe_additive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_additive_id_seq OWNED BY public.recipe_additive.id;


--
-- Name: recipe_batch; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    tag character varying(16) NOT NULL,
    production_date timestamp with time zone NOT NULL,
    sellable_date timestamp with time zone NOT NULL,
    note text NOT NULL,
    lipid_weight double precision NOT NULL,
    production_weight double precision NOT NULL,
    cured_weight double precision NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.recipe_batch OWNER TO soap;

--
-- Name: recipe_batch_additive; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_additive (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    cost double precision NOT NULL,
    additive_id integer NOT NULL,
    batch_id integer NOT NULL
);


ALTER TABLE public.recipe_batch_additive OWNER TO soap;

--
-- Name: recipe_batch_additive_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_additive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_batch_additive_id_seq OWNER TO soap;

--
-- Name: recipe_batch_additive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_additive_id_seq OWNED BY public.recipe_batch_additive.id;


--
-- Name: recipe_batch_fragrance; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_fragrance (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    cost double precision NOT NULL,
    fragrance_id integer NOT NULL,
    batch_id integer NOT NULL
);


ALTER TABLE public.recipe_batch_fragrance OWNER TO soap;

--
-- Name: recipe_batch_fragrance_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_fragrance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_batch_fragrance_id_seq OWNER TO soap;

--
-- Name: recipe_batch_fragrance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_fragrance_id_seq OWNED BY public.recipe_batch_fragrance.id;


--
-- Name: recipe_batch_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_batch_id_seq OWNER TO soap;

--
-- Name: recipe_batch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_id_seq OWNED BY public.recipe_batch.id;


--
-- Name: recipe_batch_lipid; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_lipid (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    cost double precision NOT NULL,
    lipid_id integer NOT NULL,
    batch_id integer NOT NULL
);


ALTER TABLE public.recipe_batch_lipid OWNER TO soap;

--
-- Name: recipe_batch_lipid_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_lipid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_batch_lipid_id_seq OWNER TO soap;

--
-- Name: recipe_batch_lipid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_lipid_id_seq OWNED BY public.recipe_batch_lipid.id;


--
-- Name: recipe_batch_lye; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_lye (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    discount double precision NOT NULL,
    cost double precision NOT NULL,
    lye_id integer NOT NULL,
    batch_id integer NOT NULL
);


ALTER TABLE public.recipe_batch_lye OWNER TO soap;

--
-- Name: recipe_batch_lye_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_lye_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_batch_lye_id_seq OWNER TO soap;

--
-- Name: recipe_batch_lye_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_lye_id_seq OWNED BY public.recipe_batch_lye.id;


--
-- Name: recipe_batch_note; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_note (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    note text NOT NULL,
    link character varying(255) NOT NULL,
    batch_id integer NOT NULL
);


ALTER TABLE public.recipe_batch_note OWNER TO soap;

--
-- Name: recipe_batch_note_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_note_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_batch_note_id_seq OWNER TO soap;

--
-- Name: recipe_batch_note_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_note_id_seq OWNED BY public.recipe_batch_note.id;


--
-- Name: recipe_fragrance; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_fragrance (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    percentage double precision NOT NULL,
    fragrance_id integer NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.recipe_fragrance OWNER TO soap;

--
-- Name: recipe_fragrance_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_fragrance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_fragrance_id_seq OWNER TO soap;

--
-- Name: recipe_fragrance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_fragrance_id_seq OWNED BY public.recipe_fragrance.id;


--
-- Name: recipe_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_id_seq OWNER TO soap;

--
-- Name: recipe_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_id_seq OWNED BY public.recipe.id;


--
-- Name: recipe_lipid; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_lipid (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    percentage double precision NOT NULL,
    lipid_id integer NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.recipe_lipid OWNER TO soap;

--
-- Name: recipe_lipid_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_lipid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_lipid_id_seq OWNER TO soap;

--
-- Name: recipe_lipid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_lipid_id_seq OWNED BY public.recipe_lipid.id;


--
-- Name: recipe_step; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_step (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    num integer NOT NULL,
    note text NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.recipe_step OWNER TO soap;

--
-- Name: recipe_step_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_step_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_step_id_seq OWNER TO soap;

--
-- Name: recipe_step_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_step_id_seq OWNED BY public.recipe_step.id;


--
-- Name: supplier; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.supplier (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    website character varying(255) NOT NULL,
    note text NOT NULL
);


ALTER TABLE public.supplier OWNER TO soap;

--
-- Name: supplier_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.supplier_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.supplier_id_seq OWNER TO soap;

--
-- Name: supplier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.supplier_id_seq OWNED BY public.supplier.id;


--
-- Name: additive id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive ALTER COLUMN id SET DEFAULT nextval('public.additive_id_seq'::regclass);


--
-- Name: additive_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory ALTER COLUMN id SET DEFAULT nextval('public.additive_inventory_id_seq'::regclass);


--
-- Name: auth_group id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group ALTER COLUMN id SET DEFAULT nextval('public.auth_group_id_seq'::regclass);


--
-- Name: auth_group_permissions id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_group_permissions_id_seq'::regclass);


--
-- Name: auth_permission id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission ALTER COLUMN id SET DEFAULT nextval('public.auth_permission_id_seq'::regclass);


--
-- Name: auth_user id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user ALTER COLUMN id SET DEFAULT nextval('public.auth_user_id_seq'::regclass);


--
-- Name: auth_user_groups id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups ALTER COLUMN id SET DEFAULT nextval('public.auth_user_groups_id_seq'::regclass);


--
-- Name: auth_user_user_permissions id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_user_user_permissions_id_seq'::regclass);


--
-- Name: django_admin_log id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log ALTER COLUMN id SET DEFAULT nextval('public.django_admin_log_id_seq'::regclass);


--
-- Name: django_content_type id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);


--
-- Name: django_migrations id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_migrations ALTER COLUMN id SET DEFAULT nextval('public.django_migrations_id_seq'::regclass);


--
-- Name: fragrance id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance ALTER COLUMN id SET DEFAULT nextval('public.fragrance_id_seq'::regclass);


--
-- Name: fragrance_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory ALTER COLUMN id SET DEFAULT nextval('public.fragrance_inventory_id_seq'::regclass);


--
-- Name: lipid id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid ALTER COLUMN id SET DEFAULT nextval('public.lipid_id_seq'::regclass);


--
-- Name: lipid_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory ALTER COLUMN id SET DEFAULT nextval('public.lipid_inventory_id_seq'::regclass);


--
-- Name: lye id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye ALTER COLUMN id SET DEFAULT nextval('public.lye_id_seq'::regclass);


--
-- Name: lye_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory ALTER COLUMN id SET DEFAULT nextval('public.lye_inventory_id_seq'::regclass);


--
-- Name: recipe id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe ALTER COLUMN id SET DEFAULT nextval('public.recipe_id_seq'::regclass);


--
-- Name: recipe_additive id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive ALTER COLUMN id SET DEFAULT nextval('public.recipe_additive_id_seq'::regclass);


--
-- Name: recipe_batch id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_id_seq'::regclass);


--
-- Name: recipe_batch_additive id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_additive_id_seq'::regclass);


--
-- Name: recipe_batch_fragrance id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_fragrance_id_seq'::regclass);


--
-- Name: recipe_batch_lipid id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_lipid_id_seq'::regclass);


--
-- Name: recipe_batch_lye id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_lye_id_seq'::regclass);


--
-- Name: recipe_batch_note id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_note ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_note_id_seq'::regclass);


--
-- Name: recipe_fragrance id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance ALTER COLUMN id SET DEFAULT nextval('public.recipe_fragrance_id_seq'::regclass);


--
-- Name: recipe_lipid id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid ALTER COLUMN id SET DEFAULT nextval('public.recipe_lipid_id_seq'::regclass);


--
-- Name: recipe_step id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_step ALTER COLUMN id SET DEFAULT nextval('public.recipe_step_id_seq'::regclass);


--
-- Name: supplier id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.supplier ALTER COLUMN id SET DEFAULT nextval('public.supplier_id_seq'::regclass);


--
-- Data for Name: additive; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.additive (id, created_at, updated_at, deleted_at, name, note) FROM stdin;
\.


--
-- Data for Name: additive_inventory; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.additive_inventory (id, created_at, updated_at, deleted_at, purchase_date, expiry_date, cost, weight, additive_id, supplier_id) FROM stdin;
\.


--
-- Data for Name: auth_group; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.auth_group (id, name) FROM stdin;
\.


--
-- Data for Name: auth_group_permissions; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.auth_group_permissions (id, group_id, permission_id) FROM stdin;
\.


--
-- Data for Name: auth_permission; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.auth_permission (id, name, content_type_id, codename) FROM stdin;
1	Can add log entry	1	add_logentry
2	Can change log entry	1	change_logentry
3	Can delete log entry	1	delete_logentry
4	Can view log entry	1	view_logentry
5	Can add permission	2	add_permission
6	Can change permission	2	change_permission
7	Can delete permission	2	delete_permission
8	Can view permission	2	view_permission
9	Can add group	3	add_group
10	Can change group	3	change_group
11	Can delete group	3	delete_group
12	Can view group	3	view_group
13	Can add user	4	add_user
14	Can change user	4	change_user
15	Can delete user	4	delete_user
16	Can view user	4	view_user
17	Can add content type	5	add_contenttype
18	Can change content type	5	change_contenttype
19	Can delete content type	5	delete_contenttype
20	Can view content type	5	view_contenttype
21	Can add session	6	add_session
22	Can change session	6	change_session
23	Can delete session	6	delete_session
24	Can view session	6	view_session
25	Can add additive	7	add_additive
26	Can change additive	7	change_additive
27	Can delete additive	7	delete_additive
28	Can view additive	7	view_additive
29	Can add fragrance	8	add_fragrance
30	Can change fragrance	8	change_fragrance
31	Can delete fragrance	8	delete_fragrance
32	Can view fragrance	8	view_fragrance
33	Can add lipid	9	add_lipid
34	Can change lipid	9	change_lipid
35	Can delete lipid	9	delete_lipid
36	Can view lipid	9	view_lipid
37	Can add lye	10	add_lye
38	Can change lye	10	change_lye
39	Can delete lye	10	delete_lye
40	Can view lye	10	view_lye
41	Can add recipe	11	add_recipe
42	Can change recipe	11	change_recipe
43	Can delete recipe	11	delete_recipe
44	Can view recipe	11	view_recipe
45	Can add supplier	12	add_supplier
46	Can change supplier	12	change_supplier
47	Can delete supplier	12	delete_supplier
48	Can view supplier	12	view_supplier
49	Can add recipe lipid	13	add_recipelipid
50	Can change recipe lipid	13	change_recipelipid
51	Can delete recipe lipid	13	delete_recipelipid
52	Can view recipe lipid	13	view_recipelipid
53	Can add recipe fragrance	14	add_recipefragrance
54	Can change recipe fragrance	14	change_recipefragrance
55	Can delete recipe fragrance	14	delete_recipefragrance
56	Can view recipe fragrance	14	view_recipefragrance
57	Can add recipe batch lye	15	add_recipebatchlye
58	Can change recipe batch lye	15	change_recipebatchlye
59	Can delete recipe batch lye	15	delete_recipebatchlye
60	Can view recipe batch lye	15	view_recipebatchlye
61	Can add recipe batch lipid	16	add_recipebatchlipid
62	Can change recipe batch lipid	16	change_recipebatchlipid
63	Can delete recipe batch lipid	16	delete_recipebatchlipid
64	Can view recipe batch lipid	16	view_recipebatchlipid
65	Can add recipe batch fragrance	17	add_recipebatchfragrance
66	Can change recipe batch fragrance	17	change_recipebatchfragrance
67	Can delete recipe batch fragrance	17	delete_recipebatchfragrance
68	Can view recipe batch fragrance	17	view_recipebatchfragrance
69	Can add recipe batch additive	18	add_recipebatchadditive
70	Can change recipe batch additive	18	change_recipebatchadditive
71	Can delete recipe batch additive	18	delete_recipebatchadditive
72	Can view recipe batch additive	18	view_recipebatchadditive
73	Can add recipe batch	19	add_recipebatch
74	Can change recipe batch	19	change_recipebatch
75	Can delete recipe batch	19	delete_recipebatch
76	Can view recipe batch	19	view_recipebatch
77	Can add recipe additive	20	add_recipeadditive
78	Can change recipe additive	20	change_recipeadditive
79	Can delete recipe additive	20	delete_recipeadditive
80	Can view recipe additive	20	view_recipeadditive
81	Can add lye inventory	21	add_lyeinventory
82	Can change lye inventory	21	change_lyeinventory
83	Can delete lye inventory	21	delete_lyeinventory
84	Can view lye inventory	21	view_lyeinventory
85	Can add lipid inventory	22	add_lipidinventory
86	Can change lipid inventory	22	change_lipidinventory
87	Can delete lipid inventory	22	delete_lipidinventory
88	Can view lipid inventory	22	view_lipidinventory
89	Can add fragrance inventory	23	add_fragranceinventory
90	Can change fragrance inventory	23	change_fragranceinventory
91	Can delete fragrance inventory	23	delete_fragranceinventory
92	Can view fragrance inventory	23	view_fragranceinventory
93	Can add additive inventory	24	add_additiveinventory
94	Can change additive inventory	24	change_additiveinventory
95	Can delete additive inventory	24	delete_additiveinventory
96	Can view additive inventory	24	view_additiveinventory
97	Can add recipe batch note	25	add_recipebatchnote
98	Can change recipe batch note	25	change_recipebatchnote
99	Can delete recipe batch note	25	delete_recipebatchnote
100	Can view recipe batch note	25	view_recipebatchnote
101	Can add recipe step	26	add_recipestep
102	Can change recipe step	26	change_recipestep
103	Can delete recipe step	26	delete_recipestep
104	Can view recipe step	26	view_recipestep
\.


--
-- Data for Name: auth_user; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.auth_user (id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined) FROM stdin;
\.


--
-- Data for Name: auth_user_groups; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.auth_user_groups (id, user_id, group_id) FROM stdin;
\.


--
-- Data for Name: auth_user_user_permissions; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.auth_user_user_permissions (id, user_id, permission_id) FROM stdin;
\.


--
-- Data for Name: django_admin_log; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.django_admin_log (id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id) FROM stdin;
\.


--
-- Data for Name: django_content_type; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.django_content_type (id, app_label, model) FROM stdin;
1	admin	logentry
2	auth	permission
3	auth	group
4	auth	user
5	contenttypes	contenttype
6	sessions	session
7	soap	additive
8	soap	fragrance
9	soap	lipid
10	soap	lye
11	soap	recipe
12	soap	supplier
13	soap	recipelipid
14	soap	recipefragrance
15	soap	recipebatchlye
16	soap	recipebatchlipid
17	soap	recipebatchfragrance
18	soap	recipebatchadditive
19	soap	recipebatch
20	soap	recipeadditive
21	soap	lyeinventory
22	soap	lipidinventory
23	soap	fragranceinventory
24	soap	additiveinventory
25	soap	recipebatchnote
26	soap	recipestep
\.


--
-- Data for Name: django_migrations; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.django_migrations (id, app, name, applied) FROM stdin;
1	contenttypes	0001_initial	2020-10-10 20:29:04.607923+00
2	auth	0001_initial	2020-10-10 20:29:04.658612+00
3	admin	0001_initial	2020-10-10 20:29:04.733599+00
4	admin	0002_logentry_remove_auto_add	2020-10-10 20:29:04.759445+00
5	admin	0003_logentry_add_action_flag_choices	2020-10-10 20:29:04.774969+00
6	contenttypes	0002_remove_content_type_name	2020-10-10 20:29:04.800332+00
7	auth	0002_alter_permission_name_max_length	2020-10-10 20:29:04.814947+00
8	auth	0003_alter_user_email_max_length	2020-10-10 20:29:04.831717+00
9	auth	0004_alter_user_username_opts	2020-10-10 20:29:04.893625+00
10	auth	0005_alter_user_last_login_null	2020-10-10 20:29:04.910987+00
11	auth	0006_require_contenttypes_0002	2020-10-10 20:29:04.916982+00
12	auth	0007_alter_validators_add_error_messages	2020-10-10 20:29:04.931065+00
13	auth	0008_alter_user_username_max_length	2020-10-10 20:29:04.954527+00
14	auth	0009_alter_user_last_name_max_length	2020-10-10 20:29:04.970573+00
15	auth	0010_alter_group_name_max_length	2020-10-10 20:29:04.987518+00
16	auth	0011_update_proxy_permissions	2020-10-10 20:29:05.001512+00
17	auth	0012_alter_user_first_name_max_length	2020-10-10 20:29:05.020011+00
18	sessions	0001_initial	2020-10-10 20:29:05.035901+00
19	soap	0001_initial	2020-10-10 20:29:05.259783+00
20	soap	0002_auto_20201010_2118	2020-10-10 21:18:04.8612+00
21	soap	0003_auto_20201010_2304	2020-10-10 23:04:11.895255+00
22	soap	0004_auto_20201011_0529	2020-10-11 05:29:38.686915+00
23	soap	0005_auto_20201011_1931	2020-10-11 19:31:38.285098+00
24	soap	0006_auto_20201011_2001	2020-10-11 20:01:13.596017+00
25	soap	0007_auto_20201012_1859	2020-10-12 18:59:35.285019+00
26	soap	0008_auto_20201012_2139	2020-10-12 21:39:15.535626+00
27	soap	0009_recipebatchnote_recipestep	2020-10-12 23:05:03.074883+00
\.


--
-- Data for Name: django_session; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.django_session (session_key, session_data, expire_date) FROM stdin;
\.


--
-- Data for Name: fragrance; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.fragrance (id, created_at, updated_at, deleted_at, name, note) FROM stdin;
\.


--
-- Data for Name: fragrance_inventory; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.fragrance_inventory (id, created_at, updated_at, deleted_at, purchase_date, expiry_date, cost, weight, fragrance_id, supplier_id) FROM stdin;
\.


--
-- Data for Name: lipid; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.lipid (id, created_at, updated_at, deleted_at, name, lauric, myristic, palmitic, stearic, ricinoleic, oleic, linoleic, linolenic, hardness, cleansing, conditioning, bubbly, creamy, iodine, ins, inci_name, family, naoh) FROM stdin;
\.


--
-- Data for Name: lipid_inventory; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.lipid_inventory (id, created_at, updated_at, deleted_at, purchase_date, expiry_date, cost, weight, sap, naoh, koh, grams_per_liter, lipid_id, supplier_id) FROM stdin;
\.


--
-- Data for Name: lye; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.lye (id, created_at, updated_at, deleted_at, kind, name, note) FROM stdin;
\.


--
-- Data for Name: lye_inventory; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.lye_inventory (id, created_at, updated_at, deleted_at, purchase_date, expiry_date, cost, weight, concentration, lye_id, supplier_id) FROM stdin;
\.


--
-- Data for Name: recipe; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe (id, created_at, updated_at, deleted_at, name, note) FROM stdin;
\.


--
-- Data for Name: recipe_additive; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_additive (id, created_at, updated_at, deleted_at, percentage, additive_id, recipe_id) FROM stdin;
\.


--
-- Data for Name: recipe_batch; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_batch (id, created_at, updated_at, deleted_at, tag, production_date, sellable_date, note, lipid_weight, production_weight, cured_weight, recipe_id) FROM stdin;
\.


--
-- Data for Name: recipe_batch_additive; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_batch_additive (id, created_at, updated_at, deleted_at, weight, cost, additive_id, batch_id) FROM stdin;
\.


--
-- Data for Name: recipe_batch_fragrance; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_batch_fragrance (id, created_at, updated_at, deleted_at, weight, cost, fragrance_id, batch_id) FROM stdin;
\.


--
-- Data for Name: recipe_batch_lipid; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_batch_lipid (id, created_at, updated_at, deleted_at, weight, cost, lipid_id, batch_id) FROM stdin;
\.


--
-- Data for Name: recipe_batch_lye; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_batch_lye (id, created_at, updated_at, deleted_at, weight, discount, cost, lye_id, batch_id) FROM stdin;
\.


--
-- Data for Name: recipe_batch_note; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_batch_note (id, created_at, updated_at, deleted_at, note, link, batch_id) FROM stdin;
\.


--
-- Data for Name: recipe_fragrance; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_fragrance (id, created_at, updated_at, deleted_at, percentage, fragrance_id, recipe_id) FROM stdin;
\.


--
-- Data for Name: recipe_lipid; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_lipid (id, created_at, updated_at, deleted_at, percentage, lipid_id, recipe_id) FROM stdin;
\.


--
-- Data for Name: recipe_step; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.recipe_step (id, created_at, updated_at, deleted_at, num, note, recipe_id) FROM stdin;
\.


--
-- Data for Name: supplier; Type: TABLE DATA; Schema: public; Owner: soap
--

COPY public.supplier (id, created_at, updated_at, deleted_at, name, website, note) FROM stdin;
\.


--
-- Name: additive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.additive_id_seq', 1, false);


--
-- Name: additive_inventory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.additive_inventory_id_seq', 1, false);


--
-- Name: auth_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.auth_group_id_seq', 1, false);


--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.auth_group_permissions_id_seq', 1, false);


--
-- Name: auth_permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.auth_permission_id_seq', 104, true);


--
-- Name: auth_user_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.auth_user_groups_id_seq', 1, false);


--
-- Name: auth_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.auth_user_id_seq', 1, false);


--
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.auth_user_user_permissions_id_seq', 1, false);


--
-- Name: django_admin_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.django_admin_log_id_seq', 1, false);


--
-- Name: django_content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.django_content_type_id_seq', 26, true);


--
-- Name: django_migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.django_migrations_id_seq', 27, true);


--
-- Name: fragrance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.fragrance_id_seq', 1, false);


--
-- Name: fragrance_inventory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.fragrance_inventory_id_seq', 1, false);


--
-- Name: lipid_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.lipid_id_seq', 1, false);


--
-- Name: lipid_inventory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.lipid_inventory_id_seq', 1, false);


--
-- Name: lye_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.lye_id_seq', 1, false);


--
-- Name: lye_inventory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.lye_inventory_id_seq', 1, false);


--
-- Name: recipe_additive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_additive_id_seq', 1, false);


--
-- Name: recipe_batch_additive_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_batch_additive_id_seq', 1, false);


--
-- Name: recipe_batch_fragrance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_batch_fragrance_id_seq', 1, false);


--
-- Name: recipe_batch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_batch_id_seq', 1, false);


--
-- Name: recipe_batch_lipid_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_batch_lipid_id_seq', 1, false);


--
-- Name: recipe_batch_lye_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_batch_lye_id_seq', 1, false);


--
-- Name: recipe_batch_note_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_batch_note_id_seq', 1, false);


--
-- Name: recipe_fragrance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_fragrance_id_seq', 1, false);


--
-- Name: recipe_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_id_seq', 1, false);


--
-- Name: recipe_lipid_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_lipid_id_seq', 1, false);


--
-- Name: recipe_step_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.recipe_step_id_seq', 1, false);


--
-- Name: supplier_id_seq; Type: SEQUENCE SET; Schema: public; Owner: soap
--

SELECT pg_catalog.setval('public.supplier_id_seq', 1, false);


--
-- Name: additive_inventory additive_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory
    ADD CONSTRAINT additive_inventory_pkey PRIMARY KEY (id);


--
-- Name: additive additive_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive
    ADD CONSTRAINT additive_pkey PRIMARY KEY (id);


--
-- Name: auth_group auth_group_name_key; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_name_key UNIQUE (name);


--
-- Name: auth_group_permissions auth_group_permissions_group_id_permission_id_0cd325b0_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_permission_id_0cd325b0_uniq UNIQUE (group_id, permission_id);


--
-- Name: auth_group_permissions auth_group_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);


--
-- Name: auth_group auth_group_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);


--
-- Name: auth_permission auth_permission_content_type_id_codename_01ab375a_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_codename_01ab375a_uniq UNIQUE (content_type_id, codename);


--
-- Name: auth_permission auth_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);


--
-- Name: auth_user_groups auth_user_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_pkey PRIMARY KEY (id);


--
-- Name: auth_user_groups auth_user_groups_user_id_group_id_94350c0c_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_group_id_94350c0c_uniq UNIQUE (user_id, group_id);


--
-- Name: auth_user auth_user_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_pkey PRIMARY KEY (id);


--
-- Name: auth_user_user_permissions auth_user_user_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_pkey PRIMARY KEY (id);


--
-- Name: auth_user_user_permissions auth_user_user_permissions_user_id_permission_id_14a6b632_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_permission_id_14a6b632_uniq UNIQUE (user_id, permission_id);


--
-- Name: auth_user auth_user_username_key; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_username_key UNIQUE (username);


--
-- Name: django_admin_log django_admin_log_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);


--
-- Name: django_content_type django_content_type_app_label_model_76bd3d3b_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_app_label_model_76bd3d3b_uniq UNIQUE (app_label, model);


--
-- Name: django_content_type django_content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);


--
-- Name: django_migrations django_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_migrations
    ADD CONSTRAINT django_migrations_pkey PRIMARY KEY (id);


--
-- Name: django_session django_session_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);


--
-- Name: fragrance_inventory fragrance_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory
    ADD CONSTRAINT fragrance_inventory_pkey PRIMARY KEY (id);


--
-- Name: fragrance fragrance_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance
    ADD CONSTRAINT fragrance_pkey PRIMARY KEY (id);


--
-- Name: lipid_inventory lipid_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory
    ADD CONSTRAINT lipid_inventory_pkey PRIMARY KEY (id);


--
-- Name: lipid lipid_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid
    ADD CONSTRAINT lipid_pkey PRIMARY KEY (id);


--
-- Name: lye_inventory lye_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory
    ADD CONSTRAINT lye_inventory_pkey PRIMARY KEY (id);


--
-- Name: lye lye_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye
    ADD CONSTRAINT lye_pkey PRIMARY KEY (id);


--
-- Name: recipe_additive recipe_additive_additive_id_676ee4ae_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_additive_id_676ee4ae_uniq UNIQUE (additive_id);


--
-- Name: recipe_additive recipe_additive_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_additive recipe_batch_additive_additive_id_2361c625_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_additive_id_2361c625_uniq UNIQUE (additive_id);


--
-- Name: recipe_batch_additive recipe_batch_additive_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_fragrance_id_c6546fe6_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_fragrance_id_c6546fe6_uniq UNIQUE (fragrance_id);


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_lipid recipe_batch_lipid_lipid_id_d8453e8a_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_lipid_id_d8453e8a_uniq UNIQUE (lipid_id);


--
-- Name: recipe_batch_lipid recipe_batch_lipid_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_lye recipe_batch_lye_lye_id_784d24d9_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_lye_id_784d24d9_uniq UNIQUE (lye_id);


--
-- Name: recipe_batch_lye recipe_batch_lye_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_note recipe_batch_note_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_note
    ADD CONSTRAINT recipe_batch_note_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch recipe_batch_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch
    ADD CONSTRAINT recipe_batch_pkey PRIMARY KEY (id);


--
-- Name: recipe_fragrance recipe_fragrance_fragrance_id_0dcf3ef6_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_fragrance_id_0dcf3ef6_uniq UNIQUE (fragrance_id);


--
-- Name: recipe_fragrance recipe_fragrance_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_pkey PRIMARY KEY (id);


--
-- Name: recipe_lipid recipe_lipid_lipid_id_76650ba8_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_lipid_id_76650ba8_uniq UNIQUE (lipid_id);


--
-- Name: recipe_lipid recipe_lipid_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_pkey PRIMARY KEY (id);


--
-- Name: recipe recipe_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe
    ADD CONSTRAINT recipe_pkey PRIMARY KEY (id);


--
-- Name: recipe_step recipe_step_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_step
    ADD CONSTRAINT recipe_step_pkey PRIMARY KEY (id);


--
-- Name: supplier supplier_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.supplier
    ADD CONSTRAINT supplier_pkey PRIMARY KEY (id);


--
-- Name: additive_inventory_additive_id_390dfc35; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX additive_inventory_additive_id_390dfc35 ON public.additive_inventory USING btree (additive_id);


--
-- Name: additive_inventory_supplier_id_dc5c2c7b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX additive_inventory_supplier_id_dc5c2c7b ON public.additive_inventory USING btree (supplier_id);


--
-- Name: auth_group_name_a6ea08ec_like; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_group_name_a6ea08ec_like ON public.auth_group USING btree (name varchar_pattern_ops);


--
-- Name: auth_group_permissions_group_id_b120cbf9; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_group_permissions_group_id_b120cbf9 ON public.auth_group_permissions USING btree (group_id);


--
-- Name: auth_group_permissions_permission_id_84c5c92e; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_group_permissions_permission_id_84c5c92e ON public.auth_group_permissions USING btree (permission_id);


--
-- Name: auth_permission_content_type_id_2f476e4b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_permission_content_type_id_2f476e4b ON public.auth_permission USING btree (content_type_id);


--
-- Name: auth_user_groups_group_id_97559544; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_groups_group_id_97559544 ON public.auth_user_groups USING btree (group_id);


--
-- Name: auth_user_groups_user_id_6a12ed8b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_groups_user_id_6a12ed8b ON public.auth_user_groups USING btree (user_id);


--
-- Name: auth_user_user_permissions_permission_id_1fbb5f2c; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_user_permissions_permission_id_1fbb5f2c ON public.auth_user_user_permissions USING btree (permission_id);


--
-- Name: auth_user_user_permissions_user_id_a95ead1b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_user_permissions_user_id_a95ead1b ON public.auth_user_user_permissions USING btree (user_id);


--
-- Name: auth_user_username_6821ab7c_like; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_username_6821ab7c_like ON public.auth_user USING btree (username varchar_pattern_ops);


--
-- Name: django_admin_log_content_type_id_c4bce8eb; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_admin_log_content_type_id_c4bce8eb ON public.django_admin_log USING btree (content_type_id);


--
-- Name: django_admin_log_user_id_c564eba6; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_admin_log_user_id_c564eba6 ON public.django_admin_log USING btree (user_id);


--
-- Name: django_session_expire_date_a5c62663; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_session_expire_date_a5c62663 ON public.django_session USING btree (expire_date);


--
-- Name: django_session_session_key_c0390e0f_like; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_session_session_key_c0390e0f_like ON public.django_session USING btree (session_key varchar_pattern_ops);


--
-- Name: fragrance_inventory_fragrance_id_9f202030; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX fragrance_inventory_fragrance_id_9f202030 ON public.fragrance_inventory USING btree (fragrance_id);


--
-- Name: fragrance_inventory_supplier_id_90b4caaf; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX fragrance_inventory_supplier_id_90b4caaf ON public.fragrance_inventory USING btree (supplier_id);


--
-- Name: lipid_inventory_lipid_id_4c032624; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lipid_inventory_lipid_id_4c032624 ON public.lipid_inventory USING btree (lipid_id);


--
-- Name: lipid_inventory_supplier_id_62914da1; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lipid_inventory_supplier_id_62914da1 ON public.lipid_inventory USING btree (supplier_id);


--
-- Name: lye_inventory_lye_id_5e9da65f; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lye_inventory_lye_id_5e9da65f ON public.lye_inventory USING btree (lye_id);


--
-- Name: lye_inventory_supplier_id_78e9941f; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lye_inventory_supplier_id_78e9941f ON public.lye_inventory USING btree (supplier_id);


--
-- Name: recipe_additive_recipe_id_48b68995; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_additive_recipe_id_48b68995 ON public.recipe_additive USING btree (recipe_id);


--
-- Name: recipe_batch_additive_batch_id_d26265ff; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_additive_batch_id_d26265ff ON public.recipe_batch_additive USING btree (batch_id);


--
-- Name: recipe_batch_fragrance_batch_id_a36d24d9; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_fragrance_batch_id_a36d24d9 ON public.recipe_batch_fragrance USING btree (batch_id);


--
-- Name: recipe_batch_lipid_batch_id_b292008e; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_lipid_batch_id_b292008e ON public.recipe_batch_lipid USING btree (batch_id);


--
-- Name: recipe_batch_lye_batch_id_db6fa60b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_lye_batch_id_db6fa60b ON public.recipe_batch_lye USING btree (batch_id);


--
-- Name: recipe_batch_note_batch_id_46a82deb; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_note_batch_id_46a82deb ON public.recipe_batch_note USING btree (batch_id);


--
-- Name: recipe_batch_recipe_id_60edb3ae; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_recipe_id_60edb3ae ON public.recipe_batch USING btree (recipe_id);


--
-- Name: recipe_fragrance_recipe_id_404ec2b0; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_fragrance_recipe_id_404ec2b0 ON public.recipe_fragrance USING btree (recipe_id);


--
-- Name: recipe_lipid_recipe_id_d2a52df1; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_lipid_recipe_id_d2a52df1 ON public.recipe_lipid USING btree (recipe_id);


--
-- Name: recipe_step_recipe_id_bb16b3a0; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_step_recipe_id_bb16b3a0 ON public.recipe_step USING btree (recipe_id);


--
-- Name: additive_inventory additive_inventory_additive_id_390dfc35_fk_additive_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory
    ADD CONSTRAINT additive_inventory_additive_id_390dfc35_fk_additive_id FOREIGN KEY (additive_id) REFERENCES public.additive(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: additive_inventory additive_inventory_supplier_id_dc5c2c7b_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory
    ADD CONSTRAINT additive_inventory_supplier_id_dc5c2c7b_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_group_permissions auth_group_permissio_permission_id_84c5c92e_fk_auth_perm; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissio_permission_id_84c5c92e_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_group_permissions auth_group_permissions_group_id_b120cbf9_fk_auth_group_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_b120cbf9_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_permission auth_permission_content_type_id_2f476e4b_fk_django_co; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_2f476e4b_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_groups auth_user_groups_group_id_97559544_fk_auth_group_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_group_id_97559544_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_groups auth_user_groups_user_id_6a12ed8b_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_6a12ed8b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_user_permissions auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_user_permissions auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: django_admin_log django_admin_log_content_type_id_c4bce8eb_fk_django_co; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_content_type_id_c4bce8eb_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: django_admin_log django_admin_log_user_id_c564eba6_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_user_id_c564eba6_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: fragrance_inventory fragrance_inventory_fragrance_id_9f202030_fk_fragrance_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory
    ADD CONSTRAINT fragrance_inventory_fragrance_id_9f202030_fk_fragrance_id FOREIGN KEY (fragrance_id) REFERENCES public.fragrance(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: fragrance_inventory fragrance_inventory_supplier_id_90b4caaf_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory
    ADD CONSTRAINT fragrance_inventory_supplier_id_90b4caaf_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lipid_inventory lipid_inventory_lipid_id_4c032624_fk_lipid_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory
    ADD CONSTRAINT lipid_inventory_lipid_id_4c032624_fk_lipid_id FOREIGN KEY (lipid_id) REFERENCES public.lipid(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lipid_inventory lipid_inventory_supplier_id_62914da1_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory
    ADD CONSTRAINT lipid_inventory_supplier_id_62914da1_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lye_inventory lye_inventory_lye_id_5e9da65f_fk_lye_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory
    ADD CONSTRAINT lye_inventory_lye_id_5e9da65f_fk_lye_id FOREIGN KEY (lye_id) REFERENCES public.lye(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lye_inventory lye_inventory_supplier_id_78e9941f_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory
    ADD CONSTRAINT lye_inventory_supplier_id_78e9941f_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_additive recipe_additive_additive_id_676ee4ae_fk_additive_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_additive_id_676ee4ae_fk_additive_id FOREIGN KEY (additive_id) REFERENCES public.additive(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_additive recipe_additive_recipe_id_48b68995_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_recipe_id_48b68995_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_additive recipe_batch_additive_additive_id_2361c625_fk_additive_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_additive_id_2361c625_fk_additive_id FOREIGN KEY (additive_id) REFERENCES public.additive(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_additive recipe_batch_additive_batch_id_d26265ff_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_batch_id_d26265ff_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_batch_id_a36d24d9_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_batch_id_a36d24d9_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_fragrance_id_c6546fe6_fk_fragrance_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_fragrance_id_c6546fe6_fk_fragrance_id FOREIGN KEY (fragrance_id) REFERENCES public.fragrance(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lipid recipe_batch_lipid_batch_id_b292008e_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_batch_id_b292008e_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lipid recipe_batch_lipid_lipid_id_d8453e8a_fk_lipid_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_lipid_id_d8453e8a_fk_lipid_id FOREIGN KEY (lipid_id) REFERENCES public.lipid(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lye recipe_batch_lye_batch_id_db6fa60b_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_batch_id_db6fa60b_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lye recipe_batch_lye_lye_id_784d24d9_fk_lye_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_lye_id_784d24d9_fk_lye_id FOREIGN KEY (lye_id) REFERENCES public.lye(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_note recipe_batch_note_batch_id_46a82deb_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_note
    ADD CONSTRAINT recipe_batch_note_batch_id_46a82deb_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch recipe_batch_recipe_id_60edb3ae_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch
    ADD CONSTRAINT recipe_batch_recipe_id_60edb3ae_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_fragrance recipe_fragrance_fragrance_id_0dcf3ef6_fk_fragrance_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_fragrance_id_0dcf3ef6_fk_fragrance_id FOREIGN KEY (fragrance_id) REFERENCES public.fragrance(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_fragrance recipe_fragrance_recipe_id_404ec2b0_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_recipe_id_404ec2b0_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_lipid recipe_lipid_lipid_id_76650ba8_fk_lipid_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_lipid_id_76650ba8_fk_lipid_id FOREIGN KEY (lipid_id) REFERENCES public.lipid(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_lipid recipe_lipid_recipe_id_d2a52df1_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_recipe_id_d2a52df1_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_step recipe_step_recipe_id_bb16b3a0_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_step
    ADD CONSTRAINT recipe_step_recipe_id_bb16b3a0_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- PostgreSQL database dump complete
--

