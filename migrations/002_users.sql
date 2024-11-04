-- migrate:up
INSERT INTO user_ (id, pub_key, ip) VALUES (1, '64c5b2ea9b34bb7f1e863681cadf4e6937eb196aa2a70abb70cff6f64927cd15', 'http://localhost:8080');
INSERT INTO user_ (id, pub_key, ip) VALUES (2, '0b5716b05d0821bb6543dce44f5ee8eb804de29e5d5524360f6dd5149efe9586', 'http://localhost:8081');
INSERT INTO user_ (id, pub_key, ip) VALUES (3, '7f196b34b8752e02e2ce626ddc557ddbe86884fdc25dd790cea062d54bbb3888', 'http://localhost:8082');
INSERT INTO user_ (id, pub_key, ip) VALUES (4, 'd902aa80477960ff5db1d4d7967c00a063899378e6a11399f6df4630d4667589', 'http://localhost:8083');

-- migrate:down
