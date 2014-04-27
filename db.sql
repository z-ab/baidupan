CREATE TABLE `buk` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  `uk` varchar(15) DEFAULT NULL,
  `state` int(3) unsigned DEFAULT '0',
  `fans` int(11) unsigned DEFAULT '0' COMMENT '粉丝数',
  `album` mediumint(5) unsigned DEFAULT '0' COMMENT '专辑数',
  `follow` mediumint(5) unsigned DEFAULT '0' COMMENT '关注数',
  `share` int(11) unsigned DEFAULT '0',
  `updatetime` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk` (`uk`),
  KEY `share` (`share`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `ip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `area` varchar(64) DEFAULT '',
  `operator` varchar(64) DEFAULT '',
  `status` tinyint(4) DEFAULT '1',
  `addtime` int(11) unsigned DEFAULT '0',
  `updatetime` int(11) unsigned DEFAULT '0',
  `consume` int(11) unsigned DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip` (`ip`),
  KEY `status` (`status`),
  KEY `addtime` (`addtime`),
  KEY `consume` (`consume`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `res` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uk` bigint(11) unsigned DEFAULT '0',
  `shareid` bigint(11) unsigned DEFAULT '0',
  `fid` bigint(11) unsigned DEFAULT '0',
  `category` tinyint(2) unsigned DEFAULT '0',
  `title` varchar(256) DEFAULT '',
  `username` varchar(128) DEFAULT '',
  `dateid` varchar(64) DEFAULT '',
  `file_name` varchar(256) DEFAULT NULL,
  `size` bigint(11) unsigned DEFAULT '0',
  `path` text,
  `md5` varchar(32) DEFAULT '',
  `thumburl` text,
  `sign` varchar(64) DEFAULT '',
  `t_stamp` int(11) unsigned DEFAULT '0',
  `shorturl` varchar(64) DEFAULT '',
  `addtime` int(11) unsigned DEFAULT '0',
  `updatetime` int(11) unsigned DEFAULT '0',
  `source_uid` int(11) unsigned DEFAULT '0',
  `source_id` int(11) unsigned DEFAULT '0',
  `vcnt` int(11) unsigned DEFAULT '0',
  `dcnt` int(11) unsigned DEFAULT '0',
  `tcnt` int(11) DEFAULT NULL,
  `like_status` tinyint(3) DEFAULT '0',
  `like_count` tinyint(3) DEFAULT '0',
  `comment_count` int(11) DEFAULT '0',
  `status` tinyint(3) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `usf` (`uk`,`shareid`,`fid`) USING BTREE,
  KEY `md5` (`md5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;