LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;

INSERT INTO `books` (`id`, `title`, `author`, `price`)
VALUES
	(1,'Gin-Golang','google',2000),
	(2,'scala','scala',2000),
	(3,'ruby','ruby',100),
	(4,'python','python',2000),
	(5,'rust','rust',4000);

/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;