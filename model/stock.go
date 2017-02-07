package model

import (
	"time"
)

type StockList struct {
	ID             int
	StockName      string
	StockId        string
	STOCKCHINANAME string
	CREATEDAT      time.Time
	UPDATEAT       time.Time
}

//`
//CREATE TABLE `stockLists` (
//  `ID` int(100) NOT NULL AUTO_INCREMENT,
//  `STOCKID` varchar(20) NOT NULL DEFAULT ' ',
//  `STOCKCHINANAME` varchar(20) DEFAULT NULL,
//  `STOCKNAME` varchar(20) DEFAULT NULL,
//  `STOCKUNIQUE` varchar(30) NOT NULL DEFAULT ' ',
//  `STOCKCON` varchar(10) NOT NULL DEFAULT ' ',
//  `STOCKCONSHORT` varchar(10) NOT NULL DEFAULT ' ',
//  `STOCKORG` varchar(10) NOT NULL DEFAULT ' ',
//  `UPDATEAT` datetime DEFAULT CURRENT_TIMESTAMP ,
//  PRIMARY KEY (`ID`) ,
//  UNIQUE KEY (`STOCKUNIQUE`)
//) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
//`
//
//
//`CREATE TABLE `stockCollections` (
//  `ID` int(100) NOT NULL AUTO_INCREMENT,
//  `STOCKID` varchar(20) NOT NULL DEFAULT ' ',
//  `OPENATCASH` float NOT NULL DEFAULT '0',
//  `CLOSEATCASH` float NOT NULL DEFAULT '0',
//  `MAXATCASH` float NOT NULL DEFAULT '0',
//  `MINATCASH` float NOT NULL DEFAULT '0',
//  `TRADECOUNT` decimal(18,2) NOT NULL DEFAULT '0.00',
//  `DATE` int(64) NOT NULL DEFAULT '0',
//  `STOCKUNIQUE` varchar(50) NOT NULL DEFAULT ' ',
//  `STOCKCOLLECTIONUNIQUE` varchar(100) NOT NULL DEFAULT ' ',
//  `UPDATEAT` datetime DEFAULT CURRENT_TIMESTAMP,
//  PRIMARY KEY (`ID`),
//  UNIQUE KEY `STOCKCOLLECTIONUNIQUE` (`STOCKCOLLECTIONUNIQUE`),
//  KEY `STOCKID` (`STOCKID`),
//  KEY `DATE` (`DATE`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
//`


//INSERT ignore INTO stockLists(STOCKID, STOCKCHINANAME, STOCKNAME, STOCKUNIQUE, STOCKCON, STOCKCONSHORT, STOCKORG) values('000001','上证指数','shzs','shStock000001','股票','Stock','sh')


//SELECT * FROM `stockCollections` s WHERE s.DATE BETWEEN '2015-01-01' AND '2016-01-01' AND s.STOCKUNIQUE = (SELECT STOCKUNIQUE FROM `stockLists` sl WHERE sl.STOCKID = '600123' AND sl.STOCKORG = 'sh' AND sl.STOCKCONSHORT = 'Stock')

// 5 days  http://web.ifzq.gtimg.cn/appstock/app/day/query?_var=fdays_data_sh000043&code=sh000043&r=0.6637066903920745
// http://web.ifzq.gtimg.cn/appstock/app/minute/query?_var=min_data_sh000043&code=sh000043&r=0.21536652074876295



