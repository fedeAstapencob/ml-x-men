/* the unique key is creating an index */
CREATE TABLE `person` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                          `is_mutant` tinyint(1) unsigned NOT NULL DEFAULT '0',
                          `dna` varchar(255) NOT NULL,
                          `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          UNIQUE (`dna`)
);


