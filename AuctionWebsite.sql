PGDMP  "                    |            AuctionWebsite     16.1 (Ubuntu 16.1-1.pgdg22.04+1)     16.1 (Ubuntu 16.1-1.pgdg22.04+1)     M           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            N           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            O           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            P           1262    16460    AuctionWebsite    DATABASE     v   CREATE DATABASE "AuctionWebsite" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_IN';
     DROP DATABASE "AuctionWebsite";
                postgres    false            �            1259    16461    artworks    TABLE     B  CREATE TABLE public.artworks (
    id uuid NOT NULL,
    highest_bid uuid,
    name character varying,
    image character varying,
    starting_price double precision,
    category_id uuid,
    closing_time character varying,
    owner_id uuid,
    created_at time without time zone,
    description character varying
);
    DROP TABLE public.artworks;
       public         heap    postgres    false            �            1259    16466    bids    TABLE     �   CREATE TABLE public.bids (
    id uuid NOT NULL,
    artwork_id uuid,
    amount double precision,
    status uuid,
    bidder_id uuid,
    created_at time without time zone
);
    DROP TABLE public.bids;
       public         heap    postgres    false            �            1259    16469 	   bidstatus    TABLE     T   CREATE TABLE public.bidstatus (
    id uuid NOT NULL,
    name character varying
);
    DROP TABLE public.bidstatus;
       public         heap    postgres    false            �            1259    16474    category    TABLE     S   CREATE TABLE public.category (
    id uuid NOT NULL,
    name character varying
);
    DROP TABLE public.category;
       public         heap    postgres    false            �            1259    16479    role    TABLE     X   CREATE TABLE public.role (
    id uuid NOT NULL,
    name character varying NOT NULL
);
    DROP TABLE public.role;
       public         heap    postgres    false            �            1259    16484    users    TABLE       CREATE TABLE public.users (
    id uuid NOT NULL,
    firstname character varying NOT NULL,
    lastname character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    created_at timestamp without time zone,
    role_id uuid
);
    DROP TABLE public.users;
       public         heap    postgres    false            E          0    16461    artworks 
   TABLE DATA           �   COPY public.artworks (id, highest_bid, name, image, starting_price, category_id, closing_time, owner_id, created_at, description) FROM stdin;
    public          postgres    false    215   �       F          0    16466    bids 
   TABLE DATA           U   COPY public.bids (id, artwork_id, amount, status, bidder_id, created_at) FROM stdin;
    public          postgres    false    216   �       G          0    16469 	   bidstatus 
   TABLE DATA           -   COPY public.bidstatus (id, name) FROM stdin;
    public          postgres    false    217   �        H          0    16474    category 
   TABLE DATA           ,   COPY public.category (id, name) FROM stdin;
    public          postgres    false    218   {!       I          0    16479    role 
   TABLE DATA           (   COPY public.role (id, name) FROM stdin;
    public          postgres    false    219   �!       J          0    16484    users 
   TABLE DATA           ^   COPY public.users (id, firstname, lastname, email, password, created_at, role_id) FROM stdin;
    public          postgres    false    220   u"       �           2606    16490    artworks artworks_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.artworks
    ADD CONSTRAINT artworks_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.artworks DROP CONSTRAINT artworks_pkey;
       public            postgres    false    215            �           2606    16492    bids bids_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.bids
    ADD CONSTRAINT bids_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.bids DROP CONSTRAINT bids_pkey;
       public            postgres    false    216            �           2606    16494    bidstatus bidstatus_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.bidstatus
    ADD CONSTRAINT bidstatus_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.bidstatus DROP CONSTRAINT bidstatus_pkey;
       public            postgres    false    217            �           2606    16496    category category_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.category DROP CONSTRAINT category_pkey;
       public            postgres    false    218            �           2606    16498    role role_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.role DROP CONSTRAINT role_pkey;
       public            postgres    false    219            �           2606    16500    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    220            �           2606    16501 "   artworks artworks_highest_bid_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.artworks
    ADD CONSTRAINT artworks_highest_bid_fkey FOREIGN KEY (highest_bid) REFERENCES public.bids(id);
 L   ALTER TABLE ONLY public.artworks DROP CONSTRAINT artworks_highest_bid_fkey;
       public          postgres    false    3240    216    215            �           2606    16506    artworks artworks_owner_id_fkey    FK CONSTRAINT        ALTER TABLE ONLY public.artworks
    ADD CONSTRAINT artworks_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES public.users(id);
 I   ALTER TABLE ONLY public.artworks DROP CONSTRAINT artworks_owner_id_fkey;
       public          postgres    false    215    3248    220            �           2606    16511    bids bids_artwork_id_fkey    FK CONSTRAINT     ~   ALTER TABLE ONLY public.bids
    ADD CONSTRAINT bids_artwork_id_fkey FOREIGN KEY (artwork_id) REFERENCES public.artworks(id);
 C   ALTER TABLE ONLY public.bids DROP CONSTRAINT bids_artwork_id_fkey;
       public          postgres    false    215    216    3238            �           2606    16516    bids bids_bidder_id_fkey    FK CONSTRAINT     y   ALTER TABLE ONLY public.bids
    ADD CONSTRAINT bids_bidder_id_fkey FOREIGN KEY (bidder_id) REFERENCES public.users(id);
 B   ALTER TABLE ONLY public.bids DROP CONSTRAINT bids_bidder_id_fkey;
       public          postgres    false    216    3248    220            �           2606    16521    users users_role_id_fkey    FK CONSTRAINT     v   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.role(id);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_role_id_fkey;
       public          postgres    false    219    220    3246            E   ^  x�Ց�n�@E�ٯP� �r^�� q�f�^%�V���{l8�-Ҹ4A� x��C��63��
[k p����9��-�`]%"�.'`Eb�5��2����_e��&��Q��[ô��,�Ǳ��0$�a�װ���a+}������4�`�i�,�����\`�;��z���M	�V^�r���Bba��!�`�a��H(�l�Uk�Z(T�@r'�(i$�5�9F��GBQ1s6QA��C�|cd�
F�wRH}����<�]���H���<m���y�):t\k�����8e��%oR���e�Q�X���S��	>G���]�H�5d�FԽRF)o��p�S����S:�� :��%      F   �   x���ˍ1D�ݹ��|�ݱ���?��V��RH�B�vf�#q�ز�6m)s�}��m�.Kav�̱҄/�xi,ܺ��@��T�RM��j6�VjS�%a �8�C�N^�����/���>rw+�8V45@\&��j�Gw�~��!U>�Y�G��x���4�E�Ϥ�"��i;�4rz��?U�Uv"��)��/I)X��C�.�m�#a�EL�9��%�W�n���y����\~      G   r   x��1�0й�˯��N⻰��#�B��?==�R��v�q��b,�Ib��i{޿YN���� ����X*^O�<"_��%1�h]�R�8��<2��h�흏�K��R���$z      H   f   x��Q�0 ��q+��;���&��x~�3S6>){bUe4���-�{[6��7���7�U7�Z�c㒘�W�]d��l9}����� �z?      I   t   x�%�A
! ���%E����BIb=��]��-�>�BDL�%��p��\Z3�k�t�~�=�N��Q~rA�`^bђ���}�����A��0Wݰ�\�֨��>~<��~�9r�&�      J   �  x���IS"KF��+z�6˜VO�Af�ё��L1
�����nU��|����"D���b��my��#���V~��s��"�9׿����m�ǳ�.��o��N�y��ʒ�a��TMm�[�*�I{�ͻe�F��Fe����q��0�@��X��	�TB��ib����R!���2<�-"(	�	Pi0�*@�W�X���ͅ�}��K�Ef����?���Sv(z/�-l��iZ����0aV��>�ceԚ%u��#\�8�Lq����Q����:^�C%Q��g�8�3������b}����z��o��6�CFR<j[��[{h�[���d���tT�K|���$�5c�3���$�x���(�P�:���jD�%��i`bt>�8��عO|�Q����R��f��ԩ�n)K��s8֓1�/�z{*�*�T���cy8���>�D�)���K˴�؁�KG�]4�􌡳?6|7�_���~Q�WF���pt{����6kn��}�n�kÓ��Z�7���r��j��aU�d' �_��?AJH,�3��H�=���9��8E<d�;<��'���TH�Ud+[���q7C��t�衂���}�xx'�4���d�^�B� a˄rս�F�� ��E�����`,a��ϛ�a����=�{�r���̖u;�o^K����fx���k����|~k�bh^�	��.�J0�$��U�P�5�xk,TYtpw�:cTT�o�~������X^&�v�D=�y_�~:��[ϿN�ȼ��tߨ�v�ޯ&��i\���K@XO"���["5���'�D;�`@�y�\ip޻`-�O�s<�
�h�q{��È���_�/~�^-��0���M�m�/%ty��'��J��]e�N�P��������o* �p_>y]��|>����     