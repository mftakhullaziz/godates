# Add to packages

INSERT INTO packages
(package_name, description, package_duration_in_monthly, price, unlimited_swipes, status)
VALUES ('Basic Package', 'Access to basic features for one month', 1, 99999, 1, 1),
       ('Standard Package', 'Access to standard features for three months', 3, 249999, 1, 1),
       ('Premium Package', 'Access to all features including unlimited swipes for six months', 6, 499999, 1, 1)