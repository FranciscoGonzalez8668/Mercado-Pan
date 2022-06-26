-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: localhost    Database: pan
-- ------------------------------------------------------
-- Server version	5.7.37-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `adresses`
--

DROP TABLE IF EXISTS `adresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `adresses` (
  `id_adress` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `street_name` varchar(350) NOT NULL,
  `number` int(11) DEFAULT NULL,
  `neighborhood` varchar(350) DEFAULT NULL,
  `city` varchar(350) NOT NULL,
  PRIMARY KEY (`id_adress`,`id_user`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `adresses_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `adresses`
--

LOCK TABLES `adresses` WRITE;
/*!40000 ALTER TABLE `adresses` DISABLE KEYS */;
INSERT INTO `adresses` VALUES (1,1,'Santa Rosa',1228,'Alberdi','Cordoba'),(2,2,'Av Republica de China',1056,'Los sueños, Valle Escondido','Villa Allende'),(3,3,'Caseros',3803,'San Salvador','Cordoba');
/*!40000 ALTER TABLE `adresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `detalles`
--

DROP TABLE IF EXISTS `detalles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `detalles` (
  `id_order` int(11) NOT NULL,
  `id_detalle` int(11) NOT NULL,
  `id_product` int(11) DEFAULT NULL,
  `cantidad` int(11) NOT NULL,
  `precio_moment` float NOT NULL,
  PRIMARY KEY (`id_order`,`id_detalle`),
  KEY `id_product` (`id_product`),
  CONSTRAINT `detalles_ibfk_1` FOREIGN KEY (`id_order`) REFERENCES `orders` (`id_order`),
  CONSTRAINT `detalles_ibfk_2` FOREIGN KEY (`id_product`) REFERENCES `products` (`id_product`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `detalles`
--

LOCK TABLES `detalles` WRITE;
/*!40000 ALTER TABLE `detalles` DISABLE KEYS */;
INSERT INTO `detalles` VALUES (1,2,1,1,170000),(1,3,2,1,70000),(2,4,2,2,70000),(2,5,5,1,1500),(2,6,6,2,1000),(2,7,7,1,1200),(3,4,16,1,500),(4,4,2,2,70000),(5,4,1,3,170000),(5,5,2,2,70000),(6,4,1,3,170000),(6,5,2,2,70000),(6,6,3,1,800),(7,4,1,1,170000),(7,5,2,1,70000),(8,4,2,1,70000),(9,4,19,4,600),(9,5,20,2,700);
/*!40000 ALTER TABLE `detalles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id_order` int(11) NOT NULL,
  `id_user` int(11) DEFAULT NULL,
  `total` float NOT NULL,
  PRIMARY KEY (`id_order`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,1,240000),(2,1,144700),(3,1,500),(4,1,140000),(5,1,650000),(6,1,650800),(7,2,240000),(8,1,70000),(9,1,3800);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id_product` int(11) NOT NULL,
  `name_product` varchar(350) NOT NULL,
  `price` float NOT NULL,
  `stock` int(11) NOT NULL,
  `category` varchar(350) NOT NULL,
  `description` varchar(3000) NOT NULL,
  PRIMARY KEY (`id_product`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'ThinkPad E15 2da Gen (15.6 Intel)',170000,12,'Tecnologia','Computadora gamer nashe'),(2,'Jordan Hombres Air Jordan 1 Mid DD1649 001 All-Star 2021',70000,9,'ropa','Altas zapatillas'),(3,'Iphone 12',800,19,'Tecnologia','516gb Negro'),(4,'Nikon Coolpix B500',1200,20,'Camaras','Compacta Avanzada color negro'),(5,'Canon EOS T100',1500,19,'Camaras','55mm DSLR negro'),(6,'Canon EOS M200',1000,18,'Camaras','Sin espejo'),(7,'Xbox Series X',1200,19,'Consolas','1TB color negro'),(8,'Sony Playstation 5',1200,20,'Consolas','Standard'),(9,'Nintendo Switch',800,20,'Consolas','32gb standard'),(10,'Nvidia GeForce 1060',1700,20,'Computacion','8gb Gigabyte'),(11,'Intel I3 10100',900,20,'Computacion','4 nucleos'),(12,'AMD Ryzen 7 5700G',2000,20,'Computacion','8 nucleos'),(13,'Mini drone DJI Mavic',800,20,'Drones','Light gray 3 baterias'),(14,'Drone Samsung A9',400,20,'Drones','+Info Ig MatiasLessio'),(15,'Drone Gadnic',120,20,'Drones','Camara HD'),(16,'Smart TV Samsung',500,19,'Televisores','LED 4K'),(17,'LG Smart TV',400,20,'Televisores','4K 55pulgadas'),(18,'Smart TV Noblex',200,20,'Televisores','HD 32pulgadas'),(19,'Iphone 10',600,16,'Tecnologia','128gb Azul'),(20,'Iphone 11',700,18,'Tecnologia','256gb Verde');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `name` varchar(350) NOT NULL,
  `Lname` varchar(350) NOT NULL,
  `email` varchar(500) NOT NULL,
  `password` varchar(350) NOT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Francisco','Gonzalez','ohcnap_2000','root'),(2,'Tomas','Apparicio','appa','jipjip'),(3,'Nicolas','Prandi','gendarmeria123','glock123');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-23 21:57:09
