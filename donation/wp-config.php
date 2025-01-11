<?php
/**
 * The base configuration for WordPress
 *
 * The wp-config.php creation script uses this file during the installation.
 * You don't have to use the web site, you can copy this file to "wp-config.php"
 * and fill in the values.
 *
 * This file contains the following configurations:
 *
 * * Database settings
 * * Secret keys
 * * Database table prefix
 * * ABSPATH
 *
 * @link https://wordpress.org/documentation/article/editing-wp-config-php/
 *
 * @package WordPress
 */

// ** Database settings - You can get this info from your web host ** //
/** The name of the database for WordPress */
define( 'DB_NAME', 'donation' );

/** Database username */
define( 'DB_USER', 'root' );

/** Database password */
define( 'DB_PASSWORD', '' );

/** Database hostname */
define( 'DB_HOST', 'localhost' );

/** Database charset to use in creating database tables. */
define( 'DB_CHARSET', 'utf8mb4' );

/** The database collate type. Don't change this if in doubt. */
define( 'DB_COLLATE', '' );

/**#@+
 * Authentication unique keys and salts.
 *
 * Change these to different unique phrases! You can generate these using
 * the {@link https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service}.
 *
 * You can change these at any point in time to invalidate all existing cookies.
 * This will force all users to have to log in again.
 *
 * @since 2.6.0
 */
define( 'AUTH_KEY',         '>^FU+E!^,s8,,zxCv8bX<%]$(WdBQiQ0c-v{.Y4X%i@`pHP ?~D!8rPt_k0}r%?~' );
define( 'SECURE_AUTH_KEY',  'ruRKNuCx1u5 M.1dgUM<A/pgb*YwklM?88rM9`ojZp@B4ZgGkVya4=._([PSsGw*' );
define( 'LOGGED_IN_KEY',    '{;5[I]V^444i,4G|U@X{^H,1}c)2wK>D35yCo_7a//)I[.zLC?OZ+n0476BMAeFA' );
define( 'NONCE_KEY',        'I}A8M~6R1Q4bi>>-=kK?|BleA&~lU@SqoF<ms*0hgvW5[hoSkpAuciog`#poIEBn' );
define( 'AUTH_SALT',        'E)E/0%7[#v&J1~8FYhdISs>9w88DJ?_FI*-ME.[$uohLe6Scc;v=~%+x4Uz)ry+|' );
define( 'SECURE_AUTH_SALT', 'EpLh2ICWE#)PEa91`@KnfWt3/fn/56%%[0#iChWxzo*a(rd!gp[,Kgg)V0N,A!c0' );
define( 'LOGGED_IN_SALT',   '[Voy;#iREsu_/: 40!$+Q<ZosL?%hs8.^u8)_R3ZZV1z8#)L$6^;$!S|)WrQ&X<m' );
define( 'NONCE_SALT',       'whd0{}vUU|dOi{{fOJ2@BBg:P*Ysc#!=a2 !=*R{X%_d$8Sk+`L_ Q%i5$C!7SQU' );

/**#@-*/

/**
 * WordPress database table prefix.
 *
 * You can have multiple installations in one database if you give each
 * a unique prefix. Only numbers, letters, and underscores please!
 */
$table_prefix = 'dnn_';

/**
 * For developers: WordPress debugging mode.
 *
 * Change this to true to enable the display of notices during development.
 * It is strongly recommended that plugin and theme developers use WP_DEBUG
 * in their development environments.
 *
 * For information on other constants that can be used for debugging,
 * visit the documentation.
 *
 * @link https://wordpress.org/documentation/article/debugging-in-wordpress/
 */
define( 'WP_DEBUG', false );

/* Add any custom values between this line and the "stop editing" line. */



/* That's all, stop editing! Happy publishing. */

/** Absolute path to the WordPress directory. */
if ( ! defined( 'ABSPATH' ) ) {
	define( 'ABSPATH', __DIR__ . '/' );
}

/** Sets up WordPress vars and included files. */
require_once ABSPATH . 'wp-settings.php';
