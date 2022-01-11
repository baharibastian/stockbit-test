SELECT `ID`, `UserName`, (SELECT pu.`UserName` FROM `USER` pu WHERE pu.`ID` = u.`Parent`) AS `ParentUserName`
FROM `USER` u;